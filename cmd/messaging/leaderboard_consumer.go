package messaging

import (
	"context"
	"encoding/json"
	"kaiquecaires/real-time-leaderboard/cmd/db"
	"kaiquecaires/real-time-leaderboard/cmd/models"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type LeaderboardConsumer struct {
	userScoreStore   db.UserScoreStore
	leaderboardCache db.LeaderboardCache
	userStore        db.UserStore
}

func NewLeaderboardConsumer(userScoreStore db.UserScoreStore, leaderboardCache db.LeaderboardCache, userStore db.UserStore) *LeaderboardConsumer {
	return &LeaderboardConsumer{userScoreStore: userScoreStore, leaderboardCache: leaderboardCache, userStore: userStore}
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
			continue
		}

		user, err := c.userStore.GetById(createUserScore.UserId)

		if err != nil {
			log.Printf("error to get user: %v\n", err)
			continue
		}

		if err := c.leaderboardCache.Insert(context.Background(), models.Leaderboard{
			Username: user.Username,
			Score:    createUserScore.Score,
		}); err != nil {
			log.Printf("error inserting on leaderboard cache: %v\n", err)
		}
	}
}
