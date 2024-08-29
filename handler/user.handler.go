package handler

import (
	"fiber-api/database"
	"fiber-api/model/entity"
	"fiber-api/model/request"
	"fiber-api/model/response"
	"fiber-api/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// GET: Mendapatkan seluruh data user
func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)
	// Find -> Select data where DELETED_AT is NULL

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}

// POST: Membuat data user
func UserHandlerCreate(ctx *fiber.Ctx) error  {
	user := new(request.UserCreateRequest)

	// Parser
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Adress: user.Adress,
		Phone: user.Phone,
	}

	// Hashing
	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": newUser,
	})
}

// GET: Get data user by id
func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Membuat format response baru
	userResponse := response.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Adress: user.Adress,
		Phone: user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": userResponse,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User

	userId := ctx.Params("id")
	// Check available user
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Update user data
	// Kondisi agar nama tidak kosong
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Adress = userRequest.Adress
	user.Phone = userRequest.Phone
	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// Membuat format response baru
	userResponse := response.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Adress: user.Adress,
		Phone: user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return  ctx.JSON(fiber.Map{
		"message": "success",
		"data": userResponse,
	})
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User
	var isEmailUserExist entity.User
	userId := ctx.Params("id")

	// Check available user
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Check available email
	errCheckEmail := database.DB.First(&isEmailUserExist, "email = ?", userRequest.Email).Error
	if errCheckEmail == nil {
		return ctx.Status(402).JSON(fiber.Map{
			"message": "Email already used!",
		})
	}

	// Update email
	user.Email = userRequest.Email

	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// Membuat format response baru
	userResponse := response.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Adress: user.Adress,
		Phone: user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	// Mengembalikan hasil
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": userResponse,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	// Mengambil id user
	userId := ctx.Params("id")
	// Untuk menampung data user
	var user entity.User
	// Check available user, ditampung di err
	err := database.DB.Debug().First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	// Delete user data, tambah function error, tampung di err
	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	// Jika berhasil
	return ctx.JSON(fiber.Map{
		"message": "User was deleted",
	})
}
