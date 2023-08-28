# IoK8sApiAuthorizationV1SubjectRulesReviewStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluationError** | Pointer to **string** | EvaluationError can appear in combination with Rules. It indicates an error occurred during rule evaluation, such as an authorizer that doesn&#39;t support rule evaluation, and that ResourceRules and/or NonResourceRules may be incomplete. | [optional] 
**Incomplete** | **bool** | Incomplete is true when the rules returned by this call are incomplete. This is most commonly encountered when an authorizer, such as an external authorizer, doesn&#39;t support rules evaluation. | 
**NonResourceRules** | [**[]IoK8sApiAuthorizationV1NonResourceRule**](IoK8sApiAuthorizationV1NonResourceRule.md) | NonResourceRules is the list of actions the subject is allowed to perform on non-resources. The list ordering isn&#39;t significant, may contain duplicates, and possibly be incomplete. | 
**ResourceRules** | [**[]IoK8sApiAuthorizationV1ResourceRule**](IoK8sApiAuthorizationV1ResourceRule.md) | ResourceRules is the list of actions the subject is allowed to perform on resources. The list ordering isn&#39;t significant, may contain duplicates, and possibly be incomplete. | 

## Methods

### NewIoK8sApiAuthorizationV1SubjectRulesReviewStatus

`func NewIoK8sApiAuthorizationV1SubjectRulesReviewStatus(incomplete bool, nonResourceRules []IoK8sApiAuthorizationV1NonResourceRule, resourceRules []IoK8sApiAuthorizationV1ResourceRule, ) *IoK8sApiAuthorizationV1SubjectRulesReviewStatus`

NewIoK8sApiAuthorizationV1SubjectRulesReviewStatus instantiates a new IoK8sApiAuthorizationV1SubjectRulesReviewStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthorizationV1SubjectRulesReviewStatusWithDefaults

`func NewIoK8sApiAuthorizationV1SubjectRulesReviewStatusWithDefaults() *IoK8sApiAuthorizationV1SubjectRulesReviewStatus`

NewIoK8sApiAuthorizationV1SubjectRulesReviewStatusWithDefaults instantiates a new IoK8sApiAuthorizationV1SubjectRulesReviewStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluationError

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) GetEvaluationError() string`

GetEvaluationError returns the EvaluationError field if non-nil, zero value otherwise.

### GetEvaluationErrorOk

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) GetEvaluationErrorOk() (*string, bool)`

GetEvaluationErrorOk returns a tuple with the EvaluationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationError

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) SetEvaluationError(v string)`

SetEvaluationError sets EvaluationError field to given value.

### HasEvaluationError

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) HasEvaluationError() bool`

HasEvaluationError returns a boolean if a field has been set.

### GetIncomplete

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) GetIncomplete() bool`

GetIncomplete returns the Incomplete field if non-nil, zero value otherwise.

### GetIncompleteOk

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) GetIncompleteOk() (*bool, bool)`

GetIncompleteOk returns a tuple with the Incomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomplete

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) SetIncomplete(v bool)`

SetIncomplete sets Incomplete field to given value.


### GetNonResourceRules

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) GetNonResourceRules() []IoK8sApiAuthorizationV1NonResourceRule`

GetNonResourceRules returns the NonResourceRules field if non-nil, zero value otherwise.

### GetNonResourceRulesOk

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) GetNonResourceRulesOk() (*[]IoK8sApiAuthorizationV1NonResourceRule, bool)`

GetNonResourceRulesOk returns a tuple with the NonResourceRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceRules

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) SetNonResourceRules(v []IoK8sApiAuthorizationV1NonResourceRule)`

SetNonResourceRules sets NonResourceRules field to given value.


### GetResourceRules

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) GetResourceRules() []IoK8sApiAuthorizationV1ResourceRule`

GetResourceRules returns the ResourceRules field if non-nil, zero value otherwise.

### GetResourceRulesOk

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) GetResourceRulesOk() (*[]IoK8sApiAuthorizationV1ResourceRule, bool)`

GetResourceRulesOk returns a tuple with the ResourceRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRules

`func (o *IoK8sApiAuthorizationV1SubjectRulesReviewStatus) SetResourceRules(v []IoK8sApiAuthorizationV1ResourceRule)`

SetResourceRules sets ResourceRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


