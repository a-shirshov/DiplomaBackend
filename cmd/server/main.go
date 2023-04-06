package main

import (
	"Diploma/internal/middleware/middleware"
	"Diploma/internal/router"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	userDelivery "Diploma/internal/microservices/user/delivery"
	userRepo "Diploma/internal/microservices/user/repository"
	userUsecase "Diploma/internal/microservices/user/usecase"

	eventV2Delivery "Diploma/internal/microservices/event_v2/delivery"
	eventV2Repo "Diploma/internal/microservices/event_v2/repository"
	eventV2Usecase "Diploma/internal/microservices/event_v2/usecase"

	authDelivery "Diploma/internal/microservices/auth/delivery"
	authRepo "Diploma/internal/microservices/auth/repository"
	authUsecase "Diploma/internal/microservices/auth/usecase"

	log "Diploma/pkg/logger"
	"Diploma/pkg/passwordHasher"
	"Diploma/pkg/tokenManager"
	"Diploma/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"golang.org/x/net/context"

	_ "Diploma/docs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title           Diploma API
// @version         1.0.0
// @description     Documentation for Diploma Api
// @description 	For Authorization:
// @description 	Put Access token in ApiKey with Bearer. Example: "Bearer access_token"

// @host      45.141.102.243:8080
// @BasePath  /api

// @accept json
// @produce json

// @schemes http

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

const logMessage = "cmd:server:main:"

func main() {
	logLevel := logrus.DebugLevel
	log.Init(logLevel)
	log.Info(logMessage + "started")

	viper.AddConfigPath("../../config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error(logMessage + "Config yaml is not found")
		os.Exit(1)
	}

	viper.SetConfigFile("../../.env")
	err = viper.MergeInConfig()
	if err != nil {
		log.Error(logMessage + "Config env is not found")
		os.Exit(1)
	}

	postgresDB, err := utils.InitPostgres()
	if err != nil {
		log.Error(logMessage + "Coundn't connect to postgres")
		os.Exit(1)
	}

	redisDB, err := utils.InitRedisDB()
	if err != nil {
		log.Error(logMessage + "Coudn't connect to redis")
		os.Exit(1)	
	}

	conn, err := grpc.Dial("recomendation:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("did not connect: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	passwordHasher := passwordHasher.NewPasswordHasher()
	tokenManager := tokenManager.NewTokenManager()

	userR := userRepo.NewUserRepository(postgresDB)
	sessionR := authRepo.NewSessionRepository(redisDB)
	eventV2R := eventV2Repo.NewEventRepositoryV2(postgresDB)
	authR := authRepo.NewAuthRepository(postgresDB)

	userU := userUsecase.NewUserUsecase(userR)
	eventV2U := eventV2Usecase.NewEventUsecaseV2(eventV2R, conn)
	authU := authUsecase.NewAuthUsecase(authR, sessionR, passwordHasher, tokenManager)

	userD := userDelivery.NewUserDelivery(userU)
	eventV2D := eventV2Delivery.NewEventDeliveryV2(eventV2U)
	authD := authDelivery.NewAuthDelivery(authU)

	

	mws := middleware.NewMiddleware(sessionR, tokenManager)

	baseRouter := gin.New()
	baseRouter.Use(gin.Logger())
	baseRouter.Use(gin.Recovery())
	baseRouter.Use(mws.CORSMiddleware())
	baseRouter.MaxMultipartMemory = 8 << 20
	baseRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routerAPI := baseRouter.Group("/api")

	authRouter := routerAPI.Group("/auth")
	router.AuthEndpoints(authRouter, mws, authD)

	userRouter := routerAPI.Group("/users")
	router.UserEndpoints(userRouter, mws, userD)

	eventV2Router := routerAPI.Group ("/events")
	router.EventV2Endpoints(eventV2Router, mws, eventV2D)

	port := viper.GetString("server.port")
	
	server := &http.Server{
		Addr: ":"+port,
		Handler: baseRouter,
		IdleTimeout: 10 * time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Error(logMessage + "Coundn't start server")
			os.Exit(1)
		}
	}()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <- sigChan
	log.Error("Graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	err = server.Shutdown(timeoutContext)
	if err != nil {
		log.Error(logMessage + "Graceful shutdown is not successful")
		os.Exit(1)
	}
}