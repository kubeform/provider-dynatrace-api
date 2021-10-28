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
	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/host/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AnomaliesLister helps list Anomalieses.
// All objects returned here must be treated as read-only.
type AnomaliesLister interface {
	// List lists all Anomalieses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Anomalies, err error)
	// Anomalieses returns an object that can list and get Anomalieses.
	Anomalieses(namespace string) AnomaliesNamespaceLister
	AnomaliesListerExpansion
}

// anomaliesLister implements the AnomaliesLister interface.
type anomaliesLister struct {
	indexer cache.Indexer
}

// NewAnomaliesLister returns a new AnomaliesLister.
func NewAnomaliesLister(indexer cache.Indexer) AnomaliesLister {
	return &anomaliesLister{indexer: indexer}
}

// List lists all Anomalieses in the indexer.
func (s *anomaliesLister) List(selector labels.Selector) (ret []*v1alpha1.Anomalies, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Anomalies))
	})
	return ret, err
}

// Anomalieses returns an object that can list and get Anomalieses.
func (s *anomaliesLister) Anomalieses(namespace string) AnomaliesNamespaceLister {
	return anomaliesNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AnomaliesNamespaceLister helps list and get Anomalieses.
// All objects returned here must be treated as read-only.
type AnomaliesNamespaceLister interface {
	// List lists all Anomalieses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Anomalies, err error)
	// Get retrieves the Anomalies from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Anomalies, error)
	AnomaliesNamespaceListerExpansion
}

// anomaliesNamespaceLister implements the AnomaliesNamespaceLister
// interface.
type anomaliesNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Anomalieses in the indexer for a given namespace.
func (s anomaliesNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Anomalies, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Anomalies))
	})
	return ret, err
}

// Get retrieves the Anomalies from the indexer for a given namespace and name.
func (s anomaliesNamespaceLister) Get(name string) (*v1alpha1.Anomalies, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("anomalies"), name)
	}
	return obj.(*v1alpha1.Anomalies), nil
}