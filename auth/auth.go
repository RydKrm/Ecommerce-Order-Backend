package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization");
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Authorization header required"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ");
		secret := os.Getenv("TOKEN_SECRET")
		// parse and validate jwt token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]);
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid token"});
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			role := claims["role"].(string)
			id := claims["id"].(float64);

			// check if the role is  user or not
			for ro := range roles {
				if role == roles[ro] {
					c.Set("user_id", id)
					c.Set("role", role)
					c.Next()
					return;
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid role"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid token"})
			c.Abort()
			return
		}

	}
}