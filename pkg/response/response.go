package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorInternalServerResponse(c *gin.Context, code int) {
	c.JSON(http.StatusInternalServerError, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

func ErrorBadRequestResponse(c *gin.Context, code int) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

func ErrorUnauthorizedResponse(c *gin.Context, code int) {
	c.JSON(http.StatusUnauthorized, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
