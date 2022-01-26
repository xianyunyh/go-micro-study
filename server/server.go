package main

import (
	h "go-micro-study/hello/handler"
	pb "go-micro-study/hello/proto"
	"log"

	"go-micro.dev/v4"
)

const (
	service = "rpc"
)

func main() {
	srv := micro.NewService(
		micro.Name(service),
		micro.Address(":9092"),
	)
	srv.Init()

	// Register handler
	_ = pb.RegisterHelloHandler(srv.Server(), new(h.Hello))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
