# IoK8sApiCoreV1GCEPersistentDiskVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | Pointer to **string** | fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | [optional] 
**Partition** | Pointer to **int32** | partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \&quot;1\&quot;. Similarly, the volume partition for /dev/sda is \&quot;0\&quot; (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | [optional] 
**PdName** | **string** | pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | 
**ReadOnly** | Pointer to **bool** | readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | [optional] 

## Methods

### NewIoK8sApiCoreV1GCEPersistentDiskVolumeSource

`func NewIoK8sApiCoreV1GCEPersistentDiskVolumeSource(pdName string, ) *IoK8sApiCoreV1GCEPersistentDiskVolumeSource`

NewIoK8sApiCoreV1GCEPersistentDiskVolumeSource instantiates a new IoK8sApiCoreV1GCEPersistentDiskVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1GCEPersistentDiskVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1GCEPersistentDiskVolumeSourceWithDefaults() *IoK8sApiCoreV1GCEPersistentDiskVolumeSource`

NewIoK8sApiCoreV1GCEPersistentDiskVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1GCEPersistentDiskVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsType

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetPartition

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPdName

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetPdName() string`

GetPdName returns the PdName field if non-nil, zero value otherwise.

### GetPdNameOk

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetPdNameOk() (*string, bool)`

GetPdNameOk returns a tuple with the PdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdName

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) SetPdName(v string)`

SetPdName sets PdName field to given value.


### GetReadOnly

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


