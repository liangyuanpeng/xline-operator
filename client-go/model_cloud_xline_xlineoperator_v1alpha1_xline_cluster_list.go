/*
Kubernetes

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.27.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// checks if the CloudXlineXlineoperatorV1alpha1XlineClusterList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudXlineXlineoperatorV1alpha1XlineClusterList{}

// CloudXlineXlineoperatorV1alpha1XlineClusterList XlineClusterList is a list of XlineCluster
type CloudXlineXlineoperatorV1alpha1XlineClusterList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// List of xlineclusters. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Items []CloudXlineXlineoperatorV1alpha1XlineCluster `json:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	metav1.ListMeta `json:"metadata,omitempty"`
}

// NewCloudXlineXlineoperatorV1alpha1XlineClusterList instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudXlineXlineoperatorV1alpha1XlineClusterList(items []CloudXlineXlineoperatorV1alpha1XlineCluster) *CloudXlineXlineoperatorV1alpha1XlineClusterList {
	this := CloudXlineXlineoperatorV1alpha1XlineClusterList{}
	this.Items = items
	return &this
}

// NewCloudXlineXlineoperatorV1alpha1XlineClusterListWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudXlineXlineoperatorV1alpha1XlineClusterListWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterList {
	this := CloudXlineXlineoperatorV1alpha1XlineClusterList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetItems() []CloudXlineXlineoperatorV1alpha1XlineCluster {
	if o == nil {
		var ret []CloudXlineXlineoperatorV1alpha1XlineCluster
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetItemsOk() ([]CloudXlineXlineoperatorV1alpha1XlineCluster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) SetItems(v []CloudXlineXlineoperatorV1alpha1XlineCluster) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) SetKind(v string) {
	o.Kind = &v
}

func (o CloudXlineXlineoperatorV1alpha1XlineClusterList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudXlineXlineoperatorV1alpha1XlineClusterList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	toSerialize["items"] = o.Items
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	// if !IsNil(o.Metadata) {
	// 	toSerialize["metadata"] = o.Metadata
	// }
	return toSerialize, nil
}

type NullableCloudXlineXlineoperatorV1alpha1XlineClusterList struct {
	value *CloudXlineXlineoperatorV1alpha1XlineClusterList
	isSet bool
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterList) Get() *CloudXlineXlineoperatorV1alpha1XlineClusterList {
	return v.value
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterList) Set(val *CloudXlineXlineoperatorV1alpha1XlineClusterList) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterList) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudXlineXlineoperatorV1alpha1XlineClusterList(val *CloudXlineXlineoperatorV1alpha1XlineClusterList) *NullableCloudXlineXlineoperatorV1alpha1XlineClusterList {
	return &NullableCloudXlineXlineoperatorV1alpha1XlineClusterList{value: val, isSet: true}
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


