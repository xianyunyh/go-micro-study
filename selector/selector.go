package selector

import (
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

func New(reg registry.Registry)  selector.Selector {

	return selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.Random))

}

func testSelector()  {
	s := New(registry.NewRegistry())
	//选择一个实例
	s.Select("demo")

}