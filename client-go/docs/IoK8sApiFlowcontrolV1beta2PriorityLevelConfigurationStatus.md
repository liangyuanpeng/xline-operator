# IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition**](IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition.md) | &#x60;conditions&#x60; is the current state of \&quot;request-priority\&quot;. | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus

`func NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus() *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus`

NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus instantiates a new IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatusWithDefaults

`func NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatusWithDefaults() *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus`

NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatusWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus) GetConditions() []IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus) GetConditionsOk() (*[]IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus) SetConditions(v []IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


