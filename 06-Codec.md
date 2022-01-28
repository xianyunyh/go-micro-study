## Codec编码器

编码器主要负责数据的编码，默认使用protobuf格式。框架内置了常用的编码格式
```go
var DefaultCodecs = map[string]codec.NewCodec{
		"application/grpc":         grpc.NewCodec,
		"application/grpc+json":    grpc.NewCodec,
		"application/grpc+proto":   grpc.NewCodec,
		"application/json":         json.NewCodec,
		"application/json-rpc":     jsonrpc.NewCodec,
		"application/protobuf":     proto.NewCodec,
		"application/proto-rpc":    protorpc.NewCodec,
		"application/octet-stream": raw.NewCodec,
	}
```


### 服务端使用

```go
serv2 := server.NewServer(server.Address(":9099").,server.Codec("application/json",json.NewCodec))
	var service = micro.NewService(
		micro.Server(serv2),
	)

```

### 客户端使用
```go
cli := client.NewClient(client.Codec("application/json",json.NewCodec))
```