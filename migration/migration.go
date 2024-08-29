package migration

import (
	"fiber-api/database"
	"fiber-api/model/entity"
	"fmt"
	"log"
)

func RunMigration()  {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
