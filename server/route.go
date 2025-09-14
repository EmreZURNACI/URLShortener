package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/EmreZURNACI/url-shortener/app"
	"github.com/EmreZURNACI/url-shortener/controller"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Start(repo app.Repository) {

	app := fiber.New()

	_controller := controller.NewRepository(repo)
	app.Post("/shorter", _controller.Shortener)
	app.Get("/:link", _controller.Redirect)

	zap.L().Sugar().Info("server is running")
	if err := app.Listen(":8080"); err != nil {
		zap.L().Sugar().Info("server failed")
	}

	GracefulShutdown(app)

}

func GracefulShutdown(app *fiber.App) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan

	zap.L().Sugar().Info("Shutting down server")

	if err := app.Shutdown(); err != nil {
		zap.L().Sugar().Error("Shutting down server. : %s", zap.Error(err))
	}
	zap.L().Sugar().Info("Server gracefully stopped")

}
