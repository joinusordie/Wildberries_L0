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

	addOrderQuery := fmt.Sprintf("INSERT INTO %s VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)",
	orderTable)
	_, err := r.db.Exec(addOrderQuery, 
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		order.Delivery,
		order.Payment,
		order.Items,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.ShardKey,
		order.SmId,
		order.DateCreated,
		order.OofShard)
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
	rows, err := r.db.Queryx("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	var res []models.Order
	for rows.Next() {
		var order models.Order
		if err:= rows.StructScan(&order); err != nil {
			return nil, err
		}
		res = append(res, order)
	}

	return &res, nil
}