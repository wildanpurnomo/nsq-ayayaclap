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
	nsqutil "github.com/wildanpurnomo/nsq-ayayaclap/smtp-service/nsq"
)

func main() {
	consumer, err := nsqutil.CreateNewConsumer()
	if err != nil {
		log.Fatal(err)
	}
	nsqutil.InjectConsumer(consumer)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if os.Getenv("ENV_SCHEMA") == "https" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Service is currently running"})
	})

	srv := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Shutting down gracefully. Press CTRL+C to force")

	nsqutil.StopConsumer()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server is forced shutdown: ", err)
	}

	log.Println("Server is exiting...")
}
