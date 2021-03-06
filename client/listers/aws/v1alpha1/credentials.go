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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CredentialsLister helps list Credentialses.
// All objects returned here must be treated as read-only.
type CredentialsLister interface {
	// List lists all Credentialses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Credentials, err error)
	// Credentialses returns an object that can list and get Credentialses.
	Credentialses(namespace string) CredentialsNamespaceLister
	CredentialsListerExpansion
}

// credentialsLister implements the CredentialsLister interface.
type credentialsLister struct {
	indexer cache.Indexer
}

// NewCredentialsLister returns a new CredentialsLister.
func NewCredentialsLister(indexer cache.Indexer) CredentialsLister {
	return &credentialsLister{indexer: indexer}
}

// List lists all Credentialses in the indexer.
func (s *credentialsLister) List(selector labels.Selector) (ret []*v1alpha1.Credentials, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Credentials))
	})
	return ret, err
}

// Credentialses returns an object that can list and get Credentialses.
func (s *credentialsLister) Credentialses(namespace string) CredentialsNamespaceLister {
	return credentialsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CredentialsNamespaceLister helps list and get Credentialses.
// All objects returned here must be treated as read-only.
type CredentialsNamespaceLister interface {
	// List lists all Credentialses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Credentials, err error)
	// Get retrieves the Credentials from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Credentials, error)
	CredentialsNamespaceListerExpansion
}

// credentialsNamespaceLister implements the CredentialsNamespaceLister
// interface.
type credentialsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Credentialses in the indexer for a given namespace.
func (s credentialsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Credentials, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Credentials))
	})
	return ret, err
}

// Get retrieves the Credentials from the indexer for a given namespace and name.
func (s credentialsNamespaceLister) Get(name string) (*v1alpha1.Credentials, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("credentials"), name)
	}
	return obj.(*v1alpha1.Credentials), nil
}
