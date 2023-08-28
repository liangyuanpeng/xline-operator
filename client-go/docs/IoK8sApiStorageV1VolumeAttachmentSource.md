# IoK8sApiStorageV1VolumeAttachmentSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlineVolumeSpec** | Pointer to [**IoK8sApiCoreV1PersistentVolumeSpec**](IoK8sApiCoreV1PersistentVolumeSpec.md) |  | [optional] 
**PersistentVolumeName** | Pointer to **string** | persistentVolumeName represents the name of the persistent volume to attach. | [optional] 

## Methods

### NewIoK8sApiStorageV1VolumeAttachmentSource

`func NewIoK8sApiStorageV1VolumeAttachmentSource() *IoK8sApiStorageV1VolumeAttachmentSource`

NewIoK8sApiStorageV1VolumeAttachmentSource instantiates a new IoK8sApiStorageV1VolumeAttachmentSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiStorageV1VolumeAttachmentSourceWithDefaults

`func NewIoK8sApiStorageV1VolumeAttachmentSourceWithDefaults() *IoK8sApiStorageV1VolumeAttachmentSource`

NewIoK8sApiStorageV1VolumeAttachmentSourceWithDefaults instantiates a new IoK8sApiStorageV1VolumeAttachmentSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInlineVolumeSpec

`func (o *IoK8sApiStorageV1VolumeAttachmentSource) GetInlineVolumeSpec() IoK8sApiCoreV1PersistentVolumeSpec`

GetInlineVolumeSpec returns the InlineVolumeSpec field if non-nil, zero value otherwise.

### GetInlineVolumeSpecOk

`func (o *IoK8sApiStorageV1VolumeAttachmentSource) GetInlineVolumeSpecOk() (*IoK8sApiCoreV1PersistentVolumeSpec, bool)`

GetInlineVolumeSpecOk returns a tuple with the InlineVolumeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineVolumeSpec

`func (o *IoK8sApiStorageV1VolumeAttachmentSource) SetInlineVolumeSpec(v IoK8sApiCoreV1PersistentVolumeSpec)`

SetInlineVolumeSpec sets InlineVolumeSpec field to given value.

### HasInlineVolumeSpec

`func (o *IoK8sApiStorageV1VolumeAttachmentSource) HasInlineVolumeSpec() bool`

HasInlineVolumeSpec returns a boolean if a field has been set.

### GetPersistentVolumeName

`func (o *IoK8sApiStorageV1VolumeAttachmentSource) GetPersistentVolumeName() string`

GetPersistentVolumeName returns the PersistentVolumeName field if non-nil, zero value otherwise.

### GetPersistentVolumeNameOk

`func (o *IoK8sApiStorageV1VolumeAttachmentSource) GetPersistentVolumeNameOk() (*string, bool)`

GetPersistentVolumeNameOk returns a tuple with the PersistentVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeName

`func (o *IoK8sApiStorageV1VolumeAttachmentSource) SetPersistentVolumeName(v string)`

SetPersistentVolumeName sets PersistentVolumeName field to given value.

### HasPersistentVolumeName

`func (o *IoK8sApiStorageV1VolumeAttachmentSource) HasPersistentVolumeName() bool`

HasPersistentVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


