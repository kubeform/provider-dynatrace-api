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

type EntryPoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EntryPointSpec   `json:"spec,omitempty"`
	Status            EntryPointStatus `json:"status,omitempty"`
}

type EntryPointSpecMatchesMatch struct {
	// Whether to match strings case sensitively or not
	// +optional
	CaseSensitive *bool `json:"caseSensitive,omitempty" tf:"case_sensitive"`
	// Possible values are `EQUALS`, `CONTAINS`, `STARTS_WITH`, `ENDS_WITH`, `DOES_NOT_EQUAL`, `DOES_NOT_CONTAIN`, `DOES_NOT_START_WITH` and `DOES_NOT_END_WITH`.
	Comparison *string `json:"comparison" tf:"comparison"`
	// The name of the attribute if `source` is `ATTRIBUTE`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// What to match against. Possible values are `SPAN_NAME`, `SPAN_KIND`, `ATTRIBUTE`, `INSTRUMENTATION_LIBRARY_NAME` and `INSTRUMENTATION_LIBRARY_VERSION`
	Source *string `json:"source" tf:"source"`
	// The value to compare against. When `source` is `SPAN_KIND` the only allowed values are `INTERNAL`, `SERVER`, `CLIENT`, `PRODUCER` and `CONSUMER`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type EntryPointSpecMatches struct {
	// Matching strategies for the Span
	Match []EntryPointSpecMatchesMatch `json:"match" tf:"match"`
}

type EntryPointSpec struct {
	State *EntryPointSpecResource `json:"state,omitempty" tf:"-"`

	Resource EntryPointSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EntryPointSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to create an entry point or not
	Action *string `json:"action" tf:"action"`
	// Matching strategies for the Span
	Matches *EntryPointSpecMatches `json:"matches" tf:"matches"`
	// The name of the rule
	Name *string `json:"name" tf:"name"`
}

type EntryPointStatus struct {
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

// EntryPointList is a list of EntryPoints
type EntryPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EntryPoint CRD objects
	Items []EntryPoint `json:"items,omitempty"`
}
