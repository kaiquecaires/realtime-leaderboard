package messaging

import (
	"encoding/json"
	"kaiquecaires/real-time-leaderboard/cmd/models"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type UserScorePublisher interface {
	NewScore(models.CreateUserScoreParams) error
}

type KafkaUserScorePublisher struct {
	producer *kafka.Producer
}

func NewKafkaUserScorePublisher(producer *kafka.Producer) *KafkaUserScorePublisher {
	return &KafkaUserScorePublisher{producer: producer}
}

func (k *KafkaUserScorePublisher) NewScore(params models.CreateUserScoreParams) error {
	messageBytes, err := json.Marshal(params)

	if err != nil {
		return err
	}

	topic := "leaderboard"

	m := &kafka.Message{
		Value:          messageBytes,
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	}

	err = k.producer.Produce(m, nil)
	return err
}
