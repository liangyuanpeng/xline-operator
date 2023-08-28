# IoK8sApiCoordinationV1LeaseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquireTime** | Pointer to **time.Time** | MicroTime is version of Time with microsecond level precision. | [optional] 
**HolderIdentity** | Pointer to **string** | holderIdentity contains the identity of the holder of a current lease. | [optional] 
**LeaseDurationSeconds** | Pointer to **int32** | leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed renewTime. | [optional] 
**LeaseTransitions** | Pointer to **int32** | leaseTransitions is the number of transitions of a lease between holders. | [optional] 
**RenewTime** | Pointer to **time.Time** | MicroTime is version of Time with microsecond level precision. | [optional] 

## Methods

### NewIoK8sApiCoordinationV1LeaseSpec

`func NewIoK8sApiCoordinationV1LeaseSpec() *IoK8sApiCoordinationV1LeaseSpec`

NewIoK8sApiCoordinationV1LeaseSpec instantiates a new IoK8sApiCoordinationV1LeaseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoordinationV1LeaseSpecWithDefaults

`func NewIoK8sApiCoordinationV1LeaseSpecWithDefaults() *IoK8sApiCoordinationV1LeaseSpec`

NewIoK8sApiCoordinationV1LeaseSpecWithDefaults instantiates a new IoK8sApiCoordinationV1LeaseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquireTime

`func (o *IoK8sApiCoordinationV1LeaseSpec) GetAcquireTime() time.Time`

GetAcquireTime returns the AcquireTime field if non-nil, zero value otherwise.

### GetAcquireTimeOk

`func (o *IoK8sApiCoordinationV1LeaseSpec) GetAcquireTimeOk() (*time.Time, bool)`

GetAcquireTimeOk returns a tuple with the AcquireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquireTime

`func (o *IoK8sApiCoordinationV1LeaseSpec) SetAcquireTime(v time.Time)`

SetAcquireTime sets AcquireTime field to given value.

### HasAcquireTime

`func (o *IoK8sApiCoordinationV1LeaseSpec) HasAcquireTime() bool`

HasAcquireTime returns a boolean if a field has been set.

### GetHolderIdentity

`func (o *IoK8sApiCoordinationV1LeaseSpec) GetHolderIdentity() string`

GetHolderIdentity returns the HolderIdentity field if non-nil, zero value otherwise.

### GetHolderIdentityOk

`func (o *IoK8sApiCoordinationV1LeaseSpec) GetHolderIdentityOk() (*string, bool)`

GetHolderIdentityOk returns a tuple with the HolderIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderIdentity

`func (o *IoK8sApiCoordinationV1LeaseSpec) SetHolderIdentity(v string)`

SetHolderIdentity sets HolderIdentity field to given value.

### HasHolderIdentity

`func (o *IoK8sApiCoordinationV1LeaseSpec) HasHolderIdentity() bool`

HasHolderIdentity returns a boolean if a field has been set.

### GetLeaseDurationSeconds

`func (o *IoK8sApiCoordinationV1LeaseSpec) GetLeaseDurationSeconds() int32`

GetLeaseDurationSeconds returns the LeaseDurationSeconds field if non-nil, zero value otherwise.

### GetLeaseDurationSecondsOk

`func (o *IoK8sApiCoordinationV1LeaseSpec) GetLeaseDurationSecondsOk() (*int32, bool)`

GetLeaseDurationSecondsOk returns a tuple with the LeaseDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDurationSeconds

`func (o *IoK8sApiCoordinationV1LeaseSpec) SetLeaseDurationSeconds(v int32)`

SetLeaseDurationSeconds sets LeaseDurationSeconds field to given value.

### HasLeaseDurationSeconds

`func (o *IoK8sApiCoordinationV1LeaseSpec) HasLeaseDurationSeconds() bool`

HasLeaseDurationSeconds returns a boolean if a field has been set.

### GetLeaseTransitions

`func (o *IoK8sApiCoordinationV1LeaseSpec) GetLeaseTransitions() int32`

GetLeaseTransitions returns the LeaseTransitions field if non-nil, zero value otherwise.

### GetLeaseTransitionsOk

`func (o *IoK8sApiCoordinationV1LeaseSpec) GetLeaseTransitionsOk() (*int32, bool)`

GetLeaseTransitionsOk returns a tuple with the LeaseTransitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTransitions

`func (o *IoK8sApiCoordinationV1LeaseSpec) SetLeaseTransitions(v int32)`

SetLeaseTransitions sets LeaseTransitions field to given value.

### HasLeaseTransitions

`func (o *IoK8sApiCoordinationV1LeaseSpec) HasLeaseTransitions() bool`

HasLeaseTransitions returns a boolean if a field has been set.

### GetRenewTime

`func (o *IoK8sApiCoordinationV1LeaseSpec) GetRenewTime() time.Time`

GetRenewTime returns the RenewTime field if non-nil, zero value otherwise.

### GetRenewTimeOk

`func (o *IoK8sApiCoordinationV1LeaseSpec) GetRenewTimeOk() (*time.Time, bool)`

GetRenewTimeOk returns a tuple with the RenewTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewTime

`func (o *IoK8sApiCoordinationV1LeaseSpec) SetRenewTime(v time.Time)`

SetRenewTime sets RenewTime field to given value.

### HasRenewTime

`func (o *IoK8sApiCoordinationV1LeaseSpec) HasRenewTime() bool`

HasRenewTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


