package handler

import (
	"context"
	"github.com/micro/go-micro/util/log"
	"github.com/dexinq/sdn_global/srv/dataservice/proto"
)

type DataHandler struct {
}

func (dh *DataHandler) DataServe(ctx context.Context, req *dataservice.DataAcquire, rsp *dataservice.Response) error {
	data, err := GetDataConfig(req)
	if err!=nil{
		rsp.Msg = err.Error()
		rsp.Status = 0
		rsp.Do = []*dataservice.DataObject{}
		return nil
	}
	rsp.Msg = "success"
	rsp.Status = 1
	rsp.Do = data
	return nil
}

func GetDataConfig(acquire *dataservice.DataAcquire) ([]*dataservice.DataObject, error) {
	switch acquire.DataSource {
	case "mongo":
		log.Log("do mongo data access") //do mongo data access
		res,err := GetMongoData(acquire)
		if err != nil {
			return []*dataservice.DataObject{}, err
		}
		return res, err
	case "mysql":
		log.Log("do mysql data access") //do mysql data access
		res,err := GetMysqlData(acquire)
		if err != nil {
			return []*dataservice.DataObject{}, err
		}
		return res, err
	case "es":
		log.Log("do es data access") //do elasticsearch access
		res,err := GetESData(acquire)
		if err != nil {
			return []*dataservice.DataObject{}, err
		}
		return res, err
	default:
		log.Fatalf("do not support data source")
		return []*dataservice.DataObject{}, nil   // TODO: define a self error
	}

}

func GetMongoData(acquire *dataservice.DataAcquire) ([]*dataservice.DataObject, error) {
	return []*dataservice.DataObject{}, nil
}

func GetMysqlData(acquire *dataservice.DataAcquire) ([]*dataservice.DataObject,error) {
	return []*dataservice.DataObject{}, nil
}

func GetESData(acquire *dataservice.DataAcquire) ([]*dataservice.DataObject, error) {
	return []*dataservice.DataObject{}, nil
}
