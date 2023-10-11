package router

import (
	"log"
	"net/http"

	"github.com/Ma1y0/gist-clone/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type registerInput struct {
	Username string
	Password string
	Email    string
}

// Register a new user
func handleUserRegisterRoute(c *gin.Context) {
	var body registerInput
	// Validate request
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hash, err := hashPassword([]byte(body.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser := model.UserModel{
		ID:       uuid.NewString(),
		Username: body.Username,
		Email:    body.Email,
		Password: hash,
	}

	if result := model.DB.Create(&newUser); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to regiter user",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User was successfully registered",
	})
}

func hashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	return string(hash), err
}

type logInInput struct {
	Email    string
	Password string
}

// Log In
func handleLogInUserRoute(c *gin.Context) {
	var body logInInput

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user model.UserModel
	if result := model.DB.Where("email = ?", body.Email).Find(&user); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"match": comapreHashAndPassword([]byte(body.Password), []byte(user.Password)),
	})
}

func comapreHashAndPassword(password []byte, hash []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hash, password); err != nil {
		return false
	}

	return true
}
