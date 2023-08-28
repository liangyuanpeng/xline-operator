# IoK8sApiCoreV1NodeAffinity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreferredDuringSchedulingIgnoredDuringExecution** | Pointer to [**[]IoK8sApiCoreV1PreferredSchedulingTerm**](IoK8sApiCoreV1PreferredSchedulingTerm.md) | The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \&quot;weight\&quot; to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred. | [optional] 
**RequiredDuringSchedulingIgnoredDuringExecution** | Pointer to [**IoK8sApiCoreV1NodeSelector**](IoK8sApiCoreV1NodeSelector.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1NodeAffinity

`func NewIoK8sApiCoreV1NodeAffinity() *IoK8sApiCoreV1NodeAffinity`

NewIoK8sApiCoreV1NodeAffinity instantiates a new IoK8sApiCoreV1NodeAffinity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1NodeAffinityWithDefaults

`func NewIoK8sApiCoreV1NodeAffinityWithDefaults() *IoK8sApiCoreV1NodeAffinity`

NewIoK8sApiCoreV1NodeAffinityWithDefaults instantiates a new IoK8sApiCoreV1NodeAffinity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferredDuringSchedulingIgnoredDuringExecution

`func (o *IoK8sApiCoreV1NodeAffinity) GetPreferredDuringSchedulingIgnoredDuringExecution() []IoK8sApiCoreV1PreferredSchedulingTerm`

GetPreferredDuringSchedulingIgnoredDuringExecution returns the PreferredDuringSchedulingIgnoredDuringExecution field if non-nil, zero value otherwise.

### GetPreferredDuringSchedulingIgnoredDuringExecutionOk

`func (o *IoK8sApiCoreV1NodeAffinity) GetPreferredDuringSchedulingIgnoredDuringExecutionOk() (*[]IoK8sApiCoreV1PreferredSchedulingTerm, bool)`

GetPreferredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the PreferredDuringSchedulingIgnoredDuringExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDuringSchedulingIgnoredDuringExecution

`func (o *IoK8sApiCoreV1NodeAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []IoK8sApiCoreV1PreferredSchedulingTerm)`

SetPreferredDuringSchedulingIgnoredDuringExecution sets PreferredDuringSchedulingIgnoredDuringExecution field to given value.

### HasPreferredDuringSchedulingIgnoredDuringExecution

`func (o *IoK8sApiCoreV1NodeAffinity) HasPreferredDuringSchedulingIgnoredDuringExecution() bool`

HasPreferredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.

### GetRequiredDuringSchedulingIgnoredDuringExecution

`func (o *IoK8sApiCoreV1NodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecution() IoK8sApiCoreV1NodeSelector`

GetRequiredDuringSchedulingIgnoredDuringExecution returns the RequiredDuringSchedulingIgnoredDuringExecution field if non-nil, zero value otherwise.

### GetRequiredDuringSchedulingIgnoredDuringExecutionOk

`func (o *IoK8sApiCoreV1NodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecutionOk() (*IoK8sApiCoreV1NodeSelector, bool)`

GetRequiredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the RequiredDuringSchedulingIgnoredDuringExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDuringSchedulingIgnoredDuringExecution

`func (o *IoK8sApiCoreV1NodeAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v IoK8sApiCoreV1NodeSelector)`

SetRequiredDuringSchedulingIgnoredDuringExecution sets RequiredDuringSchedulingIgnoredDuringExecution field to given value.

### HasRequiredDuringSchedulingIgnoredDuringExecution

`func (o *IoK8sApiCoreV1NodeAffinity) HasRequiredDuringSchedulingIgnoredDuringExecution() bool`

HasRequiredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


