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
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	maindb "github.com/wildanpurnomo/nsq-ayayaclap/main-service/db/migrations/postgre"
	repositories "github.com/wildanpurnomo/nsq-ayayaclap/main-service/db/repositories"
	gqlschema "github.com/wildanpurnomo/nsq-ayayaclap/main-service/gql/schema"
	"github.com/wildanpurnomo/nsq-ayayaclap/main-service/libs"
	nsqutil "github.com/wildanpurnomo/nsq-ayayaclap/main-service/nsq"
	restcontrollers "github.com/wildanpurnomo/nsq-ayayaclap/main-service/rest-controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	maindb.InitPostgre()
	maindb.MigratePostgre()
	repositories.InitRepository(maindb.PostgresDB)

	producer, err := nsqutil.CreateNewProducer()
	if err != nil {
		log.Fatal(err)
	}
	nsqutil.InjectProducer(producer)

	if os.Getenv("ENV_SCHEMA") == "https" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.Use(libs.CORSMiddleware())

	graphqlSchema, err := gqlschema.InitGQLSchema()
	if err != nil {
		log.Fatal(err)
	}
	gqlHandler := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		Pretty:   true,
		GraphiQL: true,
	})
	gqlHandlerFunc := gin.HandlerFunc(func(c *gin.Context) {
		gqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/gql", gqlHandlerFunc)
	r.POST("/gql", gqlHandlerFunc)
	r.GET("/api/user/email-confirmation", restcontrollers.ConfirmUserEmail)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
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
	log.Println("Shutting down gracefully. Press CTRL+C to force")

	// do something here I guess
	maindb.PostgresDB.Close()
	repositories.Repo.CloseDB()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server is forced shutdown: ", err)
	}

	log.Println("Server is exiting...")
}
