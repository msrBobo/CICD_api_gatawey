package config

import (
	"github.com/spf13/cast"
	"os"
	"strings"
	"time"
)

const (
	OtpSecret = "some_secret"
	SignKey   = "dennic_key"
)

type webAddress struct {
	Host string
	Port string
}

type minio struct {
	Endpoint  string
	AccessKey string
	SecretKey string
}

type Config struct {
	APP         string
	Environment string
	LogLevel    string
	Server      struct {
		Host         string
		Port         string
		ReadTimeout  string
		WriteTimeout string
		IdleTimeout  string
	}
	DB struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		SSLMode  string
	}
	Context struct {
		Timeout int
	}
	Redis struct {
		Host     string
		Port     string
		Password string
		Name     string
		Time     time.Duration
	}
	Token struct {
		Secret     string
		AccessTTL  time.Duration
		RefreshTTL time.Duration
	}
	Minio struct {
		Endpoint              string
		AccessKey             string
		SecretKey             string
		Location              string
		MovieUploadBucketName string
	}
	Kafka struct {
		Address []string
		Topic   struct {
			InvestmentPaymentTransaction string
		}
	}
	BookingService    webAddress
	HealthcareService webAddress
	UserService       webAddress
	SessionService    webAddress
	OTLPCollector     webAddress

	MinioService minio
}

func NewConfig() (*Config, error) {
	var config Config

	// general configuration
	config.APP = getEnv("APP", "dennic_api_gateway")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "debug")
	config.Context.Timeout = cast.ToInt(getEnv("CONTEXT_TIMEOUT", "2"))

	// server configuration
	config.Server.Host = getEnv("SERVER_HOST", "localhost")
	config.Server.Port = getEnv("SERVER_PORT", ":9050")
	config.Server.ReadTimeout = getEnv("SERVER_READ_TIMEOUT", "10s")
	config.Server.WriteTimeout = getEnv("SERVER_WRITE_TIMEOUT", "10s")
	config.Server.IdleTimeout = getEnv("SERVER_IDLE_TIMEOUT", "120s")

	// db configuration
	config.DB.Host = getEnv("POSTGRES_HOST", "localhost")
	config.DB.Port = getEnv("POSTGRES_PORT", "5432")
	config.DB.Name = getEnv("POSTGRES_DATABASE", "dennic")
	config.DB.User = getEnv("POSTGRES_USER", "postgres")
	config.DB.Password = getEnv("POSTGRES_PASSWORD", "2003")
	config.DB.SSLMode = getEnv("POSTGRES_SSLMODE", "disable")

	// redis configuration
	config.Redis.Host = getEnv("REDIS_HOST", "redisdb")
	config.Redis.Port = getEnv("REDIS_PORT", "6379")
	config.Redis.Password = getEnv("REDIS_PASSWORD", "")
	config.Redis.Name = getEnv("REDIS_DATABASE", "0")
	config.Redis.Time = time.Minute

	config.BookingService.Host = getEnv("BOOKING_SERVICE_GRPC_HOST", "localhost")
	config.BookingService.Port = getEnv("BOOKING_SERVICE_GRPC_PORT", ":9090")

	config.HealthcareService.Host = getEnv("HEALTHCARE_SERVICE_GRPC_HOST", "localhost")
	config.HealthcareService.Port = getEnv("HEALTHCARE_SERVICE_GRPC_PORT", ":9080")

	config.SessionService.Host = getEnv("SESSION_SERVICE_GRPC_HOST", "localhost")
	config.SessionService.Port = getEnv("SESSION_SERVICE_GRPC_PORT", ":9060")

	config.UserService.Host = getEnv("USER_SERVICE_GRPC_HOST", "localhost")
	config.UserService.Port = getEnv("USER_SERVICE_GRPC_PORT", ":9070")

	// token configuration
	config.Token.Secret = getEnv("TOKEN_SECRET", "token_secret")

	// access ttl parse
	accessTTl, err := time.ParseDuration(getEnv("TOKEN_ACCESS_TTL", "1h"))
	if err != nil {
		return nil, err
	}
	// refresh ttl parse
	refreshTTL, err := time.ParseDuration(getEnv("TOKEN_REFRESH_TTL", "24h"))
	if err != nil {
		return nil, err
	}
	config.Token.AccessTTL = accessTTl
	config.Token.RefreshTTL = refreshTTL

	// otlp collector configuration
	config.OTLPCollector.Host = getEnv("OTLP_COLLECTOR_HOST", "otel-collector")
	config.OTLPCollector.Port = getEnv("OTLP_COLLECTOR_PORT", ":4317")

	// kafka configuration
	config.Kafka.Address = strings.Split(getEnv("KAFKA_ADDRESS", "localhost:29092"), ",")
	config.Kafka.Topic.InvestmentPaymentTransaction = getEnv("KAFKA_TOPIC_INVESTMENT_PAYMENT_TRANSACTION", "investment.payment.transaction")

	// model_minio configuration
	config.MinioService.Endpoint = getEnv("MINIO_SERVICE_ENDPOINT", "dennic.uz:9000")
	config.MinioService.AccessKey = getEnv("MINIO_SERVICE_ACCESS_KEY", "dennic")
	config.MinioService.SecretKey = getEnv("MINIO_SERVICE_SECRET_KEY", "dennic_service")

	return &config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
