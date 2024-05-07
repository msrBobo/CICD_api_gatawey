package api

import (
	// "github.com/casbin/casbin/v2"
	_ "dennic_api_gateway/api/docs"
	"dennic_api_gateway/api/middleware/casbin"
	"dennic_api_gateway/internal/pkg/redis"
	"time"

	v1 "dennic_api_gateway/api/handlers/v1"
	"dennic_api_gateway/api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	grpcClients "dennic_api_gateway/internal/infrastructure/grpc_service_client"
	"dennic_api_gateway/internal/pkg/config"
)

type RouteOption struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	Redis          *redis.RedisDB
	//BrokerProducer event.BrokerProducer

}

// @title API
// @version 1.7
// @host localhost:9050

// NewRoute
// @securityDefinitions.apikey ApiKeycustomer
// @in header
// @name customerorization
func NewRoute(option RouteOption) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	HandlerV1 := v1.New(&v1.HandlerV1Config{
		Config:         option.Config,
		Logger:         option.Logger,
		ContextTimeout: option.ContextTimeout,
		Service:        option.Service,
		Redis:          option.Redis,

		//BrokerProducer: option.BrokerProducer,
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	router.Use(middleware.GinTracing())

	router.Use(casbin.NewAuthorizer())

	api := router.Group("/v1")

	// customer
	customer := api.Group("/customer")
	customer.POST("/register", HandlerV1.Register)
	customer.POST("/verify", HandlerV1.Verify)
	customer.POST("/forget_password", HandlerV1.ForgetPassword)
	customer.POST("/forget_password_verify", HandlerV1.ForgetPasswordVerify)
	customer.POST("/login", HandlerV1.Login)
	customer.GET("/logout", HandlerV1.LogOut)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
