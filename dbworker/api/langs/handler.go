package langs

import (
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

	group := app.Group("/langs")

	return res, group
}

func RegisterHandlers(app *fiber.App) {
	res, group := getGroupAndResources(app)

	group.Get("/getAvailable", res.getAvailable)
}

func (res *resource) getAvailable(c *fiber.Ctx) error {
	available := []string{"GoLang 1.19", "Python 3.11"}

	return c.Status(200).JSON(fiber.Map{
		"data": available,
	})
}
