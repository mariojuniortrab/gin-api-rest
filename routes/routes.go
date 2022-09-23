package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mariojuniortrab/gin-api-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.DetailStudent)
	r.GET("/students/cpf/:cpf", controllers.GetByCpf)
	r.POST("/students", controllers.RegisterStudent)
	r.GET("/:name", controllers.Greeting)
	r.DELETE("/students/:id", controllers.RemoveStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.Run()
}
