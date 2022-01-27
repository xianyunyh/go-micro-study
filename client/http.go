package main

import (
	"context"
	"fmt"
	"log"

	"go-micro.dev/v4/client"

	httpClient "github.com/asim/go-micro/plugins/client/http/v4"
)

//go:generate go get github.com/asim/go-micro/plugins/client/http/v4

func main() {
	cli := httpClient.NewClient(client.ContentType("application/json"))
	req := cli.NewRequest("http_demo", "/test", []byte(`{"id":1}`))
	resp := make(map[string]interface{})
	err := cli.Call(context.TODO(), req, &resp)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(resp)
}
