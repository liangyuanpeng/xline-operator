# IoK8sApiCoreV1LocalVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | Pointer to **string** | fsType is the filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. The default value is to auto-select a filesystem if unspecified. | [optional] 
**Path** | **string** | path of the full path to the volume on the node. It can be either a directory or block device (disk, partition, ...). | 

## Methods

### NewIoK8sApiCoreV1LocalVolumeSource

`func NewIoK8sApiCoreV1LocalVolumeSource(path string, ) *IoK8sApiCoreV1LocalVolumeSource`

NewIoK8sApiCoreV1LocalVolumeSource instantiates a new IoK8sApiCoreV1LocalVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1LocalVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1LocalVolumeSourceWithDefaults() *IoK8sApiCoreV1LocalVolumeSource`

NewIoK8sApiCoreV1LocalVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1LocalVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsType

`func (o *IoK8sApiCoreV1LocalVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1LocalVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1LocalVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1LocalVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetPath

`func (o *IoK8sApiCoreV1LocalVolumeSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoK8sApiCoreV1LocalVolumeSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoK8sApiCoreV1LocalVolumeSource) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


