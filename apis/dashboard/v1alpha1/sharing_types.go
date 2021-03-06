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

type Sharing struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SharingSpec   `json:"spec,omitempty"`
	Status            SharingStatus `json:"status,omitempty"`
}

type SharingSpecPermissionsPermission struct {
	// The ID of the user or group to whom the permission is granted.
	//
	// Not applicable if the **type** is set to `ALL`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// The level of the permission:
	//
	// * `VIEW`: The dashboard is shared with read permission.
	// * `EDIT`: The dashboard is shared with edit permission
	Level *string `json:"level" tf:"level"`
	// The type of the permission:
	//
	// * `USER`: The dashboard is shared with the specified user.
	// * `GROUP`: The dashboard is shared with all users of the specified group.
	// * `ALL`: The dashboard is shared via link. Any authenticated user with the link can view the dashboard
	Type *string `json:"type" tf:"type"`
}

type SharingSpecPermissions struct {
	// Access permissions of the dashboard
	// +optional
	// +kubebuilder:validation:MinItems=1
	Permission []SharingSpecPermissionsPermission `json:"permission,omitempty" tf:"permission"`
}

type SharingSpecPublic struct {
	// A list of management zones that can display data on the publicly shared dashboard.
	//
	// Specify management zone IDs here. For each management zone you specify Dynatrace generates an access link. To share the dashboard with its default management zone, use the `default` value
	// +kubebuilder:validation:MinItems=1
	ManagementZones []string `json:"managementZones" tf:"management_zones"`
}

type SharingSpec struct {
	State *SharingSpecResource `json:"state,omitempty" tf:"-"`

	Resource SharingSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SharingSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The Dynatrace entity ID of the dashboard
	DashboardID *string `json:"dashboardID" tf:"dashboard_id"`
	// The dashboard is shared (`true`) or private (`false`)
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// Access permissions of the dashboard
	// +optional
	Permissions *SharingSpecPermissions `json:"permissions,omitempty" tf:"permissions"`
	// If `true` the dashboard will be marked as preset
	// +optional
	Preset *bool `json:"preset,omitempty" tf:"preset"`
	// Configuration of the [anonymous access](https://dt-url.net/ov03sf1) to the dashboard
	// +optional
	Public *SharingSpecPublic `json:"public,omitempty" tf:"public"`
}

type SharingStatus struct {
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

// SharingList is a list of Sharings
type SharingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Sharing CRD objects
	Items []Sharing `json:"items,omitempty"`
}
