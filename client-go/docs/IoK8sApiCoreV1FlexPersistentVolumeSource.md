# IoK8sApiCoreV1FlexPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Driver** | **string** | driver is the name of the driver to use for this volume. | 
**FsType** | Pointer to **string** | fsType is the Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. The default filesystem depends on FlexVolume script. | [optional] 
**Options** | Pointer to **map[string]string** | options is Optional: this field holds extra command options if any. | [optional] 
**ReadOnly** | Pointer to **bool** | readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] 
**SecretRef** | Pointer to [**IoK8sApiCoreV1SecretReference**](IoK8sApiCoreV1SecretReference.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1FlexPersistentVolumeSource

`func NewIoK8sApiCoreV1FlexPersistentVolumeSource(driver string, ) *IoK8sApiCoreV1FlexPersistentVolumeSource`

NewIoK8sApiCoreV1FlexPersistentVolumeSource instantiates a new IoK8sApiCoreV1FlexPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1FlexPersistentVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1FlexPersistentVolumeSourceWithDefaults() *IoK8sApiCoreV1FlexPersistentVolumeSource`

NewIoK8sApiCoreV1FlexPersistentVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1FlexPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriver

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) SetDriver(v string)`

SetDriver sets Driver field to given value.


### GetFsType

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetOptions

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSecretRef

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) GetSecretRef() IoK8sApiCoreV1SecretReference`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) GetSecretRefOk() (*IoK8sApiCoreV1SecretReference, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) SetSecretRef(v IoK8sApiCoreV1SecretReference)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *IoK8sApiCoreV1FlexPersistentVolumeSource) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


