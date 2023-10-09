package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlePingRoute(c *gin.Context) {
	c.Status(http.StatusOK)
}
