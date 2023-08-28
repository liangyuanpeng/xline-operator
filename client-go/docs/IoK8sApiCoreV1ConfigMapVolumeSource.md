# IoK8sApiCoreV1ConfigMapVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMode** | Pointer to **int32** | defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. | [optional] 
**Items** | Pointer to [**[]IoK8sApiCoreV1KeyToPath**](IoK8sApiCoreV1KeyToPath.md) | items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the &#39;..&#39; path or start with &#39;..&#39;. | [optional] 
**Name** | Pointer to **string** | Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [optional] 
**Optional** | Pointer to **bool** | optional specify whether the ConfigMap or its keys must be defined | [optional] 

## Methods

### NewIoK8sApiCoreV1ConfigMapVolumeSource

`func NewIoK8sApiCoreV1ConfigMapVolumeSource() *IoK8sApiCoreV1ConfigMapVolumeSource`

NewIoK8sApiCoreV1ConfigMapVolumeSource instantiates a new IoK8sApiCoreV1ConfigMapVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ConfigMapVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1ConfigMapVolumeSourceWithDefaults() *IoK8sApiCoreV1ConfigMapVolumeSource`

NewIoK8sApiCoreV1ConfigMapVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1ConfigMapVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMode

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) GetDefaultMode() int32`

GetDefaultMode returns the DefaultMode field if non-nil, zero value otherwise.

### GetDefaultModeOk

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) GetDefaultModeOk() (*int32, bool)`

GetDefaultModeOk returns a tuple with the DefaultMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMode

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) SetDefaultMode(v int32)`

SetDefaultMode sets DefaultMode field to given value.

### HasDefaultMode

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) HasDefaultMode() bool`

HasDefaultMode returns a boolean if a field has been set.

### GetItems

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) GetItems() []IoK8sApiCoreV1KeyToPath`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) GetItemsOk() (*[]IoK8sApiCoreV1KeyToPath, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) SetItems(v []IoK8sApiCoreV1KeyToPath)`

SetItems sets Items field to given value.

### HasItems

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptional

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *IoK8sApiCoreV1ConfigMapVolumeSource) HasOptional() bool`

HasOptional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


