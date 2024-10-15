package messaging

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"kaiquecaires/real-time-leaderboard/cmd/databases"
	"kaiquecaires/real-time-leaderboard/cmd/models"
	"log"
)

type LeaderboardConsumer struct {
	userScoreStore databases.UserScoreStore
}

func NewLeaderboardConsumer(userScoreStore databases.UserScoreStore) *LeaderboardConsumer {
	return &LeaderboardConsumer{userScoreStore: userScoreStore}
}

func (c *LeaderboardConsumer) Consume(clientId string, groupId string) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9094",
		"client.id":         clientId, // change for more workers
		"group.id":          groupId,  // change for new group
		"auto.offset.reset": "latest",
	}

	kafkaConsumer, err := kafka.NewConsumer(configMap)

	if err != nil {
		log.Fatalf("error on consumer: %v", err)
		return
	}

	topics := []string{"leaderboard"}
	kafkaConsumer.SubscribeTopics(topics, nil)

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)

		if err != nil {
			log.Printf("error reading message: %v\n", err)
			continue
		}

		log.Println(string(msg.Value), msg.TopicPartition)

		var createUserScore models.CreateUserScoreParams

		if err := json.Unmarshal(msg.Value, &createUserScore); err != nil {
			log.Printf("error Unmarshaling value: %v\n", err)
			continue
		}

		if err := c.userScoreStore.Insert(createUserScore); err != nil {
			log.Printf("error inserting on user score: %v\n", err)
		}
	}
}
