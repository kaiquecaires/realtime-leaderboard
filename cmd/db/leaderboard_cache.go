package db

import (
	"context"
	"kaiquecaires/real-time-leaderboard/cmd/models"
	"log"

	"github.com/redis/go-redis/v9"
)

type LeaderboardCache interface {
	Insert(context.Context, models.Leaderboard) error
	Get(context.Context) ([]models.Leaderboard, error)
}

type RedisLeaderboardCache struct {
	client         *redis.Client
	userScoreStore UserScoreStore
}

func NewRedisLeaderboardCache(client *redis.Client, userScoreStore UserScoreStore) *RedisLeaderboardCache {
	return &RedisLeaderboardCache{client: client, userScoreStore: userScoreStore}
}

func (c *RedisLeaderboardCache) Insert(ctx context.Context, params models.Leaderboard) error {
	err := c.client.ZAdd(ctx, "leaderboard", redis.Z{
		Score:  float64(params.Score),
		Member: params.Username,
	}).Err()

	return err
}

func (c *RedisLeaderboardCache) Get(ctx context.Context) ([]models.Leaderboard, error) {
	users, err := c.client.ZRangeWithScores(ctx, "leaderboard", 0, -1).Result()

	if err != nil {
		return nil, err
	}

	var leaderboard []models.Leaderboard

	for _, user := range users {
		username := user.Member.(string)
		leaderboard = append(leaderboard, models.Leaderboard{
			Username: username,
			Score:    int(user.Score),
		})
	}

	return leaderboard, nil
}

func (c *RedisLeaderboardCache) Populate() {
	leaderboard, err := c.userScoreStore.GetLeaderboard(models.GetLeaderboardParams{})

	if err != nil {
		log.Fatalf("Failed to get leaderboard on populate redis cache: %s", err)
		return
	}

	for _, l := range leaderboard {
		err := c.Insert(context.Background(), l)

		if err != nil {
			log.Fatalf("Failed to pre populate redis cache: %s", err)
		}
	}
}
