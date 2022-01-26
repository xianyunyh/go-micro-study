## Server

服务器是编写服务的构建基块. 在这里, 您可以命名您的服务, 注册请求处理程序, 添加中间件等. 该服务基于上述包, 为服务请求提供统一接口. 内置服务器是 RPC 系统. 将来可能还会有其他实现. 服务器还允许您定义多个编解码器以服务不同的编码消息. 内置了rpc服务

### Server接口

```go
type Server interface {
	Init(...Option) error
	Options() Options
	Handle(Handler) error
	NewHandler(interface{}, ...HandlerOption) Handler
	NewSubscriber(string, interface{}, ...SubscriberOption) Subscriber
	Subscribe(Subscriber) error
	Start() error
	Stop() error
	String() string
}
```

服务器插件
- github.com/asim/go-micro/plugins/server/http/v4
- github.com/asim/go-micro/plugins/server/grpc/v4

### 简单使用
完整代码![server/http.go](server/http.go)
```go
import (
	"log"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	"net/http"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
)
const (
    serviceName = "http_demo"
    serviceAddr = ":9099"
)
func main() {
    var serv = httpServer.NewServer(
    server.Address(serviceAddr),
    server.Name(serviceName),
    )
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    _, _ = w.Write([]byte(`{"data":"hello world"}`))
    })
    h := serv.NewHandler(mux)
    _ = serv.Handle(h)
    var service = micro.NewService(
        micro.Server(serv),
    )
    service.Init()
    //运行服务
    err := service.Run()
    if err != nil {
        log.Fatal(err.Error())
    }
}
```