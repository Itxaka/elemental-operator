/*
Copyright 2022 SUSE LLC

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

package fake

import (
	"context"

	rancheroscattleiov1 "github.com/rancher-sandbox/rancheros-operator/pkg/apis/rancheros.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagedOSVersionChannels implements ManagedOSVersionChannelInterface
type FakeManagedOSVersionChannels struct {
	Fake *FakeRancherosV1
	ns   string
}

var managedosversionchannelsResource = schema.GroupVersionResource{Group: "rancheros.cattle.io", Version: "v1", Resource: "managedosversionchannels"}

var managedosversionchannelsKind = schema.GroupVersionKind{Group: "rancheros.cattle.io", Version: "v1", Kind: "ManagedOSVersionChannel"}

// Get takes name of the managedOSVersionChannel, and returns the corresponding managedOSVersionChannel object, and an error if there is any.
func (c *FakeManagedOSVersionChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *rancheroscattleiov1.ManagedOSVersionChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedosversionchannelsResource, c.ns, name), &rancheroscattleiov1.ManagedOSVersionChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rancheroscattleiov1.ManagedOSVersionChannel), err
}

// List takes label and field selectors, and returns the list of ManagedOSVersionChannels that match those selectors.
func (c *FakeManagedOSVersionChannels) List(ctx context.Context, opts v1.ListOptions) (result *rancheroscattleiov1.ManagedOSVersionChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedosversionchannelsResource, managedosversionchannelsKind, c.ns, opts), &rancheroscattleiov1.ManagedOSVersionChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rancheroscattleiov1.ManagedOSVersionChannelList{ListMeta: obj.(*rancheroscattleiov1.ManagedOSVersionChannelList).ListMeta}
	for _, item := range obj.(*rancheroscattleiov1.ManagedOSVersionChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedOSVersionChannels.
func (c *FakeManagedOSVersionChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedosversionchannelsResource, c.ns, opts))

}

// Create takes the representation of a managedOSVersionChannel and creates it.  Returns the server's representation of the managedOSVersionChannel, and an error, if there is any.
func (c *FakeManagedOSVersionChannels) Create(ctx context.Context, managedOSVersionChannel *rancheroscattleiov1.ManagedOSVersionChannel, opts v1.CreateOptions) (result *rancheroscattleiov1.ManagedOSVersionChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedosversionchannelsResource, c.ns, managedOSVersionChannel), &rancheroscattleiov1.ManagedOSVersionChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rancheroscattleiov1.ManagedOSVersionChannel), err
}

// Update takes the representation of a managedOSVersionChannel and updates it. Returns the server's representation of the managedOSVersionChannel, and an error, if there is any.
func (c *FakeManagedOSVersionChannels) Update(ctx context.Context, managedOSVersionChannel *rancheroscattleiov1.ManagedOSVersionChannel, opts v1.UpdateOptions) (result *rancheroscattleiov1.ManagedOSVersionChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedosversionchannelsResource, c.ns, managedOSVersionChannel), &rancheroscattleiov1.ManagedOSVersionChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rancheroscattleiov1.ManagedOSVersionChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedOSVersionChannels) UpdateStatus(ctx context.Context, managedOSVersionChannel *rancheroscattleiov1.ManagedOSVersionChannel, opts v1.UpdateOptions) (*rancheroscattleiov1.ManagedOSVersionChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedosversionchannelsResource, "status", c.ns, managedOSVersionChannel), &rancheroscattleiov1.ManagedOSVersionChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rancheroscattleiov1.ManagedOSVersionChannel), err
}

// Delete takes name of the managedOSVersionChannel and deletes it. Returns an error if one occurs.
func (c *FakeManagedOSVersionChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managedosversionchannelsResource, c.ns, name), &rancheroscattleiov1.ManagedOSVersionChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedOSVersionChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedosversionchannelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &rancheroscattleiov1.ManagedOSVersionChannelList{})
	return err
}

// Patch applies the patch and returns the patched managedOSVersionChannel.
func (c *FakeManagedOSVersionChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *rancheroscattleiov1.ManagedOSVersionChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedosversionchannelsResource, c.ns, name, pt, data, subresources...), &rancheroscattleiov1.ManagedOSVersionChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rancheroscattleiov1.ManagedOSVersionChannel), err
}
