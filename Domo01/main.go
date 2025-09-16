package main //简易的任务管理器实现增删改查

import (
	"gocode/Domo01/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
