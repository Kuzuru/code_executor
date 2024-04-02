package history

import (
	historyModel "dbworker/models/history"
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

	group := app.Group("/history")

	return res, group
}

func RegisterProtectedHandlers(app *fiber.App) {
	res, group := getGroupAndResources(app)

	group.Post("/get", res.getSourceHistoryPoints)
}

func (res *resource) getSourceHistoryPoints(c *fiber.Ctx) error {
	type Request struct {
		SourceID string `json:"sources"`
		Limit    int64  `json:"limit"`
		Offset   int64  `json:"offset"`
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	history, err := historyModel.GetUserHistoryPoints(req.SourceID, req.Limit, req.Offset)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Не удалось получить историю исходников",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"history": history,
	})
}
