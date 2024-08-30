package handler

import (
	"fiber-api/database"
	"fiber-api/model/entity"
	"fiber-api/model/request"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// POST : Create book
func BookHandlerCreate(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)

	// Parser
	if err := ctx.BodyParser(book); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	// Validasi request
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	// Handle jika file "required"
	var filenameString string

	// Get filename
	filename := ctx.Locals("filename")
	log.Println("filename = ", filename)

	if filename == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image cover is required.",
		})
	} else {
		// Mengubah filename menjadi string
		filenameString = fmt.Sprintf("%v", filename)
	}

	// New book
	newBook := entity.Book{
		Title: 	book.Title,
		Author: book.Author,
		Cover: 	filenameString,
	}

	errCreateBook := database.DB.Create(&newBook).Error
	if errCreateBook != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": newBook,
	})


}
