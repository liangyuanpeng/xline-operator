# IoK8sApiCoreV1ConfigMapProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]IoK8sApiCoreV1KeyToPath**](IoK8sApiCoreV1KeyToPath.md) | items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the &#39;..&#39; path or start with &#39;..&#39;. | [optional] 
**Name** | Pointer to **string** | Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [optional] 
**Optional** | Pointer to **bool** | optional specify whether the ConfigMap or its keys must be defined | [optional] 

## Methods

### NewIoK8sApiCoreV1ConfigMapProjection

`func NewIoK8sApiCoreV1ConfigMapProjection() *IoK8sApiCoreV1ConfigMapProjection`

NewIoK8sApiCoreV1ConfigMapProjection instantiates a new IoK8sApiCoreV1ConfigMapProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ConfigMapProjectionWithDefaults

`func NewIoK8sApiCoreV1ConfigMapProjectionWithDefaults() *IoK8sApiCoreV1ConfigMapProjection`

NewIoK8sApiCoreV1ConfigMapProjectionWithDefaults instantiates a new IoK8sApiCoreV1ConfigMapProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *IoK8sApiCoreV1ConfigMapProjection) GetItems() []IoK8sApiCoreV1KeyToPath`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoK8sApiCoreV1ConfigMapProjection) GetItemsOk() (*[]IoK8sApiCoreV1KeyToPath, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoK8sApiCoreV1ConfigMapProjection) SetItems(v []IoK8sApiCoreV1KeyToPath)`

SetItems sets Items field to given value.

### HasItems

`func (o *IoK8sApiCoreV1ConfigMapProjection) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiCoreV1ConfigMapProjection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1ConfigMapProjection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1ConfigMapProjection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiCoreV1ConfigMapProjection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptional

`func (o *IoK8sApiCoreV1ConfigMapProjection) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *IoK8sApiCoreV1ConfigMapProjection) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *IoK8sApiCoreV1ConfigMapProjection) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *IoK8sApiCoreV1ConfigMapProjection) HasOptional() bool`

HasOptional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


