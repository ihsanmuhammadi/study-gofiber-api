package handler

import (
	"fiber-api/database"
	"fiber-api/model/entity"
	"fiber-api/model/request"
	"fiber-api/utils"

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
	// var filenameString string

	// Get filename
	filenames := ctx.Locals("filenames")
	log.Println("filename = ", filenames)

	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image cover is required.",
		})
	} else {
		// Mengubah filename menjadi string
		// filenameString = fmt.Sprintf("%v", filenames)
		filenamesData := filenames.([]string)
		for _, filename := range filenamesData {
			// New photo
			newPhoto := entity.Photo{
				Image: filename,
				CategoryID: photo.CategoryId,
			}

			errCreatePhoto := database.DB.Create(&newPhoto).Error
			if errCreatePhoto != nil {
				log.Println("Some data not saved properly.")
			}
		}
	}

	// log.Println("filenames :: ", filenameString)

	// // New photo
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

func PhotoHandlerDelete(ctx *fiber.Ctx) error {
	photoId := ctx.Params("id")

	var photo entity.Photo

	// Check availability photo
	err := database.DB.Debug().First(&photo, "id=?", photoId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "photo not found",
		})
	}

	// Delete file di directory
	errDeleteFile := utils.HandleRemoveFile(photo.Image)

	if errDeleteFile != nil {
		log.Println("Fail to delete some file")
	}

	// Delete file di database
	errDelete := database.DB.Debug().Delete(&photo).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "photo was deleted",
	})
}
