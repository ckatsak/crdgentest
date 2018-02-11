//
// CHRISTOULAS MATHAFAQUERS
//
package fake

import (
	v1alpha1 "github.com/ckatsak/crdgentest/pkg/client/clientset/versioned/typed/tests.ckatsak/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTestsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeTestsV1alpha1) Disttates(namespace string) v1alpha1.DisttateInterface {
	return &FakeDisttates{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTestsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
