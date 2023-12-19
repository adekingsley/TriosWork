package server

import (
	"fmt"
	"goworktoday/models"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	errMalformedToken          = "Malformed token"
	errExpiredOrNotValidYet    = "Token is either expired or not active yet"
	errAuthorizationHeader     = "authorization header is missing"
	errInvalidAuthorizationFmt = "invalid authorization header format"
)

func handleTokenParseError(err error, c *gin.Context) {
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": errMalformedToken})
			return
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": errExpiredOrNotValidYet})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidToken})
}

// Helper function to extract the token from the request
func extractTokenFromRequest(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf(errAuthorizationHeader)
	}

	// Assuming the token is in the format "Bearer <token>"
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", fmt.Errorf(errInvalidAuthorizationFmt)
	}

	return tokenParts[1], nil
}

// Helper function to add the token to the blacklist (could use a database or cache)
func addToTokenBlacklist(db *gorm.DB, token string) error {
	// Check if the token is already in the blacklist
	var existingToken models.Blacklist
	if err := db.Where("token = ?", token).First(&existingToken).Error; err == nil {
		// Token is already blacklisted, no need to add again
		return nil
	}

	// Add the token to the blacklist in the database
	newToken := models.Blacklist{Token: token}
	if err := db.Create(&newToken).Error; err != nil {
		return err
	}

	return nil
}
