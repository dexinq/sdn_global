package handler

import (
	"github.com/google/uuid"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

const OvsManageTopic = "app.controller.manage"

type OvsManager struct {
	broker broker.Broker
}

func (om *OvsManager)AddNetflowRule() error{

	msg := broker.Message{
		Header: map[string]string{"id": uuid.New().String()},
		Body:[]byte{},
	}

	if err:=broker.Publish(OvsManageTopic, &msg);err!=nil{
		log.Fatal("Publish msg error %v", err)
	}
	return nil
}
