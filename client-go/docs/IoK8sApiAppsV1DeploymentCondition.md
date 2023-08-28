# IoK8sApiAppsV1DeploymentCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**LastUpdateTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | A human readable message indicating details about the transition. | [optional] 
**Reason** | Pointer to **string** | The reason for the condition&#39;s last transition. | [optional] 
**Status** | **string** | Status of the condition, one of True, False, Unknown. | 
**Type** | **string** | Type of deployment condition. | 

## Methods

### NewIoK8sApiAppsV1DeploymentCondition

`func NewIoK8sApiAppsV1DeploymentCondition(status string, type_ string, ) *IoK8sApiAppsV1DeploymentCondition`

NewIoK8sApiAppsV1DeploymentCondition instantiates a new IoK8sApiAppsV1DeploymentCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1DeploymentConditionWithDefaults

`func NewIoK8sApiAppsV1DeploymentConditionWithDefaults() *IoK8sApiAppsV1DeploymentCondition`

NewIoK8sApiAppsV1DeploymentConditionWithDefaults instantiates a new IoK8sApiAppsV1DeploymentCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *IoK8sApiAppsV1DeploymentCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *IoK8sApiAppsV1DeploymentCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *IoK8sApiAppsV1DeploymentCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *IoK8sApiAppsV1DeploymentCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *IoK8sApiAppsV1DeploymentCondition) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *IoK8sApiAppsV1DeploymentCondition) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *IoK8sApiAppsV1DeploymentCondition) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *IoK8sApiAppsV1DeploymentCondition) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiAppsV1DeploymentCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiAppsV1DeploymentCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiAppsV1DeploymentCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiAppsV1DeploymentCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiAppsV1DeploymentCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiAppsV1DeploymentCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiAppsV1DeploymentCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiAppsV1DeploymentCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiAppsV1DeploymentCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiAppsV1DeploymentCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiAppsV1DeploymentCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *IoK8sApiAppsV1DeploymentCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiAppsV1DeploymentCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiAppsV1DeploymentCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


