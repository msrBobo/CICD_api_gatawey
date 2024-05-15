package v1

import (
	grpc_service_clients "dennic_api_gateway/internal/infrastructure/grpc_service_client"
	"dennic_api_gateway/internal/pkg/config"
	"dennic_api_gateway/internal/pkg/redis"
	token "dennic_api_gateway/internal/pkg/tokens"
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"time"
)

const (
	SERVICE_ERROR            = "something went wrong on our side, try again later"
	INVALID_REQUET_BODY      = "invalid request body"
	NOT_REGISTERED           = "you have not registered before"
	INVALID_PHONE_NUMBER     = "invalid phone number"
	CODE_EXPIRATION_NOT_OVER = "code expiration is not over yet, please wait"
	INVALID_CODE             = "incorrect code entered"
)

type HandlerV1 struct {
	ContextTimeout time.Duration
	jwthandler     token.JWTHandler
	log            *zap.Logger
	serviceManager grpc_service_clients.ServiceClient
	cfg            *config.Config
	redis          *redis.RedisDB
	//BrokerProducer event.BrokerProducer
	//kafka          *kafka.Produce
}

// HandlerV1Config ...
type HandlerV1Config struct {
	ContextTimeout time.Duration
	Jwthandler     token.JWTHandler
	Logger         *zap.Logger
	Service        grpc_service_clients.ServiceClient
	Config         *config.Config
	Enforcer       casbin.Enforcer
	Redis          *redis.RedisDB

	//BrokerProducer event.BrokerProducer
	//Kafka          *kafka.Produce
}

// New ...
func New(c *HandlerV1Config) *HandlerV1 {
	return &HandlerV1{
		jwthandler:     c.Jwthandler,
		log:            c.Logger,
		serviceManager: c.Service,
		cfg:            c.Config,
		redis:          c.Redis,
		ContextTimeout: c.ContextTimeout,

		//BrokerProducer: c.BrokerProducer,
	}
}
