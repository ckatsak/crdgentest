//
// CHRISTOULAS MATHAFAQUERS
//
package v1alpha1

import (
	v1alpha1 "github.com/ckatsak/crdgentest/pkg/apis/tests.ckatsak/v1alpha1"
	"github.com/ckatsak/crdgentest/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type TestsV1alpha1Interface interface {
	RESTClient() rest.Interface
	DisttatesGetter
}

// TestsV1alpha1Client is used to interact with features provided by the tests.ckatsak group.
type TestsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *TestsV1alpha1Client) Disttates(namespace string) DisttateInterface {
	return newDisttates(c, namespace)
}

// NewForConfig creates a new TestsV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*TestsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TestsV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new TestsV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TestsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TestsV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *TestsV1alpha1Client {
	return &TestsV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TestsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
