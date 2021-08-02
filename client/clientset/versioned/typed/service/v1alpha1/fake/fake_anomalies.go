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

	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/service/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAnomalieses implements AnomaliesInterface
type FakeAnomalieses struct {
	Fake *FakeServiceV1alpha1
	ns   string
}

var anomaliesesResource = schema.GroupVersionResource{Group: "service.dynatrace.kubeform.com", Version: "v1alpha1", Resource: "anomalieses"}

var anomaliesesKind = schema.GroupVersionKind{Group: "service.dynatrace.kubeform.com", Version: "v1alpha1", Kind: "Anomalies"}

// Get takes name of the anomalies, and returns the corresponding anomalies object, and an error if there is any.
func (c *FakeAnomalieses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Anomalies, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(anomaliesesResource, c.ns, name), &v1alpha1.Anomalies{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Anomalies), err
}

// List takes label and field selectors, and returns the list of Anomalieses that match those selectors.
func (c *FakeAnomalieses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AnomaliesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(anomaliesesResource, anomaliesesKind, c.ns, opts), &v1alpha1.AnomaliesList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AnomaliesList{ListMeta: obj.(*v1alpha1.AnomaliesList).ListMeta}
	for _, item := range obj.(*v1alpha1.AnomaliesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested anomalieses.
func (c *FakeAnomalieses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(anomaliesesResource, c.ns, opts))

}

// Create takes the representation of a anomalies and creates it.  Returns the server's representation of the anomalies, and an error, if there is any.
func (c *FakeAnomalieses) Create(ctx context.Context, anomalies *v1alpha1.Anomalies, opts v1.CreateOptions) (result *v1alpha1.Anomalies, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(anomaliesesResource, c.ns, anomalies), &v1alpha1.Anomalies{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Anomalies), err
}

// Update takes the representation of a anomalies and updates it. Returns the server's representation of the anomalies, and an error, if there is any.
func (c *FakeAnomalieses) Update(ctx context.Context, anomalies *v1alpha1.Anomalies, opts v1.UpdateOptions) (result *v1alpha1.Anomalies, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(anomaliesesResource, c.ns, anomalies), &v1alpha1.Anomalies{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Anomalies), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAnomalieses) UpdateStatus(ctx context.Context, anomalies *v1alpha1.Anomalies, opts v1.UpdateOptions) (*v1alpha1.Anomalies, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(anomaliesesResource, "status", c.ns, anomalies), &v1alpha1.Anomalies{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Anomalies), err
}

// Delete takes name of the anomalies and deletes it. Returns an error if one occurs.
func (c *FakeAnomalieses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(anomaliesesResource, c.ns, name), &v1alpha1.Anomalies{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAnomalieses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(anomaliesesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AnomaliesList{})
	return err
}

// Patch applies the patch and returns the patched anomalies.
func (c *FakeAnomalieses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Anomalies, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(anomaliesesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Anomalies{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Anomalies), err
}
