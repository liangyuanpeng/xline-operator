# IoK8sApiCoreV1DownwardAPIVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMode** | Pointer to **int32** | Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. | [optional] 
**Items** | Pointer to [**[]IoK8sApiCoreV1DownwardAPIVolumeFile**](IoK8sApiCoreV1DownwardAPIVolumeFile.md) | Items is a list of downward API volume file | [optional] 

## Methods

### NewIoK8sApiCoreV1DownwardAPIVolumeSource

`func NewIoK8sApiCoreV1DownwardAPIVolumeSource() *IoK8sApiCoreV1DownwardAPIVolumeSource`

NewIoK8sApiCoreV1DownwardAPIVolumeSource instantiates a new IoK8sApiCoreV1DownwardAPIVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1DownwardAPIVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1DownwardAPIVolumeSourceWithDefaults() *IoK8sApiCoreV1DownwardAPIVolumeSource`

NewIoK8sApiCoreV1DownwardAPIVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1DownwardAPIVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMode

`func (o *IoK8sApiCoreV1DownwardAPIVolumeSource) GetDefaultMode() int32`

GetDefaultMode returns the DefaultMode field if non-nil, zero value otherwise.

### GetDefaultModeOk

`func (o *IoK8sApiCoreV1DownwardAPIVolumeSource) GetDefaultModeOk() (*int32, bool)`

GetDefaultModeOk returns a tuple with the DefaultMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMode

`func (o *IoK8sApiCoreV1DownwardAPIVolumeSource) SetDefaultMode(v int32)`

SetDefaultMode sets DefaultMode field to given value.

### HasDefaultMode

`func (o *IoK8sApiCoreV1DownwardAPIVolumeSource) HasDefaultMode() bool`

HasDefaultMode returns a boolean if a field has been set.

### GetItems

`func (o *IoK8sApiCoreV1DownwardAPIVolumeSource) GetItems() []IoK8sApiCoreV1DownwardAPIVolumeFile`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoK8sApiCoreV1DownwardAPIVolumeSource) GetItemsOk() (*[]IoK8sApiCoreV1DownwardAPIVolumeFile, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoK8sApiCoreV1DownwardAPIVolumeSource) SetItems(v []IoK8sApiCoreV1DownwardAPIVolumeFile)`

SetItems sets Items field to given value.

### HasItems

`func (o *IoK8sApiCoreV1DownwardAPIVolumeSource) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


