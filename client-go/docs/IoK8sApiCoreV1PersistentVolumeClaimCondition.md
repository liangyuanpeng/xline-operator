# IoK8sApiCoreV1PersistentVolumeClaimCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbeTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**LastTransitionTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | message is the human-readable message indicating details about last transition. | [optional] 
**Reason** | Pointer to **string** | reason is a unique, this should be a short, machine understandable string that gives the reason for condition&#39;s last transition. If it reports \&quot;ResizeStarted\&quot; that means the underlying persistent volume is being resized. | [optional] 
**Status** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewIoK8sApiCoreV1PersistentVolumeClaimCondition

`func NewIoK8sApiCoreV1PersistentVolumeClaimCondition(status string, type_ string, ) *IoK8sApiCoreV1PersistentVolumeClaimCondition`

NewIoK8sApiCoreV1PersistentVolumeClaimCondition instantiates a new IoK8sApiCoreV1PersistentVolumeClaimCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PersistentVolumeClaimConditionWithDefaults

`func NewIoK8sApiCoreV1PersistentVolumeClaimConditionWithDefaults() *IoK8sApiCoreV1PersistentVolumeClaimCondition`

NewIoK8sApiCoreV1PersistentVolumeClaimConditionWithDefaults instantiates a new IoK8sApiCoreV1PersistentVolumeClaimCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastProbeTime

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetLastProbeTime() time.Time`

GetLastProbeTime returns the LastProbeTime field if non-nil, zero value otherwise.

### GetLastProbeTimeOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetLastProbeTimeOk() (*time.Time, bool)`

GetLastProbeTimeOk returns a tuple with the LastProbeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProbeTime

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) SetLastProbeTime(v time.Time)`

SetLastProbeTime sets LastProbeTime field to given value.

### HasLastProbeTime

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) HasLastProbeTime() bool`

HasLastProbeTime returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiCoreV1PersistentVolumeClaimCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


