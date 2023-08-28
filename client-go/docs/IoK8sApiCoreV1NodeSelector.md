# IoK8sApiCoreV1NodeSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeSelectorTerms** | [**[]IoK8sApiCoreV1NodeSelectorTerm**](IoK8sApiCoreV1NodeSelectorTerm.md) | Required. A list of node selector terms. The terms are ORed. | 

## Methods

### NewIoK8sApiCoreV1NodeSelector

`func NewIoK8sApiCoreV1NodeSelector(nodeSelectorTerms []IoK8sApiCoreV1NodeSelectorTerm, ) *IoK8sApiCoreV1NodeSelector`

NewIoK8sApiCoreV1NodeSelector instantiates a new IoK8sApiCoreV1NodeSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1NodeSelectorWithDefaults

`func NewIoK8sApiCoreV1NodeSelectorWithDefaults() *IoK8sApiCoreV1NodeSelector`

NewIoK8sApiCoreV1NodeSelectorWithDefaults instantiates a new IoK8sApiCoreV1NodeSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeSelectorTerms

`func (o *IoK8sApiCoreV1NodeSelector) GetNodeSelectorTerms() []IoK8sApiCoreV1NodeSelectorTerm`

GetNodeSelectorTerms returns the NodeSelectorTerms field if non-nil, zero value otherwise.

### GetNodeSelectorTermsOk

`func (o *IoK8sApiCoreV1NodeSelector) GetNodeSelectorTermsOk() (*[]IoK8sApiCoreV1NodeSelectorTerm, bool)`

GetNodeSelectorTermsOk returns a tuple with the NodeSelectorTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelectorTerms

`func (o *IoK8sApiCoreV1NodeSelector) SetNodeSelectorTerms(v []IoK8sApiCoreV1NodeSelectorTerm)`

SetNodeSelectorTerms sets NodeSelectorTerms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


