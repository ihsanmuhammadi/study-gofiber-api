package handler

import (

	"fiber-api/model/request"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PhotoHandlerCreate(ctx *fiber.Ctx) error {
	photo := new(request.PhotoCreateRequest)

	// Parser
	if err := ctx.BodyParser(photo); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	// Validasi request
	validate := validator.New()
	errValidate := validate.Struct(photo)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	// Handle jika file "required"
	var filenameString string

	// Get filename
	filenames := ctx.Locals("filenames")
	log.Println("filename = ", filenames)

	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image cover is required.",
		})
	} else {
		// Mengubah filename menjadi string
		filenameString = fmt.Sprintf("%v", filenames)
	}

	log.Println(filenameString)
	// // New book
	// newPhoto := entity.Photo{
	// 	Image: filename,
	// 	CategoryID: 1,
	// }

	// errCreateBook := database.DB.Create(&newPhoto).Error
	// if errCreateBook != nil {
	// 	return ctx.Status(500).JSON(fiber.Map{
	// 		"message": "failed to store data",
	// 	})
	// }

	return ctx.JSON(fiber.Map{
		"message": "success",
		// "data": newPhoto,
	})

}
