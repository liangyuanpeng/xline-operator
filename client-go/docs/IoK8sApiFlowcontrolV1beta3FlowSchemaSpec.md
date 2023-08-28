# IoK8sApiFlowcontrolV1beta3FlowSchemaSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinguisherMethod** | Pointer to [**IoK8sApiFlowcontrolV1beta3FlowDistinguisherMethod**](IoK8sApiFlowcontrolV1beta3FlowDistinguisherMethod.md) |  | [optional] 
**MatchingPrecedence** | Pointer to **int32** | &#x60;matchingPrecedence&#x60; is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default. | [optional] 
**PriorityLevelConfiguration** | [**IoK8sApiFlowcontrolV1beta3PriorityLevelConfigurationReference**](IoK8sApiFlowcontrolV1beta3PriorityLevelConfigurationReference.md) |  | 
**Rules** | Pointer to [**[]IoK8sApiFlowcontrolV1beta3PolicyRulesWithSubjects**](IoK8sApiFlowcontrolV1beta3PolicyRulesWithSubjects.md) | &#x60;rules&#x60; describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema. | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta3FlowSchemaSpec

`func NewIoK8sApiFlowcontrolV1beta3FlowSchemaSpec(priorityLevelConfiguration IoK8sApiFlowcontrolV1beta3PriorityLevelConfigurationReference, ) *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec`

NewIoK8sApiFlowcontrolV1beta3FlowSchemaSpec instantiates a new IoK8sApiFlowcontrolV1beta3FlowSchemaSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta3FlowSchemaSpecWithDefaults

`func NewIoK8sApiFlowcontrolV1beta3FlowSchemaSpecWithDefaults() *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec`

NewIoK8sApiFlowcontrolV1beta3FlowSchemaSpecWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta3FlowSchemaSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinguisherMethod

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) GetDistinguisherMethod() IoK8sApiFlowcontrolV1beta3FlowDistinguisherMethod`

GetDistinguisherMethod returns the DistinguisherMethod field if non-nil, zero value otherwise.

### GetDistinguisherMethodOk

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) GetDistinguisherMethodOk() (*IoK8sApiFlowcontrolV1beta3FlowDistinguisherMethod, bool)`

GetDistinguisherMethodOk returns a tuple with the DistinguisherMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguisherMethod

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) SetDistinguisherMethod(v IoK8sApiFlowcontrolV1beta3FlowDistinguisherMethod)`

SetDistinguisherMethod sets DistinguisherMethod field to given value.

### HasDistinguisherMethod

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) HasDistinguisherMethod() bool`

HasDistinguisherMethod returns a boolean if a field has been set.

### GetMatchingPrecedence

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) GetMatchingPrecedence() int32`

GetMatchingPrecedence returns the MatchingPrecedence field if non-nil, zero value otherwise.

### GetMatchingPrecedenceOk

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) GetMatchingPrecedenceOk() (*int32, bool)`

GetMatchingPrecedenceOk returns a tuple with the MatchingPrecedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingPrecedence

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) SetMatchingPrecedence(v int32)`

SetMatchingPrecedence sets MatchingPrecedence field to given value.

### HasMatchingPrecedence

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) HasMatchingPrecedence() bool`

HasMatchingPrecedence returns a boolean if a field has been set.

### GetPriorityLevelConfiguration

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) GetPriorityLevelConfiguration() IoK8sApiFlowcontrolV1beta3PriorityLevelConfigurationReference`

GetPriorityLevelConfiguration returns the PriorityLevelConfiguration field if non-nil, zero value otherwise.

### GetPriorityLevelConfigurationOk

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) GetPriorityLevelConfigurationOk() (*IoK8sApiFlowcontrolV1beta3PriorityLevelConfigurationReference, bool)`

GetPriorityLevelConfigurationOk returns a tuple with the PriorityLevelConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevelConfiguration

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) SetPriorityLevelConfiguration(v IoK8sApiFlowcontrolV1beta3PriorityLevelConfigurationReference)`

SetPriorityLevelConfiguration sets PriorityLevelConfiguration field to given value.


### GetRules

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) GetRules() []IoK8sApiFlowcontrolV1beta3PolicyRulesWithSubjects`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) GetRulesOk() (*[]IoK8sApiFlowcontrolV1beta3PolicyRulesWithSubjects, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) SetRules(v []IoK8sApiFlowcontrolV1beta3PolicyRulesWithSubjects)`

SetRules sets Rules field to given value.

### HasRules

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaSpec) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


