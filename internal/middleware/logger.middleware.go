package middleware

import (
	"time"

	"simple_bank/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		statusCode := c.Writer.Status()

		var logLevel zapcore.Level
		if statusCode >= 500 {
			logLevel = zap.ErrorLevel
		} else if statusCode >= 400 {
			logLevel = zap.WarnLevel
		} else {
			logLevel = zap.InfoLevel
		}

		fields := []zap.Field{
			zap.Int("status", statusCode),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		}

		if len(c.Errors) > 0 {
			fields = append(fields, zap.String("errors", c.Errors.String()))
		}

		switch logLevel {
		case zap.ErrorLevel:
			global.Logger.Error("Server error", fields...)
		case zap.WarnLevel:
			global.Logger.Warn("Client error", fields...)
		default:
			global.Logger.Info("Request succeeded", fields...)
		}
	}
}
