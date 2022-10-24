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

	psqlURI := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbParams)
	db, err := sql.Open("postgres", psqlURI)
	if err != nil {
		log.Fatal(err)
	}

	// Bus
	var (
		queryBus = inmemory.NewQueryBus()
	)

	// Repositories
	var (
		celebrationRepository = postgresql.NewCelebrationRepository(db, cfg.DbTimeout)
	)

	// Services
	var (
		fetchingService = fetching.NewService(celebrationRepository)
	)

	// Query Handlers
	var (
		fetchingCelebrationByDateHandler = fetching.NewCelebrationByDateQueryHandler(fetchingService)
	)

	// Register query handlers
	queryBus.Register(fetching.CelebrationByDateQueryType, fetchingCelebrationByDateHandler)

	ctx, srv := http.New(context.Background(), cfg.HttpHost, cfg.HttpPort, cfg.ShutdownTimeout, queryBus)
	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

type config struct {
	// Http Server configuration
	HttpHost        string        `default:""`
	HttpPort        uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"diade"`
	DbPass    string        `default:"diade"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"5432"`
	DbName    string        `default:"diade"`
	DbParams  string        `default:""`
	DbTimeout time.Duration `default:"5s"`
}
