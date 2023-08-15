package subscribe

import (
	"github.com/joinusordie/Wildberries_L0/internal/service"
	"github.com/nats-io/stan.go"
)

type Subscribe struct {
	sc stan.Conn
	services *service.Service
}

func NewSubscribe(sc stan.Conn, services *service.Service) *Subscribe {
	st := &Subscribe{
		sc: sc,
		services: services,
	}
	st.Subscribing()
	return st
}

func (s *Subscribe) Subscribing() {
	s.getOrder()
}