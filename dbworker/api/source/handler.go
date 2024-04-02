package source

import (
	sourceModel "dbworker/models/sources"
	"github.com/gofiber/fiber/v2"
	"log"
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
	group.Post("/get", res.getUserSources)
}

func RegisterHandlers(app *fiber.App) {
	//res, group := getGroupAndResources(app)
	//
	//group.Get("/get", res.getSource)
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

	log.Println(foundSources)

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
