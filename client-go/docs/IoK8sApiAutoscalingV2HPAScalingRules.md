# IoK8sApiAutoscalingV2HPAScalingRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]IoK8sApiAutoscalingV2HPAScalingPolicy**](IoK8sApiAutoscalingV2HPAScalingPolicy.md) | policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid | [optional] 
**SelectPolicy** | Pointer to **string** | selectPolicy is used to specify which policy should be used. If not set, the default value Max is used. | [optional] 
**StabilizationWindowSeconds** | Pointer to **int32** | stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long). | [optional] 

## Methods

### NewIoK8sApiAutoscalingV2HPAScalingRules

`func NewIoK8sApiAutoscalingV2HPAScalingRules() *IoK8sApiAutoscalingV2HPAScalingRules`

NewIoK8sApiAutoscalingV2HPAScalingRules instantiates a new IoK8sApiAutoscalingV2HPAScalingRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2HPAScalingRulesWithDefaults

`func NewIoK8sApiAutoscalingV2HPAScalingRulesWithDefaults() *IoK8sApiAutoscalingV2HPAScalingRules`

NewIoK8sApiAutoscalingV2HPAScalingRulesWithDefaults instantiates a new IoK8sApiAutoscalingV2HPAScalingRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) GetPolicies() []IoK8sApiAutoscalingV2HPAScalingPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) GetPoliciesOk() (*[]IoK8sApiAutoscalingV2HPAScalingPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) SetPolicies(v []IoK8sApiAutoscalingV2HPAScalingPolicy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetSelectPolicy

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) GetSelectPolicy() string`

GetSelectPolicy returns the SelectPolicy field if non-nil, zero value otherwise.

### GetSelectPolicyOk

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) GetSelectPolicyOk() (*string, bool)`

GetSelectPolicyOk returns a tuple with the SelectPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectPolicy

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) SetSelectPolicy(v string)`

SetSelectPolicy sets SelectPolicy field to given value.

### HasSelectPolicy

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) HasSelectPolicy() bool`

HasSelectPolicy returns a boolean if a field has been set.

### GetStabilizationWindowSeconds

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) GetStabilizationWindowSeconds() int32`

GetStabilizationWindowSeconds returns the StabilizationWindowSeconds field if non-nil, zero value otherwise.

### GetStabilizationWindowSecondsOk

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) GetStabilizationWindowSecondsOk() (*int32, bool)`

GetStabilizationWindowSecondsOk returns a tuple with the StabilizationWindowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStabilizationWindowSeconds

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) SetStabilizationWindowSeconds(v int32)`

SetStabilizationWindowSeconds sets StabilizationWindowSeconds field to given value.

### HasStabilizationWindowSeconds

`func (o *IoK8sApiAutoscalingV2HPAScalingRules) HasStabilizationWindowSeconds() bool`

HasStabilizationWindowSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


