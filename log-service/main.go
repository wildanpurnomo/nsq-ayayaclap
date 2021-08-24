package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	dbrepos "github.com/wildanpurnomo/nsq-ayayaclap/log-service/db/repositories"
	nsqutil "github.com/wildanpurnomo/nsq-ayayaclap/log-service/nsq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := dbrepos.CreateNewDBClient()
	if err != nil {
		log.Fatal(err)
	}
	dbrepos.InitRepository(dbClient)
	dbrepos.ConnectDBClient()

	userRegConsumer, err := nsqutil.CreateNewConsumer(
		"test",
		"user_registration",
		&nsqutil.UserRegistrationHandler{},
	)
	if err != nil {
		log.Fatal(err)
	}

	nsqutil.InitLoggerConsumers()
	nsqutil.AppendConsumer(userRegConsumer)

	if os.Getenv("ENV_SCHEMA") == "https" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	srv := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Shutting down gracefully")

	dbrepos.DisconnectDBClient()
	nsqutil.StopConsumers()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server is forced into shutdown: ", err)
	}

	log.Println("Server is exiting")
}
