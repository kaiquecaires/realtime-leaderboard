FROM golang:1.23.0-alpine3.20 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o real_time_leaderboard cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/real_time_leaderboard .
EXPOSE 8080
CMD ["./real_time_leaderboard"]
