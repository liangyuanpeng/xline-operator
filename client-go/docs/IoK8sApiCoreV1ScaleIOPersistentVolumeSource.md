# IoK8sApiCoreV1ScaleIOPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | Pointer to **string** | fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Default is \&quot;xfs\&quot; | [optional] 
**Gateway** | **string** | gateway is the host address of the ScaleIO API Gateway. | 
**ProtectionDomain** | Pointer to **string** | protectionDomain is the name of the ScaleIO Protection Domain for the configured storage. | [optional] 
**ReadOnly** | Pointer to **bool** | readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] 
**SecretRef** | [**IoK8sApiCoreV1SecretReference**](IoK8sApiCoreV1SecretReference.md) |  | 
**SslEnabled** | Pointer to **bool** | sslEnabled is the flag to enable/disable SSL communication with Gateway, default false | [optional] 
**StorageMode** | Pointer to **string** | storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned. | [optional] 
**StoragePool** | Pointer to **string** | storagePool is the ScaleIO Storage Pool associated with the protection domain. | [optional] 
**System** | **string** | system is the name of the storage system as configured in ScaleIO. | 
**VolumeName** | Pointer to **string** | volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source. | [optional] 

## Methods

### NewIoK8sApiCoreV1ScaleIOPersistentVolumeSource

`func NewIoK8sApiCoreV1ScaleIOPersistentVolumeSource(gateway string, secretRef IoK8sApiCoreV1SecretReference, system string, ) *IoK8sApiCoreV1ScaleIOPersistentVolumeSource`

NewIoK8sApiCoreV1ScaleIOPersistentVolumeSource instantiates a new IoK8sApiCoreV1ScaleIOPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ScaleIOPersistentVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1ScaleIOPersistentVolumeSourceWithDefaults() *IoK8sApiCoreV1ScaleIOPersistentVolumeSource`

NewIoK8sApiCoreV1ScaleIOPersistentVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1ScaleIOPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsType

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetGateway

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetProtectionDomain

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetProtectionDomain() string`

GetProtectionDomain returns the ProtectionDomain field if non-nil, zero value otherwise.

### GetProtectionDomainOk

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetProtectionDomainOk() (*string, bool)`

GetProtectionDomainOk returns a tuple with the ProtectionDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDomain

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) SetProtectionDomain(v string)`

SetProtectionDomain sets ProtectionDomain field to given value.

### HasProtectionDomain

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) HasProtectionDomain() bool`

HasProtectionDomain returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSecretRef

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetSecretRef() IoK8sApiCoreV1SecretReference`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetSecretRefOk() (*IoK8sApiCoreV1SecretReference, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) SetSecretRef(v IoK8sApiCoreV1SecretReference)`

SetSecretRef sets SecretRef field to given value.


### GetSslEnabled

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### GetStorageMode

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetStorageMode() string`

GetStorageMode returns the StorageMode field if non-nil, zero value otherwise.

### GetStorageModeOk

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetStorageModeOk() (*string, bool)`

GetStorageModeOk returns a tuple with the StorageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageMode

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) SetStorageMode(v string)`

SetStorageMode sets StorageMode field to given value.

### HasStorageMode

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) HasStorageMode() bool`

HasStorageMode returns a boolean if a field has been set.

### GetStoragePool

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetStoragePool() string`

GetStoragePool returns the StoragePool field if non-nil, zero value otherwise.

### GetStoragePoolOk

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetStoragePoolOk() (*string, bool)`

GetStoragePoolOk returns a tuple with the StoragePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePool

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) SetStoragePool(v string)`

SetStoragePool sets StoragePool field to given value.

### HasStoragePool

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) HasStoragePool() bool`

HasStoragePool returns a boolean if a field has been set.

### GetSystem

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) SetSystem(v string)`

SetSystem sets System field to given value.


### GetVolumeName

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *IoK8sApiCoreV1ScaleIOPersistentVolumeSource) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


