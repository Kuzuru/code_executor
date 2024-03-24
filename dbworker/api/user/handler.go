package user

import (
	"os"
	"strconv"

	"dbworker/middleware/jwtware"
	models "dbworker/models/users"
	"dbworker/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var ()

// resource is the structure responsible for representing
// the HTTP request unit for a given package
type resource struct {
	app fiber.Router
}

func getGroupAndResources(app *fiber.App) (resource, fiber.Router) {
	res := resource{
		app: app,
	}

	group := app.Group("/user")

	return res, group
}

func RegisterHandlers(app *fiber.App) {
	res, group := getGroupAndResources(app)

	group.Post("/register", res.register)
	group.Post("/auth", res.login)
}

func (res *resource) login(c *fiber.Ctx) error {
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

	foundUser, err := models.FindUserByName(req.Name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный логин или пароль",
		})
	}

	match, err := utils.ComparePasswordAndHash(req.Password, foundUser.PasswordHash)
	if err != nil || !match {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Неверный логин или пароль",
		})
	}

	tokenExpireEnv, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRE_AFTER_IN_SECONDS"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Не удалось получить TOKEN_EXPIRE_AFTER_IN_SECONDS",
		})
	}

	token, err := jwtware.Encode(&jwt.MapClaims{
		"id":   foundUser.ID,
		"name": foundUser.Name,
		"role": foundUser.Role,
	}, int64(tokenExpireEnv))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
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

	_, err := models.FindUserByName(req.Name)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Пользователь с таким именем уже существует",
		})
	}

	hashedPassword, err := utils.GenerateHashFromPassword(req.Password, utils.GetArgonParams())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	var userModel models.User

	userModel.Name = req.Name
	userModel.PasswordHash = hashedPassword

	err = models.CreateUser(&userModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(200)
}
