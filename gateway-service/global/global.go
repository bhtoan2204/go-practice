package global

import (
	"gateway-service/pkg/logger"
	"gateway-service/pkg/settings"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Redis  *redis.Client
	MDB    *gorm.DB
)
