package routes

import (
	"github.com/gofiber/fiber/v2"
	"dafilsys/pkg/handlers"
)

func SetupRoutes(app * fiber.app){
	app.Get("", handlers.GetDocumentsList)
}