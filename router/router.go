package router

import "github.com/gin-gonic/gin"

func ConstructRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/ping", handlePingRoute)
	// User routes
	r.POST("/user/register", handleUserRegisterRoute)

	return r
}
