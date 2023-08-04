package handler

import (
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

	router.GET("/:order_uid", h.getOrderById)

	return router
}
