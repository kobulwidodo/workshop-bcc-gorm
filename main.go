package main

import (
	"fmt"
	"log"
	"net/http"
	"workshop-gorm/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := initDb()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		var users []domain.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	r.POST("/cars", func(c *gin.Context) {
		var input domain.CreateCarDto
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		newCar := domain.Car{
			Name:   input.Name,
			UserId: input.UserId,
		}
		if err := db.Create(&newCar).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
	})

	r.GET("/cars/:id", func(c *gin.Context) {
		var input domain.FindCarUri
		if err := c.ShouldBindUri(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		var car domain.Car
		if err := db.Where("id = ?", input.Id).Preload("User").Find(&car).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, car)
	})

	r.GET("/user-bahasa/:id", func(c *gin.Context) {
		var input domain.FindCarUri
		if err := c.ShouldBindUri(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		var user domain.User
		if err := db.Where("id = ?", input.Id).Preload("Languages").Find(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.POST("/users/:language_id", func(c *gin.Context) {
		var input domain.AppendLanguageUri
		if err := c.ShouldBindUri(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		var user domain.User
		if err := db.Where("id = ?", 2).Find(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		var lang domain.Language
		if err := db.Where("id = ?", input.LanguageId).Find(&lang).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		if err := db.Model(&user).Association("Languages").Append(&lang); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

	})

	r.Run()
}

func initDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"",
		"localhost",
		"workshop_gorm",
	)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := DB.AutoMigrate(&domain.Car{}, &domain.User{}, &domain.Language{}); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return DB, nil
}
