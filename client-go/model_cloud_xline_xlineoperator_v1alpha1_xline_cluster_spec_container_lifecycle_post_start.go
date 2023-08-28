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

// checks if the CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart{}

// CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
type CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart struct {
	Exec *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec `json:"exec,omitempty"`
	HttpGet *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet `json:"httpGet,omitempty"`
	TcpSocket *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartTcpSocket `json:"tcpSocket,omitempty"`
}

// NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart {
	this := CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart{}
	return &this
}

// NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart {
	this := CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart{}
	return &this
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) GetExec() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec {
	if o == nil || IsNil(o.Exec) {
		var ret CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec
		return ret
	}
	return *o.Exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) GetExecOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec, bool) {
	if o == nil || IsNil(o.Exec) {
		return nil, false
	}
	return o.Exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) HasExec() bool {
	if o != nil && !IsNil(o.Exec) {
		return true
	}

	return false
}

// SetExec gets a reference to the given CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec and assigns it to the Exec field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) SetExec(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec) {
	o.Exec = &v
}

// GetHttpGet returns the HttpGet field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) GetHttpGet() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet {
	if o == nil || IsNil(o.HttpGet) {
		var ret CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet
		return ret
	}
	return *o.HttpGet
}

// GetHttpGetOk returns a tuple with the HttpGet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) GetHttpGetOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet, bool) {
	if o == nil || IsNil(o.HttpGet) {
		return nil, false
	}
	return o.HttpGet, true
}

// HasHttpGet returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) HasHttpGet() bool {
	if o != nil && !IsNil(o.HttpGet) {
		return true
	}

	return false
}

// SetHttpGet gets a reference to the given CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet and assigns it to the HttpGet field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) SetHttpGet(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) {
	o.HttpGet = &v
}

// GetTcpSocket returns the TcpSocket field value if set, zero value otherwise.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) GetTcpSocket() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartTcpSocket {
	if o == nil || IsNil(o.TcpSocket) {
		var ret CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartTcpSocket
		return ret
	}
	return *o.TcpSocket
}

// GetTcpSocketOk returns a tuple with the TcpSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) GetTcpSocketOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartTcpSocket, bool) {
	if o == nil || IsNil(o.TcpSocket) {
		return nil, false
	}
	return o.TcpSocket, true
}

// HasTcpSocket returns a boolean if a field has been set.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) HasTcpSocket() bool {
	if o != nil && !IsNil(o.TcpSocket) {
		return true
	}

	return false
}

// SetTcpSocket gets a reference to the given CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartTcpSocket and assigns it to the TcpSocket field.
func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) SetTcpSocket(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartTcpSocket) {
	o.TcpSocket = &v
}

func (o CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exec) {
		toSerialize["exec"] = o.Exec
	}
	if !IsNil(o.HttpGet) {
		toSerialize["httpGet"] = o.HttpGet
	}
	if !IsNil(o.TcpSocket) {
		toSerialize["tcpSocket"] = o.TcpSocket
	}
	return toSerialize, nil
}

type NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart struct {
	value *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart
	isSet bool
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) Get() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart {
	return v.value
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) Set(val *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart(val *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart {
	return &NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart{value: val, isSet: true}
}

func (v NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


