package routes

import (
	"github.com/gofiber/fiber/v2"
	"dafilsys/pkg/handlers"
)

func SetupRoutes(app * fiber.app){
	app.Get("/documents-list", handlers.GetDocumentsList)
	app.Post("/document-registration", handlers.DocumentRegistration)
	app.Post("/document-analysis", handlers.DocumentSearch)
	app.Post("/document-analysis", handlers.DocumentAnalysis)
}