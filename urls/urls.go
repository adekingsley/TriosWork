package urls

import (
	"goworktoday/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all routes for the application
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	r.GET("/home", server.Endpoint1Handler)

	// Pass the *gorm.DB instance to the RegisterUser handler
	r.POST("/register", func(c *gin.Context) {
		server.RegisterUser(c, db)
	})

	r.GET("/users", func(c *gin.Context) {
		server.GetUserList(c, db)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		server.GetUserById(c, db)
	})

	r.POST("/login", func(c *gin.Context) {
		server.LoginHandler(c, db)
	})

	r.GET("/logout", func(c *gin.Context) {
		server.LogoutHandler(c, db)
	})

	// Route protected by JWT authentication
	r.GET("/protected", server.AuthMiddleware(), server.SomeProtectedRoute)
}
