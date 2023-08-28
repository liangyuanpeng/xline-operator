# IoK8sApimachineryPkgApisMetaV1Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | 
**Message** | **string** | message is a human readable message indicating details about the transition. This may be an empty string. | 
**ObservedGeneration** | Pointer to **int64** | observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance. | [optional] 
**Reason** | **string** | reason contains a programmatic identifier indicating the reason for the condition&#39;s last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty. | 
**Status** | **string** | status of the condition, one of True, False, Unknown. | 
**Type** | **string** | type of condition in CamelCase or in foo.example.com/CamelCase. | 

## Methods

### NewIoK8sApimachineryPkgApisMetaV1Condition

`func NewIoK8sApimachineryPkgApisMetaV1Condition(lastTransitionTime time.Time, message string, reason string, status string, type_ string, ) *IoK8sApimachineryPkgApisMetaV1Condition`

NewIoK8sApimachineryPkgApisMetaV1Condition instantiates a new IoK8sApimachineryPkgApisMetaV1Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApimachineryPkgApisMetaV1ConditionWithDefaults

`func NewIoK8sApimachineryPkgApisMetaV1ConditionWithDefaults() *IoK8sApimachineryPkgApisMetaV1Condition`

NewIoK8sApimachineryPkgApisMetaV1ConditionWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.


### GetMessage

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetObservedGeneration

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStatus

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApimachineryPkgApisMetaV1Condition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


