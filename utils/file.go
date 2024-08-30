package utils

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(ctx *fiber.Ctx) error {
	// Handle file
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error file = ", errFile)
	}

	var filename *string
	// Cek ada file yang diupload atau tidak
	if file != nil {
		filename = &file.Filename

		// Mengembalikan tipe data error
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", *filename))
		if errSaveFile != nil {
			log.Println("Fail to store file into public/covers directory.")
		}
	} else {
		log.Println("There is no file to be uploaded.")
	}

	// If-else untuk required image
	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	// ctx.Locals("filename", *filename)

	return ctx.Next()
}
