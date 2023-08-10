package service

import (
	"github.com/joinusordie/Wildberries_L0/internal/cache"
	"github.com/joinusordie/Wildberries_L0/internal/models"
	"github.com/joinusordie/Wildberries_L0/internal/repository"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	repo repository.Order
	cache cache.Order
}

func NewOrderService(repo repository.Order, cache cache.Order) *OrderService {
	return &OrderService{
		repo: repo,
		cache: cache,
	}
}

func (s *OrderService) AddOrder(order models.Order) error {
	err := s.repo.AddOrder(order)
	if err != nil {
		return nil
	}
	return s.cache.AddOne(&order)
}

func (s *OrderService) GetAll() (*[]models.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) GetById(orderUID string) (*models.Order, error) {
	return s.repo.GetById(orderUID)
}

func (s *OrderService) AddAllInCache() error {
	orders, err := s.GetAll()
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	err = s.cache.AddAll(orders)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	return nil
}

func (s* OrderService) GetAllFromCache() (*[]models.Order, error) {
	return s.cache.GetAll()
}

func (s* OrderService) GetOrderFromCacheById(orderUID string) (*models.Order, error) {
	return s.cache.GetById(orderUID)
}