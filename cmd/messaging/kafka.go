package messaging

import (
	"log"
	"sync"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

var once sync.Once
var kafkaProducer *KafkaProducer

func GetProducer() sarama.SyncProducer {
	once.Do(func() {
		kafkaProducer = &KafkaProducer{}
		kafkaProducer.connect()
	})
	return kafkaProducer.producer
}

func (k *KafkaProducer) connect() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all replicas to acknowledge
	config.Producer.Idempotent = true
	config.Net.MaxOpenRequests = 1

	brokers := []string{"host.docker.internal:9094"} // Your Kafka broker addresses
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	log.Println("producer connected!")
	kafkaProducer.producer = producer
}
