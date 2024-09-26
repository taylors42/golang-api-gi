package controllers

import (
	"api-gin/database"
	"api-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPeople(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Taylor",
	})
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
			"Err": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
