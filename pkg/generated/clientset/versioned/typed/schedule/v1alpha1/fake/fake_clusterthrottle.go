// Licensed to Shingo Omura under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Shingo Omura licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/everpeace/kube-throttler/pkg/apis/schedule/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterThrottles implements ClusterThrottleInterface
type FakeClusterThrottles struct {
	Fake *FakeScheduleV1alpha1
}

var clusterthrottlesResource = schema.GroupVersionResource{Group: "schedule.k8s.everpeace.github.com", Version: "v1alpha1", Resource: "clusterthrottles"}

var clusterthrottlesKind = schema.GroupVersionKind{Group: "schedule.k8s.everpeace.github.com", Version: "v1alpha1", Kind: "ClusterThrottle"}

// Get takes name of the clusterThrottle, and returns the corresponding clusterThrottle object, and an error if there is any.
func (c *FakeClusterThrottles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterThrottle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterthrottlesResource, name), &v1alpha1.ClusterThrottle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterThrottle), err
}

// List takes label and field selectors, and returns the list of ClusterThrottles that match those selectors.
func (c *FakeClusterThrottles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterThrottleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterthrottlesResource, clusterthrottlesKind, opts), &v1alpha1.ClusterThrottleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterThrottleList{ListMeta: obj.(*v1alpha1.ClusterThrottleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterThrottleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterThrottles.
func (c *FakeClusterThrottles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterthrottlesResource, opts))
}

// Create takes the representation of a clusterThrottle and creates it.  Returns the server's representation of the clusterThrottle, and an error, if there is any.
func (c *FakeClusterThrottles) Create(ctx context.Context, clusterThrottle *v1alpha1.ClusterThrottle, opts v1.CreateOptions) (result *v1alpha1.ClusterThrottle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterthrottlesResource, clusterThrottle), &v1alpha1.ClusterThrottle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterThrottle), err
}

// Update takes the representation of a clusterThrottle and updates it. Returns the server's representation of the clusterThrottle, and an error, if there is any.
func (c *FakeClusterThrottles) Update(ctx context.Context, clusterThrottle *v1alpha1.ClusterThrottle, opts v1.UpdateOptions) (result *v1alpha1.ClusterThrottle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterthrottlesResource, clusterThrottle), &v1alpha1.ClusterThrottle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterThrottle), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterThrottles) UpdateStatus(ctx context.Context, clusterThrottle *v1alpha1.ClusterThrottle, opts v1.UpdateOptions) (*v1alpha1.ClusterThrottle, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterthrottlesResource, "status", clusterThrottle), &v1alpha1.ClusterThrottle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterThrottle), err
}

// Delete takes name of the clusterThrottle and deletes it. Returns an error if one occurs.
func (c *FakeClusterThrottles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterthrottlesResource, name), &v1alpha1.ClusterThrottle{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterThrottles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterthrottlesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterThrottleList{})
	return err
}

// Patch applies the patch and returns the patched clusterThrottle.
func (c *FakeClusterThrottles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterThrottle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterthrottlesResource, name, pt, data, subresources...), &v1alpha1.ClusterThrottle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterThrottle), err
}
