package main

import (
	"context"
	"fmt"
	"os"

	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/qchart-app/service-tv-udf/internal/infrastructure/cache"
	"github.com/qchart-app/service-tv-udf/internal/infrastructure/database"
	http "github.com/qchart-app/service-tv-udf/internal/infrastructure/http"
	"github.com/qchart-app/service-tv-udf/internal/infrastructure/repository"
	"github.com/qchart-app/service-tv-udf/internal/usecase"
	"github.com/qchart-app/service-tv-udf/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	dbPort33 := flag.Int("db-port", 0, "Database port number")
	flag.Parse()
	if flag.NFlag() > 0 {
		fmt.Printf("port :%d", *dbPort33)
		os.Exit(0)
	} else {
		// Run normal code
		// ...

		env := os.Getenv("APP_ENV")
		if env == "development" {
			fmt.Println("Running in development mode")
			viper.SetConfigName("config.dev")
		} else if env == "production" {
			fmt.Println("Running in production mode")
			viper.SetConfigName("config")
		} else {

			logrus.Panic(fmt.Errorf("Unknown environment"))
		}

		viper.SetConfigType("yaml")

		viper.AddConfigPath("./configs")
		viper.AddConfigPath("./../../configs/")

		err := viper.ReadInConfig()
		if err != nil {
			logrus.Panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
}
func RegisterRedisService(client *cache.CacheClient, channel_name_pub string) {

	subscriber := cache.CacheClient.NewSubscriber(*client)

	msgService := usecase.NewRedisMSGService(subscriber)

	err := msgService.ListenForMessages(context.Background(), channel_name_pub)
	if err != nil {
		logrus.Panic(fmt.Errorf("fatal Redis sub : %s", err))
	}
}

func main() {
	dbConfig := viper.GetStringMapString("db")

	db, err := database.NewGormDB(dbConfig)

	if err != nil {

		logrus.Fatalf("Failed to create PostgresDB instance: %v", err)

	}
	//------------------------CACHE ----------------
	cache_Config := viper.GetStringMapString("redis")
	host := cache_Config["host"]
	pass := cache_Config["password"]
	channel_pub_event := cache_Config["channel_pub_event"]
	db_redis, _ := util.ToInt(cache_Config["db"])

	cache_client, erro_cache := cache.NewRedisClient(host, pass, db_redis)

	if erro_cache != nil {
		logrus.Fatalf("Failed to create Cache client  instance: %v", err)
	}
	// Register Redisx
	RegisterRedisService(&cache_client, channel_pub_event)
	//----
	userService := usecase.NewUserServiceRedis(&cache_client)
	// Create repository instance
	userRepo := repository.NewGormUserRepository(db)

	// Create use case instance
	userUseCase := usecase.NewUserUseCase(userRepo, userService)

	// Create handler instance
	userHandler := http.NewUserHandler(userUseCase)

	// Create Fiber app instance
	app := fiber.New()

	// Define routes
	userHandler.RegisterRoutes(app)

	// Start server-----
	server_Config := viper.GetStringMapString("server")
	server_host := server_Config["host"]
	server_port := server_Config["port"]
	err1 := app.Listen(fmt.Sprintf("%v:%v", server_host, server_port))
	if err1 != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
