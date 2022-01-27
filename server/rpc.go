package main

import (
	h "go-micro-study/handler"
	pb "go-micro-study/proto"
	"log"

	"go-micro.dev/v4"
)

func main() {
	//创建服务
	serv := micro.NewService(micro.Address(":9099"), micro.Name("rpc_demo"))
	//注册服务
	err := pb.RegisterHelloHandler(serv.Server(), new(h.Hello))
	if err != nil {
		log.Fatal(err.Error())
	}
	serv.Init()
	if err := serv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
