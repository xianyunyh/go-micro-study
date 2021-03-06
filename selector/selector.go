package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go-micro.dev/v4/client"

	pb "go-micro-study/proto"

	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

func main() {

	sel := selector.NewSelector(
		selector.Registry(registry.NewRegistry()),
		selector.SetStrategy(selector.Random), //Random 随机 RoundRobin：轮询
	)
	//选择rpc_demo 服务
	next, err := sel.Select("rpc_demo")
	if err != nil {
		log.Fatal(err.Error())
	}
	for i := 0; i < 10; i++ {
		node, err := next()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(node.Id, node.Address)
	}
	reg := registry.NewRegistry()
	sele := selector.NewSelector(
		selector.Registry(registry.NewRegistry()),
		selector.SetStrategy(selector.Random), //Random 随机 RoundRobin：轮询
	)
	cli := client.NewClient(client.Registry(reg), client.Selector(sele))
	for i := 0; i < 10; i++ {
		req := pb.CallRequest{Name: "hello" + strconv.Itoa(i)}
		resp := &pb.CallResponse{}

		err = cli.Call(context.TODO(), client.NewRequest("rpc_demo", "Hello.Call", &req), resp)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(resp.Msg)
	}

}
