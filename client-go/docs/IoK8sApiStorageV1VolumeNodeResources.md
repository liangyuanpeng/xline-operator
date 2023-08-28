# IoK8sApiStorageV1VolumeNodeResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | count indicates the maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared among multiple pods on the same node. If this field is not specified, then the supported number of volumes on this node is unbounded. | [optional] 

## Methods

### NewIoK8sApiStorageV1VolumeNodeResources

`func NewIoK8sApiStorageV1VolumeNodeResources() *IoK8sApiStorageV1VolumeNodeResources`

NewIoK8sApiStorageV1VolumeNodeResources instantiates a new IoK8sApiStorageV1VolumeNodeResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiStorageV1VolumeNodeResourcesWithDefaults

`func NewIoK8sApiStorageV1VolumeNodeResourcesWithDefaults() *IoK8sApiStorageV1VolumeNodeResources`

NewIoK8sApiStorageV1VolumeNodeResourcesWithDefaults instantiates a new IoK8sApiStorageV1VolumeNodeResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IoK8sApiStorageV1VolumeNodeResources) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IoK8sApiStorageV1VolumeNodeResources) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IoK8sApiStorageV1VolumeNodeResources) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IoK8sApiStorageV1VolumeNodeResources) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


