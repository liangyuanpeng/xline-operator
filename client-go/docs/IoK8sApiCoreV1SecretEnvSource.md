# IoK8sApiCoreV1SecretEnvSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [optional] 
**Optional** | Pointer to **bool** | Specify whether the Secret must be defined | [optional] 

## Methods

### NewIoK8sApiCoreV1SecretEnvSource

`func NewIoK8sApiCoreV1SecretEnvSource() *IoK8sApiCoreV1SecretEnvSource`

NewIoK8sApiCoreV1SecretEnvSource instantiates a new IoK8sApiCoreV1SecretEnvSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1SecretEnvSourceWithDefaults

`func NewIoK8sApiCoreV1SecretEnvSourceWithDefaults() *IoK8sApiCoreV1SecretEnvSource`

NewIoK8sApiCoreV1SecretEnvSourceWithDefaults instantiates a new IoK8sApiCoreV1SecretEnvSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiCoreV1SecretEnvSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1SecretEnvSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1SecretEnvSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiCoreV1SecretEnvSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptional

`func (o *IoK8sApiCoreV1SecretEnvSource) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *IoK8sApiCoreV1SecretEnvSource) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *IoK8sApiCoreV1SecretEnvSource) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *IoK8sApiCoreV1SecretEnvSource) HasOptional() bool`

HasOptional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


