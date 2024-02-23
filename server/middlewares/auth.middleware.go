package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"serveur/server/services/jwt"
)

func AuthMiddleware(c *gin.Context) {
	token := extractToken(c)

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "No token provided"})
		c.Abort()
		return
	}

	claims := services.ParseAccessToken(token)

	if claims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}

/** Extract token from the request header
 */
func extractToken(c *gin.Context) string {
	if c.Request.Header.Get("Authorization") == "" {
		return ""
	}
	authorization := c.Request.Header.Get("Authorization")
	if len(authorization) < 7 {
		return ""
	}
	return authorization[7:]
}
