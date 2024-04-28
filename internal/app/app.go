package app

import (
	"context"
	"dennic_api_gateway/api"
	grpcService "dennic_api_gateway/internal/infrastructure/grpc_service_client"
	redisrepo "dennic_api_gateway/internal/infrastructure/redis"
	"dennic_api_gateway/internal/pkg/config"
	"dennic_api_gateway/internal/pkg/logger"
	"dennic_api_gateway/internal/pkg/otlp"
	"dennic_api_gateway/internal/pkg/policy"
	"dennic_api_gateway/internal/pkg/postgres"
	"dennic_api_gateway/internal/pkg/redis"
	"fmt"
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type App struct {
	Config       *config.Config
	Logger       *zap.Logger
	DB           *postgres.PostgresDB
	RedisDB      *redis.RedisDB
	server       *http.Server
	Enforcer     *casbin.CachedEnforcer
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
	enforcer, err := policy.NewCachedEnforcer(cfg, l)
	if err != nil {
		return nil, err
	}

	enforcer.SetCache(policy.NewCache(&redisdb.Client))

	return &App{
		Config:   cfg,
		Logger:   l,
		DB:       db,
		RedisDB:  redisdb,
		Enforcer: enforcer,
		//BrokerProducer: kafkaProducer,
		ShutdownOTLP: shutdownOTLP,
		//appVersion:     appVersionUseCase,
	}, nil
}

func (a *App) Run() error {
	contextTimeout, err := time.ParseDuration(a.Config.Context.Timeout)
	if err != nil {
		return fmt.Errorf("error while parsing context timeout: %v", err)
	}

	// initialize cache
	cache := redisrepo.NewCache(a.RedisDB)

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
		Cache:          cache,
		Enforcer:       a.Enforcer,
		Service:        clients,
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
	if err := a.ShutdownOTLP(); err != nil {
		a.Logger.Error("shutdown otlp collector", zap.Error(err))
	}

	// zap logger sync
	a.Logger.Sync()
}
