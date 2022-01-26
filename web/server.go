package main

import (
	"net/http"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4/registry"

	"go-micro.dev/v4/util/log"
	"go-micro.dev/v4/web"
)

func main() {
	serv := newServer()
	serv.Init()
	err := serv.Run()
	if err != nil {
		log.Error(err.Error())
	}

}

func newServer() web.Service {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(wr http.ResponseWriter, request *http.Request) {
		wr.Write([]byte("hello test"))
	})
	reg := consul.NewRegistry(registry.Addrs(":8500"))
	serv := web.NewService(
		web.Name("demo"),
		web.Address(":9999"),
		web.Registry(reg),
		web.StaticDir("./images"),
		web.Handler(mux),
	)

	return serv
}
