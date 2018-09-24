/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package quota

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedClusterResourceQuota) DeepCopyInto(out *AppliedClusterResourceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedClusterResourceQuota.
func (in *AppliedClusterResourceQuota) DeepCopy() *AppliedClusterResourceQuota {
	if in == nil {
		return nil
	}
	out := new(AppliedClusterResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedClusterResourceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedClusterResourceQuotaList) DeepCopyInto(out *AppliedClusterResourceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppliedClusterResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedClusterResourceQuotaList.
func (in *AppliedClusterResourceQuotaList) DeepCopy() *AppliedClusterResourceQuotaList {
	if in == nil {
		return nil
	}
	out := new(AppliedClusterResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedClusterResourceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuota) DeepCopyInto(out *ClusterResourceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuota.
func (in *ClusterResourceQuota) DeepCopy() *ClusterResourceQuota {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaList) DeepCopyInto(out *ClusterResourceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaList.
func (in *ClusterResourceQuotaList) DeepCopy() *ClusterResourceQuotaList {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaSelector) DeepCopyInto(out *ClusterResourceQuotaSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AnnotationSelector != nil {
		in, out := &in.AnnotationSelector, &out.AnnotationSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaSelector.
func (in *ClusterResourceQuotaSelector) DeepCopy() *ClusterResourceQuotaSelector {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaSpec) DeepCopyInto(out *ClusterResourceQuotaSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Quota.DeepCopyInto(&out.Quota)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaSpec.
func (in *ClusterResourceQuotaSpec) DeepCopy() *ClusterResourceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaStatus) DeepCopyInto(out *ClusterResourceQuotaStatus) {
	*out = *in
	in.Total.DeepCopyInto(&out.Total)
	out.Namespaces = in.Namespaces.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaStatus.
func (in *ClusterResourceQuotaStatus) DeepCopy() *ClusterResourceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotasStatusByNamespace) DeepCopyInto(out *ResourceQuotasStatusByNamespace) {
	*out = in.DeepCopy()
	return
}