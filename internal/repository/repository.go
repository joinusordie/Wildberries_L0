package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/joinusordie/Wildberries_L0/internal/models"
)

type Order interface {
	AddOrder(models.Order) error
	GetOrderById(string) (models.Order, error)
	GetCache() error
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
