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

	"Diploma/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

func main() {

	l := log.New(os.Stdout, "Diploma-API", log.LstdFlags)

	viper.AddConfigPath("../../config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("Config isn't found 1")
		os.Exit(1)
	}
	viper.SetConfigFile("../../.env")
	err = viper.MergeInConfig()
	if err != nil {
		log.Print("Config isn't found 2")
		os.Exit(1)
	}

	postgresDB, err := utils.InitPostgresDB()
	if err != nil {
		log.Print("InitPG")
		log.Println(err)
		os.Exit(1)
	}

	err = utils.Prepare(postgresDB)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	redisDB, err := utils.InitRedisDB()
	if err != nil {
		log.Println(err)
		os.Exit(1)	
	}

	userR := userRepo.NewUserRepository(postgresDB)
	sessionR := authRepo.NewSessionRepository(redisDB)
	eventR := eventRepo.NewEventRepository(postgresDB)
	authR := authRepo.NewAuthRepository(postgresDB)
	placeR := placeRepo.NewPlaceRepository(postgresDB)

	userU := userUsecase.NewUserUsecase(userR)
	eventU := eventUsecase.NewEventUsecase(eventR)
	authU := authUsecase.NewAuthUsecase(authR, sessionR)
	placeU := placeUsecase.NewPlaceUsecase(placeR)

	userD := userDelivery.NewUserDelivery(userU)
	eventD := eventDelivery.NewEventDelivery(eventU)
	authD := authDelivery.NewAuthDelivery(authU)
	placeD := placeDelivery.NewPlaceDelivery(placeU)

	mws := middleware.NewMiddleware(sessionR)

	baseRouter := gin.New()
	baseRouter.Use(gin.Logger())
	baseRouter.Use(gin.Recovery())
	baseRouter.Use(mws.CORSMiddleware())
	baseRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routerAPI := baseRouter.Group("/api")

	authRouter := routerAPI.Group("/auth")
	router.AuthEndpoints(authRouter, mws, authD)

	userRouter := routerAPI.Group("/users")
	router.UserEndpoints(userRouter, mws, userD)

	eventRouter := routerAPI.Group("/events")
	router.EventEndpoints(eventRouter, eventD)

	placeRouter := routerAPI.Group("/places")
	router.PlaceEndpoints(placeRouter, placeD)

	port := viper.GetString("server.port")
	
	server := &http.Server{
		Addr: ":"+port,
		ErrorLog: l,
		Handler: baseRouter,
		IdleTimeout: 10 * time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <- sigChan
	log.Println("Graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(timeoutContext)
}