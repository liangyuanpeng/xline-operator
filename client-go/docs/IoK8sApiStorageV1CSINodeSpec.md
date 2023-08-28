# IoK8sApiStorageV1CSINodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drivers** | [**[]IoK8sApiStorageV1CSINodeDriver**](IoK8sApiStorageV1CSINodeDriver.md) | drivers is a list of information of all CSI Drivers existing on a node. If all drivers in the list are uninstalled, this can become empty. | 

## Methods

### NewIoK8sApiStorageV1CSINodeSpec

`func NewIoK8sApiStorageV1CSINodeSpec(drivers []IoK8sApiStorageV1CSINodeDriver, ) *IoK8sApiStorageV1CSINodeSpec`

NewIoK8sApiStorageV1CSINodeSpec instantiates a new IoK8sApiStorageV1CSINodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiStorageV1CSINodeSpecWithDefaults

`func NewIoK8sApiStorageV1CSINodeSpecWithDefaults() *IoK8sApiStorageV1CSINodeSpec`

NewIoK8sApiStorageV1CSINodeSpecWithDefaults instantiates a new IoK8sApiStorageV1CSINodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrivers

`func (o *IoK8sApiStorageV1CSINodeSpec) GetDrivers() []IoK8sApiStorageV1CSINodeDriver`

GetDrivers returns the Drivers field if non-nil, zero value otherwise.

### GetDriversOk

`func (o *IoK8sApiStorageV1CSINodeSpec) GetDriversOk() (*[]IoK8sApiStorageV1CSINodeDriver, bool)`

GetDriversOk returns a tuple with the Drivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivers

`func (o *IoK8sApiStorageV1CSINodeSpec) SetDrivers(v []IoK8sApiStorageV1CSINodeDriver)`

SetDrivers sets Drivers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


