// +build !ignore_autogenerated

/*
Copyright 2021 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	v1alpha1 "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	upgradecattleiov1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventory) DeepCopyInto(out *MachineInventory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventory.
func (in *MachineInventory) DeepCopy() *MachineInventory {
	if in == nil {
		return nil
	}
	out := new(MachineInventory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineInventory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventoryList) DeepCopyInto(out *MachineInventoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineInventory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventoryList.
func (in *MachineInventoryList) DeepCopy() *MachineInventoryList {
	if in == nil {
		return nil
	}
	out := new(MachineInventoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineInventoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventorySpec) DeepCopyInto(out *MachineInventorySpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventorySpec.
func (in *MachineInventorySpec) DeepCopy() *MachineInventorySpec {
	if in == nil {
		return nil
	}
	out := new(MachineInventorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventoryStatus) DeepCopyInto(out *MachineInventoryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventoryStatus.
func (in *MachineInventoryStatus) DeepCopy() *MachineInventoryStatus {
	if in == nil {
		return nil
	}
	out := new(MachineInventoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRuntimeConfig) DeepCopyInto(out *MachineRuntimeConfig) {
	*out = *in
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]corev1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ConfigValues != nil {
		in, out := &in.ConfigValues, &out.ConfigValues
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRuntimeConfig.
func (in *MachineRuntimeConfig) DeepCopy() *MachineRuntimeConfig {
	if in == nil {
		return nil
	}
	out := new(MachineRuntimeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSImage) DeepCopyInto(out *ManagedOSImage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSImage.
func (in *ManagedOSImage) DeepCopy() *ManagedOSImage {
	if in == nil {
		return nil
	}
	out := new(ManagedOSImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedOSImage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSImageList) DeepCopyInto(out *ManagedOSImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedOSImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSImageList.
func (in *ManagedOSImageList) DeepCopy() *ManagedOSImageList {
	if in == nil {
		return nil
	}
	out := new(ManagedOSImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedOSImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSImageSpec) DeepCopyInto(out *ManagedOSImageSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Concurrency != nil {
		in, out := &in.Concurrency, &out.Concurrency
		*out = new(int64)
		**out = **in
	}
	if in.Prepare != nil {
		in, out := &in.Prepare, &out.Prepare
		*out = new(upgradecattleiov1.ContainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Cordon != nil {
		in, out := &in.Cordon, &out.Cordon
		*out = new(bool)
		**out = **in
	}
	if in.Drain != nil {
		in, out := &in.Drain, &out.Drain
		*out = new(upgradecattleiov1.DrainSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterRolloutStrategy != nil {
		in, out := &in.ClusterRolloutStrategy, &out.ClusterRolloutStrategy
		*out = new(v1alpha1.RolloutStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]v1alpha1.BundleTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSImageSpec.
func (in *ManagedOSImageSpec) DeepCopy() *ManagedOSImageSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedOSImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSImageStatus) DeepCopyInto(out *ManagedOSImageStatus) {
	*out = *in
	in.BundleStatus.DeepCopyInto(&out.BundleStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSImageStatus.
func (in *ManagedOSImageStatus) DeepCopy() *ManagedOSImageStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedOSImageStatus)
	in.DeepCopyInto(out)
	return out
}
