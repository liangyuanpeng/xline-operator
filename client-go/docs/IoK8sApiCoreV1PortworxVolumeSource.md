# IoK8sApiCoreV1PortworxVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | Pointer to **string** | fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. | [optional] 
**ReadOnly** | Pointer to **bool** | readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] 
**VolumeID** | **string** | volumeID uniquely identifies a Portworx volume | 

## Methods

### NewIoK8sApiCoreV1PortworxVolumeSource

`func NewIoK8sApiCoreV1PortworxVolumeSource(volumeID string, ) *IoK8sApiCoreV1PortworxVolumeSource`

NewIoK8sApiCoreV1PortworxVolumeSource instantiates a new IoK8sApiCoreV1PortworxVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PortworxVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1PortworxVolumeSourceWithDefaults() *IoK8sApiCoreV1PortworxVolumeSource`

NewIoK8sApiCoreV1PortworxVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1PortworxVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsType

`func (o *IoK8sApiCoreV1PortworxVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1PortworxVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1PortworxVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1PortworxVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1PortworxVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1PortworxVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1PortworxVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1PortworxVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetVolumeID

`func (o *IoK8sApiCoreV1PortworxVolumeSource) GetVolumeID() string`

GetVolumeID returns the VolumeID field if non-nil, zero value otherwise.

### GetVolumeIDOk

`func (o *IoK8sApiCoreV1PortworxVolumeSource) GetVolumeIDOk() (*string, bool)`

GetVolumeIDOk returns a tuple with the VolumeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeID

`func (o *IoK8sApiCoreV1PortworxVolumeSource) SetVolumeID(v string)`

SetVolumeID sets VolumeID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


