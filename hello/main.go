package main

import (
	"go-micro-study/hello/handler"
	pb "go-micro-study/hello/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "hello"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	_ = pb.RegisterHelloHandler(srv.Server(), new(handler.Hello))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
