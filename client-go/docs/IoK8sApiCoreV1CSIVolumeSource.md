# IoK8sApiCoreV1CSIVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Driver** | **string** | driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster. | 
**FsType** | Pointer to **string** | fsType to mount. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply. | [optional] 
**NodePublishSecretRef** | Pointer to [**IoK8sApiCoreV1LocalObjectReference**](IoK8sApiCoreV1LocalObjectReference.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** | readOnly specifies a read-only configuration for the volume. Defaults to false (read/write). | [optional] 
**VolumeAttributes** | Pointer to **map[string]string** | volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver&#39;s documentation for supported values. | [optional] 

## Methods

### NewIoK8sApiCoreV1CSIVolumeSource

`func NewIoK8sApiCoreV1CSIVolumeSource(driver string, ) *IoK8sApiCoreV1CSIVolumeSource`

NewIoK8sApiCoreV1CSIVolumeSource instantiates a new IoK8sApiCoreV1CSIVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1CSIVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1CSIVolumeSourceWithDefaults() *IoK8sApiCoreV1CSIVolumeSource`

NewIoK8sApiCoreV1CSIVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1CSIVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriver

`func (o *IoK8sApiCoreV1CSIVolumeSource) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *IoK8sApiCoreV1CSIVolumeSource) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *IoK8sApiCoreV1CSIVolumeSource) SetDriver(v string)`

SetDriver sets Driver field to given value.


### GetFsType

`func (o *IoK8sApiCoreV1CSIVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1CSIVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1CSIVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1CSIVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetNodePublishSecretRef

`func (o *IoK8sApiCoreV1CSIVolumeSource) GetNodePublishSecretRef() IoK8sApiCoreV1LocalObjectReference`

GetNodePublishSecretRef returns the NodePublishSecretRef field if non-nil, zero value otherwise.

### GetNodePublishSecretRefOk

`func (o *IoK8sApiCoreV1CSIVolumeSource) GetNodePublishSecretRefOk() (*IoK8sApiCoreV1LocalObjectReference, bool)`

GetNodePublishSecretRefOk returns a tuple with the NodePublishSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePublishSecretRef

`func (o *IoK8sApiCoreV1CSIVolumeSource) SetNodePublishSecretRef(v IoK8sApiCoreV1LocalObjectReference)`

SetNodePublishSecretRef sets NodePublishSecretRef field to given value.

### HasNodePublishSecretRef

`func (o *IoK8sApiCoreV1CSIVolumeSource) HasNodePublishSecretRef() bool`

HasNodePublishSecretRef returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1CSIVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1CSIVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1CSIVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1CSIVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetVolumeAttributes

`func (o *IoK8sApiCoreV1CSIVolumeSource) GetVolumeAttributes() map[string]string`

GetVolumeAttributes returns the VolumeAttributes field if non-nil, zero value otherwise.

### GetVolumeAttributesOk

`func (o *IoK8sApiCoreV1CSIVolumeSource) GetVolumeAttributesOk() (*map[string]string, bool)`

GetVolumeAttributesOk returns a tuple with the VolumeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAttributes

`func (o *IoK8sApiCoreV1CSIVolumeSource) SetVolumeAttributes(v map[string]string)`

SetVolumeAttributes sets VolumeAttributes field to given value.

### HasVolumeAttributes

`func (o *IoK8sApiCoreV1CSIVolumeSource) HasVolumeAttributes() bool`

HasVolumeAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


