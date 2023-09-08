package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test/config"
	"test/internal/handler"
	"test/internal/server"
	"test/internal/service"
	"test/internal/store"
	"test/pkg/db/postgres"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("error: init config %s", err.Error())
	}

	fmt.Println(cfg)

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {

		log.Fatalf("error: init postgres %s", err.Error())
	}

	store := store.NewStore(db)
	service := service.NewService(store)
	handler := handler.NewHandler(*service)

	server := new(server.Server)

	fmt.Printf("Server run at port %s\n", cfg.ServerPort)

	down := make(chan os.Signal, 1)

	signal.Notify(down, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if err := server.Start(cfg.ServerPort, handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}
