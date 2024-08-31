package databases

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

type Database struct {
	connection *sql.DB
}

var instance *Database
var once sync.Once

func GetDBInstance() *sql.DB {
	once.Do(func() {
		instance = &Database{}
		instance.connect()
	})
	return instance.connection
}

func (db *Database) connect() {
	dsn := "user=admin password=admin dbname=realtime_leaderboard host=realtime_leaderboard_postgres port=5432 sslmode=disable"
	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	conn.SetMaxIdleConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	err = conn.Ping()

	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL with connection pool!")
	db.connection = conn
}
