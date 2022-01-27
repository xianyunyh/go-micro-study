package main

import (
	"context"
	"fmt"
	pb "go-micro-study/proto"
	"log"

	"go-micro.dev/v4/client"
)

func main() {
	cli := client.NewClient()
	req := pb.CallRequest{Name: "hello"}
	resp := &pb.CallResponse{}

	err := cli.Call(context.TODO(), client.NewRequest("rpc_demo", "Hello.Call", &req), resp)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(resp.Msg)
}
