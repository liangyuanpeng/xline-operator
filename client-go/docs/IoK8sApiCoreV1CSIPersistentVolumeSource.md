# IoK8sApiCoreV1CSIPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerExpandSecretRef** | Pointer to [**IoK8sApiCoreV1SecretReference**](IoK8sApiCoreV1SecretReference.md) |  | [optional] 
**ControllerPublishSecretRef** | Pointer to [**IoK8sApiCoreV1SecretReference**](IoK8sApiCoreV1SecretReference.md) |  | [optional] 
**Driver** | **string** | driver is the name of the driver to use for this volume. Required. | 
**FsType** | Pointer to **string** | fsType to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. | [optional] 
**NodeExpandSecretRef** | Pointer to [**IoK8sApiCoreV1SecretReference**](IoK8sApiCoreV1SecretReference.md) |  | [optional] 
**NodePublishSecretRef** | Pointer to [**IoK8sApiCoreV1SecretReference**](IoK8sApiCoreV1SecretReference.md) |  | [optional] 
**NodeStageSecretRef** | Pointer to [**IoK8sApiCoreV1SecretReference**](IoK8sApiCoreV1SecretReference.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** | readOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write). | [optional] 
**VolumeAttributes** | Pointer to **map[string]string** | volumeAttributes of the volume to publish. | [optional] 
**VolumeHandle** | **string** | volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required. | 

## Methods

### NewIoK8sApiCoreV1CSIPersistentVolumeSource

`func NewIoK8sApiCoreV1CSIPersistentVolumeSource(driver string, volumeHandle string, ) *IoK8sApiCoreV1CSIPersistentVolumeSource`

NewIoK8sApiCoreV1CSIPersistentVolumeSource instantiates a new IoK8sApiCoreV1CSIPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1CSIPersistentVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1CSIPersistentVolumeSourceWithDefaults() *IoK8sApiCoreV1CSIPersistentVolumeSource`

NewIoK8sApiCoreV1CSIPersistentVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1CSIPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerExpandSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetControllerExpandSecretRef() IoK8sApiCoreV1SecretReference`

GetControllerExpandSecretRef returns the ControllerExpandSecretRef field if non-nil, zero value otherwise.

### GetControllerExpandSecretRefOk

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetControllerExpandSecretRefOk() (*IoK8sApiCoreV1SecretReference, bool)`

GetControllerExpandSecretRefOk returns a tuple with the ControllerExpandSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerExpandSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) SetControllerExpandSecretRef(v IoK8sApiCoreV1SecretReference)`

SetControllerExpandSecretRef sets ControllerExpandSecretRef field to given value.

### HasControllerExpandSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) HasControllerExpandSecretRef() bool`

HasControllerExpandSecretRef returns a boolean if a field has been set.

### GetControllerPublishSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetControllerPublishSecretRef() IoK8sApiCoreV1SecretReference`

GetControllerPublishSecretRef returns the ControllerPublishSecretRef field if non-nil, zero value otherwise.

### GetControllerPublishSecretRefOk

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetControllerPublishSecretRefOk() (*IoK8sApiCoreV1SecretReference, bool)`

GetControllerPublishSecretRefOk returns a tuple with the ControllerPublishSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerPublishSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) SetControllerPublishSecretRef(v IoK8sApiCoreV1SecretReference)`

SetControllerPublishSecretRef sets ControllerPublishSecretRef field to given value.

### HasControllerPublishSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) HasControllerPublishSecretRef() bool`

HasControllerPublishSecretRef returns a boolean if a field has been set.

### GetDriver

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) SetDriver(v string)`

SetDriver sets Driver field to given value.


### GetFsType

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetNodeExpandSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetNodeExpandSecretRef() IoK8sApiCoreV1SecretReference`

GetNodeExpandSecretRef returns the NodeExpandSecretRef field if non-nil, zero value otherwise.

### GetNodeExpandSecretRefOk

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetNodeExpandSecretRefOk() (*IoK8sApiCoreV1SecretReference, bool)`

GetNodeExpandSecretRefOk returns a tuple with the NodeExpandSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExpandSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) SetNodeExpandSecretRef(v IoK8sApiCoreV1SecretReference)`

SetNodeExpandSecretRef sets NodeExpandSecretRef field to given value.

### HasNodeExpandSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) HasNodeExpandSecretRef() bool`

HasNodeExpandSecretRef returns a boolean if a field has been set.

### GetNodePublishSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetNodePublishSecretRef() IoK8sApiCoreV1SecretReference`

GetNodePublishSecretRef returns the NodePublishSecretRef field if non-nil, zero value otherwise.

### GetNodePublishSecretRefOk

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetNodePublishSecretRefOk() (*IoK8sApiCoreV1SecretReference, bool)`

GetNodePublishSecretRefOk returns a tuple with the NodePublishSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePublishSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) SetNodePublishSecretRef(v IoK8sApiCoreV1SecretReference)`

SetNodePublishSecretRef sets NodePublishSecretRef field to given value.

### HasNodePublishSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) HasNodePublishSecretRef() bool`

HasNodePublishSecretRef returns a boolean if a field has been set.

### GetNodeStageSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetNodeStageSecretRef() IoK8sApiCoreV1SecretReference`

GetNodeStageSecretRef returns the NodeStageSecretRef field if non-nil, zero value otherwise.

### GetNodeStageSecretRefOk

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetNodeStageSecretRefOk() (*IoK8sApiCoreV1SecretReference, bool)`

GetNodeStageSecretRefOk returns a tuple with the NodeStageSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStageSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) SetNodeStageSecretRef(v IoK8sApiCoreV1SecretReference)`

SetNodeStageSecretRef sets NodeStageSecretRef field to given value.

### HasNodeStageSecretRef

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) HasNodeStageSecretRef() bool`

HasNodeStageSecretRef returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetVolumeAttributes

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetVolumeAttributes() map[string]string`

GetVolumeAttributes returns the VolumeAttributes field if non-nil, zero value otherwise.

### GetVolumeAttributesOk

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetVolumeAttributesOk() (*map[string]string, bool)`

GetVolumeAttributesOk returns a tuple with the VolumeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAttributes

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) SetVolumeAttributes(v map[string]string)`

SetVolumeAttributes sets VolumeAttributes field to given value.

### HasVolumeAttributes

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) HasVolumeAttributes() bool`

HasVolumeAttributes returns a boolean if a field has been set.

### GetVolumeHandle

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetVolumeHandle() string`

GetVolumeHandle returns the VolumeHandle field if non-nil, zero value otherwise.

### GetVolumeHandleOk

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) GetVolumeHandleOk() (*string, bool)`

GetVolumeHandleOk returns a tuple with the VolumeHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeHandle

`func (o *IoK8sApiCoreV1CSIPersistentVolumeSource) SetVolumeHandle(v string)`

SetVolumeHandle sets VolumeHandle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


