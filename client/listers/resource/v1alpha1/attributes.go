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
	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/resource/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AttributesLister helps list Attributeses.
// All objects returned here must be treated as read-only.
type AttributesLister interface {
	// List lists all Attributeses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Attributes, err error)
	// Attributeses returns an object that can list and get Attributeses.
	Attributeses(namespace string) AttributesNamespaceLister
	AttributesListerExpansion
}

// attributesLister implements the AttributesLister interface.
type attributesLister struct {
	indexer cache.Indexer
}

// NewAttributesLister returns a new AttributesLister.
func NewAttributesLister(indexer cache.Indexer) AttributesLister {
	return &attributesLister{indexer: indexer}
}

// List lists all Attributeses in the indexer.
func (s *attributesLister) List(selector labels.Selector) (ret []*v1alpha1.Attributes, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Attributes))
	})
	return ret, err
}

// Attributeses returns an object that can list and get Attributeses.
func (s *attributesLister) Attributeses(namespace string) AttributesNamespaceLister {
	return attributesNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AttributesNamespaceLister helps list and get Attributeses.
// All objects returned here must be treated as read-only.
type AttributesNamespaceLister interface {
	// List lists all Attributeses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Attributes, err error)
	// Get retrieves the Attributes from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Attributes, error)
	AttributesNamespaceListerExpansion
}

// attributesNamespaceLister implements the AttributesNamespaceLister
// interface.
type attributesNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Attributeses in the indexer for a given namespace.
func (s attributesNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Attributes, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Attributes))
	})
	return ret, err
}

// Get retrieves the Attributes from the indexer for a given namespace and name.
func (s attributesNamespaceLister) Get(name string) (*v1alpha1.Attributes, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("attributes"), name)
	}
	return obj.(*v1alpha1.Attributes), nil
}
