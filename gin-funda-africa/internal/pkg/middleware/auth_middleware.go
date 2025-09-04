package middleware

import (
	"errors"
	"gin-funda-africa/internal/pkg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

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

func CheckOwnership (param uuid.UUID, ctx *gin.Context) (error) {
	requestUser, exists := ctx.Get("userID")
	if !exists {
		return errors.New("Requesting user not defined")
	}

	userRole, exists := ctx.Get("role")

	strValue, ok := requestUser.(string)
	if !ok {
		return errors.New("Invalid user id")
	}

	requestUserID, err := uuid.Parse(strValue)
	if err != nil {
		return errors.New(utils.InvalidUUID)
	}

	if exists && (userRole == "super_admin" || userRole == "admin"){
		return nil
	}

	if requestUserID != param { 
		return errors.New("Forbidden")
	}

	return nil
}

func GetCurrentUserID(ctx *gin.Context) (*uuid.UUID, error) {
	requestUser, exists := ctx.Get("userID")
	if !exists {
		return nil, errors.New("Requesting user not defined")
	}

	strValue, ok := requestUser.(string)
	if !ok {
		return nil, errors.New("Invalid user id")
	}

	requestUserID, err := uuid.Parse(strValue)
	if err != nil {
		return nil, errors.New(utils.InvalidUUID)
	}

	return &requestUserID, nil
}

func GetCurrentUserRole(ctx *gin.Context) (*string, error) {
	userRole, exists := ctx.Get("role")
	if !exists {
		return nil, errors.New("Requesting user role not defined")
	}

	strValue, ok := userRole.(string)
	if !ok {
		return nil, errors.New("Invalid user role")
	}
	return &strValue, nil
}
