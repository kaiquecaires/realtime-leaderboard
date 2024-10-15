package messaging

import (
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
