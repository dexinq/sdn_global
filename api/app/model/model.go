package model

type AppData struct {  // the data get app config here
	DBType string
	TableName string
	Granularity string
}

type App struct {
	Name string
	AppType string
	DataMeta AppData
}
