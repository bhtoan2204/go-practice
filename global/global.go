package global

import (
	"github.com/bhtoan2204/go-practice.git/pkg/logger"
	"github.com/bhtoan2204/go-practice.git/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Redis  *redis.Client
	MDB    *gorm.DB
)
