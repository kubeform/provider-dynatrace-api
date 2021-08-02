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
func (in *Anomalies) DeepCopyInto(out *Anomalies) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Anomalies.
func (in *Anomalies) DeepCopy() *Anomalies {
	if in == nil {
		return nil
	}
	out := new(Anomalies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Anomalies) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomaliesList) DeepCopyInto(out *AnomaliesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Anomalies, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomaliesList.
func (in *AnomaliesList) DeepCopy() *AnomaliesList {
	if in == nil {
		return nil
	}
	out := new(AnomaliesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnomaliesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomaliesSpec) DeepCopyInto(out *AnomaliesSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AnomaliesSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomaliesSpec.
func (in *AnomaliesSpec) DeepCopy() *AnomaliesSpec {
	if in == nil {
		return nil
	}
	out := new(AnomaliesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomaliesSpecDiskName) DeepCopyInto(out *AnomaliesSpecDiskName) {
	*out = *in
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomaliesSpecDiskName.
func (in *AnomaliesSpecDiskName) DeepCopy() *AnomaliesSpecDiskName {
	if in == nil {
		return nil
	}
	out := new(AnomaliesSpecDiskName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomaliesSpecResource) DeepCopyInto(out *AnomaliesSpecResource) {
	*out = *in
	if in.DiskName != nil {
		in, out := &in.DiskName, &out.DiskName
		*out = new(AnomaliesSpecDiskName)
		(*in).DeepCopyInto(*out)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.HostGroupID != nil {
		in, out := &in.HostGroupID, &out.HostGroupID
		*out = new(string)
		**out = **in
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Samples != nil {
		in, out := &in.Samples, &out.Samples
		*out = new(int64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(AnomaliesSpecTags)
		(*in).DeepCopyInto(*out)
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ViolatingSamples != nil {
		in, out := &in.ViolatingSamples, &out.ViolatingSamples
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomaliesSpecResource.
func (in *AnomaliesSpecResource) DeepCopy() *AnomaliesSpecResource {
	if in == nil {
		return nil
	}
	out := new(AnomaliesSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomaliesSpecTags) DeepCopyInto(out *AnomaliesSpecTags) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]AnomaliesSpecTagsFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomaliesSpecTags.
func (in *AnomaliesSpecTags) DeepCopy() *AnomaliesSpecTags {
	if in == nil {
		return nil
	}
	out := new(AnomaliesSpecTags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomaliesSpecTagsFilter) DeepCopyInto(out *AnomaliesSpecTagsFilter) {
	*out = *in
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomaliesSpecTagsFilter.
func (in *AnomaliesSpecTagsFilter) DeepCopy() *AnomaliesSpecTagsFilter {
	if in == nil {
		return nil
	}
	out := new(AnomaliesSpecTagsFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomaliesStatus) DeepCopyInto(out *AnomaliesStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomaliesStatus.
func (in *AnomaliesStatus) DeepCopy() *AnomaliesStatus {
	if in == nil {
		return nil
	}
	out := new(AnomaliesStatus)
	in.DeepCopyInto(out)
	return out
}
