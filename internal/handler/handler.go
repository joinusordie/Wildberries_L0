package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joinusordie/Wildberries_L0/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())

	router.GET("/:order_uid", h.getOrderByIdFromCache)
	router.GET("/", h.getAllOrderFromCache)
	router.POST("/", h.addOrder)

	return router
}
