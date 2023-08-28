# IoK8sApiCoreV1Taint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.  Possible enum values:  - &#x60;\&quot;NoExecute\&quot;&#x60; Evict any already-running pods that do not tolerate the taint. Currently enforced by NodeController.  - &#x60;\&quot;NoSchedule\&quot;&#x60; Do not allow new pods to schedule onto the node unless they tolerate the taint, but allow all pods submitted to Kubelet without going through the scheduler to start, and allow all already-running pods to continue running. Enforced by the scheduler.  - &#x60;\&quot;PreferNoSchedule\&quot;&#x60; Like TaintEffectNoSchedule, but the scheduler tries not to schedule new pods onto the node, rather than prohibiting new pods from scheduling onto the node entirely. Enforced by the scheduler. | 
**Key** | **string** | Required. The taint key to be applied to a node. | 
**TimeAdded** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Value** | Pointer to **string** | The taint value corresponding to the taint key. | [optional] 

## Methods

### NewIoK8sApiCoreV1Taint

`func NewIoK8sApiCoreV1Taint(effect string, key string, ) *IoK8sApiCoreV1Taint`

NewIoK8sApiCoreV1Taint instantiates a new IoK8sApiCoreV1Taint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1TaintWithDefaults

`func NewIoK8sApiCoreV1TaintWithDefaults() *IoK8sApiCoreV1Taint`

NewIoK8sApiCoreV1TaintWithDefaults instantiates a new IoK8sApiCoreV1Taint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *IoK8sApiCoreV1Taint) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *IoK8sApiCoreV1Taint) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *IoK8sApiCoreV1Taint) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetKey

`func (o *IoK8sApiCoreV1Taint) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IoK8sApiCoreV1Taint) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IoK8sApiCoreV1Taint) SetKey(v string)`

SetKey sets Key field to given value.


### GetTimeAdded

`func (o *IoK8sApiCoreV1Taint) GetTimeAdded() time.Time`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *IoK8sApiCoreV1Taint) GetTimeAddedOk() (*time.Time, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *IoK8sApiCoreV1Taint) SetTimeAdded(v time.Time)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *IoK8sApiCoreV1Taint) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.

### GetValue

`func (o *IoK8sApiCoreV1Taint) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IoK8sApiCoreV1Taint) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IoK8sApiCoreV1Taint) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IoK8sApiCoreV1Taint) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


