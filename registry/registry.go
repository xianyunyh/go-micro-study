package registry

import (
	"fmt"
	"go-micro.dev/v4/registry"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"log"
)

//创建consul注册中心
func NewConsulRegistry(addr string) registry.Registry {
	return consul.NewRegistry(
		registry.Addrs(addr),
		)
}

func registryAction(addr string,name string)  {
	//创建实例
	reg := NewConsulRegistry(addr)
	//获取服务
	serv ,err := reg.GetService(name)
	if err != nil {
		log.Fatal(err)
	}
	if len(serv) ==0 {
		return
	}
	fmt.Println((*serv[0]).Name)

}