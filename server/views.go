package server

import (
	"fmt"
	"goworktoday/models"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	errInvalidToken            = "Invalid token"
	errFailedToInvalidateToken = "Failed to invalidate token"
)

func Endpoint1Handler(c *gin.Context) {
	c.String(http.StatusOK, "HELLO WELCOME")

}

// RegisterUser handles user registration
func RegisterUser(c *gin.Context, db *gorm.DB) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user input
	if !ValidateUserInput(c, db, &newUser) {
		return
	}

	// Hash the password before saving to the database
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = hashedPassword

	// Create the new user to the database
	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func GetUserList(c *gin.Context, db *gorm.DB) {
	var users []models.User

	// Fetch all users from the database
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user list"})
		return
	}

	// Return the list of users as JSON
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context, db *gorm.DB) {
	// Get the user ID from the URL parameters
	userID := c.Param("id")

	var user models.User

	// Fetch the user from the database based on the ID
	if err := db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Return the user as JSON
	c.JSON(http.StatusOK, user.Username)
}

// Define a struct to represent the JSON data
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context, db *gorm.DB) {
	var loginRequest LoginRequest

	// Bind the JSON data to the struct
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	username := loginRequest.Username
	password := loginRequest.Password
	// Validate user credentials against the database
	var user models.User
	log.Println("Attempting to find user with username:", username)
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username credentials"})
		return
	}

	// Compare the hashed password with the provided password using bcrypt
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password credentials"})
		return
	}

	// Generate JWT for the user
	token, err := GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}

	// Return the user ID and token in the response
	c.JSON(http.StatusOK, gin.H{"userId": user.ID, "token": token})
}

// LogoutHandler handles user logout
func LogoutHandler(c *gin.Context, db *gorm.DB) {
	// Extract the token from the request header or wherever you store it
	tokenString, err := extractTokenFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidToken})
		return
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method and get the key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil // Replace with your actual secret key
	})

	if err != nil {
		handleTokenParseError(err, c)
		return
	}

	if !token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidToken})
		return
	}

	// Add the token to the blacklist (or some persistent storage)
	err = addToTokenBlacklist(db, tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFailedToInvalidateToken})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
