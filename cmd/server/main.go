// Package classification Diploma API
//
// Documentation for Diploma API
//
//	Schemes: http
//  Host: 45.141.102.243
//  Version: 1.0.0
//
//  Security:
//  - access_token:
//
//  SecurityDefinitions:
//  access_token:
//    type: apiKey
//    name: Authorization
//    in: header
//
//    Consumes:
//    - application/json
//
//    Produces:
//    - application/json
//
// swagger:meta
package main

import (
	"Diploma/internal/router"
	"net/http"

	userDelivery "Diploma/internal/server/delivery"
	userRepo "Diploma/internal/server/repository"
	userUsecase "Diploma/internal/server/usecase"

	eventDelivery "Diploma/internal/event/delivery"
	eventRepo "Diploma/internal/event/repository"
	eventUsecase "Diploma/internal/event/usecase"

	"Diploma/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	docMiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/spf13/viper"
)

func main() {

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
	sessionR := userRepo.NewSessionRepository(redisDB)
	eventR := eventRepo.NewEventRepository(postgresDB)

	userU := userUsecase.NewUserUsecase(userR, sessionR)
	eventU := eventUsecase.NewEventUsecase(eventR)

	userD := userDelivery.NewUserDelivery(userU)
	eventD := eventDelivery.NewEventDelivery(eventU)

	baseRouter := gin.Default()
	routerAPI := baseRouter.Group("/api")

	opts := docMiddleware.RedocOpts{
		SpecURL: "swagger.yaml",
	}
	docsHandler := docMiddleware.Redoc(opts, nil)
	 
	baseRouter.GET("/docs", gin.WrapH(docsHandler))
	baseRouter.GET("/swagger.yaml", gin.WrapH(http.FileServer(http.Dir("../../"))))

	userRouter := routerAPI.Group("/user")
	router.UserEndpoints(userRouter, userD)

	eventRouter := routerAPI.Group("/events")
	router.EventEndpoints(eventRouter, eventD)

	port := viper.GetString("server.port")
	baseRouter.Run(":"+port)
}