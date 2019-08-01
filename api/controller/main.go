package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"github.com/sdn_global/api/controller/handler"
	"github.com/sdn_global/api/controller/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.v1.app"),
	)

	service.Init()

	if err:=broker.Init(); err!=nil{
		log.Fatalf(err.Error())
	}

	if err:=broker.Connect(); err!=nil{
		log.Fatalf(err.Error())
	}
	conhdl := handler.NewControllerHandler(service.Client(), service.Server().Options().Broker)

	if err:=controller.RegisterControllerHandler(service.Server(), conhdl); err!=nil{
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
