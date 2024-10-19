# Real-time leaderboard
Inspired by https://roadmap.sh/projects/realtime-leaderboard-system.
A real-time leaderboard application written in go.

# How to run?
First, run the dependencies through docker compose using the command:
```bash
docker-compose up -d
```

Run the go project:
```bash
go run cmd/main.go
```

## Done
- [x] User sign up
- [x] Send User Score Data
- [x] Postgresql User Score Consumer
- [x] Create a game
- [x] Replace Sarama for Confluent Kafka library
- [x] Login route
- [x] Middleware for authentication
- [x] Improve the consumer logic by adding consumer groups and multiple workers
- [x] Route to get leaderboard

## Todo
- [ ] Handle with idempotency on user score
- [ ] Draw the System Design
- [ ] Implement redis as cache
- [ ] Implement Lazy loading on redis
- [ ] Handle with retries on save on redis

## Things to study
- [ ] What is the best practices when saving leaderboard on redis?
- [ ] What is the best practices when working with kafka?

# References
https://medium.com/@mayilb77/design-a-real-time-leaderboard-system-for-millions-of-users-08b96b4b64ce
