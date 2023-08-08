package service

import (
	"github.com/joinusordie/Wildberries_L0/internal/models"
	"github.com/joinusordie/Wildberries_L0/internal/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) AddOrder(order models.Order) error {
	return s.repo.AddOrder(order)
}

func (s *OrderService) GetCache() error {
	return s.repo.GetCache()
}

func (s *OrderService) GetOrderById(orderUID string) (models.Order, error) {
	return s.repo.GetOrderById(orderUID)
}
