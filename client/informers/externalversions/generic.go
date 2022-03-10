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
	"fmt"

	v1alpha1 "kubeform.dev/provider-dynatrace-api/apis/alerting/v1alpha1"
	applicationv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/application/v1alpha1"
	autotagv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/autotag/v1alpha1"
	awsv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/aws/v1alpha1"
	azurev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/azure/v1alpha1"
	browserv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/browser/v1alpha1"
	calculatedv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/calculated/v1alpha1"
	customv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/custom/v1alpha1"
	dashboardv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/dashboard/v1alpha1"
	databasev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/database/v1alpha1"
	diskv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/disk/v1alpha1"
	environmentv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/environment/v1alpha1"
	hostv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/host/v1alpha1"
	httpv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/http/v1alpha1"
	k8sv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/k8s/v1alpha1"
	keyv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/key/v1alpha1"
	maintenancev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/maintenance/v1alpha1"
	managementv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/management/v1alpha1"
	mobilev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/mobile/v1alpha1"
	notificationv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/notification/v1alpha1"
	processgroupv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/processgroup/v1alpha1"
	requestv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/request/v1alpha1"
	resourcev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/resource/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/service/v1alpha1"
	slov1alpha1 "kubeform.dev/provider-dynatrace-api/apis/slo/v1alpha1"
	spanv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/span/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/user/v1alpha1"
	webv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/web/v1alpha1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=alerting.dynatrace.kubeform.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("profiles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Alerting().V1alpha1().Profiles().Informer()}, nil

		// Group=application.dynatrace.kubeform.com, Version=v1alpha1
	case applicationv1alpha1.SchemeGroupVersion.WithResource("anomalieses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Application().V1alpha1().Anomalieses().Informer()}, nil
	case applicationv1alpha1.SchemeGroupVersion.WithResource("dataprivacies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Application().V1alpha1().DataPrivacies().Informer()}, nil
	case applicationv1alpha1.SchemeGroupVersion.WithResource("errorruleses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Application().V1alpha1().ErrorRuleses().Informer()}, nil

		// Group=autotag.dynatrace.kubeform.com, Version=v1alpha1
	case autotagv1alpha1.SchemeGroupVersion.WithResource("autotags"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Autotag().V1alpha1().Autotags().Informer()}, nil

		// Group=aws.dynatrace.kubeform.com, Version=v1alpha1
	case awsv1alpha1.SchemeGroupVersion.WithResource("credentialses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Aws().V1alpha1().Credentialses().Informer()}, nil

		// Group=azure.dynatrace.kubeform.com, Version=v1alpha1
	case azurev1alpha1.SchemeGroupVersion.WithResource("credentialses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Azure().V1alpha1().Credentialses().Informer()}, nil

		// Group=browser.dynatrace.kubeform.com, Version=v1alpha1
	case browserv1alpha1.SchemeGroupVersion.WithResource("monitors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Browser().V1alpha1().Monitors().Informer()}, nil

		// Group=calculated.dynatrace.kubeform.com, Version=v1alpha1
	case calculatedv1alpha1.SchemeGroupVersion.WithResource("servicemetrics"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Calculated().V1alpha1().ServiceMetrics().Informer()}, nil

		// Group=custom.dynatrace.kubeform.com, Version=v1alpha1
	case customv1alpha1.SchemeGroupVersion.WithResource("anomalieses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Custom().V1alpha1().Anomalieses().Informer()}, nil
	case customv1alpha1.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Custom().V1alpha1().Services().Informer()}, nil

		// Group=dashboard.dynatrace.kubeform.com, Version=v1alpha1
	case dashboardv1alpha1.SchemeGroupVersion.WithResource("dashboards"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dashboard().V1alpha1().Dashboards().Informer()}, nil
	case dashboardv1alpha1.SchemeGroupVersion.WithResource("sharings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dashboard().V1alpha1().Sharings().Informer()}, nil

		// Group=database.dynatrace.kubeform.com, Version=v1alpha1
	case databasev1alpha1.SchemeGroupVersion.WithResource("anomalieses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Database().V1alpha1().Anomalieses().Informer()}, nil

		// Group=disk.dynatrace.kubeform.com, Version=v1alpha1
	case diskv1alpha1.SchemeGroupVersion.WithResource("anomalieses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Disk().V1alpha1().Anomalieses().Informer()}, nil

		// Group=environment.dynatrace.kubeform.com, Version=v1alpha1
	case environmentv1alpha1.SchemeGroupVersion.WithResource("environments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Environment().V1alpha1().Environments().Informer()}, nil

		// Group=host.dynatrace.kubeform.com, Version=v1alpha1
	case hostv1alpha1.SchemeGroupVersion.WithResource("anomalieses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Host().V1alpha1().Anomalieses().Informer()}, nil
	case hostv1alpha1.SchemeGroupVersion.WithResource("namings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Host().V1alpha1().Namings().Informer()}, nil

		// Group=http.dynatrace.kubeform.com, Version=v1alpha1
	case httpv1alpha1.SchemeGroupVersion.WithResource("monitors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Http().V1alpha1().Monitors().Informer()}, nil

		// Group=k8s.dynatrace.kubeform.com, Version=v1alpha1
	case k8sv1alpha1.SchemeGroupVersion.WithResource("credentialses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.K8s().V1alpha1().Credentialses().Informer()}, nil

		// Group=key.dynatrace.kubeform.com, Version=v1alpha1
	case keyv1alpha1.SchemeGroupVersion.WithResource("requestses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Key().V1alpha1().Requestses().Informer()}, nil

		// Group=maintenance.dynatrace.kubeform.com, Version=v1alpha1
	case maintenancev1alpha1.SchemeGroupVersion.WithResource("windows"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Maintenance().V1alpha1().Windows().Informer()}, nil

		// Group=management.dynatrace.kubeform.com, Version=v1alpha1
	case managementv1alpha1.SchemeGroupVersion.WithResource("zones"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1alpha1().Zones().Informer()}, nil

		// Group=mobile.dynatrace.kubeform.com, Version=v1alpha1
	case mobilev1alpha1.SchemeGroupVersion.WithResource("applications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mobile().V1alpha1().Applications().Informer()}, nil

		// Group=notification.dynatrace.kubeform.com, Version=v1alpha1
	case notificationv1alpha1.SchemeGroupVersion.WithResource("notifications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Notification().V1alpha1().Notifications().Informer()}, nil

		// Group=processgroup.dynatrace.kubeform.com, Version=v1alpha1
	case processgroupv1alpha1.SchemeGroupVersion.WithResource("namings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Processgroup().V1alpha1().Namings().Informer()}, nil

		// Group=request.dynatrace.kubeform.com, Version=v1alpha1
	case requestv1alpha1.SchemeGroupVersion.WithResource("attributes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Request().V1alpha1().Attributes().Informer()}, nil
	case requestv1alpha1.SchemeGroupVersion.WithResource("namings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Request().V1alpha1().Namings().Informer()}, nil
	case requestv1alpha1.SchemeGroupVersion.WithResource("namingses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Request().V1alpha1().Namingses().Informer()}, nil

		// Group=resource.dynatrace.kubeform.com, Version=v1alpha1
	case resourcev1alpha1.SchemeGroupVersion.WithResource("attributeses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1alpha1().Attributeses().Informer()}, nil

		// Group=service.dynatrace.kubeform.com, Version=v1alpha1
	case servicev1alpha1.SchemeGroupVersion.WithResource("anomalieses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Service().V1alpha1().Anomalieses().Informer()}, nil
	case servicev1alpha1.SchemeGroupVersion.WithResource("namings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Service().V1alpha1().Namings().Informer()}, nil

		// Group=slo.dynatrace.kubeform.com, Version=v1alpha1
	case slov1alpha1.SchemeGroupVersion.WithResource("slos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Slo().V1alpha1().Slos().Informer()}, nil

		// Group=span.dynatrace.kubeform.com, Version=v1alpha1
	case spanv1alpha1.SchemeGroupVersion.WithResource("attributes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Span().V1alpha1().Attributes().Informer()}, nil
	case spanv1alpha1.SchemeGroupVersion.WithResource("capturerules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Span().V1alpha1().CaptureRules().Informer()}, nil
	case spanv1alpha1.SchemeGroupVersion.WithResource("contextpropagations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Span().V1alpha1().ContextPropagations().Informer()}, nil
	case spanv1alpha1.SchemeGroupVersion.WithResource("entrypoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Span().V1alpha1().EntryPoints().Informer()}, nil

		// Group=user.dynatrace.kubeform.com, Version=v1alpha1
	case userv1alpha1.SchemeGroupVersion.WithResource("groups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.User().V1alpha1().Groups().Informer()}, nil
	case userv1alpha1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.User().V1alpha1().Users().Informer()}, nil

		// Group=web.dynatrace.kubeform.com, Version=v1alpha1
	case webv1alpha1.SchemeGroupVersion.WithResource("applications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Web().V1alpha1().Applications().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
