# IoK8sApiCoreV1EnvFromSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMapRef** | Pointer to [**IoK8sApiCoreV1ConfigMapEnvSource**](IoK8sApiCoreV1ConfigMapEnvSource.md) |  | [optional] 
**Prefix** | Pointer to **string** | An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER. | [optional] 
**SecretRef** | Pointer to [**IoK8sApiCoreV1SecretEnvSource**](IoK8sApiCoreV1SecretEnvSource.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1EnvFromSource

`func NewIoK8sApiCoreV1EnvFromSource() *IoK8sApiCoreV1EnvFromSource`

NewIoK8sApiCoreV1EnvFromSource instantiates a new IoK8sApiCoreV1EnvFromSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1EnvFromSourceWithDefaults

`func NewIoK8sApiCoreV1EnvFromSourceWithDefaults() *IoK8sApiCoreV1EnvFromSource`

NewIoK8sApiCoreV1EnvFromSourceWithDefaults instantiates a new IoK8sApiCoreV1EnvFromSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMapRef

`func (o *IoK8sApiCoreV1EnvFromSource) GetConfigMapRef() IoK8sApiCoreV1ConfigMapEnvSource`

GetConfigMapRef returns the ConfigMapRef field if non-nil, zero value otherwise.

### GetConfigMapRefOk

`func (o *IoK8sApiCoreV1EnvFromSource) GetConfigMapRefOk() (*IoK8sApiCoreV1ConfigMapEnvSource, bool)`

GetConfigMapRefOk returns a tuple with the ConfigMapRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapRef

`func (o *IoK8sApiCoreV1EnvFromSource) SetConfigMapRef(v IoK8sApiCoreV1ConfigMapEnvSource)`

SetConfigMapRef sets ConfigMapRef field to given value.

### HasConfigMapRef

`func (o *IoK8sApiCoreV1EnvFromSource) HasConfigMapRef() bool`

HasConfigMapRef returns a boolean if a field has been set.

### GetPrefix

`func (o *IoK8sApiCoreV1EnvFromSource) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IoK8sApiCoreV1EnvFromSource) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IoK8sApiCoreV1EnvFromSource) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *IoK8sApiCoreV1EnvFromSource) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSecretRef

`func (o *IoK8sApiCoreV1EnvFromSource) GetSecretRef() IoK8sApiCoreV1SecretEnvSource`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *IoK8sApiCoreV1EnvFromSource) GetSecretRefOk() (*IoK8sApiCoreV1SecretEnvSource, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *IoK8sApiCoreV1EnvFromSource) SetSecretRef(v IoK8sApiCoreV1SecretEnvSource)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *IoK8sApiCoreV1EnvFromSource) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


