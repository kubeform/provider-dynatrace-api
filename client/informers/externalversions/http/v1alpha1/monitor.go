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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	httpv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/http/v1alpha1"
	versioned "kubeform.dev/provider-dynatrace-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-dynatrace-api/client/listers/http/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MonitorInformer provides access to a shared informer and lister for
// Monitors.
type MonitorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MonitorLister
}

type monitorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMonitorInformer constructs a new informer for Monitor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMonitorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMonitorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMonitorInformer constructs a new informer for Monitor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMonitorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HttpV1alpha1().Monitors(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HttpV1alpha1().Monitors(namespace).Watch(context.TODO(), options)
			},
		},
		&httpv1alpha1.Monitor{},
		resyncPeriod,
		indexers,
	)
}

func (f *monitorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMonitorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *monitorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&httpv1alpha1.Monitor{}, f.defaultInformer)
}

func (f *monitorInformer) Lister() v1alpha1.MonitorLister {
	return v1alpha1.NewMonitorLister(f.Informer().GetIndexer())
}