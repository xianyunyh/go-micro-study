package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
)

func main() {
	reg := consul.NewRegistry(registry.Addrs(":8500"))
	serv := web.NewService(
		web.Registry(reg),
		web.Name("demo"),
	)
	resp, err := serv.Client().Get("http://demo/test")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
	_ = resp
}
