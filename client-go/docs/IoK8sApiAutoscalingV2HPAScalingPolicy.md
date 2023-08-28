# IoK8sApiAutoscalingV2HPAScalingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodSeconds** | **int32** | periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min). | 
**Type** | **string** | type is used to specify the scaling policy. | 
**Value** | **int32** | value contains the amount of change which is permitted by the policy. It must be greater than zero | 

## Methods

### NewIoK8sApiAutoscalingV2HPAScalingPolicy

`func NewIoK8sApiAutoscalingV2HPAScalingPolicy(periodSeconds int32, type_ string, value int32, ) *IoK8sApiAutoscalingV2HPAScalingPolicy`

NewIoK8sApiAutoscalingV2HPAScalingPolicy instantiates a new IoK8sApiAutoscalingV2HPAScalingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2HPAScalingPolicyWithDefaults

`func NewIoK8sApiAutoscalingV2HPAScalingPolicyWithDefaults() *IoK8sApiAutoscalingV2HPAScalingPolicy`

NewIoK8sApiAutoscalingV2HPAScalingPolicyWithDefaults instantiates a new IoK8sApiAutoscalingV2HPAScalingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodSeconds

`func (o *IoK8sApiAutoscalingV2HPAScalingPolicy) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *IoK8sApiAutoscalingV2HPAScalingPolicy) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *IoK8sApiAutoscalingV2HPAScalingPolicy) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.


### GetType

`func (o *IoK8sApiAutoscalingV2HPAScalingPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiAutoscalingV2HPAScalingPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiAutoscalingV2HPAScalingPolicy) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *IoK8sApiAutoscalingV2HPAScalingPolicy) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IoK8sApiAutoscalingV2HPAScalingPolicy) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IoK8sApiAutoscalingV2HPAScalingPolicy) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


