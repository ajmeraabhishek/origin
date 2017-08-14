package fake

import (
	v1 "github.com/openshift/origin/pkg/user/generated/clientset/typed/user/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeUserV1 struct {
	*testing.Fake
}

func (c *FakeUserV1) Identities() v1.IdentityInterface {
	return &FakeIdentities{c}
}

func (c *FakeUserV1) Users() v1.UserResourceInterface {
	return &FakeUsers{c}
}

func (c *FakeUserV1) UserIdentityMappings() v1.UserIdentityMappingInterface {
	return &FakeUserIdentityMappings{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeUserV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
