# IoK8sApiCoreV1RBDPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | Pointer to **string** | fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd | [optional] 
**Image** | **string** | image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | 
**Keyring** | Pointer to **string** | keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] 
**Monitors** | **[]string** | monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | 
**Pool** | Pointer to **string** | pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] 
**ReadOnly** | Pointer to **bool** | readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] 
**SecretRef** | Pointer to [**IoK8sApiCoreV1SecretReference**](IoK8sApiCoreV1SecretReference.md) |  | [optional] 
**User** | Pointer to **string** | user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] 

## Methods

### NewIoK8sApiCoreV1RBDPersistentVolumeSource

`func NewIoK8sApiCoreV1RBDPersistentVolumeSource(image string, monitors []string, ) *IoK8sApiCoreV1RBDPersistentVolumeSource`

NewIoK8sApiCoreV1RBDPersistentVolumeSource instantiates a new IoK8sApiCoreV1RBDPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1RBDPersistentVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1RBDPersistentVolumeSourceWithDefaults() *IoK8sApiCoreV1RBDPersistentVolumeSource`

NewIoK8sApiCoreV1RBDPersistentVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1RBDPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsType

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetImage

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) SetImage(v string)`

SetImage sets Image field to given value.


### GetKeyring

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetKeyring() string`

GetKeyring returns the Keyring field if non-nil, zero value otherwise.

### GetKeyringOk

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetKeyringOk() (*string, bool)`

GetKeyringOk returns a tuple with the Keyring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyring

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) SetKeyring(v string)`

SetKeyring sets Keyring field to given value.

### HasKeyring

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) HasKeyring() bool`

HasKeyring returns a boolean if a field has been set.

### GetMonitors

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetMonitors() []string`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetMonitorsOk() (*[]string, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) SetMonitors(v []string)`

SetMonitors sets Monitors field to given value.


### GetPool

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSecretRef

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetSecretRef() IoK8sApiCoreV1SecretReference`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetSecretRefOk() (*IoK8sApiCoreV1SecretReference, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) SetSecretRef(v IoK8sApiCoreV1SecretReference)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.

### GetUser

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IoK8sApiCoreV1RBDPersistentVolumeSource) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


