# IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | message is a human-readable explanation containing details about the transition | [optional] 
**Reason** | Pointer to **string** | reason is the reason for the condition&#39;s last transition. | [optional] 
**Status** | **string** | status is the status of the condition (True, False, Unknown) | 
**Type** | **string** | type describes the current condition | 

## Methods

### NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition

`func NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition(status string, type_ string, ) *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition`

NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition instantiates a new IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerConditionWithDefaults

`func NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerConditionWithDefaults() *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition`

NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerConditionWithDefaults instantiates a new IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


