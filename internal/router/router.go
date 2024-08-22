package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hazaloolu/blog-api/internal/auth"
	"github.com/hazaloolu/blog-api/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Login)

	// protected routes

	authenticated := r.Group("/")
	authenticated.Use(AuthMiddleware())
	{
		authenticated.POST("/create-post", handler.CreatePost)
		authenticated.PUT("/update-post/:id", handler.UpdatePost)
		authenticated.GET("/get-post/:id", handler.GetPost)
		authenticated.DELETE("/delete-post/:id", handler.DeletePost)
	}

	return r
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Remove 'Bearer ' prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims, err := auth.ValidateJwt(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set userID in context
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
