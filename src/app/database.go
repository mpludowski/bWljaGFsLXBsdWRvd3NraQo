package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func ConnectDb() *pgxpool.Pool {
	dbURL := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	log.Print("Database connected.")

	conn, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	initDb(conn)

	return conn
}

func CloseDb(conn *pgxpool.Pool) {
	conn.Close()
}

func initDb(conn *pgxpool.Pool) {
	_, err := conn.Query(context.Background(),"SELECT * FROM urls")

	if err != nil {
		conn.Exec(
			context.Background(),
			"CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"",
		)
		conn.Exec(
			context.Background(),
			"CREATE TABLE urls (id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), url TEXT, interval INTEGER)",
		)
		conn.Exec(
			context.Background(),
			"CREATE TABLE history (id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), url_id UUID, response TEXT, duration NUMERIC, created_at timestamp DEFAULT NOW())",
		)

		log.Print("Created database schema.")
	}
}