package main

import (
	"github.com/ievseev/fibonacci-slicer/internal/handler"
	"github.com/ievseev/fibonacci-slicer/internal/server"
	"github.com/ievseev/fibonacci-slicer/internal/service"
	"log"
)

func main() {
	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("80", handlers.InitRoutes()); err != nil {
		log.Fatalln("error occured while running http server: %s", err.Error())
	}
}
