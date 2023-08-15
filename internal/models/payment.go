package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

func (p Payment) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Payment) Scan(value interface{}) error {
	b, err := value.([]byte)
	if !err {
		logrus.Fatalf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &p)
}