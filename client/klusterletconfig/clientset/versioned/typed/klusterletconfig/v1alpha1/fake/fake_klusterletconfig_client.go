// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/stolostron/cluster-lifecycle-api/client/klusterletconfig/clientset/versioned/typed/klusterletconfig/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeConfigV1alpha1 struct {
	*testing.Fake
}

func (c *FakeConfigV1alpha1) KlusterletConfigs() v1alpha1.KlusterletConfigInterface {
	return newFakeKlusterletConfigs(c)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeConfigV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
