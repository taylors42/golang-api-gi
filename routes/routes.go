package routes

import (
	"api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/student", controllers.GetStudent)
	r.GET("/student/:id", controllers.GetStudentById)
	r.GET("/:name", controllers.ReturnName)
	r.POST("/student", controllers.AddStudent)
	r.DELETE("/student/:id", controllers.DeleteStudent)
	r.PATCH("/student/:id", controllers.EditStudent)
	r.Run()
}
