package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/rfdez/diade/internal/fetching"
	"github.com/rfdez/diade/internal/platform/bus/inmemory"
	"github.com/rfdez/diade/internal/platform/server/http"
	"github.com/rfdez/diade/internal/platform/storage/postgresql"
)

func main() {
	var cfg config
	err := envconfig.Process("diade", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	psqlURI := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBParams)
	db, err := sql.Open("postgres", psqlURI)
	if err != nil {
		log.Fatal(err)
	}

	// Bus

	queryBus := inmemory.NewQueryBus()

	// Repositories

	celebrationRepository := postgresql.NewCelebrationRepository(db, cfg.DBTimeout)

	// Services

	fetchingService := fetching.NewService(celebrationRepository)

	// Query Handlers

	fetchingCelebrationByDateHandler := fetching.NewCelebrationByDateQueryHandler(fetchingService)

	// Register query handlers
	queryBus.Register(fetching.CelebrationByDateQueryType, fetchingCelebrationByDateHandler)

	ctx, srv := http.New(context.Background(), cfg.HTTPHost, cfg.HTTPPort, cfg.ShutdownTimeout, queryBus)
	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

type config struct {
	// Http Server configuration
	HTTPHost        string        `default:""`
	HTTPPort        uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DBUser    string        `default:"diade"`
	DBPass    string        `default:"diade"`
	DBHost    string        `default:"localhost"`
	DBPort    uint          `default:"5432"`
	DBName    string        `default:"diade"`
	DBParams  string        `default:""`
	DBTimeout time.Duration `default:"5s"`
}
