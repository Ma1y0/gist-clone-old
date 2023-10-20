package router

import (
	"net/http"

	"github.com/Ma1y0/gist-clone/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type newGistInput struct {
	Title       string `binding:"required"`
	Description string
	Code        string `binding:"required"`
}

func handleGistNewRoute(c *gin.Context) {
	user_id_any, exists := c.Get("ID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User is not found",
		})
		return
	}

	user_id, ok := user_id_any.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User ID is not a string",
		})
		return
	}

	// Validates input
	var body newGistInput
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong request",
			"error":   err.Error(),
		})
		return
	}

	// Insert new gist into the dadabase
	newGist := model.GistModel{
		ID:          uuid.NewString(),
		Title:       body.Title,
		Description: body.Description,
		Code:        body.Code,
		OwnerID:     user_id,
	}

	if result := model.DB.Create(&newGist); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create a new gist",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created a new gist",
		"gist":    newGist,
	})
}

// GEt gist by ID
func HandleGetGistByIdRoute(c *gin.Context) {
	id := c.Param("id")

	var gist model.GistModel
	if result := model.DB.First(&gist, "id = ?", id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Gist not found",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"message": "Gist successfully found",
		"gist":    gist,
	})
}
