package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func CheckRoleMiddleware(role string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("user_role")
        if !exists || userRole != role {
            c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
            c.Abort()
            return
        }
        c.Next()
    }
}