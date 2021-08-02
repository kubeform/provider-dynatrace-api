// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credentials) DeepCopyInto(out *Credentials) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credentials.
func (in *Credentials) DeepCopy() *Credentials {
	if in == nil {
		return nil
	}
	out := new(Credentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Credentials) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsList) DeepCopyInto(out *CredentialsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Credentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsList.
func (in *CredentialsList) DeepCopy() *CredentialsList {
	if in == nil {
		return nil
	}
	out := new(CredentialsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CredentialsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsSpec) DeepCopyInto(out *CredentialsSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CredentialsSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsSpec.
func (in *CredentialsSpec) DeepCopy() *CredentialsSpec {
	if in == nil {
		return nil
	}
	out := new(CredentialsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsSpecEventsFieldSelectors) DeepCopyInto(out *CredentialsSpecEventsFieldSelectors) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.FieldSelector != nil {
		in, out := &in.FieldSelector, &out.FieldSelector
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Unknowns != nil {
		in, out := &in.Unknowns, &out.Unknowns
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsSpecEventsFieldSelectors.
func (in *CredentialsSpecEventsFieldSelectors) DeepCopy() *CredentialsSpecEventsFieldSelectors {
	if in == nil {
		return nil
	}
	out := new(CredentialsSpecEventsFieldSelectors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsSpecResource) DeepCopyInto(out *CredentialsSpecResource) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.AuthToken != nil {
		in, out := &in.AuthToken, &out.AuthToken
		*out = new(string)
		**out = **in
	}
	if in.CertificateCheckEnabled != nil {
		in, out := &in.CertificateCheckEnabled, &out.CertificateCheckEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DavisEventsIntegrationEnabled != nil {
		in, out := &in.DavisEventsIntegrationEnabled, &out.DavisEventsIntegrationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EndpointURL != nil {
		in, out := &in.EndpointURL, &out.EndpointURL
		*out = new(string)
		**out = **in
	}
	if in.EventsFieldSelectors != nil {
		in, out := &in.EventsFieldSelectors, &out.EventsFieldSelectors
		*out = make([]CredentialsSpecEventsFieldSelectors, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventsIntegrationEnabled != nil {
		in, out := &in.EventsIntegrationEnabled, &out.EventsIntegrationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HostnameVerification != nil {
		in, out := &in.HostnameVerification, &out.HostnameVerification
		*out = new(bool)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.PrometheusExporters != nil {
		in, out := &in.PrometheusExporters, &out.PrometheusExporters
		*out = new(bool)
		**out = **in
	}
	if in.Unknowns != nil {
		in, out := &in.Unknowns, &out.Unknowns
		*out = new(string)
		**out = **in
	}
	if in.WorkloadIntegrationEnabled != nil {
		in, out := &in.WorkloadIntegrationEnabled, &out.WorkloadIntegrationEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsSpecResource.
func (in *CredentialsSpecResource) DeepCopy() *CredentialsSpecResource {
	if in == nil {
		return nil
	}
	out := new(CredentialsSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsStatus) DeepCopyInto(out *CredentialsStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsStatus.
func (in *CredentialsStatus) DeepCopy() *CredentialsStatus {
	if in == nil {
		return nil
	}
	out := new(CredentialsStatus)
	in.DeepCopyInto(out)
	return out
}
