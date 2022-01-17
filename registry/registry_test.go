package registry

import "testing"

func Test_Consul(t *testing.T)  {
	registryAction(":8500","go.micro.server")
}
