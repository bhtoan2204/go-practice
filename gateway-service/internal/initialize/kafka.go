package initialize

import (
	"context"
	"gateway-service/global"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

// initKafkaProducer initializes the Kafka producer
func initKafkaProducer() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP(global.Config.KafkaConfig.Broker),
		Topic:    global.Config.KafkaConfig.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	global.Logger.Info("Kafka producer initialized")
}

// ProduceMessage sends a message to Kafka
func ProduceMessage(key, message string) error {
	err := global.KafkaProducer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(message),
		},
	)
	if err != nil {
		global.Logger.Error("Failed to produce message to Kafka", zap.Error(err))
		return err
	}
	global.Logger.Info("Message sent to Kafka", zap.String("key", key), zap.String("value", message))
	return nil
}

// initKafkaConsumer initializes the Kafka consumer
func initKafkaConsumer() {
	global.KafkaConsumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{global.Config.KafkaConfig.Broker},
		GroupID: global.Config.KafkaConfig.GroupID,
		Topic:   global.Config.KafkaConfig.Topic,
	})

	global.Logger.Info("Kafka consumer initialized", zap.String("broker", global.Config.KafkaConfig.Broker), zap.String("group_id", global.Config.KafkaConfig.GroupID), zap.String("topic", global.Config.KafkaConfig.Topic))
}

// consumeMessages starts consuming messages from Kafka
func consumeMessages() {
	for {
		msg, err := global.KafkaConsumer.ReadMessage(context.Background())
		if err != nil {
			global.Logger.Error("Failed to consume message from Kafka", zap.Error(err))
			continue
		}
		global.Logger.Info("Message consumed from Kafka", zap.String("key", string(msg.Key)), zap.String("value", string(msg.Value)))
	}
}

func InitKafka() {
	initKafkaProducer()
	initKafkaConsumer()
	global.Logger.Info("Kafka configuration", zap.Any("KafkaConfig", global.Config.KafkaConfig))
	// Optionally start consuming messages in a separate goroutine
	go consumeMessages()
}
