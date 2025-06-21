package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hiepphatle1104/matcha_latte_store/config"
	"net/http"
)

func JWTAuth(c *gin.Context) {
	tokenString, err := c.Cookie("auth_token")
	if err != nil || tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "message": "No token provided"})
		return
	}

	// Verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "message": "Invalid token"})
	}

	// Extract account ID from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "message": "Invalid token claims"})
		return
	}

	c.Set("account_id", claims["sub"])
	c.Next()
}
