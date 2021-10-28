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

	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/dashboard/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSharings implements SharingInterface
type FakeSharings struct {
	Fake *FakeDashboardV1alpha1
	ns   string
}

var sharingsResource = schema.GroupVersionResource{Group: "dashboard.dynatrace.kubeform.com", Version: "v1alpha1", Resource: "sharings"}

var sharingsKind = schema.GroupVersionKind{Group: "dashboard.dynatrace.kubeform.com", Version: "v1alpha1", Kind: "Sharing"}

// Get takes name of the sharing, and returns the corresponding sharing object, and an error if there is any.
func (c *FakeSharings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Sharing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharingsResource, c.ns, name), &v1alpha1.Sharing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sharing), err
}

// List takes label and field selectors, and returns the list of Sharings that match those selectors.
func (c *FakeSharings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SharingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharingsResource, sharingsKind, c.ns, opts), &v1alpha1.SharingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SharingList{ListMeta: obj.(*v1alpha1.SharingList).ListMeta}
	for _, item := range obj.(*v1alpha1.SharingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sharings.
func (c *FakeSharings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharingsResource, c.ns, opts))

}

// Create takes the representation of a sharing and creates it.  Returns the server's representation of the sharing, and an error, if there is any.
func (c *FakeSharings) Create(ctx context.Context, sharing *v1alpha1.Sharing, opts v1.CreateOptions) (result *v1alpha1.Sharing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharingsResource, c.ns, sharing), &v1alpha1.Sharing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sharing), err
}

// Update takes the representation of a sharing and updates it. Returns the server's representation of the sharing, and an error, if there is any.
func (c *FakeSharings) Update(ctx context.Context, sharing *v1alpha1.Sharing, opts v1.UpdateOptions) (result *v1alpha1.Sharing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharingsResource, c.ns, sharing), &v1alpha1.Sharing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sharing), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSharings) UpdateStatus(ctx context.Context, sharing *v1alpha1.Sharing, opts v1.UpdateOptions) (*v1alpha1.Sharing, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sharingsResource, "status", c.ns, sharing), &v1alpha1.Sharing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sharing), err
}

// Delete takes name of the sharing and deletes it. Returns an error if one occurs.
func (c *FakeSharings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sharingsResource, c.ns, name), &v1alpha1.Sharing{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSharings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sharingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SharingList{})
	return err
}

// Patch applies the patch and returns the patched sharing.
func (c *FakeSharings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Sharing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Sharing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sharing), err
}
