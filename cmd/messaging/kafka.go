package messaging

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

var once sync.Once
var kafkaProducer *KafkaProducer

func GetProducer() *kafka.Producer {
	once.Do(func() {
		kafkaProducer = &KafkaProducer{}
		kafkaProducer.connect()
	})
	return kafkaProducer.producer
}

func (k *KafkaProducer) connect() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "localhost:9094",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}

	p, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Fatalf("error on new producer: %v", err)
	}

	kafkaProducer.producer = p
}

func CreateTopic() {
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "localhost:9094"})

	if err != nil {
		log.Fatalf("Failed to create Admin client: %s\n", err)
		return
	}

	defer adminClient.Close()

	topics, err := adminClient.GetMetadata(nil, true, 5000)
	if err != nil {
		fmt.Printf("Failed to get metadata: %s\n", err)
		return
	}

	topicName := "leaderboard"
	numPartitions := 1
	replicationFactor := 1

	// Check if topic already exists
	if _, exists := topics.Topics[topicName]; exists {
		fmt.Printf("Topic %s already exists, skipping creation.\n", topicName)
		return
	}

	results, err := adminClient.CreateTopics(
		context.Background(),
		[]kafka.TopicSpecification{
			{
				Topic:             topicName,
				NumPartitions:     numPartitions,
				ReplicationFactor: replicationFactor,
			},
		},
	)

	if err != nil {
		log.Fatalf("Failed to create topic: %s\n", err)
		return
	}

	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			fmt.Printf("Failed to create topic %s: %v\n", result.Topic, result.Error)
		} else {
			fmt.Printf("Topic %s created successfully!\n", result.Topic)
		}
	}
}
