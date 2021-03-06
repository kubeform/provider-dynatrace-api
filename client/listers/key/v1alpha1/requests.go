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
	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/key/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RequestsLister helps list Requestses.
// All objects returned here must be treated as read-only.
type RequestsLister interface {
	// List lists all Requestses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Requests, err error)
	// Requestses returns an object that can list and get Requestses.
	Requestses(namespace string) RequestsNamespaceLister
	RequestsListerExpansion
}

// requestsLister implements the RequestsLister interface.
type requestsLister struct {
	indexer cache.Indexer
}

// NewRequestsLister returns a new RequestsLister.
func NewRequestsLister(indexer cache.Indexer) RequestsLister {
	return &requestsLister{indexer: indexer}
}

// List lists all Requestses in the indexer.
func (s *requestsLister) List(selector labels.Selector) (ret []*v1alpha1.Requests, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Requests))
	})
	return ret, err
}

// Requestses returns an object that can list and get Requestses.
func (s *requestsLister) Requestses(namespace string) RequestsNamespaceLister {
	return requestsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RequestsNamespaceLister helps list and get Requestses.
// All objects returned here must be treated as read-only.
type RequestsNamespaceLister interface {
	// List lists all Requestses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Requests, err error)
	// Get retrieves the Requests from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Requests, error)
	RequestsNamespaceListerExpansion
}

// requestsNamespaceLister implements the RequestsNamespaceLister
// interface.
type requestsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Requestses in the indexer for a given namespace.
func (s requestsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Requests, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Requests))
	})
	return ret, err
}

// Get retrieves the Requests from the indexer for a given namespace and name.
func (s requestsNamespaceLister) Get(name string) (*v1alpha1.Requests, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("requests"), name)
	}
	return obj.(*v1alpha1.Requests), nil
}
