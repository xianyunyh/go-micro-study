## 客户端
开发完服务端就可以使用客户端进行调用了，`go-micro`封装了客户端的接口，可以使用客户端方便的调用服务

### 客户端接口
```go

type Client interface {
	Init(...Option) error
	Options() Options
	NewMessage(topic string, msg interface{}, opts ...MessageOption) Message
	NewRequest(service, endpoint string, req interface{}, reqOpts ...RequestOption) Request
	Call(ctx context.Context, req Request, rsp interface{}, opts ...CallOption) error
	Stream(ctx context.Context, req Request, opts ...CallOption) (Stream, error)
	Publish(ctx context.Context, msg Message, opts ...PublishOption) error
	String() string
}
```

### http客户端
完整代码 [`client/http.go`](client/http.go)

http客户端需要用使用插件的方式
```shell
$ go get  github.com/asim/go-micro/plugins/client/http/v4
```
```go
import (
    httpClient  "github.com/asim/go-micro/plugins/client/http/v4"
)
func  main() {
    cli := httpClient.NewClient(client.ContentType("application/json"))
    //http_demo 服务名 
    // /test 接口名 默认使用post请求
    req := cli.NewRequest("http_demo", "/test", []byte(`{"id":1}`))
    resp := make(map[string]interface{})
    err := cli.Call(context.TODO(), req, &resp)
    if err != nil {
        log.Fatal(err.Error())
    }
    fmt.Println(resp)
}
```

### RPC客户端
完整代码 [client/rpc.go](client/rpc.go)

```go
//创建客户端
cli := client.NewClient()
//请求对象
req := pb.CallRequest{Name: "hello"}
resp := &pb.CallResponse{}
// Hello.Call gRPC 服务方法
err := cli.Call(context.TODO(), client.NewRequest("rpc_demo", "Hello.Call", &req), resp)
if err != nil {
    log.Fatal(err.Error())
}
```