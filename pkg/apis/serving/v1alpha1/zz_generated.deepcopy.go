// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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
	apis "knative.dev/pkg/apis"
	v1 "knative.dev/pkg/apis/duck/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CannotConvertError) DeepCopyInto(out *CannotConvertError) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CannotConvertError.
func (in *CannotConvertError) DeepCopy() *CannotConvertError {
	if in == nil {
		return nil
	}
	out := new(CannotConvertError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainMapping) DeepCopyInto(out *DomainMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainMapping.
func (in *DomainMapping) DeepCopy() *DomainMapping {
	if in == nil {
		return nil
	}
	out := new(DomainMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainMappingList) DeepCopyInto(out *DomainMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DomainMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainMappingList.
func (in *DomainMappingList) DeepCopy() *DomainMappingList {
	if in == nil {
		return nil
	}
	out := new(DomainMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainMappingSpec) DeepCopyInto(out *DomainMappingSpec) {
	*out = *in
	out.Ref = in.Ref
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainMappingSpec.
func (in *DomainMappingSpec) DeepCopy() *DomainMappingSpec {
	if in == nil {
		return nil
	}
	out := new(DomainMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainMappingStatus) DeepCopyInto(out *DomainMappingStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(v1.Addressable)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainMappingStatus.
func (in *DomainMappingStatus) DeepCopy() *DomainMappingStatus {
	if in == nil {
		return nil
	}
	out := new(DomainMappingStatus)
	in.DeepCopyInto(out)
	return out
}