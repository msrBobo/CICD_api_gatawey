package api

import (
	_ "dennic_api_gateway/api/docs"
	v1 "dennic_api_gateway/api/handlers/v1/booking_service"
	redisrepo "dennic_api_gateway/internal/infrastructure/redis"
	"github.com/casbin/casbin/v2"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
	"time"

	"dennic_api_gateway/api/handlers"
	"dennic_api_gateway/api/middleware"
	grpcClients "dennic_api_gateway/internal/infrastructure/grpc_service_client"
	"dennic_api_gateway/internal/pkg/config"
	//redisrepo "dennic_api_gateway/internal/pkg/redis"
)

type RouteOption struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Cache          redisrepo.Cache
	Enforcer       *casbin.CachedEnforcer
	Service        grpcClients.ServiceClient
	//RefreshToken   refresh_token.RefreshToken
	//BrokerProducer event.BrokerProducer
	//AppVersion     app_version.AppVersion
}

func NewRoute(option RouteOption) http.Handler {
	handleOption := &handlers.HandlerOption{
		Config:         option.Config,
		Logger:         option.Logger,
		ContextTimeout: option.ContextTimeout,
		Cache:          option.Cache,
		Enforcer:       option.Enforcer,
		Service:        option.Service,
		//RefreshToken:   option.RefreshToken,
		//AppVersion:     option.AppVersion,
	}

	router := chi.NewRouter()
	//router.Use(chimiddleware.RealIP, chimiddleware.Logger, chimiddleware.Recoverer)
	router.Use(chimiddleware.Timeout(option.ContextTimeout))
	router.Use(middleware.Tracing)
	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Request-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Route("/v1", func(r chi.Router) {
		r.Use(middleware.AuthContext(option.Config.Token.Secret))
		r.Mount("/patients", v1.NewPatientHandler(handleOption))

	})

	// declare swagger api route
	router.Get("/swagger/*", httpSwagger.Handler())
	return router
}
