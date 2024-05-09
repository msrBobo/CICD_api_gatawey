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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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

	api.POST("/file-upload", HandlerV1.UploadFile)

	// customer
	customer := api.Group("/customer")
	customer.POST("/register", HandlerV1.Register)
	customer.POST("/verify", HandlerV1.Verify)
	customer.POST("/forget_password", HandlerV1.ForgetPassword)
	customer.POST("/forget_password_verify", HandlerV1.ForgetPasswordVerify)
	customer.POST("/login", HandlerV1.Login)
	customer.POST("/logout", HandlerV1.LogOut)

	// user
	user := api.Group("/user")
	user.GET("/get", HandlerV1.GetUser)
	user.GET("/", HandlerV1.ListUsers)
	user.PUT("/", HandlerV1.UpdateUser)
	user.DELETE("/", HandlerV1.DeleteUser)

	// archive
	archive := api.Group("/archive")
	archive.POST("/", HandlerV1.CreateArchive)
	archive.GET("/get", HandlerV1.GetArchive)
	archive.GET("/", HandlerV1.ListArchive)
	archive.PUT("/", HandlerV1.UpdateArchive)
	archive.DELETE("/", HandlerV1.DeleteArchive)

	// doctor notes
	doctorNote := api.Group("/doctor-notes")
	doctorNote.POST("/", HandlerV1.CreateDoctorNote)
	doctorNote.GET("/get", HandlerV1.GetDoctorNote)
	doctorNote.GET("/", HandlerV1.ListDoctorNotes)
	doctorNote.PUT("/", HandlerV1.UpdateDoctorNote)
	doctorNote.DELETE("/", HandlerV1.DeleteDoctorNote)

	// doctor time
	doctorTime := api.Group("/doctor-time")
	doctorTime.POST("/", HandlerV1.CreateDoctorTimes)
	doctorTime.GET("/get", HandlerV1.GetDoctorTimes)
	doctorTime.GET("/", HandlerV1.ListDoctorTimes)
	doctorTime.PUT("/", HandlerV1.UpdateDoctorTimes)
	doctorTime.DELETE("/", HandlerV1.DeleteDoctorTimes)

	// patient
	patient := api.Group("/patient")
	patient.POST("/", HandlerV1.CreatePatient)
	patient.GET("/get", HandlerV1.GetPatient)
	patient.GET("/", HandlerV1.ListPatient)
	patient.PUT("/", HandlerV1.UpdatePatient)
	patient.DELETE("/", HandlerV1.DeletePatient)

	// appointment
	appointment := api.Group("/appointment")
	appointment.POST("/", HandlerV1.CreateBookedAppointment)
	appointment.GET("/get", HandlerV1.GetBookedAppointment)
	appointment.GET("/", HandlerV1.ListBookedAppointments)
	appointment.PUT("/", HandlerV1.UpdateBookedAppointment)
	appointment.DELETE("/", HandlerV1.DeleteBookedAppointment)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
