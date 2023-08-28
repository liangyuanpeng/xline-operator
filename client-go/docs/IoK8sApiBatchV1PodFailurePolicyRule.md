# IoK8sApiBatchV1PodFailurePolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Specifies the action taken on a pod failure when the requirements are satisfied. Possible values are:  - FailJob: indicates that the pod&#39;s job is marked as Failed and all   running pods are terminated. - Ignore: indicates that the counter towards the .backoffLimit is not   incremented and a replacement pod is created. - Count: indicates that the pod is handled in the default way - the   counter towards the .backoffLimit is incremented. Additional values are considered to be added in the future. Clients should react to an unknown action by skipping the rule.  Possible enum values:  - &#x60;\&quot;Count\&quot;&#x60; This is an action which might be taken on a pod failure - the pod failure is handled in the default way - the counter towards .backoffLimit, represented by the job&#39;s .status.failed field, is incremented.  - &#x60;\&quot;FailJob\&quot;&#x60; This is an action which might be taken on a pod failure - mark the pod&#39;s job as Failed and terminate all running pods.  - &#x60;\&quot;Ignore\&quot;&#x60; This is an action which might be taken on a pod failure - the counter towards .backoffLimit, represented by the job&#39;s .status.failed field, is not incremented and a replacement pod is created. | 
**OnExitCodes** | Pointer to [**IoK8sApiBatchV1PodFailurePolicyOnExitCodesRequirement**](IoK8sApiBatchV1PodFailurePolicyOnExitCodesRequirement.md) |  | [optional] 
**OnPodConditions** | [**[]IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern**](IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern.md) | Represents the requirement on the pod conditions. The requirement is represented as a list of pod condition patterns. The requirement is satisfied if at least one pattern matches an actual pod condition. At most 20 elements are allowed. | 

## Methods

### NewIoK8sApiBatchV1PodFailurePolicyRule

`func NewIoK8sApiBatchV1PodFailurePolicyRule(action string, onPodConditions []IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern, ) *IoK8sApiBatchV1PodFailurePolicyRule`

NewIoK8sApiBatchV1PodFailurePolicyRule instantiates a new IoK8sApiBatchV1PodFailurePolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiBatchV1PodFailurePolicyRuleWithDefaults

`func NewIoK8sApiBatchV1PodFailurePolicyRuleWithDefaults() *IoK8sApiBatchV1PodFailurePolicyRule`

NewIoK8sApiBatchV1PodFailurePolicyRuleWithDefaults instantiates a new IoK8sApiBatchV1PodFailurePolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IoK8sApiBatchV1PodFailurePolicyRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IoK8sApiBatchV1PodFailurePolicyRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IoK8sApiBatchV1PodFailurePolicyRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetOnExitCodes

`func (o *IoK8sApiBatchV1PodFailurePolicyRule) GetOnExitCodes() IoK8sApiBatchV1PodFailurePolicyOnExitCodesRequirement`

GetOnExitCodes returns the OnExitCodes field if non-nil, zero value otherwise.

### GetOnExitCodesOk

`func (o *IoK8sApiBatchV1PodFailurePolicyRule) GetOnExitCodesOk() (*IoK8sApiBatchV1PodFailurePolicyOnExitCodesRequirement, bool)`

GetOnExitCodesOk returns a tuple with the OnExitCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnExitCodes

`func (o *IoK8sApiBatchV1PodFailurePolicyRule) SetOnExitCodes(v IoK8sApiBatchV1PodFailurePolicyOnExitCodesRequirement)`

SetOnExitCodes sets OnExitCodes field to given value.

### HasOnExitCodes

`func (o *IoK8sApiBatchV1PodFailurePolicyRule) HasOnExitCodes() bool`

HasOnExitCodes returns a boolean if a field has been set.

### GetOnPodConditions

`func (o *IoK8sApiBatchV1PodFailurePolicyRule) GetOnPodConditions() []IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern`

GetOnPodConditions returns the OnPodConditions field if non-nil, zero value otherwise.

### GetOnPodConditionsOk

`func (o *IoK8sApiBatchV1PodFailurePolicyRule) GetOnPodConditionsOk() (*[]IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern, bool)`

GetOnPodConditionsOk returns a tuple with the OnPodConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPodConditions

`func (o *IoK8sApiBatchV1PodFailurePolicyRule) SetOnPodConditions(v []IoK8sApiBatchV1PodFailurePolicyOnPodConditionsPattern)`

SetOnPodConditions sets OnPodConditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


