package handler

import (
	"encoding/json"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	api_controller "github.com/sdn_global/api/controller/proto"
	"github.com/sdn_global/srv/ovncontroller/controller"
)

func SyncMessageHandler(publication broker.Publication) error{

	msg := publication.Message().Body
	obj := new(api_controller.ControllerObj)
	if err:=json.Unmarshal(msg, obj); err!=nil{
		log.Fatal("Unmarshal msg encounter error %v", err)
	}

	ovnController := controller.NewOvn(obj.Ovn)

	if err:=ovnController.SyncData(); err!=nil{
		log.Fatalf("err fuck this damn code, I do not want to make up err msg again %v", err)
	}

	return nil
}
