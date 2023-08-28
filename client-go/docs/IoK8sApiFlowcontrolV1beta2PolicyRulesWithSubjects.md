# IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NonResourceRules** | Pointer to [**[]IoK8sApiFlowcontrolV1beta2NonResourcePolicyRule**](IoK8sApiFlowcontrolV1beta2NonResourcePolicyRule.md) | &#x60;nonResourceRules&#x60; is a list of NonResourcePolicyRules that identify matching requests according to their verb and the target non-resource URL. | [optional] 
**ResourceRules** | Pointer to [**[]IoK8sApiFlowcontrolV1beta2ResourcePolicyRule**](IoK8sApiFlowcontrolV1beta2ResourcePolicyRule.md) | &#x60;resourceRules&#x60; is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target resource. At least one of &#x60;resourceRules&#x60; and &#x60;nonResourceRules&#x60; has to be non-empty. | [optional] 
**Subjects** | [**[]IoK8sApiFlowcontrolV1beta2Subject**](IoK8sApiFlowcontrolV1beta2Subject.md) | subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches every request. Required. | 

## Methods

### NewIoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects

`func NewIoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects(subjects []IoK8sApiFlowcontrolV1beta2Subject, ) *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects`

NewIoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects instantiates a new IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjectsWithDefaults

`func NewIoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjectsWithDefaults() *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects`

NewIoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjectsWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonResourceRules

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) GetNonResourceRules() []IoK8sApiFlowcontrolV1beta2NonResourcePolicyRule`

GetNonResourceRules returns the NonResourceRules field if non-nil, zero value otherwise.

### GetNonResourceRulesOk

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) GetNonResourceRulesOk() (*[]IoK8sApiFlowcontrolV1beta2NonResourcePolicyRule, bool)`

GetNonResourceRulesOk returns a tuple with the NonResourceRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceRules

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) SetNonResourceRules(v []IoK8sApiFlowcontrolV1beta2NonResourcePolicyRule)`

SetNonResourceRules sets NonResourceRules field to given value.

### HasNonResourceRules

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) HasNonResourceRules() bool`

HasNonResourceRules returns a boolean if a field has been set.

### GetResourceRules

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) GetResourceRules() []IoK8sApiFlowcontrolV1beta2ResourcePolicyRule`

GetResourceRules returns the ResourceRules field if non-nil, zero value otherwise.

### GetResourceRulesOk

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) GetResourceRulesOk() (*[]IoK8sApiFlowcontrolV1beta2ResourcePolicyRule, bool)`

GetResourceRulesOk returns a tuple with the ResourceRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRules

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) SetResourceRules(v []IoK8sApiFlowcontrolV1beta2ResourcePolicyRule)`

SetResourceRules sets ResourceRules field to given value.

### HasResourceRules

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) HasResourceRules() bool`

HasResourceRules returns a boolean if a field has been set.

### GetSubjects

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) GetSubjects() []IoK8sApiFlowcontrolV1beta2Subject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) GetSubjectsOk() (*[]IoK8sApiFlowcontrolV1beta2Subject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *IoK8sApiFlowcontrolV1beta2PolicyRulesWithSubjects) SetSubjects(v []IoK8sApiFlowcontrolV1beta2Subject)`

SetSubjects sets Subjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


