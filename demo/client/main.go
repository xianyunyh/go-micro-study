package main

import (
	"context"
	"fmt"
	httpClient "github.com/asim/go-micro/plugins/client/http/v4"
	"go-micro-study/registry"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/selector"
)

func main()  {
	_ = callApi()
	
}

func callApi() error  {
	reg := registry.NewConsulRegistry(":8500")
	//创建选择器
	var sel = selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.Random),//随机选择
		)
	cli := httpClient.NewClient(client.Selector(sel),
		client.ContentType("application/json"))
	req := cli.NewRequest("demo","/","")
	var resp map[string]string
	_ = cli.Call(context.TODO(), req, &resp)
	fmt.Print(resp["data"])
	return nil
}