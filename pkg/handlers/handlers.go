package handlers

import (
	"github.com/gofiber/fiber/v2"
	"dafilsys/internal/models"
)

func GetDocumentsList(c *fiber.Ctx) error{
	
	docs := []models.Document{
		{
			Id: 1,
			Name: "Document 1",
			Data: []byte{0x1, 0x2},
		},
		{
			Id: 2,
			Name: "Document 2",
			Data: []byte{0x3, 0x4},
		},
		{
			Id: 3,
			Name: "Document 3",
			Data: []byte{0x5, 0x6},
		},
	}

	return c.JSON(docs)

}

func DocumentAnalysis(c *fiber.Ctx){
	// TODO
}

func DocumentSearch(c *fiber.Ctx) {
	// TODO
}

func DocumentRegistration(c *fiber.Ctx){
	// TODO
}