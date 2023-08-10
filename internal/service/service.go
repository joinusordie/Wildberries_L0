package service

import (
	"github.com/joinusordie/Wildberries_L0/internal/cache"
	"github.com/joinusordie/Wildberries_L0/internal/models"
	"github.com/joinusordie/Wildberries_L0/internal/repository"
)

type Order interface {
	AddOrder(order models.Order) error
	GetById(orderUID string) (*models.Order, error)
	GetAll() (*[]models.Order, error)

	AddAllInCache() error
	GetAllFromCache() (*[]models.Order, error)
	GetOrderFromCacheById(id string) (*models.Order, error)
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Order: NewOrderService(repos.Order, cache.Order),
	}
}
