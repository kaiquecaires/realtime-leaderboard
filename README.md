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

## Todo
- [ ] Middleware for authentication
- [ ] Redis User Score Consumer
- [ ] Route to get leader board
- [ ] Implement Websocket
- [ ] Improve the consumer logic by adding consumer groups and multiple workers
- [ ] Handle with idempotency on user score
- [ ] Handle with retries
- [ ] Draw the System Design
- [ ] Improve folders' architecture

# References
https://medium.com/@mayilb77/design-a-real-time-leaderboard-system-for-millions-of-users-08b96b4b64ce
