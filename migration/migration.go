package migration

import (
	"fiber-api/database"
	"fiber-api/model/entity"
	"fmt"
	"log"
)

func RunMigration()  {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Category{}, &entity.Photo{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
