# IoK8sApiAuthorizationV1SubjectAccessReviewStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** | Allowed is required. True if the action would be allowed, false otherwise. | 
**Denied** | Pointer to **bool** | Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true. | [optional] 
**EvaluationError** | Pointer to **string** | EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request. | [optional] 
**Reason** | Pointer to **string** | Reason is optional.  It indicates why a request was allowed or denied. | [optional] 

## Methods

### NewIoK8sApiAuthorizationV1SubjectAccessReviewStatus

`func NewIoK8sApiAuthorizationV1SubjectAccessReviewStatus(allowed bool, ) *IoK8sApiAuthorizationV1SubjectAccessReviewStatus`

NewIoK8sApiAuthorizationV1SubjectAccessReviewStatus instantiates a new IoK8sApiAuthorizationV1SubjectAccessReviewStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthorizationV1SubjectAccessReviewStatusWithDefaults

`func NewIoK8sApiAuthorizationV1SubjectAccessReviewStatusWithDefaults() *IoK8sApiAuthorizationV1SubjectAccessReviewStatus`

NewIoK8sApiAuthorizationV1SubjectAccessReviewStatusWithDefaults instantiates a new IoK8sApiAuthorizationV1SubjectAccessReviewStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetDenied

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) GetDenied() bool`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) GetDeniedOk() (*bool, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) SetDenied(v bool)`

SetDenied sets Denied field to given value.

### HasDenied

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) HasDenied() bool`

HasDenied returns a boolean if a field has been set.

### GetEvaluationError

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) GetEvaluationError() string`

GetEvaluationError returns the EvaluationError field if non-nil, zero value otherwise.

### GetEvaluationErrorOk

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) GetEvaluationErrorOk() (*string, bool)`

GetEvaluationErrorOk returns a tuple with the EvaluationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationError

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) SetEvaluationError(v string)`

SetEvaluationError sets EvaluationError field to given value.

### HasEvaluationError

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) HasEvaluationError() bool`

HasEvaluationError returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


