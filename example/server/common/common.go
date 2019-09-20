package common

/**
请求公共参数
*/
type CommonRequest struct {
	Token  string // 请求权限校验参数
	Device Device // 设备标识
}

type Device struct {
	Channel string // 来源渠道
}

/*
返回公共参数
*/
type CommonResponse struct {
	RequestId string // 请求唯一id
	Code      string // 请求结果状态码
	Message   string // 请求结果描述信息
}
