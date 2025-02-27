package main

import (
	server "PGBridge"
	"PGBridge/internal/controller"
	"PGBridge/internal/router"
	"PGBridge/internal/repository"
	"PGBridge/internal/repository/postgres"
	"PGBridge/internal/service"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// TODO: init log

	// TODO: init .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	// TODO: init db
	db, err := postgres.NewPostgresRepository()
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}

	//TODO: start server
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	controller := controller.NewController(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), router.InitRoutes(controller)); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("Server is running")

	//TODO: Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Server Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		fmt.Errorf("error occured on db connection close: %s", err.Error())
	}
}
