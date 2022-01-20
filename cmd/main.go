package main

import (
	"github.com/ievseev/fibonacci-slicer"
	"github.com/ievseev/fibonacci-slicer/pkg/handler"
	"github.com/ievseev/fibonacci-slicer/pkg/service"
	"log"
)

func main() {
	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(fiboslicer.Server)
	if err := srv.Run("80", handlers.InitRoutes()); err != nil {
		log.Fatalln("error occured while running http server: %s", err.Error())
	}
}
