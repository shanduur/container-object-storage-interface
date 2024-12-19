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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/container-object-storage-interface/client/apis/objectstorage/v1alpha1"
)

// BucketLister helps list Buckets.
// All objects returned here must be treated as read-only.
type BucketLister interface {
	// List lists all Buckets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Bucket, err error)
	// Get retrieves the Bucket from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Bucket, error)
	BucketListerExpansion
}

// bucketLister implements the BucketLister interface.
type bucketLister struct {
	listers.ResourceIndexer[*v1alpha1.Bucket]
}

// NewBucketLister returns a new BucketLister.
func NewBucketLister(indexer cache.Indexer) BucketLister {
	return &bucketLister{listers.New[*v1alpha1.Bucket](indexer, v1alpha1.Resource("bucket"))}
}
