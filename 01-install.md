安装go-micro
最新的go-micro已经移到 `https://github.com/asim/go-micro` 包名已经变成了`go-micro.dev/v4`
```shell
#安装go-micro
$ go get go-micro.dev/v4

#安装micro-cli
$ go install go-micro.dev/v4/cmd/micro@master

#安装grpc依赖

$ go get -u google.golang.org/protobuf/proto
$ go install github.com/golang/protobuf/protoc-gen-go@latest
$ go install go-micro.dev/v4/cmd/protoc-gen-micro@latest

# 安装protoc-micro
$ go install go-micro.dev/v4/cmd/protoc-gen-micro@v4
```

## 快速开始

```go
//创建服务
package main

import (
	"context"
	"log"
	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Address(":9090"),
	)

	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

```