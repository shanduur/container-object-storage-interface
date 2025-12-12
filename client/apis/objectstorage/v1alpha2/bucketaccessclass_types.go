/*
Copyright 2025 The Kubernetes Authors.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BucketAccessClassSpec defines the desired state of BucketAccessClass
type BucketAccessClassSpec struct {
	// driverName is the name of the driver that fulfills requests for this BucketAccessClass.
	// +required
	// +kubebuilder:validation:MinLength=1
	DriverName string `json:"driverName,omitempty"`

	// authenticationType specifies which authentication mechanism is used bucket access.
	// Possible values:
	//  - Key: The driver should generate a protocol-appropriate access key that clients can use to
	//    authenticate to the backend object store.
	//  - ServiceAccount: The driver should configure the system such that Pods using the given
	//    ServiceAccount authenticate to the backend object store automatically.
	// +required
	AuthenticationType BucketAccessAuthenticationType `json:"authenticationType,omitempty"`

	// parameters is an opaque map of driver-specific configuration items passed to the driver that
	// fulfills requests for this BucketAccessClass.
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`

	// featureOptions can be used to adjust various COSI access provisioning behaviors.
	// If specified, at least one option must be set.
	// +optional
	FeatureOptions BucketAccessFeatureOptions `json:"featureOptions,omitzero"`
}

// BucketAccessFeatureOptions defines various COSI access provisioning behaviors.
// +kubebuilder:validation:MinProperties=1
type BucketAccessFeatureOptions struct {
	// disallowedBucketAccessModes is a list of disallowed Read/Write access modes. A BucketAccess
	// using this class will not be allowed to request access to a BucketClaim with any access mode
	// listed here.
	// +optional
	// +listType=set
	DisallowedBucketAccessModes []BucketAccessMode `json:"disallowedBucketAccessModes,omitempty"`

	// disallowMultiBucketAccess disables the ability for a BucketAccess to reference multiple
	// BucketClaims when set.
	// +optional
	DisallowMultiBucketAccess *bool `json:"disallowMultiBucketAccess,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:metadata:annotations="api-approved.kubernetes.io=unapproved, experimental v1alpha2 changes"

// BucketAccessClass is the Schema for the bucketaccessclasses API
type BucketAccessClass struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of BucketAccessClass
	// +required
	// +kubebuilder:validation:XValidation:message="BucketAccessClass spec is immutable",rule="self == oldSelf"
	Spec BucketAccessClassSpec `json:"spec,omitzero"`
}

// +kubebuilder:object:root=true

// BucketAccessClassList contains a list of BucketAccessClass
type BucketAccessClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketAccessClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BucketAccessClass{}, &BucketAccessClassList{})
}
