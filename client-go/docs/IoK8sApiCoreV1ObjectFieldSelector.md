# IoK8sApiCoreV1ObjectFieldSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | Version of the schema the FieldPath is written in terms of, defaults to \&quot;v1\&quot;. | [optional] 
**FieldPath** | **string** | Path of the field to select in the specified API version. | 

## Methods

### NewIoK8sApiCoreV1ObjectFieldSelector

`func NewIoK8sApiCoreV1ObjectFieldSelector(fieldPath string, ) *IoK8sApiCoreV1ObjectFieldSelector`

NewIoK8sApiCoreV1ObjectFieldSelector instantiates a new IoK8sApiCoreV1ObjectFieldSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ObjectFieldSelectorWithDefaults

`func NewIoK8sApiCoreV1ObjectFieldSelectorWithDefaults() *IoK8sApiCoreV1ObjectFieldSelector`

NewIoK8sApiCoreV1ObjectFieldSelectorWithDefaults instantiates a new IoK8sApiCoreV1ObjectFieldSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiCoreV1ObjectFieldSelector) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiCoreV1ObjectFieldSelector) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiCoreV1ObjectFieldSelector) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiCoreV1ObjectFieldSelector) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetFieldPath

`func (o *IoK8sApiCoreV1ObjectFieldSelector) GetFieldPath() string`

GetFieldPath returns the FieldPath field if non-nil, zero value otherwise.

### GetFieldPathOk

`func (o *IoK8sApiCoreV1ObjectFieldSelector) GetFieldPathOk() (*string, bool)`

GetFieldPathOk returns a tuple with the FieldPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldPath

`func (o *IoK8sApiCoreV1ObjectFieldSelector) SetFieldPath(v string)`

SetFieldPath sets FieldPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


