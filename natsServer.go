package l0

import (
	"encoding/json"

	"github.com/joinusordie/Wildberries_L0/internal/models"
	"github.com/joinusordie/Wildberries_L0/internal/service"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type NServer struct {
	natsConn stan.Conn
	natsSub  stan.Subscription
	serv     service.Order
}

func NatsServer(natsURL string, serv service.Service, nutsClusterID string, nutsSubj string) error {

	s := &NServer{}

	sc, err := stan.Connect(nutsClusterID, "subscriber", stan.NatsURL(natsURL))
	if err != nil {
		return err
	}
	s.natsConn = sc

	sub, err := sc.Subscribe(nutsSubj, s.handleNatsMessage)
	if err != nil {
		return err
	}

	s.natsSub = sub

	if err := s.serv.GetCache(); err != nil {
		return err
	}

	return nil
}

func (s *NServer) handleNatsMessage(m *stan.Msg) {
	order := models.Order{}
	if err := json.Unmarshal(m.Data, &order); err != nil {
		logrus.Error(err)
		return
	}
	if err := s.serv.AddOrder(order); err != nil {
		logrus.Error(err)
		return
	}

	logrus.Info("Added order: " + order.OrderUid)
}
