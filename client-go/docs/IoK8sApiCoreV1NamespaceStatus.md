# IoK8sApiCoreV1NamespaceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IoK8sApiCoreV1NamespaceCondition**](IoK8sApiCoreV1NamespaceCondition.md) | Represents the latest available observations of a namespace&#39;s current state. | [optional] 
**Phase** | Pointer to **string** | Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/  Possible enum values:  - &#x60;\&quot;Active\&quot;&#x60; means the namespace is available for use in the system  - &#x60;\&quot;Terminating\&quot;&#x60; means the namespace is undergoing graceful termination | [optional] 

## Methods

### NewIoK8sApiCoreV1NamespaceStatus

`func NewIoK8sApiCoreV1NamespaceStatus() *IoK8sApiCoreV1NamespaceStatus`

NewIoK8sApiCoreV1NamespaceStatus instantiates a new IoK8sApiCoreV1NamespaceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1NamespaceStatusWithDefaults

`func NewIoK8sApiCoreV1NamespaceStatusWithDefaults() *IoK8sApiCoreV1NamespaceStatus`

NewIoK8sApiCoreV1NamespaceStatusWithDefaults instantiates a new IoK8sApiCoreV1NamespaceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *IoK8sApiCoreV1NamespaceStatus) GetConditions() []IoK8sApiCoreV1NamespaceCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiCoreV1NamespaceStatus) GetConditionsOk() (*[]IoK8sApiCoreV1NamespaceCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiCoreV1NamespaceStatus) SetConditions(v []IoK8sApiCoreV1NamespaceCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiCoreV1NamespaceStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPhase

`func (o *IoK8sApiCoreV1NamespaceStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *IoK8sApiCoreV1NamespaceStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *IoK8sApiCoreV1NamespaceStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *IoK8sApiCoreV1NamespaceStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


