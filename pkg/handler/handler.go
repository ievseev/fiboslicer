package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ievseev/fibonacci-slicer/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		fibonacciSlice := api.Group("/fibonacci-slice")
		{
			fibonacciSlice.POST("/", h.getFibonacciSlice)
		}
	}

	return router
}
