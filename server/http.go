package main

import (
	"log"
	"net/http"

	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
)

const (
	serviceName = "http_demo"
	serviceAddr = ":9099"
)

func main() {
	serv := httpServer.NewServer(server.Address(serviceAddr), server.Name(serviceName))
	h := serv.NewHandler(handler())
	_ = serv.Handle(h)
	service := micro.NewService(micro.Server(serv))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(resp http.ResponseWriter, request *http.Request) {
		_, _ = resp.Write([]byte(`{"code":1}`))
	})
	return mux
}
