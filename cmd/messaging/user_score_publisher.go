package messaging

import (
	"encoding/json"
	"kaiquecaires/real-time-leaderboard/cmd/models"
	"log"

	"github.com/IBM/sarama"
)

type UserScorePublisher interface {
	NewScore(models.CreateUserScoreParams) error
}

type KafkaUserScorePublisher struct {
	producer sarama.SyncProducer
}

func NewKafkaUserScorePublisher(producer sarama.SyncProducer) *KafkaUserScorePublisher {
	return &KafkaUserScorePublisher{producer: producer}
}

func (k *KafkaUserScorePublisher) NewScore(params models.CreateUserScoreParams) error {
	messageBytes, err := json.Marshal(params)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "user_score",
		Value: sarama.ByteEncoder(messageBytes),
	}

	partition, offset, err := k.producer.SendMessage(msg)

	if err != nil {
		return err
	}

	log.Printf("Message sent to partition %d with offset %d\n", partition, offset)

	return nil
}
