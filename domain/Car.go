package domain

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Name   string
	UserId uint
	User   User
}

type CreateCarDto struct {
	Name   string `binding:"required"`
	UserId uint   `binding:"required" json:"user_id"`
}

type FindCarUri struct {
	Id uint `uri:"id" binding:"required"`
}
