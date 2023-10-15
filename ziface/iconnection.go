package ziface

import "net"

type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 从当前连接获取原始的 socket TCPConn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接 ID
	GetConnID() uint32
	// 获取远程客户端地址信息
	RemoteAddr() net.Addr
}

type HandFunc func(*net.TCPConn, []byte, int) error
