package cache

import (
	"errors"
	"fmt"

	"github.com/joinusordie/Wildberries_L0/internal/models"
)

type CacheOrder struct {
	order *[]models.Order
}

func NewCacheOrder() *CacheOrder {
	return &CacheOrder{
		order: new([]models.Order),
	}
}

func (c *CacheOrder) AddAll(order *[]models.Order) error {
	*c.order = *order
	return nil
}

func (c *CacheOrder) AddOne(order *models.Order) error {
	*c.order = append(*c.order, *order)
	return nil
}

func (c *CacheOrder) GetAll() (*[]models.Order, error) {
	return c.order, nil
}

func (c *CacheOrder) GetById(orderUID string) (*models.Order, error) {
	for _, order := range *c.order {
		if order.Order_uid == orderUID {
			return &order, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Not found order with order_uid = %s", orderUID))
}