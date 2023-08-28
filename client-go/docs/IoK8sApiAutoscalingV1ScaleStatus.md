# IoK8sApiAutoscalingV1ScaleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | **int32** | replicas is the actual number of observed instances of the scaled object. | 
**Selector** | Pointer to **string** | selector is the label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ | [optional] 

## Methods

### NewIoK8sApiAutoscalingV1ScaleStatus

`func NewIoK8sApiAutoscalingV1ScaleStatus(replicas int32, ) *IoK8sApiAutoscalingV1ScaleStatus`

NewIoK8sApiAutoscalingV1ScaleStatus instantiates a new IoK8sApiAutoscalingV1ScaleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV1ScaleStatusWithDefaults

`func NewIoK8sApiAutoscalingV1ScaleStatusWithDefaults() *IoK8sApiAutoscalingV1ScaleStatus`

NewIoK8sApiAutoscalingV1ScaleStatusWithDefaults instantiates a new IoK8sApiAutoscalingV1ScaleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicas

`func (o *IoK8sApiAutoscalingV1ScaleStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoK8sApiAutoscalingV1ScaleStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoK8sApiAutoscalingV1ScaleStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetSelector

`func (o *IoK8sApiAutoscalingV1ScaleStatus) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiAutoscalingV1ScaleStatus) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiAutoscalingV1ScaleStatus) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *IoK8sApiAutoscalingV1ScaleStatus) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


