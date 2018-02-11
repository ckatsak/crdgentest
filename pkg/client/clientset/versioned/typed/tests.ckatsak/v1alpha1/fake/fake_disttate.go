//
// CHRISTOULAS MATHAFAQUERS
//
package fake

import (
	v1alpha1 "github.com/ckatsak/crdgentest/pkg/apis/tests.ckatsak/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDisttates implements DisttateInterface
type FakeDisttates struct {
	Fake *FakeTestsV1alpha1
	ns   string
}

var disttatesResource = schema.GroupVersionResource{Group: "tests.ckatsak", Version: "v1alpha1", Resource: "disttates"}

var disttatesKind = schema.GroupVersionKind{Group: "tests.ckatsak", Version: "v1alpha1", Kind: "Disttate"}

// Get takes name of the disttate, and returns the corresponding disttate object, and an error if there is any.
func (c *FakeDisttates) Get(name string, options v1.GetOptions) (result *v1alpha1.Disttate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(disttatesResource, c.ns, name), &v1alpha1.Disttate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disttate), err
}

// List takes label and field selectors, and returns the list of Disttates that match those selectors.
func (c *FakeDisttates) List(opts v1.ListOptions) (result *v1alpha1.DisttateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(disttatesResource, disttatesKind, c.ns, opts), &v1alpha1.DisttateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DisttateList{}
	for _, item := range obj.(*v1alpha1.DisttateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested disttates.
func (c *FakeDisttates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(disttatesResource, c.ns, opts))

}

// Create takes the representation of a disttate and creates it.  Returns the server's representation of the disttate, and an error, if there is any.
func (c *FakeDisttates) Create(disttate *v1alpha1.Disttate) (result *v1alpha1.Disttate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(disttatesResource, c.ns, disttate), &v1alpha1.Disttate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disttate), err
}

// Update takes the representation of a disttate and updates it. Returns the server's representation of the disttate, and an error, if there is any.
func (c *FakeDisttates) Update(disttate *v1alpha1.Disttate) (result *v1alpha1.Disttate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(disttatesResource, c.ns, disttate), &v1alpha1.Disttate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disttate), err
}

// Delete takes name of the disttate and deletes it. Returns an error if one occurs.
func (c *FakeDisttates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(disttatesResource, c.ns, name), &v1alpha1.Disttate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDisttates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(disttatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DisttateList{})
	return err
}

// Patch applies the patch and returns the patched disttate.
func (c *FakeDisttates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Disttate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(disttatesResource, c.ns, name, data, subresources...), &v1alpha1.Disttate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disttate), err
}
