# IoK8sApiFlowcontrolV1beta2FlowSchemaCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | &#x60;message&#x60; is a human-readable message indicating details about last transition. | [optional] 
**Reason** | Pointer to **string** | &#x60;reason&#x60; is a unique, one-word, CamelCase reason for the condition&#39;s last transition. | [optional] 
**Status** | Pointer to **string** | &#x60;status&#x60; is the status of the condition. Can be True, False, Unknown. Required. | [optional] 
**Type** | Pointer to **string** | &#x60;type&#x60; is the type of the condition. Required. | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta2FlowSchemaCondition

`func NewIoK8sApiFlowcontrolV1beta2FlowSchemaCondition() *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition`

NewIoK8sApiFlowcontrolV1beta2FlowSchemaCondition instantiates a new IoK8sApiFlowcontrolV1beta2FlowSchemaCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta2FlowSchemaConditionWithDefaults

`func NewIoK8sApiFlowcontrolV1beta2FlowSchemaConditionWithDefaults() *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition`

NewIoK8sApiFlowcontrolV1beta2FlowSchemaConditionWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta2FlowSchemaCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiFlowcontrolV1beta2FlowSchemaCondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


