package main

import (
	// "market_go_apigateway/api"
	// "market_go_apigateway/config"
	// "market_go_apigateway/grpc/client"
	// "market_go_apigateway/pkg/logger"
	// "market_go_apigateway/api/handlers"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-psql/api"
	handlers "github.com/go-psql/api/handlers"
	"github.com/go-psql/config"
	"github.com/go-psql/pkg/logger"
	"github.com/go-psql/storage/postgresql"
)

func main() {
	// get data from config
	cfg := config.Load()

	// set logger level and gin mode
	var loggerLevel string

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	// create logger
	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			log.Error("!!!Cleanup->Error-->>>", logger.Error(err))
		}
	}()

	pgStore, err := postgresql.NewPostrgesqlConnection(context.Background(), cfg)
	if err != nil {
		log.Error("!!!NewPostrgesqlConnection->Error-->>>", logger.Error(err))
	}

	defer pgStore.CloseDB()

	h := handlers.NewHandler(cfg, log, pgStore)

	r := api.SetUpRouter(h, cfg)

	log.Info("HTTP: Server being started...", logger.String("port", cfg.HTTPPort))

	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
