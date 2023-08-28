# IoK8sApiCoreV1NodeSelectorTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchExpressions** | Pointer to [**[]IoK8sApiCoreV1NodeSelectorRequirement**](IoK8sApiCoreV1NodeSelectorRequirement.md) | A list of node selector requirements by node&#39;s labels. | [optional] 
**MatchFields** | Pointer to [**[]IoK8sApiCoreV1NodeSelectorRequirement**](IoK8sApiCoreV1NodeSelectorRequirement.md) | A list of node selector requirements by node&#39;s fields. | [optional] 

## Methods

### NewIoK8sApiCoreV1NodeSelectorTerm

`func NewIoK8sApiCoreV1NodeSelectorTerm() *IoK8sApiCoreV1NodeSelectorTerm`

NewIoK8sApiCoreV1NodeSelectorTerm instantiates a new IoK8sApiCoreV1NodeSelectorTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1NodeSelectorTermWithDefaults

`func NewIoK8sApiCoreV1NodeSelectorTermWithDefaults() *IoK8sApiCoreV1NodeSelectorTerm`

NewIoK8sApiCoreV1NodeSelectorTermWithDefaults instantiates a new IoK8sApiCoreV1NodeSelectorTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchExpressions

`func (o *IoK8sApiCoreV1NodeSelectorTerm) GetMatchExpressions() []IoK8sApiCoreV1NodeSelectorRequirement`

GetMatchExpressions returns the MatchExpressions field if non-nil, zero value otherwise.

### GetMatchExpressionsOk

`func (o *IoK8sApiCoreV1NodeSelectorTerm) GetMatchExpressionsOk() (*[]IoK8sApiCoreV1NodeSelectorRequirement, bool)`

GetMatchExpressionsOk returns a tuple with the MatchExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpressions

`func (o *IoK8sApiCoreV1NodeSelectorTerm) SetMatchExpressions(v []IoK8sApiCoreV1NodeSelectorRequirement)`

SetMatchExpressions sets MatchExpressions field to given value.

### HasMatchExpressions

`func (o *IoK8sApiCoreV1NodeSelectorTerm) HasMatchExpressions() bool`

HasMatchExpressions returns a boolean if a field has been set.

### GetMatchFields

`func (o *IoK8sApiCoreV1NodeSelectorTerm) GetMatchFields() []IoK8sApiCoreV1NodeSelectorRequirement`

GetMatchFields returns the MatchFields field if non-nil, zero value otherwise.

### GetMatchFieldsOk

`func (o *IoK8sApiCoreV1NodeSelectorTerm) GetMatchFieldsOk() (*[]IoK8sApiCoreV1NodeSelectorRequirement, bool)`

GetMatchFieldsOk returns a tuple with the MatchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchFields

`func (o *IoK8sApiCoreV1NodeSelectorTerm) SetMatchFields(v []IoK8sApiCoreV1NodeSelectorRequirement)`

SetMatchFields sets MatchFields field to given value.

### HasMatchFields

`func (o *IoK8sApiCoreV1NodeSelectorTerm) HasMatchFields() bool`

HasMatchFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


