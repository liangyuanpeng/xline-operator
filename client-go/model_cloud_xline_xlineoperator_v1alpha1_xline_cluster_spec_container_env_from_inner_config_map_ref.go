/*
Kubernetes

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.27.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef{}

// CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef The ConfigMap to select from
type CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`
	// Specify whether the ConfigMap must be defined
	Optional *bool `json:"optional,omitempty"`
}

// NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef {
	this := CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef{}
	return &this
}

// NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRefWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRefWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef {
	this := CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) SetName(v string) {
	o.Name = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) GetOptional() bool {
	if o == nil || IsNil(o.Optional) {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) GetOptionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Optional) {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) HasOptional() bool {
	if o != nil && !IsNil(o.Optional) {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) SetOptional(v bool) {
	o.Optional = &v
}

func (o CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Optional) {
		toSerialize["optional"] = o.Optional
	}
	return toSerialize, nil
}

type NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef struct {
	value *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef
	isSet bool
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) Get() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef {
	return v.value
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) Set(val *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef(val *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef {
	return &NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef{value: val, isSet: true}
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInnerConfigMapRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


