package handler

import (
	"context"
	"github.com/sdn_global/srv/appcron/crontask"
	"github.com/sdn_global/srv/appcron/proto"
)

type AppCronHandler struct {
}

func (ac *AppCronHandler) GetCronTask(ctx context.Context, req *appcron.Empty, rsp *appcron.ListCronTask) error {
	return nil
}

func (ac *AppCronHandler) RegisterCronTask(ctx context.Context, req *appcron.CronTask, rsp *appcron.Response) error {
	task := crontask.NewCronTask(req.TaskType)
	task.InitTask(req.Tm)
	if err := task.RunTask(); err!=nil{
		rsp.Status = 0
		rsp.Msg = "something wrong"
	}

	return nil
}

func (ac *AppCronHandler) UpdateCronTask(ctx context.Context, req *appcron.CronTask, rsp *appcron.Response) error {
	return nil
}


