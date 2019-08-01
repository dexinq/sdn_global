package handler

import "github.com/micro/go-micro/broker"

type ControllerHandler struct {
	broker broker.Broker
}


func NewControllerHandler(broker broker.Broker) *ControllerHandler {
	return &ControllerHandler{broker:broker}
}




