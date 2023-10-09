package router

import "github.com/gin-gonic/gin"

func ConstructRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/ping", handlePingRoute)

	return r
}
