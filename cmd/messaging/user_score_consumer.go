package messaging

import (
	"encoding/json"
	"fmt"
	"kaiquecaires/real-time-leaderboard/cmd/databases"
	"kaiquecaires/real-time-leaderboard/cmd/models"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

type UserScoreConsumer struct {
	userScoreStore databases.UserScoreStore
}

func NewUserScoreConsumer(userScoreStore databases.UserScoreStore) *UserScoreConsumer {
	return &UserScoreConsumer{userScoreStore: userScoreStore}
}

func (c *UserScoreConsumer) Consume() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	brokers := []string{"host.docker.internal:9094"}
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("error creating consumer %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("error closing consumer: %v", err)
		}
	}()

	topic := "user_score"
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("error getting partitions: %v", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	done := make(chan bool)

	for _, partition := range partitions {
		go func(partition int32) {
			partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)

			if err != nil {
				log.Fatalf("Error consuming partition %d: %v", partition, err)
			}

			defer func() {
				if err := partitionConsumer.Close(); err != nil {
					log.Fatalf("error closing partition consumer: %v", err)
				}
			}()

			log.Println("listening for new messages...")

			for {
				select {
				case msg := <-partitionConsumer.Messages():
					fmt.Printf("message received: key=%s, value=%s, partition=%d, offset=%d\n", string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)

					var createUserScore models.CreateUserScoreParams
					if err := json.Unmarshal(msg.Value, &createUserScore); err != nil {
						log.Printf("error to unmarshal message: %v\n", err)
						continue
					}

					if err := c.userScoreStore.Insert(createUserScore); err != nil {
						log.Printf("error to insert user score store: %v\n", err)
						continue
					}
				case err := <-partitionConsumer.Errors():
					log.Printf("error in partition %d: %v", partition, err)
				case <-signals:
					done <- true
					return
				}
			}
		}(partition)
	}

	<-done
	log.Println("Shutting down")
	os.Exit(0)
}
