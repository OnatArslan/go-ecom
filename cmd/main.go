package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env variables can not load")
	}

	cfg := &config{
		addr: ":8080",
		db: dbConfig{
			dsn: os.Getenv("DB_CON"),
		},
	}
	// Structured logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, cfg.db.dsn)
	if err != nil {
		log.Fatal("database connection failed")
	}

	defer conn.Close()
	logger.Info("connected to database", "dsn", cfg.db.dsn)

	api := &application{
		config: *cfg,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
