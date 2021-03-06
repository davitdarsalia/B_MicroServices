package main

import (
	"context"
	"github.com/davitdarsalia/LendAppBackend/cache"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/handler"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// @title E_Commerce Microservices
// @version 0.1
// @description Endpoints for Authorization, Account, Settings, Payments

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	go logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := repository.NewPostgresDB()

	if err != nil {
		logrus.Fatalf("Error While Initializing DataBase Connection; %s", err.Error())
	}

	redisConn := cache.NewRedisCache(&redis.Options{
		Addr: os.Getenv("REDIS_CACHE_PORT"),
		DB:   0,
	})
	repos := repository.NewRepository(db)
	services := service.NewService(repos, redisConn)
	handlers := handler.NewHandler(services)

	server := new(entities.MainServer)

	go func() {
		if err := server.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error While Running Server On Port %s", os.Getenv("PORT"))
		}
	}()

	quitSignal := make(chan os.Signal, 1)

	signal.Notify(quitSignal, syscall.SIGTERM, syscall.SIGINT)

	<-quitSignal

	logrus.Print("Server Shut Down")

	if err := server.ShutDown(context.Background()); err != nil {
		logrus.Errorf("Failed To Shut Down Server: \n %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Failed To Close DB: \n %s", err.Error())
	}

}

func init() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}
}
