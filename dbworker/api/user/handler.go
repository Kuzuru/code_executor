package user

import (
	models "dbworker/models/users"

	"github.com/gofiber/fiber/v2"
)

// resource is the structure responsible for representing
// the HTTP request unit for a given package
type resource struct {
	app fiber.Router
}

func RegisterHandler(app *fiber.App) {
	res := resource{
		app: app,
	}

	group := app.Group("/user")

	group.Post("/register", res.register)
}

func (res *resource) register(c *fiber.Ctx) error {
	type Request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var userModel models.User

	userModel.Name = req.Name
	userModel.PasswordHash = req.Password // TODO: BCRYPT [!!!]

	_, err := models.FindUserByName(req.Name)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Пользователь с таким именем уже существует",
		})
	}

	err = models.CreateUser(&userModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(200)
}
