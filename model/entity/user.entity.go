package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID			uint			`json:"id" gorm:"primaryKey, unique"`
	Name		string			`json:"name"`
	Email		string			`json:"email"`
	// Login needs
	Password	string			`json:"-" gorm:"password"`
	Adress		string			`json:"address"`
	Phone		string			`json:"phone"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`json:"-" gorm: "index,column:deleted_at"`
}
