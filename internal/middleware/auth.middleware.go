package middleware

import (
	"simple_bank/package/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if (len(accessToken) < 2) || (accessToken[1] == "") {
			response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
