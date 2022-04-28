package domain

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Language string
}
