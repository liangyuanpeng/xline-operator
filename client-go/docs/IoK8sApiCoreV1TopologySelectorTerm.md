# IoK8sApiCoreV1TopologySelectorTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchLabelExpressions** | Pointer to [**[]IoK8sApiCoreV1TopologySelectorLabelRequirement**](IoK8sApiCoreV1TopologySelectorLabelRequirement.md) | A list of topology selector requirements by labels. | [optional] 

## Methods

### NewIoK8sApiCoreV1TopologySelectorTerm

`func NewIoK8sApiCoreV1TopologySelectorTerm() *IoK8sApiCoreV1TopologySelectorTerm`

NewIoK8sApiCoreV1TopologySelectorTerm instantiates a new IoK8sApiCoreV1TopologySelectorTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1TopologySelectorTermWithDefaults

`func NewIoK8sApiCoreV1TopologySelectorTermWithDefaults() *IoK8sApiCoreV1TopologySelectorTerm`

NewIoK8sApiCoreV1TopologySelectorTermWithDefaults instantiates a new IoK8sApiCoreV1TopologySelectorTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchLabelExpressions

`func (o *IoK8sApiCoreV1TopologySelectorTerm) GetMatchLabelExpressions() []IoK8sApiCoreV1TopologySelectorLabelRequirement`

GetMatchLabelExpressions returns the MatchLabelExpressions field if non-nil, zero value otherwise.

### GetMatchLabelExpressionsOk

`func (o *IoK8sApiCoreV1TopologySelectorTerm) GetMatchLabelExpressionsOk() (*[]IoK8sApiCoreV1TopologySelectorLabelRequirement, bool)`

GetMatchLabelExpressionsOk returns a tuple with the MatchLabelExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchLabelExpressions

`func (o *IoK8sApiCoreV1TopologySelectorTerm) SetMatchLabelExpressions(v []IoK8sApiCoreV1TopologySelectorLabelRequirement)`

SetMatchLabelExpressions sets MatchLabelExpressions field to given value.

### HasMatchLabelExpressions

`func (o *IoK8sApiCoreV1TopologySelectorTerm) HasMatchLabelExpressions() bool`

HasMatchLabelExpressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


