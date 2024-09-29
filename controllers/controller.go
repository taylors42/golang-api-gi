package controllers

import (
	"api-gin/database"
	"api-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudent(c *gin.Context) {
	var student []models.Student
	database.DB.Find(&student)
	c.JSON(200, student)
}

func GetStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "404",
		})
	}
	// c.JSON(200, student)
}

func ReturnName(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"": "" + name,
	})
}

func AddStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erroooooooooor": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "404",
		})
	} else {
		database.DB.Delete(&student, id)
		c.JSON(200, gin.H{
			"student deleted": "sucess",
		})
	}
}

func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
}
