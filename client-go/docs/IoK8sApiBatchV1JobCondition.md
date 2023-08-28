# IoK8sApiBatchV1JobCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbeTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**LastTransitionTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | Human readable message indicating details about last transition. | [optional] 
**Reason** | Pointer to **string** | (brief) reason for the condition&#39;s last transition. | [optional] 
**Status** | **string** | Status of the condition, one of True, False, Unknown. | 
**Type** | **string** | Type of job condition, Complete or Failed. | 

## Methods

### NewIoK8sApiBatchV1JobCondition

`func NewIoK8sApiBatchV1JobCondition(status string, type_ string, ) *IoK8sApiBatchV1JobCondition`

NewIoK8sApiBatchV1JobCondition instantiates a new IoK8sApiBatchV1JobCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiBatchV1JobConditionWithDefaults

`func NewIoK8sApiBatchV1JobConditionWithDefaults() *IoK8sApiBatchV1JobCondition`

NewIoK8sApiBatchV1JobConditionWithDefaults instantiates a new IoK8sApiBatchV1JobCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastProbeTime

`func (o *IoK8sApiBatchV1JobCondition) GetLastProbeTime() time.Time`

GetLastProbeTime returns the LastProbeTime field if non-nil, zero value otherwise.

### GetLastProbeTimeOk

`func (o *IoK8sApiBatchV1JobCondition) GetLastProbeTimeOk() (*time.Time, bool)`

GetLastProbeTimeOk returns a tuple with the LastProbeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProbeTime

`func (o *IoK8sApiBatchV1JobCondition) SetLastProbeTime(v time.Time)`

SetLastProbeTime sets LastProbeTime field to given value.

### HasLastProbeTime

`func (o *IoK8sApiBatchV1JobCondition) HasLastProbeTime() bool`

HasLastProbeTime returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *IoK8sApiBatchV1JobCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *IoK8sApiBatchV1JobCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *IoK8sApiBatchV1JobCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *IoK8sApiBatchV1JobCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiBatchV1JobCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiBatchV1JobCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiBatchV1JobCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiBatchV1JobCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiBatchV1JobCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiBatchV1JobCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiBatchV1JobCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiBatchV1JobCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiBatchV1JobCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiBatchV1JobCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiBatchV1JobCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *IoK8sApiBatchV1JobCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiBatchV1JobCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiBatchV1JobCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


