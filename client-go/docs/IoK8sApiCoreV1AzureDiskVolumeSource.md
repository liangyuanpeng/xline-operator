# IoK8sApiCoreV1AzureDiskVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CachingMode** | Pointer to **string** | cachingMode is the Host Caching mode: None, Read Only, Read Write.  Possible enum values:  - &#x60;\&quot;None\&quot;&#x60;  - &#x60;\&quot;ReadOnly\&quot;&#x60;  - &#x60;\&quot;ReadWrite\&quot;&#x60; | [optional] 
**DiskName** | **string** | diskName is the Name of the data disk in the blob storage | 
**DiskURI** | **string** | diskURI is the URI of data disk in the blob storage | 
**FsType** | Pointer to **string** | fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. | [optional] 
**Kind** | Pointer to **string** | kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared  Possible enum values:  - &#x60;\&quot;Dedicated\&quot;&#x60;  - &#x60;\&quot;Managed\&quot;&#x60;  - &#x60;\&quot;Shared\&quot;&#x60; | [optional] 
**ReadOnly** | Pointer to **bool** | readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] 

## Methods

### NewIoK8sApiCoreV1AzureDiskVolumeSource

`func NewIoK8sApiCoreV1AzureDiskVolumeSource(diskName string, diskURI string, ) *IoK8sApiCoreV1AzureDiskVolumeSource`

NewIoK8sApiCoreV1AzureDiskVolumeSource instantiates a new IoK8sApiCoreV1AzureDiskVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1AzureDiskVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1AzureDiskVolumeSourceWithDefaults() *IoK8sApiCoreV1AzureDiskVolumeSource`

NewIoK8sApiCoreV1AzureDiskVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1AzureDiskVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCachingMode

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetCachingMode() string`

GetCachingMode returns the CachingMode field if non-nil, zero value otherwise.

### GetCachingModeOk

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetCachingModeOk() (*string, bool)`

GetCachingModeOk returns a tuple with the CachingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingMode

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetCachingMode(v string)`

SetCachingMode sets CachingMode field to given value.

### HasCachingMode

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) HasCachingMode() bool`

HasCachingMode returns a boolean if a field has been set.

### GetDiskName

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetDiskName() string`

GetDiskName returns the DiskName field if non-nil, zero value otherwise.

### GetDiskNameOk

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetDiskNameOk() (*string, bool)`

GetDiskNameOk returns a tuple with the DiskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskName

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetDiskName(v string)`

SetDiskName sets DiskName field to given value.


### GetDiskURI

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetDiskURI() string`

GetDiskURI returns the DiskURI field if non-nil, zero value otherwise.

### GetDiskURIOk

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetDiskURIOk() (*string, bool)`

GetDiskURIOk returns a tuple with the DiskURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskURI

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetDiskURI(v string)`

SetDiskURI sets DiskURI field to given value.


### GetFsType

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1AzureDiskVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


