package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	repo "github.com/mohammedAlbalushi582/ecom/internal/adapters/postgresql/sqlc"
	"github.com/mohammedAlbalushi582/ecom/internal/env"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "my test app",
}

var populateCmd = &cobra.Command{
	Use:   "populate",
	Short: "populate dummy data",
	Run:   populateRun,
}

func init() {
	rootCmd.AddCommand(populateCmd)
}

func populateRun(cmd *cobra.Command, args []string) {
	logger := slog.Default()
	ctx := context.Background()

	cfg := config{
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	dbConfig, err := pgxpool.ParseConfig(cfg.db.dsn)
	if err != nil {
		log.Fatalf("failed to create pgx pool: %v\n", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		log.Fatalf("failed to create pgx pool: %v\n", err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("failed to ping database: %v\n", err)
	}

	logger.Info("connected to database")

	queries := repo.New(pool)

	logger.Info("Creating dummy data for products")

}
