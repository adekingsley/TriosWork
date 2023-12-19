package server

import (
	"goworktoday/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidateUserInput(c *gin.Context, db *gorm.DB, newUser *models.User) bool {
	// Check if the username is unique
	var existingUser models.User
	if db.Where("username = ?", newUser.Username).First(&existingUser).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is already taken"})
		return false
	}

	// Check if the email is unique
	if db.Where("email = ?", newUser.Email).First(&existingUser).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already exist"})
		return false
	}

	// Check if username is not less than 5 characters
	if len(newUser.Username) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username must be at least 5 characters long"})
		return false
	}

	// Check if the password meets the minimum length requirement
	if len(newUser.Password) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 5 characters long"})
		return false
	}

	// Check if the password contains both alphabet characters and numbers
	if !containsAlphabetAndNumber(newUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must contain both alphabet characters and numbers"})
		return false
	}

	// All checks passed, return true to indicate successful validation
	return true
}
