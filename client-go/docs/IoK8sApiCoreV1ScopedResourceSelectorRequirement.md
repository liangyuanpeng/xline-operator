# IoK8sApiCoreV1ScopedResourceSelectorRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Represents a scope&#39;s relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.  Possible enum values:  - &#x60;\&quot;DoesNotExist\&quot;&#x60;  - &#x60;\&quot;Exists\&quot;&#x60;  - &#x60;\&quot;In\&quot;&#x60;  - &#x60;\&quot;NotIn\&quot;&#x60; | 
**ScopeName** | **string** | The name of the scope that the selector applies to.  Possible enum values:  - &#x60;\&quot;BestEffort\&quot;&#x60; Match all pod objects that have best effort quality of service  - &#x60;\&quot;CrossNamespacePodAffinity\&quot;&#x60; Match all pod objects that have cross-namespace pod (anti)affinity mentioned.  - &#x60;\&quot;NotBestEffort\&quot;&#x60; Match all pod objects that do not have best effort quality of service  - &#x60;\&quot;NotTerminating\&quot;&#x60; Match all pod objects where spec.activeDeadlineSeconds is nil  - &#x60;\&quot;PriorityClass\&quot;&#x60; Match all pod objects that have priority class mentioned  - &#x60;\&quot;Terminating\&quot;&#x60; Match all pod objects where spec.activeDeadlineSeconds &gt;&#x3D;0 | 
**Values** | Pointer to **[]string** | An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch. | [optional] 

## Methods

### NewIoK8sApiCoreV1ScopedResourceSelectorRequirement

`func NewIoK8sApiCoreV1ScopedResourceSelectorRequirement(operator string, scopeName string, ) *IoK8sApiCoreV1ScopedResourceSelectorRequirement`

NewIoK8sApiCoreV1ScopedResourceSelectorRequirement instantiates a new IoK8sApiCoreV1ScopedResourceSelectorRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ScopedResourceSelectorRequirementWithDefaults

`func NewIoK8sApiCoreV1ScopedResourceSelectorRequirementWithDefaults() *IoK8sApiCoreV1ScopedResourceSelectorRequirement`

NewIoK8sApiCoreV1ScopedResourceSelectorRequirementWithDefaults instantiates a new IoK8sApiCoreV1ScopedResourceSelectorRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *IoK8sApiCoreV1ScopedResourceSelectorRequirement) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *IoK8sApiCoreV1ScopedResourceSelectorRequirement) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *IoK8sApiCoreV1ScopedResourceSelectorRequirement) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetScopeName

`func (o *IoK8sApiCoreV1ScopedResourceSelectorRequirement) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *IoK8sApiCoreV1ScopedResourceSelectorRequirement) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *IoK8sApiCoreV1ScopedResourceSelectorRequirement) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.


### GetValues

`func (o *IoK8sApiCoreV1ScopedResourceSelectorRequirement) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *IoK8sApiCoreV1ScopedResourceSelectorRequirement) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *IoK8sApiCoreV1ScopedResourceSelectorRequirement) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *IoK8sApiCoreV1ScopedResourceSelectorRequirement) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


