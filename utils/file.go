package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/covers/"

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

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		log.Println("Error Read Multipart Form Request, Error = ", errForm)
	}

	files := form.File["photos"]

	var filenames []string

	for i, file := range files {
		var filename string
		// Cek ada file yang diupload atau tidak
		if file != nil {
			filename = fmt.Sprintf("%d-%s", i, file.Filename)

			// Mengembalikan tipe data error
			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))
			if errSaveFile != nil {
				log.Println("Fail to store file into public/covers directory.")
			}
		} else {
			log.Println("There is no file to be uploaded.")
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
	}
	ctx.Locals("filenames", filenames)

	return ctx.Next()
}

func HandleRemoveFile(filename string, pathFile ...string) error {
	if len(pathFile) > 0 {
		err := os.Remove(pathFile[0] + filename)
		if err != nil {
			log.Println("Failed to remove file.")
			return err
		}
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("Failed to remove file.")
			return err
		}
	}

	return nil
}
