package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	Log "github.com/sirupsen/logrus"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// AuthMiddleware validates JWT from cookie (web) or Authorization header (mobile)
func AuthMiddleware(jwtSecretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string

		// Check bearer token first
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// Check cookie
			cookie, err := c.Cookie("session_token")
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
				c.Abort()
				return
			}
			tokenStr = cookie
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		Log.Debugf("Token %v", token)

		if err != nil || !token.Valid || claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Add user info to context
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
