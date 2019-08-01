package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.cron"),
	)

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
