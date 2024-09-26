package kafka

import (
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaConsumer *kafka.Writer
)

const (
	kafkaURL   = "localhost:9092"
	kafkaTopic = "user_topic_001"
)

// for producer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, // string[]{"",""}
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

// for consumer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
