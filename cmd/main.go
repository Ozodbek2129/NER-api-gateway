package main

import (
	"log"
	"gateway/api"
	"gateway/api/handler"
	"gateway/casbin"
	"gateway/config"
	"gateway/pkg/client"
	"gateway/pkg/logger"
	"gateway/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("API Gateway started successfully!")

	logger := logger.NewLogger()
	logger.Info("API Gateway started successfully!")

	enforcer, err := casbin.CasbinEnforcer(logger)
	if err != nil {
		log.Println("Error initializing casbin enforcer", "error", err.Error())
		logger.Error("Error initializing enforcer", "error", err.Error())
		return
	}

	cfg := config.Load()

	// gRPC ServiceManager
	serviceManager, err := client.NewServiceManager()
	if err != nil {
		log.Println("Error initializing service manager", "error", err.Error())
		logger.Error("Error initializing service manager", "error", err.Error())
		return
	}

	// Redis client
	redisClient, err := redis.NewRedisClient()
	if err != nil {
		log.Println("Error initializing Redis", "error", err.Error())
		logger.Error("Error initializing Redis", "error", err.Error())
		return
	}
	defer redisClient.Close()

	// Handler
	h := handler.NewHandler(
		serviceManager.UserService(),
		serviceManager.Productionservice(),
		logger,
		enforcer,
		redisClient,
	)

	controller := api.NewController(gin.Default(), redisClient, enforcer)
	controller.SetupRoutes(*h, logger)

	// Start server
	if err := controller.StartServer(cfg); err != nil {
		log.Println("Failed to start server", "error", err.Error())
		logger.Error("Failed to start server", "error", err.Error())
	}
}