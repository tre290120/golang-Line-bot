package router

import (
	"simplebot/service"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	server := gin.Default()
	server.POST("/callback", service.DoService)
	return server
}
