package crontask

type CronTask interface {
	InitTask(interface{})
	RunTask() error
	StopTask() error
}

func NewCronTask(taskType string) CronTask {
	return newAggDataTask()
}