package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rfdez/diade/internal/platform/server/http/handlers/celebrations"
	"github.com/rfdez/diade/internal/platform/server/http/handlers/status"
	"github.com/rfdez/diade/kit/query"
	"github.com/rs/zerolog"
)

type Server struct {
	httpAddr        string
	engine          *gin.Engine
	shutdownTimeout time.Duration

	// deps
	queryBys query.Bus
}

// New creates a new HTTP server.
func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, queryBys query.Bus) (context.Context, Server) {
	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		engine:          gin.New(),
		shutdownTimeout: shutdownTimeout,

		// deps
		queryBys: queryBys,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	s.engine.Use(gin.Recovery())
	s.engine.Use(logger.SetLogger(
		logger.WithLogger(func(c *gin.Context, l zerolog.Logger) zerolog.Logger {
			output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: true}
			return l.Output(output).With().
				Timestamp().
				Logger()
		}),
	))

	s.engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	s.engine.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	s.engine.HandleMethodNotAllowed = true
	s.engine.NoMethod(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	})

	s.engine.GET("/status", status.GetHandler())
	s.engine.GET("/celebrations", celebrations.GetHandler(s.queryBys))
}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("Server running on %s\n", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(ctx, s.shutdownTimeout)
	defer cancel()

	log.Println("Server shutting down...")

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
