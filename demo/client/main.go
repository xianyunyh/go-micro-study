package main

import (
	"context"
	"fmt"
	pb "go-micro-study/hello/proto"
	"go-micro-study/registry"
	"log"

	grcClient "github.com/asim/go-micro/plugins/client/grpc/v4"
	httpClient "github.com/asim/go-micro/plugins/client/http/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/selector"
)

func main() {
	err := callRpc()
	if err != nil {
		log.Fatal(err)
	}

}

func callRpc() error {
	reg := registry.NewConsulRegistry(":8500")
	var sel = selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.Random), //随机选择
	)
	_ = sel
	cli := grcClient.NewClient()
	req := cli.NewRequest("demo", "Hello.Call", &pb.CallRequest{Name: "hello world"})
	resp := pb.CallResponse{}
	err := cli.Call(context.TODO(), req, &resp)
	if err != nil {
		return err
	}
	fmt.Println(resp.Msg)
	return nil
}

func callApi() error {
	reg := registry.NewConsulRegistry(":8500")
	//创建选择器
	var sel = selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.Random), //随机选择
	)
	cli := httpClient.NewClient(client.Selector(sel),
		client.ContentType("application/json"))
	req := cli.NewRequest("demo", "/", "")
	var resp map[string]string
	_ = cli.Call(context.TODO(), req, &resp)
	fmt.Print(resp["data"])
	return nil
}
