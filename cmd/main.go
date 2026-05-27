package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/jackc/pgx/v5"
	"github.com/mohammedAlbalushi582/ecom/internal/env"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	Execute()

	// database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	logger.Info("conntected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		slog.Error("failed to execute root command", slog.String("error", err.Error()))
		os.Exit(1) // ← don't forget this!
	}
}
