package main

import (
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/server"

	"net/http"

	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
)

func main() {
	//创建http服务
	//serv := newHttpServer("demo", ":8080")
	//创建注册中心
	//reg := registry.NewConsulRegistry(":8500")
	//rpcServ := newRpcServer("demo", ":9099")
	//实例一个服务
	serv2 := server.NewServer(server.Address(":9099"))
	var service = micro.NewService(
		//micro.Server(serv),
		micro.Server(serv2),
		//micro.Server(rpcServ),
	)

	service.Init()
	//运行服务
	err := service.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func newHttpServer(name, addr string) server.Server {
	var serv = httpServer.NewServer(
		server.Address(addr),
		server.Name(name),
	)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"data":"hello world"}`))
	})
	h := serv.NewHandler(mux)
	_ = serv.Handle(h)
	return serv
}
