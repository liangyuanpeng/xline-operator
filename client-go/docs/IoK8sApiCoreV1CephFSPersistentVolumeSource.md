# IoK8sApiCoreV1CephFSPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitors** | **[]string** | monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | 
**Path** | Pointer to **string** | path is Optional: Used as the mounted root, rather than the full Ceph tree, default is / | [optional] 
**ReadOnly** | Pointer to **bool** | readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | [optional] 
**SecretFile** | Pointer to **string** | secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | [optional] 
**SecretRef** | Pointer to [**IoK8sApiCoreV1SecretReference**](IoK8sApiCoreV1SecretReference.md) |  | [optional] 
**User** | Pointer to **string** | user is Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | [optional] 

## Methods

### NewIoK8sApiCoreV1CephFSPersistentVolumeSource

`func NewIoK8sApiCoreV1CephFSPersistentVolumeSource(monitors []string, ) *IoK8sApiCoreV1CephFSPersistentVolumeSource`

NewIoK8sApiCoreV1CephFSPersistentVolumeSource instantiates a new IoK8sApiCoreV1CephFSPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1CephFSPersistentVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1CephFSPersistentVolumeSourceWithDefaults() *IoK8sApiCoreV1CephFSPersistentVolumeSource`

NewIoK8sApiCoreV1CephFSPersistentVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1CephFSPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitors

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetMonitors() []string`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetMonitorsOk() (*[]string, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) SetMonitors(v []string)`

SetMonitors sets Monitors field to given value.


### GetPath

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSecretFile

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetSecretFile() string`

GetSecretFile returns the SecretFile field if non-nil, zero value otherwise.

### GetSecretFileOk

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetSecretFileOk() (*string, bool)`

GetSecretFileOk returns a tuple with the SecretFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFile

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) SetSecretFile(v string)`

SetSecretFile sets SecretFile field to given value.

### HasSecretFile

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) HasSecretFile() bool`

HasSecretFile returns a boolean if a field has been set.

### GetSecretRef

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetSecretRef() IoK8sApiCoreV1SecretReference`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetSecretRefOk() (*IoK8sApiCoreV1SecretReference, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) SetSecretRef(v IoK8sApiCoreV1SecretReference)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.

### GetUser

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IoK8sApiCoreV1CephFSPersistentVolumeSource) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


