package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type AppendLanguageUri struct {
	LanguageId uint `uri:"language_id" binding:"required"`
}
