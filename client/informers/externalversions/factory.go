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

package externalversions

import (
	reflect "reflect"
	sync "sync"
	time "time"

	versioned "kubeform.dev/provider-dynatrace-api/client/clientset/versioned"
	alerting "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/alerting"
	application "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/application"
	autotag "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/autotag"
	aws "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/aws"
	azure "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/azure"
	browser "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/browser"
	calculated "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/calculated"
	custom "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/custom"
	dashboard "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/dashboard"
	database "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/database"
	disk "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/disk"
	environment "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/environment"
	host "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/host"
	http "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/http"
	internalinterfaces "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/internalinterfaces"
	k8s "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/k8s"
	key "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/key"
	maintenance "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/maintenance"
	management "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/management"
	mobile "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/mobile"
	notification "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/notification"
	processgroup "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/processgroup"
	request "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/request"
	resource "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/resource"
	service "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/service"
	slo "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/slo"
	span "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/span"
	user "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/user"
	web "kubeform.dev/provider-dynatrace-api/client/informers/externalversions/web"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*sharedInformerFactory) *sharedInformerFactory

type sharedInformerFactory struct {
	client           versioned.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock             sync.Mutex
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[v1.Object]time.Duration) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		for k, v := range resyncConfig {
			factory.customResync[reflect.TypeOf(k)] = v
		}
		return factory
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.tweakListOptions = tweakListOptions
		return factory
	}
}

// WithNamespace limits the SharedInformerFactory to the specified namespace.
func WithNamespace(namespace string) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.namespace = namespace
		return factory
	}
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client versioned.Interface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewFilteredSharedInformerFactory constructs a new instance of sharedInformerFactory.
// Listers obtained via this SharedInformerFactory will be subject to the same filters
// as specified here.
// Deprecated: Please use NewSharedInformerFactoryWithOptions instead
func NewFilteredSharedInformerFactory(client versioned.Interface, defaultResync time.Duration, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync, WithNamespace(namespace), WithTweakListOptions(tweakListOptions))
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client versioned.Interface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		namespace:        v1.NamespaceAll,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		factory = opt(factory)
	}

	return factory
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]cache.SharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]cache.SharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}

	resyncPeriod, exists := f.customResync[informerType]
	if !exists {
		resyncPeriod = f.defaultResync
	}

	informer = newFunc(f.client, resyncPeriod)
	f.informers[informerType] = informer

	return informer
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Alerting() alerting.Interface
	Application() application.Interface
	Autotag() autotag.Interface
	Aws() aws.Interface
	Azure() azure.Interface
	Browser() browser.Interface
	Calculated() calculated.Interface
	Custom() custom.Interface
	Dashboard() dashboard.Interface
	Database() database.Interface
	Disk() disk.Interface
	Environment() environment.Interface
	Host() host.Interface
	Http() http.Interface
	K8s() k8s.Interface
	Key() key.Interface
	Maintenance() maintenance.Interface
	Management() management.Interface
	Mobile() mobile.Interface
	Notification() notification.Interface
	Processgroup() processgroup.Interface
	Request() request.Interface
	Resource() resource.Interface
	Service() service.Interface
	Slo() slo.Interface
	Span() span.Interface
	User() user.Interface
	Web() web.Interface
}

func (f *sharedInformerFactory) Alerting() alerting.Interface {
	return alerting.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Application() application.Interface {
	return application.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Autotag() autotag.Interface {
	return autotag.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Aws() aws.Interface {
	return aws.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Azure() azure.Interface {
	return azure.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Browser() browser.Interface {
	return browser.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Calculated() calculated.Interface {
	return calculated.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Custom() custom.Interface {
	return custom.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dashboard() dashboard.Interface {
	return dashboard.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Database() database.Interface {
	return database.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Disk() disk.Interface {
	return disk.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Environment() environment.Interface {
	return environment.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Host() host.Interface {
	return host.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Http() http.Interface {
	return http.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) K8s() k8s.Interface {
	return k8s.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Key() key.Interface {
	return key.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Maintenance() maintenance.Interface {
	return maintenance.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Management() management.Interface {
	return management.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Mobile() mobile.Interface {
	return mobile.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Notification() notification.Interface {
	return notification.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Processgroup() processgroup.Interface {
	return processgroup.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Request() request.Interface {
	return request.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Resource() resource.Interface {
	return resource.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Service() service.Interface {
	return service.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Slo() slo.Interface {
	return slo.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Span() span.Interface {
	return span.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) User() user.Interface {
	return user.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Web() web.Interface {
	return web.New(f, f.namespace, f.tweakListOptions)
}
