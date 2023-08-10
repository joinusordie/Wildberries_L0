package cache

import "github.com/joinusordie/Wildberries_L0/internal/models"

type Order interface {
	AddAll(orders *[]models.Order) error
	AddOne(order *models.Order) error
	GetAll() (*[]models.Order, error)
	GetById(orderUID string) (*models.Order, error)
}

type Cache struct {
	Order
}

func NewCache() *Cache {
	return &Cache{
		Order: NewCacheOrder(),
	}
}