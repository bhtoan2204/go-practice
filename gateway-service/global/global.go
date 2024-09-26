package global

import (
	"gateway-service/pkg/logger"
	"gateway-service/pkg/settings"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        settings.Config
	Logger        *logger.LoggerZap
	Redis         *redis.Client
	MDB           *gorm.DB
	KafkaProducer *kafka.Writer
	KafkaConsumer *kafka.Reader
)
