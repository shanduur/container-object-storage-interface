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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "sigs.k8s.io/container-object-storage-interface-api/apis/objectstorage.k8s.io/v1alpha1"
	scheme "sigs.k8s.io/container-object-storage-interface-api/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BucketClassesGetter has a method to return a BucketClassInterface.
// A group's client should implement this interface.
type BucketClassesGetter interface {
	BucketClasses() BucketClassInterface
}

// BucketClassInterface has methods to work with BucketClass resources.
type BucketClassInterface interface {
	Create(ctx context.Context, bucketClass *v1alpha1.BucketClass, opts v1.CreateOptions) (*v1alpha1.BucketClass, error)
	Update(ctx context.Context, bucketClass *v1alpha1.BucketClass, opts v1.UpdateOptions) (*v1alpha1.BucketClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BucketClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BucketClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BucketClass, err error)
	BucketClassExpansion
}

// bucketClasses implements BucketClassInterface
type bucketClasses struct {
	client rest.Interface
}

// newBucketClasses returns a BucketClasses
func newBucketClasses(c *ObjectstorageV1alpha1Client) *bucketClasses {
	return &bucketClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the bucketClass, and returns the corresponding bucketClass object, and an error if there is any.
func (c *bucketClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BucketClass, err error) {
	result = &v1alpha1.BucketClass{}
	err = c.client.Get().
		Resource("bucketclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BucketClasses that match those selectors.
func (c *bucketClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BucketClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BucketClassList{}
	err = c.client.Get().
		Resource("bucketclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bucketClasses.
func (c *bucketClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("bucketclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bucketClass and creates it.  Returns the server's representation of the bucketClass, and an error, if there is any.
func (c *bucketClasses) Create(ctx context.Context, bucketClass *v1alpha1.BucketClass, opts v1.CreateOptions) (result *v1alpha1.BucketClass, err error) {
	result = &v1alpha1.BucketClass{}
	err = c.client.Post().
		Resource("bucketclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bucketClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bucketClass and updates it. Returns the server's representation of the bucketClass, and an error, if there is any.
func (c *bucketClasses) Update(ctx context.Context, bucketClass *v1alpha1.BucketClass, opts v1.UpdateOptions) (result *v1alpha1.BucketClass, err error) {
	result = &v1alpha1.BucketClass{}
	err = c.client.Put().
		Resource("bucketclasses").
		Name(bucketClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bucketClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bucketClass and deletes it. Returns an error if one occurs.
func (c *bucketClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("bucketclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bucketClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("bucketclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bucketClass.
func (c *bucketClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BucketClass, err error) {
	result = &v1alpha1.BucketClass{}
	err = c.client.Patch(pt).
		Resource("bucketclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
