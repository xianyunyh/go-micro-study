## 注册中心

`go-micro` 提供了注册中心接口。内置了mdns。接口如下

```go
type Registry interface {
	Init(...Option) error
	Options() Options
	Register(*Service, ...RegisterOption) error
	Deregister(*Service, ...DeregisterOption) error
	GetService(string, ...GetOption) ([]*Service, error)
	ListServices(...ListOption) ([]*Service, error)
	Watch(...WatchOption) (Watcher, error)
	String() string
}
```

注册中心插件在插件目录`plugins/registry`

- github.com/asim/go-micro/plugins/registry/consul/v4
- github.com/asim/go-micro/plugins/registry/etcd/v4
....
  
简单使用
```shell
$ go get github.com/asim/go-micro/plugins/registry/consul/v4
```
```go
reg := consul.NewRegistry(registry.Addrs(":8500"))
//添加到service中
serv := micro.NewService(micro.Registry(reg))


```