package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/joinusordie/Wildberries_L0/internal/models"
)

type OrderPostgres struct {
	db    *sqlx.DB
	cache map[string]*models.Order
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) AddOrder(order models.Order) error {

	_, err := r.db.Exec(
		"INSERT INTO orders VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)",
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
		order.OofShard,
	)
	if err != nil {
		return err
	}

	r.cache[order.OrderUid] = &order

	return nil
}

func (r *OrderPostgres) GetCache() error {
	var orders []models.Order

	query := fmt.Sprintf(`SELECT * FROM %s`, orderTable)
	err := r.db.Select(&orders, query)

	return err
}

func (r *OrderPostgres) GetOrderById(orderUID string) (*models.Order, error) {
	if _, ok := r.cache[orderUID]; ok {
		return r.cache[orderUID], nil
	}

	order := &models.Order{}
	if err := r.db.QueryRowx(
		"SELECT * FROM orders WHERE order_uid = $1",
		orderUID,
	).StructScan(
		order,
	); err != nil {
		return nil, err
	}

	r.cache[order.OrderUid] = order

	return order, nil
}
