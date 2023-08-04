package service

import "github.com/joinusordie/Wildberries_L0/internal/repository"

type Order interface {
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
