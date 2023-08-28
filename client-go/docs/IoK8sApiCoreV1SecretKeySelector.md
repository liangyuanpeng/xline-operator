# IoK8sApiCoreV1SecretKeySelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the secret to select from.  Must be a valid secret key. | 
**Name** | Pointer to **string** | Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [optional] 
**Optional** | Pointer to **bool** | Specify whether the Secret or its key must be defined | [optional] 

## Methods

### NewIoK8sApiCoreV1SecretKeySelector

`func NewIoK8sApiCoreV1SecretKeySelector(key string, ) *IoK8sApiCoreV1SecretKeySelector`

NewIoK8sApiCoreV1SecretKeySelector instantiates a new IoK8sApiCoreV1SecretKeySelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1SecretKeySelectorWithDefaults

`func NewIoK8sApiCoreV1SecretKeySelectorWithDefaults() *IoK8sApiCoreV1SecretKeySelector`

NewIoK8sApiCoreV1SecretKeySelectorWithDefaults instantiates a new IoK8sApiCoreV1SecretKeySelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *IoK8sApiCoreV1SecretKeySelector) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IoK8sApiCoreV1SecretKeySelector) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IoK8sApiCoreV1SecretKeySelector) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *IoK8sApiCoreV1SecretKeySelector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1SecretKeySelector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1SecretKeySelector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiCoreV1SecretKeySelector) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptional

`func (o *IoK8sApiCoreV1SecretKeySelector) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *IoK8sApiCoreV1SecretKeySelector) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *IoK8sApiCoreV1SecretKeySelector) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *IoK8sApiCoreV1SecretKeySelector) HasOptional() bool`

HasOptional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


