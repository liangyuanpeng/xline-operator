# IoK8sApiCoreV1CinderPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | Pointer to **string** | fsType Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md | [optional] 
**ReadOnly** | Pointer to **bool** | readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md | [optional] 
**SecretRef** | Pointer to [**IoK8sApiCoreV1SecretReference**](IoK8sApiCoreV1SecretReference.md) |  | [optional] 
**VolumeID** | **string** | volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md | 

## Methods

### NewIoK8sApiCoreV1CinderPersistentVolumeSource

`func NewIoK8sApiCoreV1CinderPersistentVolumeSource(volumeID string, ) *IoK8sApiCoreV1CinderPersistentVolumeSource`

NewIoK8sApiCoreV1CinderPersistentVolumeSource instantiates a new IoK8sApiCoreV1CinderPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1CinderPersistentVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1CinderPersistentVolumeSourceWithDefaults() *IoK8sApiCoreV1CinderPersistentVolumeSource`

NewIoK8sApiCoreV1CinderPersistentVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1CinderPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsType

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSecretRef

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) GetSecretRef() IoK8sApiCoreV1SecretReference`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) GetSecretRefOk() (*IoK8sApiCoreV1SecretReference, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) SetSecretRef(v IoK8sApiCoreV1SecretReference)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.

### GetVolumeID

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) GetVolumeID() string`

GetVolumeID returns the VolumeID field if non-nil, zero value otherwise.

### GetVolumeIDOk

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) GetVolumeIDOk() (*string, bool)`

GetVolumeIDOk returns a tuple with the VolumeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeID

`func (o *IoK8sApiCoreV1CinderPersistentVolumeSource) SetVolumeID(v string)`

SetVolumeID sets VolumeID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


