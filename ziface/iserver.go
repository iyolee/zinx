package ziface

type IServer interface {
	// 启动通信服务
	Start()
	// 停止服务
	Stop()
	// 启动整体业务服务
	Serve()
}
