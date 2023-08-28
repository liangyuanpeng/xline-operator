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

// checks if the CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef{}

// CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
type CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef struct {
	// Container name: required for volumes, optional for env vars
	ContainerName *string `json:"containerName,omitempty"`
	// Specifies the output format of the exposed resources, defaults to \"1\"
	Divisor *string `json:"divisor,omitempty"`
	// Required: resource to select
	Resource string `json:"resource"`
}

// NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef(resource string) *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef {
	this := CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef{}
	this.Resource = resource
	return &this
}

// NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRefWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRefWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef {
	this := CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef{}
	return &this
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) GetContainerName() string {
	if o == nil || IsNil(o.ContainerName) {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) GetContainerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) HasContainerName() bool {
	if o != nil && !IsNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given string and assigns it to the ContainerName field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) SetContainerName(v string) {
	o.ContainerName = &v
}

// GetDivisor returns the Divisor field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) GetDivisor() string {
	if o == nil || IsNil(o.Divisor) {
		var ret string
		return ret
	}
	return *o.Divisor
}

// GetDivisorOk returns a tuple with the Divisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) GetDivisorOk() (*string, bool) {
	if o == nil || IsNil(o.Divisor) {
		return nil, false
	}
	return o.Divisor, true
}

// HasDivisor returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) HasDivisor() bool {
	if o != nil && !IsNil(o.Divisor) {
		return true
	}

	return false
}

// SetDivisor gets a reference to the given string and assigns it to the Divisor field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) SetDivisor(v string) {
	o.Divisor = &v
}

// GetResource returns the Resource field value
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) SetResource(v string) {
	o.Resource = v
}

func (o CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerName) {
		toSerialize["containerName"] = o.ContainerName
	}
	if !IsNil(o.Divisor) {
		toSerialize["divisor"] = o.Divisor
	}
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

type NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef struct {
	value *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef
	isSet bool
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) Get() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef {
	return v.value
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) Set(val *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef(val *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef {
	return &NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef{value: val, isSet: true}
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFromResourceFieldRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


