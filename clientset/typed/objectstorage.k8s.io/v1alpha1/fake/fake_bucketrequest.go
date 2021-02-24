/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "sigs.k8s.io/container-object-storage-interface-api/apis/objectstorage.k8s.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBucketRequests implements BucketRequestInterface
type FakeBucketRequests struct {
	Fake *FakeObjectstorageV1alpha1
	ns   string
}

var bucketrequestsResource = schema.GroupVersionResource{Group: "objectstorage.k8s.io", Version: "v1alpha1", Resource: "bucketrequests"}

var bucketrequestsKind = schema.GroupVersionKind{Group: "objectstorage.k8s.io", Version: "v1alpha1", Kind: "BucketRequest"}

// Get takes name of the bucketRequest, and returns the corresponding bucketRequest object, and an error if there is any.
func (c *FakeBucketRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BucketRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bucketrequestsResource, c.ns, name), &v1alpha1.BucketRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketRequest), err
}

// List takes label and field selectors, and returns the list of BucketRequests that match those selectors.
func (c *FakeBucketRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BucketRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bucketrequestsResource, bucketrequestsKind, c.ns, opts), &v1alpha1.BucketRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BucketRequestList{ListMeta: obj.(*v1alpha1.BucketRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.BucketRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bucketRequests.
func (c *FakeBucketRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bucketrequestsResource, c.ns, opts))

}

// Create takes the representation of a bucketRequest and creates it.  Returns the server's representation of the bucketRequest, and an error, if there is any.
func (c *FakeBucketRequests) Create(ctx context.Context, bucketRequest *v1alpha1.BucketRequest, opts v1.CreateOptions) (result *v1alpha1.BucketRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bucketrequestsResource, c.ns, bucketRequest), &v1alpha1.BucketRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketRequest), err
}

// Update takes the representation of a bucketRequest and updates it. Returns the server's representation of the bucketRequest, and an error, if there is any.
func (c *FakeBucketRequests) Update(ctx context.Context, bucketRequest *v1alpha1.BucketRequest, opts v1.UpdateOptions) (result *v1alpha1.BucketRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bucketrequestsResource, c.ns, bucketRequest), &v1alpha1.BucketRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBucketRequests) UpdateStatus(ctx context.Context, bucketRequest *v1alpha1.BucketRequest, opts v1.UpdateOptions) (*v1alpha1.BucketRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bucketrequestsResource, "status", c.ns, bucketRequest), &v1alpha1.BucketRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketRequest), err
}

// Delete takes name of the bucketRequest and deletes it. Returns an error if one occurs.
func (c *FakeBucketRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bucketrequestsResource, c.ns, name), &v1alpha1.BucketRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBucketRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bucketrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BucketRequestList{})
	return err
}

// Patch applies the patch and returns the patched bucketRequest.
func (c *FakeBucketRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BucketRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bucketrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BucketRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketRequest), err
}
