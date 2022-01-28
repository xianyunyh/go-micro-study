## 选择器
选择器是在客户端调用服务端的时候使用的，客户端可以使用选择器从注册中心根据对应的算法选择对应的节点
`go-micro` 默认提供了 随机**Random** 和 轮询 **RoundRobin**

```go
sel := selector.NewSelector(
		selector.Registry(registry.NewRegistry()),
		selector.SetStrategy(selector.Random), //Random 随机 RoundRobin：轮询
	)
next, err := sel.Select("rpc_demo")
if err != nil {
    log.Fatal(err.Error())
}
fmt.Println(node.Id, node.Address)
```


配合客户端使用

```go
    //创建注册中心实例
    reg := registry.NewRegistry()
    //选择器
	sele := selector.NewSelector(
		selector.Registry(registry.NewRegistry()),
		selector.SetStrategy(selector.Random), //Random 随机 RoundRobin：轮询
	)
	//调用服务
	cli := client.NewClient(client.Registry(reg), client.Selector(sele))
	for i := 0; i < 10; i++ {
		req := pb.CallRequest{Name: "hello" + strconv.Itoa(i)}
		resp := &pb.CallResponse{}

		err = cli.Call(context.TODO(), client.NewRequest("rpc_demo", "Hello.Call", &req), resp)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(resp.Msg)
	}
```