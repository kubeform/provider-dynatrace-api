/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/host/v1alpha1"
	scheme "kubeform.dev/provider-dynatrace-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NamingsGetter has a method to return a NamingInterface.
// A group's client should implement this interface.
type NamingsGetter interface {
	Namings(namespace string) NamingInterface
}

// NamingInterface has methods to work with Naming resources.
type NamingInterface interface {
	Create(ctx context.Context, naming *v1alpha1.Naming, opts v1.CreateOptions) (*v1alpha1.Naming, error)
	Update(ctx context.Context, naming *v1alpha1.Naming, opts v1.UpdateOptions) (*v1alpha1.Naming, error)
	UpdateStatus(ctx context.Context, naming *v1alpha1.Naming, opts v1.UpdateOptions) (*v1alpha1.Naming, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Naming, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NamingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Naming, err error)
	NamingExpansion
}

// namings implements NamingInterface
type namings struct {
	client rest.Interface
	ns     string
}

// newNamings returns a Namings
func newNamings(c *HostV1alpha1Client, namespace string) *namings {
	return &namings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the naming, and returns the corresponding naming object, and an error if there is any.
func (c *namings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Naming, err error) {
	result = &v1alpha1.Naming{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("namings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Namings that match those selectors.
func (c *namings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NamingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NamingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("namings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested namings.
func (c *namings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("namings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a naming and creates it.  Returns the server's representation of the naming, and an error, if there is any.
func (c *namings) Create(ctx context.Context, naming *v1alpha1.Naming, opts v1.CreateOptions) (result *v1alpha1.Naming, err error) {
	result = &v1alpha1.Naming{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("namings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(naming).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a naming and updates it. Returns the server's representation of the naming, and an error, if there is any.
func (c *namings) Update(ctx context.Context, naming *v1alpha1.Naming, opts v1.UpdateOptions) (result *v1alpha1.Naming, err error) {
	result = &v1alpha1.Naming{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("namings").
		Name(naming.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(naming).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *namings) UpdateStatus(ctx context.Context, naming *v1alpha1.Naming, opts v1.UpdateOptions) (result *v1alpha1.Naming, err error) {
	result = &v1alpha1.Naming{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("namings").
		Name(naming.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(naming).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the naming and deletes it. Returns an error if one occurs.
func (c *namings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("namings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *namings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("namings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched naming.
func (c *namings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Naming, err error) {
	result = &v1alpha1.Naming{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("namings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
