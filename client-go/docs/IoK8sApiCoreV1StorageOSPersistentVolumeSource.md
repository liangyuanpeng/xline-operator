# IoK8sApiCoreV1StorageOSPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | Pointer to **string** | fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. | [optional] 
**ReadOnly** | Pointer to **bool** | readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] 
**SecretRef** | Pointer to [**IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) |  | [optional] 
**VolumeName** | Pointer to **string** | volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace. | [optional] 
**VolumeNamespace** | Pointer to **string** | volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod&#39;s namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to \&quot;default\&quot; if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created. | [optional] 

## Methods

### NewIoK8sApiCoreV1StorageOSPersistentVolumeSource

`func NewIoK8sApiCoreV1StorageOSPersistentVolumeSource() *IoK8sApiCoreV1StorageOSPersistentVolumeSource`

NewIoK8sApiCoreV1StorageOSPersistentVolumeSource instantiates a new IoK8sApiCoreV1StorageOSPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1StorageOSPersistentVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1StorageOSPersistentVolumeSourceWithDefaults() *IoK8sApiCoreV1StorageOSPersistentVolumeSource`

NewIoK8sApiCoreV1StorageOSPersistentVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1StorageOSPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsType

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSecretRef

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) GetSecretRef() IoK8sApiCoreV1ObjectReference`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) GetSecretRefOk() (*IoK8sApiCoreV1ObjectReference, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) SetSecretRef(v IoK8sApiCoreV1ObjectReference)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.

### GetVolumeName

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeNamespace

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) GetVolumeNamespace() string`

GetVolumeNamespace returns the VolumeNamespace field if non-nil, zero value otherwise.

### GetVolumeNamespaceOk

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) GetVolumeNamespaceOk() (*string, bool)`

GetVolumeNamespaceOk returns a tuple with the VolumeNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeNamespace

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) SetVolumeNamespace(v string)`

SetVolumeNamespace sets VolumeNamespace field to given value.

### HasVolumeNamespace

`func (o *IoK8sApiCoreV1StorageOSPersistentVolumeSource) HasVolumeNamespace() bool`

HasVolumeNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


