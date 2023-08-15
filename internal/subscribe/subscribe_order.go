package subscribe

import (
	"encoding/json"
	"fmt"

	"github.com/joinusordie/Wildberries_L0/internal/models"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (s *Subscribe) getOrder() {
	s.sc.Subscribe(viper.GetString("nats.subject"), s.msgHandler, stan.DurableName(viper.GetString("nats.subscriber")))
}

func (s *Subscribe) msgHandler(msg *stan.Msg) {
	fmt.Println("SSSSSSSSSSSSSSSS")
	var order models.Order
	if err := json.Unmarshal(msg.Data, &order); err != nil {
		logrus.Errorf("error parse order from subcribe: %s", err.Error())
		return
	}
	if order.OrderUid == "" {
		logrus.Error("order_uid cannot be empty")
		return
	}
	if err := s.services.AddOrder(order); err != nil {
		logrus.Errorf("cannot be add order: %s", err.Error())
	}
}
