package model

type TaskModel struct {
	TaskId        string `bson:"taskid"`
	TaskName      string `bson:"taskname"`
	TaskStatus    string `bson:"taskstatus"`
	TaskStartTime int64  `bson:"taskstarttime"`
	TaskEndTime   int64  `bson:"taskendtime"`
	TaskLog       string `bson:"tasklog"`
}
