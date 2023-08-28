# IoK8sApiAutoscalingV1ScaleSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | Pointer to **int32** | replicas is the desired number of instances for the scaled object. | [optional] 

## Methods

### NewIoK8sApiAutoscalingV1ScaleSpec

`func NewIoK8sApiAutoscalingV1ScaleSpec() *IoK8sApiAutoscalingV1ScaleSpec`

NewIoK8sApiAutoscalingV1ScaleSpec instantiates a new IoK8sApiAutoscalingV1ScaleSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV1ScaleSpecWithDefaults

`func NewIoK8sApiAutoscalingV1ScaleSpecWithDefaults() *IoK8sApiAutoscalingV1ScaleSpec`

NewIoK8sApiAutoscalingV1ScaleSpecWithDefaults instantiates a new IoK8sApiAutoscalingV1ScaleSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicas

`func (o *IoK8sApiAutoscalingV1ScaleSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoK8sApiAutoscalingV1ScaleSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoK8sApiAutoscalingV1ScaleSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *IoK8sApiAutoscalingV1ScaleSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


