package routes

import (
	"gocode/Domo01/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.AddTask)
	r.PUT("/tasks/:id/complete", controllers.CompleteTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
}
