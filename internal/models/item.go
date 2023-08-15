package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type Items []Item
type Item struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func (i Items) Value() (driver.Value, error) {
	return json.Marshal(i)

}

func (i *Items) Scan(value interface{}) error {
	b, err := value.([]byte)
	if !err {
		logrus.Fatalf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &i)
}