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

// checks if the CloudXlineXlineoperatorV1alpha1XlineClusterSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudXlineXlineoperatorV1alpha1XlineClusterSpec{}

// CloudXlineXlineoperatorV1alpha1XlineClusterSpec Xline cluster specification
type CloudXlineXlineoperatorV1alpha1XlineClusterSpec struct {
	// Backup specification
	Backup map[string]interface{} `json:"backup,omitempty"`
	Container CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer `json:"container"`
	// The data PVC, if it is not specified, then use emptyDir instead
	Data map[string]interface{} `json:"data,omitempty"`
	// Some user defined persistent volume claim templates
	Pvcs map[string]interface{} `json:"pvcs,omitempty"`
	// Size of the xline cluster, less than 3 is not allowed
	Size int32 `json:"size"`
}

// NewCloudXlineXlineoperatorV1alpha1XlineClusterSpec instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpec(container CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer, size int32) *CloudXlineXlineoperatorV1alpha1XlineClusterSpec {
	this := CloudXlineXlineoperatorV1alpha1XlineClusterSpec{}
	this.Container = container
	this.Size = size
	return &this
}

// NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpec {
	this := CloudXlineXlineoperatorV1alpha1XlineClusterSpec{}
	return &this
}

// GetBackup returns the Backup field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetBackup() map[string]interface{} {
	if o == nil || IsNil(o.Backup) {
		var ret map[string]interface{}
		return ret
	}
	return o.Backup
}

// GetBackupOk returns a tuple with the Backup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetBackupOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Backup) {
		return map[string]interface{}{}, false
	}
	return o.Backup, true
}

// HasBackup returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) HasBackup() bool {
	if o != nil && !IsNil(o.Backup) {
		return true
	}

	return false
}

// SetBackup gets a reference to the given map[string]interface{} and assigns it to the Backup field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) SetBackup(v map[string]interface{}) {
	o.Backup = v
}

// GetContainer returns the Container field value
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetContainer() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer {
	if o == nil {
		var ret CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetContainerOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) SetContainer(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) {
	o.Container = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetPvcs returns the Pvcs field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetPvcs() map[string]interface{} {
	if o == nil || IsNil(o.Pvcs) {
		var ret map[string]interface{}
		return ret
	}
	return o.Pvcs
}

// GetPvcsOk returns a tuple with the Pvcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetPvcsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Pvcs) {
		return map[string]interface{}{}, false
	}
	return o.Pvcs, true
}

// HasPvcs returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) HasPvcs() bool {
	if o != nil && !IsNil(o.Pvcs) {
		return true
	}

	return false
}

// SetPvcs gets a reference to the given map[string]interface{} and assigns it to the Pvcs field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) SetPvcs(v map[string]interface{}) {
	o.Pvcs = v
}

// GetSize returns the Size field value
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) SetSize(v int32) {
	o.Size = v
}

func (o CloudXlineXlineoperatorV1alpha1XlineClusterSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudXlineXlineoperatorV1alpha1XlineClusterSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Backup) {
		toSerialize["backup"] = o.Backup
	}
	toSerialize["container"] = o.Container
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pvcs) {
		toSerialize["pvcs"] = o.Pvcs
	}
	toSerialize["size"] = o.Size
	return toSerialize, nil
}

type NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpec struct {
	value *CloudXlineXlineoperatorV1alpha1XlineClusterSpec
	isSet bool
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpec) Get() *CloudXlineXlineoperatorV1alpha1XlineClusterSpec {
	return v.value
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpec) Set(val *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudXlineXlineoperatorV1alpha1XlineClusterSpec(val *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpec {
	return &NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpec{value: val, isSet: true}
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


