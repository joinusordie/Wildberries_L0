package handler

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joinusordie/Wildberries_L0/internal/models"
)

type getAllModelResponse struct {
	Data *[]models.Order `json:"data"`
}

func (h *Handler) getOrderByIdFromCache(c *gin.Context) {

	id := c.Param("order_uid")

	order, err := h.services.Order.GetOrderFromCacheById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

// func (h *Handler) getAllOrderFromCache(c *gin.Context) {
// 	orders, err := h.services.Order.GetAllFromCache()
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, getAllModelResponse{
// 		Data: orders,
// 	})
// }

func (h *Handler) addOrder(c *gin.Context) {
	var input models.Order

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.AddOrder(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) handleIndex(c *gin.Context) {
		t, err := template.ParseFiles("./static/index.html")
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		}

		if err := t.Execute(c.Writer, nil); err != nil {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		}

		c.JSON(http.StatusOK, "OK")
}