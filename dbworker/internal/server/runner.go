package server

import (
	"os"

	"dbworker/internal/api/user"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Run() {
	log.Info("Starting server...")

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	})

	app.Use(cors.New())

	// Building all routes and middlewares together
	buildHandlers(app)

	address := ":" + os.Getenv("APP_PORT_HTTP")

	go func() {
		if err := app.Listen(address); err == nil {
			log.Fatal("app.Listen error", err)
		}
	}()

	if !fiber.IsChild() {
		log.Info("Server is listening", address)
	}
}

func buildHandlers(app *fiber.App) {
	user.RegisterHandler(app)
}
