package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

    Log "github.com/sirupsen/logrus"
)

func RequirePermission(enforcer *casbin.Enforcer) gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("role")
        if !exists {
            c.JSON(http.StatusForbidden, gin.H{"error": "role not found"})
            c.Abort()
            return
        }

        obj := c.FullPath() // The route path
        act := c.Request.Method // HTTP method
        Log.Debugf("path and act: %s, %s", obj, act)

        allowed, err := enforcer.Enforce(role, obj, act)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "authorization error"})
            c.Abort()
            return
        }
        if !allowed {
            c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
            c.Abort()
            return
        }

        c.Next()
    }
}

