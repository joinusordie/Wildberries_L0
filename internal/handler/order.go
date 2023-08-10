package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joinusordie/Wildberries_L0/internal/models"
)

func (h *Handler) getOrderById(c *gin.Context) {

	id := c.Param("id")

	list, err := h.services.Order.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}


func (h *Handler) addOrder(c *gin.Context) {
	var input models.Order

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Order.AddOrder(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")
}
