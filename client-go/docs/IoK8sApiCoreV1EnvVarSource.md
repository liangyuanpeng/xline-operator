# IoK8sApiCoreV1EnvVarSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMapKeyRef** | Pointer to [**IoK8sApiCoreV1ConfigMapKeySelector**](IoK8sApiCoreV1ConfigMapKeySelector.md) |  | [optional] 
**FieldRef** | Pointer to [**IoK8sApiCoreV1ObjectFieldSelector**](IoK8sApiCoreV1ObjectFieldSelector.md) |  | [optional] 
**ResourceFieldRef** | Pointer to [**IoK8sApiCoreV1ResourceFieldSelector**](IoK8sApiCoreV1ResourceFieldSelector.md) |  | [optional] 
**SecretKeyRef** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1EnvVarSource

`func NewIoK8sApiCoreV1EnvVarSource() *IoK8sApiCoreV1EnvVarSource`

NewIoK8sApiCoreV1EnvVarSource instantiates a new IoK8sApiCoreV1EnvVarSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1EnvVarSourceWithDefaults

`func NewIoK8sApiCoreV1EnvVarSourceWithDefaults() *IoK8sApiCoreV1EnvVarSource`

NewIoK8sApiCoreV1EnvVarSourceWithDefaults instantiates a new IoK8sApiCoreV1EnvVarSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMapKeyRef

`func (o *IoK8sApiCoreV1EnvVarSource) GetConfigMapKeyRef() IoK8sApiCoreV1ConfigMapKeySelector`

GetConfigMapKeyRef returns the ConfigMapKeyRef field if non-nil, zero value otherwise.

### GetConfigMapKeyRefOk

`func (o *IoK8sApiCoreV1EnvVarSource) GetConfigMapKeyRefOk() (*IoK8sApiCoreV1ConfigMapKeySelector, bool)`

GetConfigMapKeyRefOk returns a tuple with the ConfigMapKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapKeyRef

`func (o *IoK8sApiCoreV1EnvVarSource) SetConfigMapKeyRef(v IoK8sApiCoreV1ConfigMapKeySelector)`

SetConfigMapKeyRef sets ConfigMapKeyRef field to given value.

### HasConfigMapKeyRef

`func (o *IoK8sApiCoreV1EnvVarSource) HasConfigMapKeyRef() bool`

HasConfigMapKeyRef returns a boolean if a field has been set.

### GetFieldRef

`func (o *IoK8sApiCoreV1EnvVarSource) GetFieldRef() IoK8sApiCoreV1ObjectFieldSelector`

GetFieldRef returns the FieldRef field if non-nil, zero value otherwise.

### GetFieldRefOk

`func (o *IoK8sApiCoreV1EnvVarSource) GetFieldRefOk() (*IoK8sApiCoreV1ObjectFieldSelector, bool)`

GetFieldRefOk returns a tuple with the FieldRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRef

`func (o *IoK8sApiCoreV1EnvVarSource) SetFieldRef(v IoK8sApiCoreV1ObjectFieldSelector)`

SetFieldRef sets FieldRef field to given value.

### HasFieldRef

`func (o *IoK8sApiCoreV1EnvVarSource) HasFieldRef() bool`

HasFieldRef returns a boolean if a field has been set.

### GetResourceFieldRef

`func (o *IoK8sApiCoreV1EnvVarSource) GetResourceFieldRef() IoK8sApiCoreV1ResourceFieldSelector`

GetResourceFieldRef returns the ResourceFieldRef field if non-nil, zero value otherwise.

### GetResourceFieldRefOk

`func (o *IoK8sApiCoreV1EnvVarSource) GetResourceFieldRefOk() (*IoK8sApiCoreV1ResourceFieldSelector, bool)`

GetResourceFieldRefOk returns a tuple with the ResourceFieldRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceFieldRef

`func (o *IoK8sApiCoreV1EnvVarSource) SetResourceFieldRef(v IoK8sApiCoreV1ResourceFieldSelector)`

SetResourceFieldRef sets ResourceFieldRef field to given value.

### HasResourceFieldRef

`func (o *IoK8sApiCoreV1EnvVarSource) HasResourceFieldRef() bool`

HasResourceFieldRef returns a boolean if a field has been set.

### GetSecretKeyRef

`func (o *IoK8sApiCoreV1EnvVarSource) GetSecretKeyRef() IoK8sApiCoreV1SecretKeySelector`

GetSecretKeyRef returns the SecretKeyRef field if non-nil, zero value otherwise.

### GetSecretKeyRefOk

`func (o *IoK8sApiCoreV1EnvVarSource) GetSecretKeyRefOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSecretKeyRefOk returns a tuple with the SecretKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyRef

`func (o *IoK8sApiCoreV1EnvVarSource) SetSecretKeyRef(v IoK8sApiCoreV1SecretKeySelector)`

SetSecretKeyRef sets SecretKeyRef field to given value.

### HasSecretKeyRef

`func (o *IoK8sApiCoreV1EnvVarSource) HasSecretKeyRef() bool`

HasSecretKeyRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


