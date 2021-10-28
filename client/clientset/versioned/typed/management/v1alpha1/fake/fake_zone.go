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

	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/management/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeZones implements ZoneInterface
type FakeZones struct {
	Fake *FakeManagementV1alpha1
	ns   string
}

var zonesResource = schema.GroupVersionResource{Group: "management.dynatrace.kubeform.com", Version: "v1alpha1", Resource: "zones"}

var zonesKind = schema.GroupVersionKind{Group: "management.dynatrace.kubeform.com", Version: "v1alpha1", Kind: "Zone"}

// Get takes name of the zone, and returns the corresponding zone object, and an error if there is any.
func (c *FakeZones) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Zone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(zonesResource, c.ns, name), &v1alpha1.Zone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Zone), err
}

// List takes label and field selectors, and returns the list of Zones that match those selectors.
func (c *FakeZones) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ZoneList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(zonesResource, zonesKind, c.ns, opts), &v1alpha1.ZoneList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ZoneList{ListMeta: obj.(*v1alpha1.ZoneList).ListMeta}
	for _, item := range obj.(*v1alpha1.ZoneList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested zones.
func (c *FakeZones) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(zonesResource, c.ns, opts))

}

// Create takes the representation of a zone and creates it.  Returns the server's representation of the zone, and an error, if there is any.
func (c *FakeZones) Create(ctx context.Context, zone *v1alpha1.Zone, opts v1.CreateOptions) (result *v1alpha1.Zone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(zonesResource, c.ns, zone), &v1alpha1.Zone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Zone), err
}

// Update takes the representation of a zone and updates it. Returns the server's representation of the zone, and an error, if there is any.
func (c *FakeZones) Update(ctx context.Context, zone *v1alpha1.Zone, opts v1.UpdateOptions) (result *v1alpha1.Zone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(zonesResource, c.ns, zone), &v1alpha1.Zone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Zone), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeZones) UpdateStatus(ctx context.Context, zone *v1alpha1.Zone, opts v1.UpdateOptions) (*v1alpha1.Zone, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(zonesResource, "status", c.ns, zone), &v1alpha1.Zone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Zone), err
}

// Delete takes name of the zone and deletes it. Returns an error if one occurs.
func (c *FakeZones) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(zonesResource, c.ns, name), &v1alpha1.Zone{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeZones) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(zonesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ZoneList{})
	return err
}

// Patch applies the patch and returns the patched zone.
func (c *FakeZones) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Zone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(zonesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Zone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Zone), err
}