# IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | &#x60;message&#x60; is a human-readable message indicating details about last transition. | [optional] 
**Reason** | Pointer to **string** | &#x60;reason&#x60; is a unique, one-word, CamelCase reason for the condition&#39;s last transition. | [optional] 
**Status** | Pointer to **string** | &#x60;status&#x60; is the status of the condition. Can be True, False, Unknown. Required. | [optional] 
**Type** | Pointer to **string** | &#x60;type&#x60; is the type of the condition. Required. | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition

`func NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition() *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition`

NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition instantiates a new IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationConditionWithDefaults

`func NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationConditionWithDefaults() *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition`

NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationConditionWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


