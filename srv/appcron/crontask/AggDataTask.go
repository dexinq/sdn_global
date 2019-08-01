package crontask

type AggDataTask struct {

}

func newAggDataTask() AggDataTask{
	return AggDataTask{}
}


func (a AggDataTask) InitTask(interface{}) {

}

func (a AggDataTask) RunTask() error {
	return nil
}

func (a AggDataTask) StopTask() error {

}
