package service

import (
	"github.com/joinusordie/Wildberries_L0/internal/models"
	"github.com/joinusordie/Wildberries_L0/internal/repository"
)

type Order interface {
	AddOrder(models.Order) error
	GetOrderById(string) (models.Order, error)
	GetCache() error
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
