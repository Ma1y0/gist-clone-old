package router

import (
	"github.com/Ma1y0/gist-clone/router/middleware"
	"github.com/gin-gonic/gin"
)

func ConstructRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/ping", handlePingRoute)
	r.GET("/secret", middleware.AuthMiddleware, secrets)
	// User routes
	r.POST("/user/register", handleUserRegisterRoute)
	r.POST("/user/login", handleLogInUserRoute)
	r.GET("/user/:id", handleGetUsersGistsRoute)
	// GIST routes
	r.POST("/gist", middleware.AuthMiddleware, handleGistNewRoute)
	r.GET("/gist/:id", HandleGetGistByIdRoute)

	return r
}

func secrets(c *gin.Context) {
	c.JSON(200, gin.H{
		"secret": "Hi",
	})
}

// func testA(c *gin.Context) {
// 	var user model.UserModel
// 	if result := model.DB.Where("email = ?", "matyas.barr@gmail.com").Find(&user); result.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "User not found",
// 		})
// 	}
//
// 	a, _ := helpers.GenerateJWT(user)
//
// 	c.JSON(http.StatusOK, gin.H{
// 		"a": a,
// 	})
// }
//
// func testB(c *gin.Context) {
// 	token := c.Query("token")
// 	valid, err := helpers.VerifyJWT(token)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"err": err.Error(),
// 		})
// 		return
// 	}
//
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": valid,
// 	})
// }
