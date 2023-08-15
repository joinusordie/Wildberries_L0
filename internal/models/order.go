package models

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
)

type Order struct {
	OrderUid          string    `json:"order_uid" db:"order_uid"`
	TrackNumber       string    `json:"track_number" db:"track_number"`
	Entry             string    `json:"entry" db:"entry"`
	Delivery          []byte  `json:"delivery" db:"delivery"`
	Payment           []byte  `json:"payment" db:"payment"`
	Items             []byte    `json:"items" db:"items"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerId        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	ShardKey          string    `json:"shardkey" db:"shardkey"`
	SmId              int       `json:"sm_id" db:"sm_id"`
	DateCreated       time.Time `json:"date_created" db:"date_created"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`
}

func (o *Order) Scan(value interface{}) error {
	b, err := value.([]byte)
	if !err {
		logrus.Fatalf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &o)
}