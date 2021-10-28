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

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/span/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeContextPropagations implements ContextPropagationInterface
type FakeContextPropagations struct {
	Fake *FakeSpanV1alpha1
	ns   string
}

var contextpropagationsResource = schema.GroupVersionResource{Group: "span.dynatrace.kubeform.com", Version: "v1alpha1", Resource: "contextpropagations"}

var contextpropagationsKind = schema.GroupVersionKind{Group: "span.dynatrace.kubeform.com", Version: "v1alpha1", Kind: "ContextPropagation"}

// Get takes name of the contextPropagation, and returns the corresponding contextPropagation object, and an error if there is any.
func (c *FakeContextPropagations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ContextPropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(contextpropagationsResource, c.ns, name), &v1alpha1.ContextPropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContextPropagation), err
}

// List takes label and field selectors, and returns the list of ContextPropagations that match those selectors.
func (c *FakeContextPropagations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ContextPropagationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(contextpropagationsResource, contextpropagationsKind, c.ns, opts), &v1alpha1.ContextPropagationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ContextPropagationList{ListMeta: obj.(*v1alpha1.ContextPropagationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ContextPropagationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested contextPropagations.
func (c *FakeContextPropagations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(contextpropagationsResource, c.ns, opts))

}

// Create takes the representation of a contextPropagation and creates it.  Returns the server's representation of the contextPropagation, and an error, if there is any.
func (c *FakeContextPropagations) Create(ctx context.Context, contextPropagation *v1alpha1.ContextPropagation, opts v1.CreateOptions) (result *v1alpha1.ContextPropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(contextpropagationsResource, c.ns, contextPropagation), &v1alpha1.ContextPropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContextPropagation), err
}

// Update takes the representation of a contextPropagation and updates it. Returns the server's representation of the contextPropagation, and an error, if there is any.
func (c *FakeContextPropagations) Update(ctx context.Context, contextPropagation *v1alpha1.ContextPropagation, opts v1.UpdateOptions) (result *v1alpha1.ContextPropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(contextpropagationsResource, c.ns, contextPropagation), &v1alpha1.ContextPropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContextPropagation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContextPropagations) UpdateStatus(ctx context.Context, contextPropagation *v1alpha1.ContextPropagation, opts v1.UpdateOptions) (*v1alpha1.ContextPropagation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(contextpropagationsResource, "status", c.ns, contextPropagation), &v1alpha1.ContextPropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContextPropagation), err
}

// Delete takes name of the contextPropagation and deletes it. Returns an error if one occurs.
func (c *FakeContextPropagations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(contextpropagationsResource, c.ns, name), &v1alpha1.ContextPropagation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContextPropagations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(contextpropagationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ContextPropagationList{})
	return err
}

// Patch applies the patch and returns the patched contextPropagation.
func (c *FakeContextPropagations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ContextPropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(contextpropagationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ContextPropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContextPropagation), err
}