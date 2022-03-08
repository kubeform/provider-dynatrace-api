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

type Application struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec,omitempty"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

type ApplicationSpecApdex struct {
	// Apdex **frustrated** threshold, in milliseconds: a duration greater than or equal to this value is considered frustrated
	Frustrated *int64 `json:"frustrated" tf:"frustrated"`
	// Apdex error condition: if `true` the user session is considered frustrated when an error is reported
	// +optional
	FrustratedOnError *bool `json:"frustratedOnError,omitempty" tf:"frustrated_on_error"`
	// Apdex **tolerable** threshold, in milliseconds: a duration greater than or equal to this value is considered tolerable
	Tolerable *int64 `json:"tolerable" tf:"tolerable"`
}

type ApplicationSpecPropertiesApiValue struct {
	// The aggregation type of the property. It defines how multiple values of the property are aggregated. Possible values are `SUM`, `MIN`, `MAX`, `FIRST` and `LAST`
	// +optional
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation"`
	// The cleanup rule of the property. Defines how to extract the data you need from a string value. Specify the [regular expression](https://dt-url.net/k9e0iaq) for the data you need there
	// +optional
	CleanupRule *string `json:"cleanupRule,omitempty" tf:"cleanup_rule"`
	// The display name of the property
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// The unique key of the mobile session or user action property
	Key *string `json:"key" tf:"key"`
	// The name of the reported value
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// If `true`, the property is stored as a session property
	// +optional
	StoreAsSessionProperty *bool `json:"storeAsSessionProperty,omitempty" tf:"store_as_session_property"`
	// If `true`, the property is stored as a user action property
	// +optional
	StoreAsUserActionProperty *bool `json:"storeAsUserActionProperty,omitempty" tf:"store_as_user_action_property"`
	// The data type of the property. Possible values are `DOUBLE`, `LONG` and `STRING`
	Type *string `json:"type" tf:"type"`
}

type ApplicationSpecPropertiesRequestAttribute struct {
	// The aggregation type of the property. It defines how multiple values of the property are aggregated. Possible values are `SUM`, `MIN`, `MAX`, `FIRST` and `LAST`
	// +optional
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation"`
	// The cleanup rule of the property. Defines how to extract the data you need from a string value. Specify the [regular expression](https://dt-url.net/k9e0iaq) for the data you need there
	// +optional
	CleanupRule *string `json:"cleanupRule,omitempty" tf:"cleanup_rule"`
	// The display name of the property
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// The ID of the request attribute
	ID *string `json:"ID" tf:"id"`
	// The unique key of the mobile session or user action property
	Key *string `json:"key" tf:"key"`
	// If `true`, the property is stored as a session property
	// +optional
	StoreAsSessionProperty *bool `json:"storeAsSessionProperty,omitempty" tf:"store_as_session_property"`
	// If `true`, the property is stored as a user action property
	// +optional
	StoreAsUserActionProperty *bool `json:"storeAsUserActionProperty,omitempty" tf:"store_as_user_action_property"`
	// The data type of the property. Possible values are `DOUBLE`, `LONG` and `STRING`. The value MUST match the data type of the Request Attribute.
	Type *string `json:"type" tf:"type"`
}

type ApplicationSpecProperties struct {
	// A User Action / Session Property based on a value reported by the API
	// +optional
	ApiValue []ApplicationSpecPropertiesApiValue `json:"apiValue,omitempty" tf:"api_value"`
	// A User Action / Session Property based on a Server Side Request Attribute
	// +optional
	RequestAttribute []ApplicationSpecPropertiesRequestAttribute `json:"requestAttribute,omitempty" tf:"request_attribute"`
}

type ApplicationSpec struct {
	State *ApplicationSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApplicationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ApplicationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Apdex configuration of a mobile application. A duration less than the **tolerable** threshold is considered satisfied
	Apdex *ApplicationSpecApdex `json:"apdex" tf:"apdex"`
	// The UUID of the application.
	//
	// It is used only by OneAgent to send data to Dynatrace. If not specified it will get generated.
	// +optional
	ApplicationID *string `json:"applicationID,omitempty" tf:"application_id"`
	// The type of the beacon endpoint. Possible values are `CLUSTER_ACTIVE_GATE`, `ENVIRONMENT_ACTIVE_GATE` and `INSTRUMENTED_WEB_SERVER`.
	BeaconEndpointType *string `json:"beaconEndpointType" tf:"beacon_endpoint_type"`
	// The URL of the beacon endpoint.
	//
	// Only applicable when the **beacon_endpoint_type** is set to `ENVIRONMENT_ACTIVE_GATE` or `INSTRUMENTED_WEB_SERVER`
	// +optional
	BeaconEndpointURL *string `json:"beaconEndpointURL,omitempty" tf:"beacon_endpoint_url"`
	// User Action names to be flagged as Key User Actions
	// +optional
	KeyUserActions []string `json:"keyUserActions,omitempty" tf:"key_user_actions"`
	// The name of the application
	Name *string `json:"name" tf:"name"`
	// The opt-in mode is enabled (`true`) or disabled (`false`)
	// +optional
	OptInMode *bool `json:"optInMode,omitempty" tf:"opt_in_mode"`
	// User Action and Session Properties
	// +optional
	Properties *ApplicationSpecProperties `json:"properties,omitempty" tf:"properties"`
	// The session replay is enabled (`true`) or disabled (`false`).
	// +optional
	SessionReplay *bool `json:"sessionReplay,omitempty" tf:"session_replay"`
	// The session replay on crash is enabled (`true`) or disabled (`false`).
	//
	// Enabling requires both **sessionReplayEnabled** and **optInModeEnabled** values set to `true`.
	// +optional
	SessionReplayOnCrash *bool `json:"sessionReplayOnCrash,omitempty" tf:"session_replay_on_crash"`
	// The percentage of user sessions to be analyzed
	// +optional
	UserSessionPercentage *int64 `json:"userSessionPercentage,omitempty" tf:"user_session_percentage"`
}

type ApplicationStatus struct {
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

// ApplicationList is a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Application CRD objects
	Items []Application `json:"items,omitempty"`
}
