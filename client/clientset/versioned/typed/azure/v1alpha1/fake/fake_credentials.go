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

	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/azure/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCredentialses implements CredentialsInterface
type FakeCredentialses struct {
	Fake *FakeAzureV1alpha1
	ns   string
}

var credentialsesResource = schema.GroupVersionResource{Group: "azure.dynatrace.kubeform.com", Version: "v1alpha1", Resource: "credentialses"}

var credentialsesKind = schema.GroupVersionKind{Group: "azure.dynatrace.kubeform.com", Version: "v1alpha1", Kind: "Credentials"}

// Get takes name of the credentials, and returns the corresponding credentials object, and an error if there is any.
func (c *FakeCredentialses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Credentials, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(credentialsesResource, c.ns, name), &v1alpha1.Credentials{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credentials), err
}

// List takes label and field selectors, and returns the list of Credentialses that match those selectors.
func (c *FakeCredentialses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CredentialsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(credentialsesResource, credentialsesKind, c.ns, opts), &v1alpha1.CredentialsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CredentialsList{ListMeta: obj.(*v1alpha1.CredentialsList).ListMeta}
	for _, item := range obj.(*v1alpha1.CredentialsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested credentialses.
func (c *FakeCredentialses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(credentialsesResource, c.ns, opts))

}

// Create takes the representation of a credentials and creates it.  Returns the server's representation of the credentials, and an error, if there is any.
func (c *FakeCredentialses) Create(ctx context.Context, credentials *v1alpha1.Credentials, opts v1.CreateOptions) (result *v1alpha1.Credentials, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(credentialsesResource, c.ns, credentials), &v1alpha1.Credentials{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credentials), err
}

// Update takes the representation of a credentials and updates it. Returns the server's representation of the credentials, and an error, if there is any.
func (c *FakeCredentialses) Update(ctx context.Context, credentials *v1alpha1.Credentials, opts v1.UpdateOptions) (result *v1alpha1.Credentials, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(credentialsesResource, c.ns, credentials), &v1alpha1.Credentials{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credentials), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCredentialses) UpdateStatus(ctx context.Context, credentials *v1alpha1.Credentials, opts v1.UpdateOptions) (*v1alpha1.Credentials, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(credentialsesResource, "status", c.ns, credentials), &v1alpha1.Credentials{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credentials), err
}

// Delete takes name of the credentials and deletes it. Returns an error if one occurs.
func (c *FakeCredentialses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(credentialsesResource, c.ns, name), &v1alpha1.Credentials{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCredentialses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(credentialsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CredentialsList{})
	return err
}

// Patch applies the patch and returns the patched credentials.
func (c *FakeCredentialses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Credentials, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(credentialsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Credentials{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credentials), err
}
