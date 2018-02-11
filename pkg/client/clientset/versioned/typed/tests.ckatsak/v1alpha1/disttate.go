//
// CHRISTOULAS MATHAFAQUERS
//
package v1alpha1

import (
	v1alpha1 "github.com/ckatsak/crdgentest/pkg/apis/tests.ckatsak/v1alpha1"
	scheme "github.com/ckatsak/crdgentest/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DisttatesGetter has a method to return a DisttateInterface.
// A group's client should implement this interface.
type DisttatesGetter interface {
	Disttates(namespace string) DisttateInterface
}

// DisttateInterface has methods to work with Disttate resources.
type DisttateInterface interface {
	Create(*v1alpha1.Disttate) (*v1alpha1.Disttate, error)
	Update(*v1alpha1.Disttate) (*v1alpha1.Disttate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Disttate, error)
	List(opts v1.ListOptions) (*v1alpha1.DisttateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Disttate, err error)
	DisttateExpansion
}

// disttates implements DisttateInterface
type disttates struct {
	client rest.Interface
	ns     string
}

// newDisttates returns a Disttates
func newDisttates(c *TestsV1alpha1Client, namespace string) *disttates {
	return &disttates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the disttate, and returns the corresponding disttate object, and an error if there is any.
func (c *disttates) Get(name string, options v1.GetOptions) (result *v1alpha1.Disttate, err error) {
	result = &v1alpha1.Disttate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("disttates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Disttates that match those selectors.
func (c *disttates) List(opts v1.ListOptions) (result *v1alpha1.DisttateList, err error) {
	result = &v1alpha1.DisttateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("disttates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested disttates.
func (c *disttates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("disttates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a disttate and creates it.  Returns the server's representation of the disttate, and an error, if there is any.
func (c *disttates) Create(disttate *v1alpha1.Disttate) (result *v1alpha1.Disttate, err error) {
	result = &v1alpha1.Disttate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("disttates").
		Body(disttate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a disttate and updates it. Returns the server's representation of the disttate, and an error, if there is any.
func (c *disttates) Update(disttate *v1alpha1.Disttate) (result *v1alpha1.Disttate, err error) {
	result = &v1alpha1.Disttate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("disttates").
		Name(disttate.Name).
		Body(disttate).
		Do().
		Into(result)
	return
}

// Delete takes name of the disttate and deletes it. Returns an error if one occurs.
func (c *disttates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("disttates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *disttates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("disttates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched disttate.
func (c *disttates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Disttate, err error) {
	result = &v1alpha1.Disttate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("disttates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
