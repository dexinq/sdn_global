package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/sdn_global/srv/dataservice/handler"
	"github.com/sdn_global/srv/dataservice/proto"
)


func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.dataservice"),
	)
	service.Init()

	chdl := handler.NewControllerHandler()

	if err:=dataservice.RegisterControllerServiceHandler(service.Server(), chdl);err!=nil{
		log.Fatalf(err.Error())
	}


	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
