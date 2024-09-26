package routes

import (
	"api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/people", controllers.GetPeople)
	r.GET("/:name", controllers.ReturnName)
	r.GET("/student", controllers.AddStudent)
	r.Run()
}
