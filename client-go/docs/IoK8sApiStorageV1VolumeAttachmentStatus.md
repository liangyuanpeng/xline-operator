# IoK8sApiStorageV1VolumeAttachmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachError** | Pointer to [**IoK8sApiStorageV1VolumeError**](IoK8sApiStorageV1VolumeError.md) |  | [optional] 
**Attached** | **bool** | attached indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher. | 
**AttachmentMetadata** | Pointer to **map[string]string** | attachmentMetadata is populated with any information returned by the attach operation, upon successful attach, that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher. | [optional] 
**DetachError** | Pointer to [**IoK8sApiStorageV1VolumeError**](IoK8sApiStorageV1VolumeError.md) |  | [optional] 

## Methods

### NewIoK8sApiStorageV1VolumeAttachmentStatus

`func NewIoK8sApiStorageV1VolumeAttachmentStatus(attached bool, ) *IoK8sApiStorageV1VolumeAttachmentStatus`

NewIoK8sApiStorageV1VolumeAttachmentStatus instantiates a new IoK8sApiStorageV1VolumeAttachmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiStorageV1VolumeAttachmentStatusWithDefaults

`func NewIoK8sApiStorageV1VolumeAttachmentStatusWithDefaults() *IoK8sApiStorageV1VolumeAttachmentStatus`

NewIoK8sApiStorageV1VolumeAttachmentStatusWithDefaults instantiates a new IoK8sApiStorageV1VolumeAttachmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachError

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) GetAttachError() IoK8sApiStorageV1VolumeError`

GetAttachError returns the AttachError field if non-nil, zero value otherwise.

### GetAttachErrorOk

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) GetAttachErrorOk() (*IoK8sApiStorageV1VolumeError, bool)`

GetAttachErrorOk returns a tuple with the AttachError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachError

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) SetAttachError(v IoK8sApiStorageV1VolumeError)`

SetAttachError sets AttachError field to given value.

### HasAttachError

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) HasAttachError() bool`

HasAttachError returns a boolean if a field has been set.

### GetAttached

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) GetAttached() bool`

GetAttached returns the Attached field if non-nil, zero value otherwise.

### GetAttachedOk

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) GetAttachedOk() (*bool, bool)`

GetAttachedOk returns a tuple with the Attached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttached

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) SetAttached(v bool)`

SetAttached sets Attached field to given value.


### GetAttachmentMetadata

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) GetAttachmentMetadata() map[string]string`

GetAttachmentMetadata returns the AttachmentMetadata field if non-nil, zero value otherwise.

### GetAttachmentMetadataOk

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) GetAttachmentMetadataOk() (*map[string]string, bool)`

GetAttachmentMetadataOk returns a tuple with the AttachmentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentMetadata

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) SetAttachmentMetadata(v map[string]string)`

SetAttachmentMetadata sets AttachmentMetadata field to given value.

### HasAttachmentMetadata

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) HasAttachmentMetadata() bool`

HasAttachmentMetadata returns a boolean if a field has been set.

### GetDetachError

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) GetDetachError() IoK8sApiStorageV1VolumeError`

GetDetachError returns the DetachError field if non-nil, zero value otherwise.

### GetDetachErrorOk

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) GetDetachErrorOk() (*IoK8sApiStorageV1VolumeError, bool)`

GetDetachErrorOk returns a tuple with the DetachError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachError

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) SetDetachError(v IoK8sApiStorageV1VolumeError)`

SetDetachError sets DetachError field to given value.

### HasDetachError

`func (o *IoK8sApiStorageV1VolumeAttachmentStatus) HasDetachError() bool`

HasDetachError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


