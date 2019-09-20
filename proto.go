package main

/*
MicroService 用于存储解析后的文件原信息
*/
type MicroService struct {
	PackageName string
	Imports     []string
	ImportTime  bool
	Service     Service
	Messages    []Message
}

/*
Service 服务信息
*/
type Service struct {
	Name             string            // 服务名称
	PackageName      string            // 包名
	ServiceFunctions []ServiceFunction // 方法列表
}

/*
ServiceFunction 接口定义的方法
*/
type ServiceFunction struct {
	Name       string
	ParamType  string
	ResultType string
	Comment    string
	Stream     bool
	PingPong   bool
}

/*
Message 接口定义的结构体
*/
type Message struct {
	Name          string
	MessageFields []MessageField
}

/*
MessageField 结构体字段属性
*/
type MessageField struct {
	Index     int
	FieldName string
	FieldType string
	Comment   string
}
