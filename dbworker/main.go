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
)

func main() {
	buildAndListen()

	// Waiting for exit signal from OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	<-quit
}

func buildAndListen() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	})

	app.Use(cors.New())

	// Building all routes and middlewares together
	buildHandlers(app)

	address := "0.0.0.0:" + os.Getenv("DBWORKER_PORT")

	go func() {
		if err := app.Listen(address); err == nil {
			log.Fatal("app.Listen error", err)
		}
	}()

	if !fiber.IsChild() {
		log.Info("Server is listening on: ", address)
	}
}

func buildHandlers(app *fiber.App) {
	user.RegisterHandler(app)
}
