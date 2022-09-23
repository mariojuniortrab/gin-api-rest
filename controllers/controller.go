package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariojuniortrab/gin-api-rest/database"
	"github.com/mariojuniortrab/gin-api-rest/models"
	"gorm.io/gorm"
)

func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func RegisterStudent(c *gin.Context) {
	var student models.Student

	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	database.DB.Create(&student)

	c.JSON(http.StatusCreated, student)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"Api says: ": "Hello " + name + ", How are you?",
		"test":       "test",
	})
}

func DetailStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	err := database.DB.First(&student, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func RemoveStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	err := database.DB.First(&student, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student was removed",
	})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	err := database.DB.First(&student, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	err = c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, student)
}

func GetByCpf(c *gin.Context) {
	var student models.Student
	cpf := c.Params.ByName("cpf")
	err := database.DB.Where(&models.Student{CPF: cpf}).First(&student).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
	}

	c.JSON(http.StatusOK, student)
}
