package controllers

import (
	"auth/models"
	"auth/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AssignRoleToUser(c *gin.Context) {
	var requestBody struct {
		UserID uint
		RoleID uint
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user := &models.User{}
	role := &models.Role{}

	// Find the user and role by their IDs
	if err := pkg.DB.First(user, requestBody.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := pkg.DB.First(role, requestBody.RoleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	// Assign the role to the user
	if err := pkg.DB.Model(user).Association("Roles").Append(role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role assigned successfully"})
}
