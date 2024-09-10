package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "No authorization header"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, gin.Error{
                    Err:  gin.Error{},
                    Type: gin.ErrorTypePrivate,
                }
            }
            return []byte("secret"), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
            c.Abort()
            return
        }

        claims, _ := token.Claims.(jwt.MapClaims)
        userID := claims["user_id"].(float64)
        role := claims["role"].(string)

        c.Set("user_id", userID)
        c.Set("user_role", role)

        c.Next()
    }
}
