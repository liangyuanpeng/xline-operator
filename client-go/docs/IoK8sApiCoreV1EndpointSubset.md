# IoK8sApiCoreV1EndpointSubset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]IoK8sApiCoreV1EndpointAddress**](IoK8sApiCoreV1EndpointAddress.md) | IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize. | [optional] 
**NotReadyAddresses** | Pointer to [**[]IoK8sApiCoreV1EndpointAddress**](IoK8sApiCoreV1EndpointAddress.md) | IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check. | [optional] 
**Ports** | Pointer to [**[]IoK8sApiCoreV1EndpointPort**](IoK8sApiCoreV1EndpointPort.md) | Port numbers available on the related IP addresses. | [optional] 

## Methods

### NewIoK8sApiCoreV1EndpointSubset

`func NewIoK8sApiCoreV1EndpointSubset() *IoK8sApiCoreV1EndpointSubset`

NewIoK8sApiCoreV1EndpointSubset instantiates a new IoK8sApiCoreV1EndpointSubset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1EndpointSubsetWithDefaults

`func NewIoK8sApiCoreV1EndpointSubsetWithDefaults() *IoK8sApiCoreV1EndpointSubset`

NewIoK8sApiCoreV1EndpointSubsetWithDefaults instantiates a new IoK8sApiCoreV1EndpointSubset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *IoK8sApiCoreV1EndpointSubset) GetAddresses() []IoK8sApiCoreV1EndpointAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *IoK8sApiCoreV1EndpointSubset) GetAddressesOk() (*[]IoK8sApiCoreV1EndpointAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *IoK8sApiCoreV1EndpointSubset) SetAddresses(v []IoK8sApiCoreV1EndpointAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *IoK8sApiCoreV1EndpointSubset) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetNotReadyAddresses

`func (o *IoK8sApiCoreV1EndpointSubset) GetNotReadyAddresses() []IoK8sApiCoreV1EndpointAddress`

GetNotReadyAddresses returns the NotReadyAddresses field if non-nil, zero value otherwise.

### GetNotReadyAddressesOk

`func (o *IoK8sApiCoreV1EndpointSubset) GetNotReadyAddressesOk() (*[]IoK8sApiCoreV1EndpointAddress, bool)`

GetNotReadyAddressesOk returns a tuple with the NotReadyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotReadyAddresses

`func (o *IoK8sApiCoreV1EndpointSubset) SetNotReadyAddresses(v []IoK8sApiCoreV1EndpointAddress)`

SetNotReadyAddresses sets NotReadyAddresses field to given value.

### HasNotReadyAddresses

`func (o *IoK8sApiCoreV1EndpointSubset) HasNotReadyAddresses() bool`

HasNotReadyAddresses returns a boolean if a field has been set.

### GetPorts

`func (o *IoK8sApiCoreV1EndpointSubset) GetPorts() []IoK8sApiCoreV1EndpointPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *IoK8sApiCoreV1EndpointSubset) GetPortsOk() (*[]IoK8sApiCoreV1EndpointPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *IoK8sApiCoreV1EndpointSubset) SetPorts(v []IoK8sApiCoreV1EndpointPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *IoK8sApiCoreV1EndpointSubset) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


