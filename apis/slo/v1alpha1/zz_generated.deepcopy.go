//go:build !ignore_autogenerated
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
func (in *Slo) DeepCopyInto(out *Slo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Slo.
func (in *Slo) DeepCopy() *Slo {
	if in == nil {
		return nil
	}
	out := new(Slo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Slo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SloList) DeepCopyInto(out *SloList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Slo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SloList.
func (in *SloList) DeepCopy() *SloList {
	if in == nil {
		return nil
	}
	out := new(SloList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SloList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SloSpec) DeepCopyInto(out *SloSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SloSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SloSpec.
func (in *SloSpec) DeepCopy() *SloSpec {
	if in == nil {
		return nil
	}
	out := new(SloSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SloSpecResource) DeepCopyInto(out *SloSpecResource) {
	*out = *in
	if in.Denominator != nil {
		in, out := &in.Denominator, &out.Denominator
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.Evaluation != nil {
		in, out := &in.Evaluation, &out.Evaluation
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Numerator != nil {
		in, out := &in.Numerator, &out.Numerator
		*out = new(string)
		**out = **in
	}
	if in.Rate != nil {
		in, out := &in.Rate, &out.Rate
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(float64)
		**out = **in
	}
	if in.Timeframe != nil {
		in, out := &in.Timeframe, &out.Timeframe
		*out = new(string)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SloSpecResource.
func (in *SloSpecResource) DeepCopy() *SloSpecResource {
	if in == nil {
		return nil
	}
	out := new(SloSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SloStatus) DeepCopyInto(out *SloStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SloStatus.
func (in *SloStatus) DeepCopy() *SloStatus {
	if in == nil {
		return nil
	}
	out := new(SloStatus)
	in.DeepCopyInto(out)
	return out
}
