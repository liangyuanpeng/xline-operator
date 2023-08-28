# IoK8sApiCoreV1ProjectedVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMode** | Pointer to **int32** | defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. | [optional] 
**Sources** | Pointer to [**[]IoK8sApiCoreV1VolumeProjection**](IoK8sApiCoreV1VolumeProjection.md) | sources is the list of volume projections | [optional] 

## Methods

### NewIoK8sApiCoreV1ProjectedVolumeSource

`func NewIoK8sApiCoreV1ProjectedVolumeSource() *IoK8sApiCoreV1ProjectedVolumeSource`

NewIoK8sApiCoreV1ProjectedVolumeSource instantiates a new IoK8sApiCoreV1ProjectedVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ProjectedVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1ProjectedVolumeSourceWithDefaults() *IoK8sApiCoreV1ProjectedVolumeSource`

NewIoK8sApiCoreV1ProjectedVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1ProjectedVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMode

`func (o *IoK8sApiCoreV1ProjectedVolumeSource) GetDefaultMode() int32`

GetDefaultMode returns the DefaultMode field if non-nil, zero value otherwise.

### GetDefaultModeOk

`func (o *IoK8sApiCoreV1ProjectedVolumeSource) GetDefaultModeOk() (*int32, bool)`

GetDefaultModeOk returns a tuple with the DefaultMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMode

`func (o *IoK8sApiCoreV1ProjectedVolumeSource) SetDefaultMode(v int32)`

SetDefaultMode sets DefaultMode field to given value.

### HasDefaultMode

`func (o *IoK8sApiCoreV1ProjectedVolumeSource) HasDefaultMode() bool`

HasDefaultMode returns a boolean if a field has been set.

### GetSources

`func (o *IoK8sApiCoreV1ProjectedVolumeSource) GetSources() []IoK8sApiCoreV1VolumeProjection`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *IoK8sApiCoreV1ProjectedVolumeSource) GetSourcesOk() (*[]IoK8sApiCoreV1VolumeProjection, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *IoK8sApiCoreV1ProjectedVolumeSource) SetSources(v []IoK8sApiCoreV1VolumeProjection)`

SetSources sets Sources field to given value.

### HasSources

`func (o *IoK8sApiCoreV1ProjectedVolumeSource) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


