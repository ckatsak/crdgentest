//
// CHRISTOULAS MATHAFAQUERS
//
package fake

import (
	clientset "github.com/ckatsak/crdgentest/pkg/client/clientset/versioned"
	testsv1alpha1 "github.com/ckatsak/crdgentest/pkg/client/clientset/versioned/typed/tests.ckatsak/v1alpha1"
	faketestsv1alpha1 "github.com/ckatsak/crdgentest/pkg/client/clientset/versioned/typed/tests.ckatsak/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// TestsV1alpha1 retrieves the TestsV1alpha1Client
func (c *Clientset) TestsV1alpha1() testsv1alpha1.TestsV1alpha1Interface {
	return &faketestsv1alpha1.FakeTestsV1alpha1{Fake: &c.Fake}
}

// Tests retrieves the TestsV1alpha1Client
func (c *Clientset) Tests() testsv1alpha1.TestsV1alpha1Interface {
	return &faketestsv1alpha1.FakeTestsV1alpha1{Fake: &c.Fake}
}
