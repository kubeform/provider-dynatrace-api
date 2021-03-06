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
	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/autotag/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AutotagLister helps list Autotags.
// All objects returned here must be treated as read-only.
type AutotagLister interface {
	// List lists all Autotags in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Autotag, err error)
	// Autotags returns an object that can list and get Autotags.
	Autotags(namespace string) AutotagNamespaceLister
	AutotagListerExpansion
}

// autotagLister implements the AutotagLister interface.
type autotagLister struct {
	indexer cache.Indexer
}

// NewAutotagLister returns a new AutotagLister.
func NewAutotagLister(indexer cache.Indexer) AutotagLister {
	return &autotagLister{indexer: indexer}
}

// List lists all Autotags in the indexer.
func (s *autotagLister) List(selector labels.Selector) (ret []*v1alpha1.Autotag, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Autotag))
	})
	return ret, err
}

// Autotags returns an object that can list and get Autotags.
func (s *autotagLister) Autotags(namespace string) AutotagNamespaceLister {
	return autotagNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AutotagNamespaceLister helps list and get Autotags.
// All objects returned here must be treated as read-only.
type AutotagNamespaceLister interface {
	// List lists all Autotags in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Autotag, err error)
	// Get retrieves the Autotag from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Autotag, error)
	AutotagNamespaceListerExpansion
}

// autotagNamespaceLister implements the AutotagNamespaceLister
// interface.
type autotagNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Autotags in the indexer for a given namespace.
func (s autotagNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Autotag, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Autotag))
	})
	return ret, err
}

// Get retrieves the Autotag from the indexer for a given namespace and name.
func (s autotagNamespaceLister) Get(name string) (*v1alpha1.Autotag, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("autotag"), name)
	}
	return obj.(*v1alpha1.Autotag), nil
}
