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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Credentials struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CredentialsSpec   `json:"spec,omitempty"`
	Status            CredentialsStatus `json:"status,omitempty"`
}

type CredentialsSpecEventsFieldSelectors struct {
	// Whether subscription to this events field selector is enabled (value set to `true`). If disabled (value set to `false`), Dynatrace will stop fetching events from the Kubernetes API for this events field selector
	Active *bool `json:"active" tf:"active"`
	// The field selector string (url decoding is applied) when storing it.
	FieldSelector *string `json:"fieldSelector" tf:"field_selector"`
	// A label of the events field selector.
	Label *string `json:"label" tf:"label"`
	// Any attributes that aren't yet supported by this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
}

type CredentialsSpec struct {
	State *CredentialsSpecResource `json:"state,omitempty" tf:"-"`

	Resource CredentialsSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CredentialsSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Monitoring is enabled (`true`) or disabled (`false`) for given credentials configuration.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	// +optional
	Active *bool `json:"active,omitempty" tf:"active"`
	// The service account bearer token for the Kubernetes API server.  Submit your token on creation or update of the configuration. For security reasons, GET requests return this field as `null`.  If the field is omitted during an update, the old value remains unaffected.
	// +optional
	AuthToken *string `json:"authToken,omitempty" tf:"auth_token"`
	// The check of SSL certificates is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	// +optional
	CertificateCheckEnabled *bool `json:"certificateCheckEnabled,omitempty" tf:"certificate_check_enabled"`
	// Inclusion of all Davis relevant events is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected
	// +optional
	DavisEventsIntegrationEnabled *bool `json:"davisEventsIntegrationEnabled,omitempty" tf:"davis_events_integration_enabled"`
	// The URL of the Kubernetes API server.  It must be unique within a Dynatrace environment.  The URL must valid according to RFC 2396. Leading or trailing whitespaces are not allowed.
	// +optional
	EndpointURL *string `json:"endpointURL,omitempty" tf:"endpoint_url"`
	// The check of SSL certificates is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	// +optional
	EventsFieldSelectors []CredentialsSpecEventsFieldSelectors `json:"eventsFieldSelectors,omitempty" tf:"events_field_selectors"`
	// Monitoring of events is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. Event monitoring depends on the active state of this configuration to be true.  If not set on creation, the `false` value is used.  If the field is omitted during an update, the old value remains unaffected.
	// +optional
	EventsIntegrationEnabled *bool `json:"eventsIntegrationEnabled,omitempty" tf:"events_integration_enabled"`
	// Verify hostname in certificate against Kubernetes API URL
	// +optional
	HostnameVerification *bool `json:"hostnameVerification,omitempty" tf:"hostname_verification"`
	// The name of the Kubernetes credentials configuration.  Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`. Leading or trailing whitespace is not allowed.
	Label *string `json:"label" tf:"label"`
	// Prometheus exporters integration is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.If the field is omitted during an update, the old value remains unaffected
	// +optional
	PrometheusExporters *bool `json:"prometheusExporters,omitempty" tf:"prometheus_exporters"`
	// Any attributes that aren't yet supported by this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
	// Workload and cloud application processing is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected.
	// +optional
	WorkloadIntegrationEnabled *bool `json:"workloadIntegrationEnabled,omitempty" tf:"workload_integration_enabled"`
}

type CredentialsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CredentialsList is a list of Credentialss
type CredentialsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Credentials CRD objects
	Items []Credentials `json:"items,omitempty"`
}
