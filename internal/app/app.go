package app

import (
	"CICD_api_gatawey/api"
	grpcService "CICD_api_gatawey/internal/infrastructure/grpc_service_client"
	"CICD_api_gatawey/internal/pkg/config"
	"CICD_api_gatawey/internal/pkg/logger"
	"CICD_api_gatawey/internal/pkg/otlp"
	"CICD_api_gatawey/internal/pkg/postgres"
	"CICD_api_gatawey/internal/pkg/redis"
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type App struct {
	Config       *config.Config
	Logger       *zap.Logger
	DB           *postgres.PostgresDB
	RedisDB      *redis.RedisDB
	server       *http.Server
	Clients      grpcService.ServiceClient
	ShutdownOTLP func() error
	//BrokerProducer event.BrokerProducer
}

func NewApp(cfg *config.Config) (*App, error) {
	// l init
	l, err := logger.New(cfg.LogLevel, cfg.Environment, cfg.APP+".log")
	if err != nil {
		return nil, err
	}

	// kafka producer init
	//kafkaProducer := kafka.NewProducer(&cfg, logger)

	// postgres init
	db, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}

	// redis init
	redisdb, err := redis.New(cfg)
	if err != nil {
		return nil, err
	}

	// otlp collector init
	shutdownOTLP, err := otlp.InitOTLPProvider(cfg)
	if err != nil {
		return nil, err
	}

	// initialization enforcer

	return &App{
		Config:  cfg,
		Logger:  l,
		DB:      db,
		RedisDB: redisdb,
		//BrokerProducer: kafkaProducer,
		ShutdownOTLP: shutdownOTLP,
		//appVersion:     appVersionUseCase,
	}, nil
}

func (a *App) Run() error {
	contextTimeout, err := time.ParseDuration("2s")
	if err != nil {
		return fmt.Errorf("error while parsing context timeout: %v", err)
	}

	// initialize cache

	clients, err := grpcService.New(a.Config)
	if err != nil {
		return err
	}
	a.Clients = clients

	// api init
	handler := api.NewRoute(api.RouteOption{
		Config:         a.Config,
		Logger:         a.Logger,
		ContextTimeout: contextTimeout,
		Service:        clients,
		Redis:          a.RedisDB,
		//BrokerProducer: a.BrokerProducer,
	})

	// server init
	a.server, err = api.NewServer(a.Config, handler)
	if err != nil {
		return fmt.Errorf("error while initializing server: %v", err)
	}

	return a.server.ListenAndServe()
}

func (a *App) Stop() {

	// close database
	a.DB.Close()

	// close grpc connections
	a.Clients.Close()

	// shutdown server http
	if err := a.server.Shutdown(context.Background()); err != nil {
		a.Logger.Error("shutdown server http ", zap.Error(err))
	}

	// shutdown otlp collector

	// zap logger sync
	a.Logger.Sync()
}
