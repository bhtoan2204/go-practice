package middleware

import (
	"net/http"
	"strings"

	"github.com/bhtoan2204/go-practice.git/utils"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(clientToken) != 2 {
			c.JSON(400, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}
		tokenString := clientToken[1]
		claims, err := utils.ValidateToken(tokenString)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("is_admin", claims.IsAdmin)
		c.Set("user_id", claims.ID)
		c.Set("firstname", claims.Firstname)
		c.Set("lastname", claims.Lastname)

		c.Next()
	}
}
