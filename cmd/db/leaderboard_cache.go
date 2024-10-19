package db

import (
	"context"
	"kaiquecaires/real-time-leaderboard/cmd/models"

	"github.com/redis/go-redis/v9"
)

type LeaderboardCache interface {
	Insert(context.Context, models.Leaderboard) error
	Get(context.Context) ([]models.Leaderboard, error)
}

type RedisLeaderboardCache struct {
	client *redis.Client
}

func NewRedisLeaderboardCache(client *redis.Client) *RedisLeaderboardCache {
	return &RedisLeaderboardCache{client: client}
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
