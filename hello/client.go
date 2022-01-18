package main

import (
	"context"
	pb "go-micro-study/hello/proto"
	"time"

	"go-micro.dev/v4"

	log "go-micro.dev/v4/logger"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init()

	// Create client
	c := pb.NewHelloService(service, srv.Client())

	for {
		// Call service
		rsp, err := c.Call(context.Background(), &pb.CallRequest{Name: "John"})
		if err != nil {
			log.Fatal(err)
		}

		log.Info(rsp)

		time.Sleep(1 * time.Second)
	}
}
