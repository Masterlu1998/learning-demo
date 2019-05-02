package service

import (
	"net/rpc"
	"net/http"
	"net"
	"fmt"
)

// 定义RPC入参
type Args struct {
	A int
	B int
}

// 自定义可调用的RPC接口
type Arith int

// 定义接口可调用方法
func (t *Arith) GetMessage(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func StartRPCService() {
	arith := new(Arith)
	// 注册接口
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1122")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	go http.Serve(l, nil)
}

