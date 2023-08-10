package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/joinusordie/Wildberries_L0/internal/models"
)

type OrderPostgres struct {
	db    *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) AddOrder(order models.Order) error {

	addOrderQuery := fmt.Sprintf("INSERT INTO %s (model) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)",
	orderTable)

	_, err := r.db.Exec(addOrderQuery, 
		order.Order_uid,
		order.Track_number,
		order.Entry,
		order.Delivery,
		order.Payment,
		order.Items,
		order.Locale,
		order.Internal_signature,
		order.Customer_id,
		order.Delivery_service,
		order.Shardkey,
		order.Sm_id,
		order.Date_created,
		order.Oof_shard)
	
	return err
}

func (r *OrderPostgres) GetById(orderUID string) (*models.Order, error) {
	getOrderQuery := fmt.Sprintf("SELECT * FROM %s WHERE order_uid'= $1", orderTable)

	var order models.Order
	
	err := r.db.QueryRow(getOrderQuery, orderUID).Scan(&order)
	if err != nil {
		return nil, err 
	}

	return &order, nil
}

func (r *OrderPostgres) GetAll() (*[]models.Order, error) {
	getAllOrderQuery := fmt.Sprintf("SELECT * FROM %s", orderTable)
	var order []models.Order

	err := r.db.Select(&order, getAllOrderQuery)
	if err != nil {
		return nil, err
	}

	return &order, nil
}