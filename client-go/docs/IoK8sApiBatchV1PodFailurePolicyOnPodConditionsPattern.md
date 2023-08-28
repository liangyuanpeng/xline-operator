# IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Specifies the required Pod condition status. To match a pod condition it is required that the specified status equals the pod condition status. Defaults to True. | 
**Type** | **string** | Specifies the required Pod condition type. To match a pod condition it is required that specified type equals the pod condition type. | 

## Methods

### NewIoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern

`func NewIoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern(status string, type_ string, ) *IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern`

NewIoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern instantiates a new IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiBatchV1PodFailurePolicyOnPodConditionsPatternWithDefaults

`func NewIoK8sApiBatchV1PodFailurePolicyOnPodConditionsPatternWithDefaults() *IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern`

NewIoK8sApiBatchV1PodFailurePolicyOnPodConditionsPatternWithDefaults instantiates a new IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


