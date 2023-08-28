# IoK8sApiCoreV1ScopeSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchExpressions** | Pointer to [**[]IoK8sApiCoreV1ScopedResourceSelectorRequirement**](IoK8sApiCoreV1ScopedResourceSelectorRequirement.md) | A list of scope selector requirements by scope of the resources. | [optional] 

## Methods

### NewIoK8sApiCoreV1ScopeSelector

`func NewIoK8sApiCoreV1ScopeSelector() *IoK8sApiCoreV1ScopeSelector`

NewIoK8sApiCoreV1ScopeSelector instantiates a new IoK8sApiCoreV1ScopeSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ScopeSelectorWithDefaults

`func NewIoK8sApiCoreV1ScopeSelectorWithDefaults() *IoK8sApiCoreV1ScopeSelector`

NewIoK8sApiCoreV1ScopeSelectorWithDefaults instantiates a new IoK8sApiCoreV1ScopeSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchExpressions

`func (o *IoK8sApiCoreV1ScopeSelector) GetMatchExpressions() []IoK8sApiCoreV1ScopedResourceSelectorRequirement`

GetMatchExpressions returns the MatchExpressions field if non-nil, zero value otherwise.

### GetMatchExpressionsOk

`func (o *IoK8sApiCoreV1ScopeSelector) GetMatchExpressionsOk() (*[]IoK8sApiCoreV1ScopedResourceSelectorRequirement, bool)`

GetMatchExpressionsOk returns a tuple with the MatchExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpressions

`func (o *IoK8sApiCoreV1ScopeSelector) SetMatchExpressions(v []IoK8sApiCoreV1ScopedResourceSelectorRequirement)`

SetMatchExpressions sets MatchExpressions field to given value.

### HasMatchExpressions

`func (o *IoK8sApiCoreV1ScopeSelector) HasMatchExpressions() bool`

HasMatchExpressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


