# IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | Human-readable message indicating details about last transition. | [optional] 
**Reason** | Pointer to **string** | Unique, one-word, CamelCase reason for the condition&#39;s last transition. | [optional] 
**Status** | **string** | Status is the status of the condition. Can be True, False, Unknown. | 
**Type** | **string** | Type is the type of the condition. | 

## Methods

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition(status string, type_ string, ) *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceConditionWithDefaults

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceConditionWithDefaults() *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceConditionWithDefaults instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


