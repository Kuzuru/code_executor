package main

import (
	"os"
	"os/signal"
	"syscall"

	"dbworker/api/user"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	log.Info(".env file loaded successfully")

	buildAndListen()

	// Waiting for exit signal from OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	<-quit
}

func buildAndListen() {
	log.Info("Starting cmd...")

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
