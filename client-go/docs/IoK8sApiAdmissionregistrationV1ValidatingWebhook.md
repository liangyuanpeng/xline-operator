# IoK8sApiAdmissionregistrationV1ValidatingWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdmissionReviewVersions** | **[]string** | AdmissionReviewVersions is an ordered list of preferred &#x60;AdmissionReview&#x60; versions the Webhook expects. API server will try to use first version in the list which it supports. If none of the versions specified in this list supported by API server, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail and be subject to the failure policy. | 
**ClientConfig** | [**IoK8sApiAdmissionregistrationV1WebhookClientConfig**](IoK8sApiAdmissionregistrationV1WebhookClientConfig.md) |  | 
**FailurePolicy** | Pointer to **string** | FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail. Defaults to Fail.  Possible enum values:  - &#x60;\&quot;Fail\&quot;&#x60; means that an error calling the webhook causes the admission to fail.  - &#x60;\&quot;Ignore\&quot;&#x60; means that an error calling the webhook is ignored. | [optional] 
**MatchConditions** | Pointer to [**[]IoK8sApiAdmissionregistrationV1MatchCondition**](IoK8sApiAdmissionregistrationV1MatchCondition.md) | MatchConditions is a list of conditions that must be met for a request to be sent to this webhook. Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed.  The exact matching logic is (in order):   1. If ANY matchCondition evaluates to FALSE, the webhook is skipped.   2. If ALL matchConditions evaluate to TRUE, the webhook is called.   3. If any matchCondition evaluates to an error (but none are FALSE):      - If failurePolicy&#x3D;Fail, reject the request      - If failurePolicy&#x3D;Ignore, the error is ignored and the webhook is skipped  This is an alpha feature and managed by the AdmissionWebhookMatchConditions feature gate. | [optional] 
**MatchPolicy** | Pointer to **string** | matchPolicy defines how the \&quot;rules\&quot; list is used to match incoming requests. Allowed values are \&quot;Exact\&quot; or \&quot;Equivalent\&quot;.  - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but \&quot;rules\&quot; only included &#x60;apiGroups:[\&quot;apps\&quot;], apiVersions:[\&quot;v1\&quot;], resources: [\&quot;deployments\&quot;]&#x60;, a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook.  - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and \&quot;rules\&quot; only included &#x60;apiGroups:[\&quot;apps\&quot;], apiVersions:[\&quot;v1\&quot;], resources: [\&quot;deployments\&quot;]&#x60;, a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the webhook.  Defaults to \&quot;Equivalent\&quot;  Possible enum values:  - &#x60;\&quot;Equivalent\&quot;&#x60; means requests should be sent to the webhook if they modify a resource listed in rules via another API group or version.  - &#x60;\&quot;Exact\&quot;&#x60; means requests should only be sent to the webhook if they exactly match a given rule. | [optional] 
**Name** | **string** | The name of the admission webhook. Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where \&quot;imagepolicy\&quot; is the name of the webhook, and kubernetes.io is the name of the organization. Required. | 
**NamespaceSelector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**ObjectSelector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**Rules** | Pointer to [**[]IoK8sApiAdmissionregistrationV1RuleWithOperations**](IoK8sApiAdmissionregistrationV1RuleWithOperations.md) | Rules describes what operations on what resources/subresources the webhook cares about. The webhook cares about an operation if it matches _any_ Rule. However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state which cannot be recovered from without completely disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects. | [optional] 
**SideEffects** | **string** | SideEffects states whether this webhook has side effects. Acceptable values are: None, NoneOnDryRun (webhooks created via v1beta1 may also specify Some or Unknown). Webhooks with side effects MUST implement a reconciliation system, since a request may be rejected by a future step in the admission chain and the side effects therefore need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a webhook with sideEffects &#x3D;&#x3D; Unknown or Some.  Possible enum values:  - &#x60;\&quot;None\&quot;&#x60; means that calling the webhook will have no side effects.  - &#x60;\&quot;NoneOnDryRun\&quot;&#x60; means that calling the webhook will possibly have side effects, but if the request being reviewed has the dry-run attribute, the side effects will be suppressed.  - &#x60;\&quot;Some\&quot;&#x60; means that calling the webhook will possibly have side effects. If a request with the dry-run attribute would trigger a call to this webhook, the request will instead fail.  - &#x60;\&quot;Unknown\&quot;&#x60; means that no information is known about the side effects of calling the webhook. If a request with the dry-run attribute would trigger a call to this webhook, the request will instead fail. | 
**TimeoutSeconds** | Pointer to **int32** | TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds. | [optional] 

## Methods

### NewIoK8sApiAdmissionregistrationV1ValidatingWebhook

`func NewIoK8sApiAdmissionregistrationV1ValidatingWebhook(admissionReviewVersions []string, clientConfig IoK8sApiAdmissionregistrationV1WebhookClientConfig, name string, sideEffects string, ) *IoK8sApiAdmissionregistrationV1ValidatingWebhook`

NewIoK8sApiAdmissionregistrationV1ValidatingWebhook instantiates a new IoK8sApiAdmissionregistrationV1ValidatingWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAdmissionregistrationV1ValidatingWebhookWithDefaults

`func NewIoK8sApiAdmissionregistrationV1ValidatingWebhookWithDefaults() *IoK8sApiAdmissionregistrationV1ValidatingWebhook`

NewIoK8sApiAdmissionregistrationV1ValidatingWebhookWithDefaults instantiates a new IoK8sApiAdmissionregistrationV1ValidatingWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmissionReviewVersions

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetAdmissionReviewVersions() []string`

GetAdmissionReviewVersions returns the AdmissionReviewVersions field if non-nil, zero value otherwise.

### GetAdmissionReviewVersionsOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetAdmissionReviewVersionsOk() (*[]string, bool)`

GetAdmissionReviewVersionsOk returns a tuple with the AdmissionReviewVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmissionReviewVersions

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetAdmissionReviewVersions(v []string)`

SetAdmissionReviewVersions sets AdmissionReviewVersions field to given value.


### GetClientConfig

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetClientConfig() IoK8sApiAdmissionregistrationV1WebhookClientConfig`

GetClientConfig returns the ClientConfig field if non-nil, zero value otherwise.

### GetClientConfigOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetClientConfigOk() (*IoK8sApiAdmissionregistrationV1WebhookClientConfig, bool)`

GetClientConfigOk returns a tuple with the ClientConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConfig

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetClientConfig(v IoK8sApiAdmissionregistrationV1WebhookClientConfig)`

SetClientConfig sets ClientConfig field to given value.


### GetFailurePolicy

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetFailurePolicy() string`

GetFailurePolicy returns the FailurePolicy field if non-nil, zero value otherwise.

### GetFailurePolicyOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetFailurePolicyOk() (*string, bool)`

GetFailurePolicyOk returns a tuple with the FailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailurePolicy

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetFailurePolicy(v string)`

SetFailurePolicy sets FailurePolicy field to given value.

### HasFailurePolicy

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) HasFailurePolicy() bool`

HasFailurePolicy returns a boolean if a field has been set.

### GetMatchConditions

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetMatchConditions() []IoK8sApiAdmissionregistrationV1MatchCondition`

GetMatchConditions returns the MatchConditions field if non-nil, zero value otherwise.

### GetMatchConditionsOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetMatchConditionsOk() (*[]IoK8sApiAdmissionregistrationV1MatchCondition, bool)`

GetMatchConditionsOk returns a tuple with the MatchConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConditions

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetMatchConditions(v []IoK8sApiAdmissionregistrationV1MatchCondition)`

SetMatchConditions sets MatchConditions field to given value.

### HasMatchConditions

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) HasMatchConditions() bool`

HasMatchConditions returns a boolean if a field has been set.

### GetMatchPolicy

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetMatchPolicy() string`

GetMatchPolicy returns the MatchPolicy field if non-nil, zero value otherwise.

### GetMatchPolicyOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetMatchPolicyOk() (*string, bool)`

GetMatchPolicyOk returns a tuple with the MatchPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPolicy

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetMatchPolicy(v string)`

SetMatchPolicy sets MatchPolicy field to given value.

### HasMatchPolicy

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) HasMatchPolicy() bool`

HasMatchPolicy returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceSelector

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetNamespaceSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetNamespaceSelector returns the NamespaceSelector field if non-nil, zero value otherwise.

### GetNamespaceSelectorOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetNamespaceSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetNamespaceSelectorOk returns a tuple with the NamespaceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceSelector

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetNamespaceSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetNamespaceSelector sets NamespaceSelector field to given value.

### HasNamespaceSelector

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) HasNamespaceSelector() bool`

HasNamespaceSelector returns a boolean if a field has been set.

### GetObjectSelector

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetObjectSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetObjectSelector returns the ObjectSelector field if non-nil, zero value otherwise.

### GetObjectSelectorOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetObjectSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetObjectSelectorOk returns a tuple with the ObjectSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSelector

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetObjectSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetObjectSelector sets ObjectSelector field to given value.

### HasObjectSelector

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) HasObjectSelector() bool`

HasObjectSelector returns a boolean if a field has been set.

### GetRules

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetRules() []IoK8sApiAdmissionregistrationV1RuleWithOperations`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetRulesOk() (*[]IoK8sApiAdmissionregistrationV1RuleWithOperations, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetRules(v []IoK8sApiAdmissionregistrationV1RuleWithOperations)`

SetRules sets Rules field to given value.

### HasRules

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSideEffects

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetSideEffects() string`

GetSideEffects returns the SideEffects field if non-nil, zero value otherwise.

### GetSideEffectsOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetSideEffectsOk() (*string, bool)`

GetSideEffectsOk returns a tuple with the SideEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideEffects

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetSideEffects(v string)`

SetSideEffects sets SideEffects field to given value.


### GetTimeoutSeconds

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *IoK8sApiAdmissionregistrationV1ValidatingWebhook) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


