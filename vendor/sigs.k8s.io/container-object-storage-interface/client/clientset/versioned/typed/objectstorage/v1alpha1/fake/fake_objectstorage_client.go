/*
Copyright 2024 The Kubernetes Authors.

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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/container-object-storage-interface/client/clientset/versioned/typed/objectstorage/v1alpha1"
)

type FakeObjectstorageV1alpha1 struct {
	*testing.Fake
}

func (c *FakeObjectstorageV1alpha1) Buckets() v1alpha1.BucketInterface {
	return &FakeBuckets{c}
}

func (c *FakeObjectstorageV1alpha1) BucketAccesses(namespace string) v1alpha1.BucketAccessInterface {
	return &FakeBucketAccesses{c, namespace}
}

func (c *FakeObjectstorageV1alpha1) BucketAccessClasses() v1alpha1.BucketAccessClassInterface {
	return &FakeBucketAccessClasses{c}
}

func (c *FakeObjectstorageV1alpha1) BucketClaims(namespace string) v1alpha1.BucketClaimInterface {
	return &FakeBucketClaims{c, namespace}
}

func (c *FakeObjectstorageV1alpha1) BucketClasses() v1alpha1.BucketClassInterface {
	return &FakeBucketClasses{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeObjectstorageV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
