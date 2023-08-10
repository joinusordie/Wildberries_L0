package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/joinusordie/Wildberries_L0/internal/models"
)

type Order interface {
	AddOrder(order models.Order) error
	GetById(orderUID string) (*models.Order, error)
	GetAll() (*[]models.Order, error)
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
	}
}
