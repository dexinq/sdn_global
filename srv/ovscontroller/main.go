package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	app_handler "github.com/sdn_global/api/ovs/handler"
	"github.com/sdn_global/srv/ovscontroller/handler"
)

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.ovscontroller"),
	)

	service.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	if _, err := broker.Subscribe(app_handler.OvsManageTopic, handler.MessageClassification); err != nil {
		log.Fatalf("Broker Subscribe topic error %v", err)
	}

	if err := service.Run(); err != nil {
		log.Fatalf("Service start failed")
	}
}
