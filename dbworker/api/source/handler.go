package source

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	historyModel "dbworker/models/history"
	sourceModel "dbworker/models/sources"

	"github.com/gofiber/fiber/v2"
)

// resource is the structure responsible for representing
// the HTTP request unit for a given package
type resource struct {
	app fiber.Router
}

func getGroupAndResources(app *fiber.App) (resource, fiber.Router) {
	res := resource{
		app: app,
	}

	group := app.Group("/source")

	return res, group
}

func RegisterProtectedHandlers(app *fiber.App) {
	res, group := getGroupAndResources(app)

	group.Post("/new", res.createNewSource)
	group.Post("/save", res.saveExistingSource)
	group.Post("/get", res.getUserSources)
	group.Post("/run", res.runSource)
}

func RegisterHandlers(app *fiber.App) {
	//res, group := getGroupAndResources(app)
	//
	//group.Get("/get", res.getSource)
}

func (res *resource) saveExistingSource(c *fiber.Ctx) error {
	type Request struct {
		SourceID string `json:"source_id"`
		FileName string `json:"filename"`
		Data     string `json:"data"`
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	foundSources, err := sourceModel.FindSourceByID(req.SourceID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Не удалось найти исходник пользователя",
			"message": err.Error(),
		})
	}

	err = foundSources.UpdateSourceUpdatedAt()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Не удалось обновить время сохранения исходника",
			"message": err.Error(),
		})
	}

	err = foundSources.UpdateSourceData(req.Data, req.FileName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Не удалось обновить новые данные исходника",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"sources": foundSources,
	})
}

func (res *resource) runSource(c *fiber.Ctx) error {
	var req sourceModel.SourceRunDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Marshal the 'files' field into a JSON array
	filesJSON, err := json.Marshal(req.Files)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create a new request body with the 'command' and 'files' JSON arrays
	reqBody := []byte(fmt.Sprintf(`{"command": "%s", "files": %s}`, req.Command, filesJSON))

	resp, err := http.Post("http://backend:3030/run", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Не удалось прочитать тело ответа сервера",
			})
		}

		var bodyMap map[string]interface{}
		if err := json.Unmarshal(body, &bodyMap); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Не удалось декодировать тело ответа сервера",
			})
		}

		return c.Status(resp.StatusCode).JSON(bodyMap)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось прочитать тело ответа сервера",
		})
	}

	// Parse the response body to SourceRunResponse struct
	var sourceRunResponse sourceModel.SourceRunResponse
	if err := json.Unmarshal(bodyBytes, &sourceRunResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось декодировать тело ответа сервера",
		})
	}

	history := new(historyModel.History)

	history.SourceId = req.SourceID
	history.ExitCode = sourceRunResponse.ExitCode
	history.Stdout = sourceRunResponse.Stdout
	history.Stderr = sourceRunResponse.Stderr
	history.OOMKilled = sourceRunResponse.OomKilled
	history.Timeout = sourceRunResponse.Timeout
	history.Duration = sourceRunResponse.Duration

	err = historyModel.CreateHistoryPoint(*history)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Не удалось создать запись истории",
			"message": err.Error(),
		})
	}

	// Updating last run
	foundSource, err := sourceModel.FindSourceByID(history.SourceId)
	if err != nil {
		return errors.New("исходника не существует")
	}

	err = foundSource.UpdateSourceLastRunAt()
	if err != nil {
		return errors.New("не удалось обновить время последнего запуска")
	}

	if len(req.Files) != 0 {
		err = foundSource.UpdateSourceData(req.Files[0].Content, req.Files[0].Name)
		if err != nil {
			return errors.New("не удалось обновить время последнего запуска")
		}
	}

	return c.JSON(sourceRunResponse)
}

func (res *resource) getUserSources(c *fiber.Ctx) error {
	type Request struct {
		UserID string `json:"user_id"`
		Limit  int64  `json:"limit"`
		Offset int64  `json:"offset"`
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	foundSources, err := sourceModel.GetUserSources(req.UserID, req.Limit, req.Offset)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Не удалось найти исходники пользователя",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"sources": foundSources,
	})
}

func (res *resource) createNewSource(c *fiber.Ctx) error {
	var req sourceModel.SourceDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := sourceModel.CreateSource(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Не удалось создать исходник",
			"message": err.Error(),
		})
	}

	return c.SendStatus(200)
}
