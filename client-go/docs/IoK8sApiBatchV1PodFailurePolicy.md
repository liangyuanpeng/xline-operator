# IoK8sApiBatchV1PodFailurePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**[]IoK8sApiBatchV1PodFailurePolicyRule**](IoK8sApiBatchV1PodFailurePolicyRule.md) | A list of pod failure policy rules. The rules are evaluated in order. Once a rule matches a Pod failure, the remaining of the rules are ignored. When no rule matches the Pod failure, the default handling applies - the counter of pod failures is incremented and it is checked against the backoffLimit. At most 20 elements are allowed. | 

## Methods

### NewIoK8sApiBatchV1PodFailurePolicy

`func NewIoK8sApiBatchV1PodFailurePolicy(rules []IoK8sApiBatchV1PodFailurePolicyRule, ) *IoK8sApiBatchV1PodFailurePolicy`

NewIoK8sApiBatchV1PodFailurePolicy instantiates a new IoK8sApiBatchV1PodFailurePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiBatchV1PodFailurePolicyWithDefaults

`func NewIoK8sApiBatchV1PodFailurePolicyWithDefaults() *IoK8sApiBatchV1PodFailurePolicy`

NewIoK8sApiBatchV1PodFailurePolicyWithDefaults instantiates a new IoK8sApiBatchV1PodFailurePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *IoK8sApiBatchV1PodFailurePolicy) GetRules() []IoK8sApiBatchV1PodFailurePolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *IoK8sApiBatchV1PodFailurePolicy) GetRulesOk() (*[]IoK8sApiBatchV1PodFailurePolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *IoK8sApiBatchV1PodFailurePolicy) SetRules(v []IoK8sApiBatchV1PodFailurePolicyRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


