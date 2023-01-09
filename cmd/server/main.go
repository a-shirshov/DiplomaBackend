package main

import (
	"Diploma/internal/middleware"
	"Diploma/internal/router"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	userDelivery "Diploma/internal/microservices/user/delivery"
	userRepo "Diploma/internal/microservices/user/repository"
	userUsecase "Diploma/internal/microservices/user/usecase"

	eventDelivery "Diploma/internal/microservices/event/delivery"
	eventRepo "Diploma/internal/microservices/event/repository"
	eventUsecase "Diploma/internal/microservices/event/usecase"

	authDelivery "Diploma/internal/microservices/auth/delivery"
	authRepo "Diploma/internal/microservices/auth/repository"
	authUsecase "Diploma/internal/microservices/auth/usecase"

	placeDelivery "Diploma/internal/microservices/place/delivery"
	placeRepo "Diploma/internal/microservices/place/repository"
	placeUsecase "Diploma/internal/microservices/place/usecase"

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

	passwordHasher := passwordHasher.NewPasswordHasher()
	tokenManager := tokenManager.NewTokenManager()

	userR := userRepo.NewUserRepository(postgresDB)
	sessionR := authRepo.NewSessionRepository(redisDB)
	eventR := eventRepo.NewEventRepository(postgresDB)
	authR := authRepo.NewAuthRepository(postgresDB)
	placeR := placeRepo.NewPlaceRepository(postgresDB)

	userU := userUsecase.NewUserUsecase(userR)
	eventU := eventUsecase.NewEventUsecase(eventR)
	authU := authUsecase.NewAuthUsecase(authR, sessionR, passwordHasher, tokenManager)
	placeU := placeUsecase.NewPlaceUsecase(placeR)

	userD := userDelivery.NewUserDelivery(userU)
	eventD := eventDelivery.NewEventDelivery(eventU)
	authD := authDelivery.NewAuthDelivery(authU)
	placeD := placeDelivery.NewPlaceDelivery(placeU)

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

	eventRouter := routerAPI.Group("/events")
	router.EventEndpoints(eventRouter, mws, eventD)

	placeRouter := routerAPI.Group("/places")
	router.PlaceEndpoints(placeRouter, placeD, eventD)

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