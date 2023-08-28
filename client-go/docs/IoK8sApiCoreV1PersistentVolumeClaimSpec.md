# IoK8sApiCoreV1PersistentVolumeClaimSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** | accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**DataSource** | Pointer to [**IoK8sApiCoreV1TypedLocalObjectReference**](IoK8sApiCoreV1TypedLocalObjectReference.md) |  | [optional] 
**DataSourceRef** | Pointer to [**IoK8sApiCoreV1TypedObjectReference**](IoK8sApiCoreV1TypedObjectReference.md) |  | [optional] 
**Resources** | Pointer to [**IoK8sApiCoreV1ResourceRequirements**](IoK8sApiCoreV1ResourceRequirements.md) |  | [optional] 
**Selector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**StorageClassName** | Pointer to **string** | storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 | [optional] 
**VolumeMode** | Pointer to **string** | volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.  Possible enum values:  - &#x60;\&quot;Block\&quot;&#x60; means the volume will not be formatted with a filesystem and will remain a raw block device.  - &#x60;\&quot;Filesystem\&quot;&#x60; means the volume will be or is formatted with a filesystem. | [optional] 
**VolumeName** | Pointer to **string** | volumeName is the binding reference to the PersistentVolume backing this claim. | [optional] 

## Methods

### NewIoK8sApiCoreV1PersistentVolumeClaimSpec

`func NewIoK8sApiCoreV1PersistentVolumeClaimSpec() *IoK8sApiCoreV1PersistentVolumeClaimSpec`

NewIoK8sApiCoreV1PersistentVolumeClaimSpec instantiates a new IoK8sApiCoreV1PersistentVolumeClaimSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PersistentVolumeClaimSpecWithDefaults

`func NewIoK8sApiCoreV1PersistentVolumeClaimSpecWithDefaults() *IoK8sApiCoreV1PersistentVolumeClaimSpec`

NewIoK8sApiCoreV1PersistentVolumeClaimSpecWithDefaults instantiates a new IoK8sApiCoreV1PersistentVolumeClaimSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetDataSource

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetDataSource() IoK8sApiCoreV1TypedLocalObjectReference`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetDataSourceOk() (*IoK8sApiCoreV1TypedLocalObjectReference, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) SetDataSource(v IoK8sApiCoreV1TypedLocalObjectReference)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataSourceRef

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetDataSourceRef() IoK8sApiCoreV1TypedObjectReference`

GetDataSourceRef returns the DataSourceRef field if non-nil, zero value otherwise.

### GetDataSourceRefOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetDataSourceRefOk() (*IoK8sApiCoreV1TypedObjectReference, bool)`

GetDataSourceRefOk returns a tuple with the DataSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRef

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) SetDataSourceRef(v IoK8sApiCoreV1TypedObjectReference)`

SetDataSourceRef sets DataSourceRef field to given value.

### HasDataSourceRef

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) HasDataSourceRef() bool`

HasDataSourceRef returns a boolean if a field has been set.

### GetResources

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetResources() IoK8sApiCoreV1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetResourcesOk() (*IoK8sApiCoreV1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) SetResources(v IoK8sApiCoreV1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSelector

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) SetSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetStorageClassName

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### GetVolumeMode

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetVolumeMode() string`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetVolumeModeOk() (*string, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) SetVolumeMode(v string)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.

### GetVolumeName

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *IoK8sApiCoreV1PersistentVolumeClaimSpec) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


