package middleware

import (
	"gateway-service/pkg/response"

	"github.com/gin-gonic/gin"
)

func ErrHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			response.ErrorInternalServerResponse(c, response.ErrorInternalServer)
		}
	}
}
