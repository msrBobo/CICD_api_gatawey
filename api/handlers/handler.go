package handlers

import (
	"context"
	"dennic_api_gateway/api/middleware"
	"dennic_api_gateway/internal/pkg/otlp"

	"time"

	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	//"golang.org/x/net/context"
	//
	//"dennic_api_gateway/api/middleware"
	grpcClients "dennic_api_gateway/internal/infrastructure/grpc_service_client"
	"dennic_api_gateway/internal/pkg/config"
	//"dennic_api_gateway/internal/pkg/otlp"
	"dennic_api_gateway/internal/infrastructure/redis"
	//appV "dennic_api_gateway/internal/usecase/app_version"
	//"dennic_api_gateway/internal/usecase/event"
	//"dennic_api_gateway/internal/usecase/refresh_token"
)

const (
	InvestorToken = "investor"
)

type HandlerOption struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Enforcer       *casbin.CachedEnforcer
	Cache          redis.Cache
	Service        grpcClients.ServiceClient
	//	RefreshToken   refresh_token.RefreshToken
	//	AppVersion     appV.AppVersion
	//	BrokerProducer event.BrokerProducer
}

type BaseHandler struct {
	Cache  redis.Cache
	Config *config.Config
	Client grpcClients.ServiceClient
}

func (h *BaseHandler) GetAuthData(ctx context.Context) (map[string]string, bool) {
	// tracing
	ctx, span := otlp.Start(ctx, "handler", "GetAuthData")
	defer span.End()

	data, ok := ctx.Value(middleware.RequestAuthCtx).(map[string]string)
	return data, ok
}
