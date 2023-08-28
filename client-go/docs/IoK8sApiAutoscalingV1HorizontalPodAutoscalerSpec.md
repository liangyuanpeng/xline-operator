# IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxReplicas** | **int32** | maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas. | 
**MinReplicas** | Pointer to **int32** | minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available. | [optional] 
**ScaleTargetRef** | [**IoK8sApiAutoscalingV1CrossVersionObjectReference**](IoK8sApiAutoscalingV1CrossVersionObjectReference.md) |  | 
**TargetCPUUtilizationPercentage** | Pointer to **int32** | targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used. | [optional] 

## Methods

### NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec

`func NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec(maxReplicas int32, scaleTargetRef IoK8sApiAutoscalingV1CrossVersionObjectReference, ) *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec`

NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec instantiates a new IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerSpecWithDefaults

`func NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerSpecWithDefaults() *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec`

NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerSpecWithDefaults instantiates a new IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxReplicas

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.


### GetMinReplicas

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetScaleTargetRef

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) GetScaleTargetRef() IoK8sApiAutoscalingV1CrossVersionObjectReference`

GetScaleTargetRef returns the ScaleTargetRef field if non-nil, zero value otherwise.

### GetScaleTargetRefOk

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) GetScaleTargetRefOk() (*IoK8sApiAutoscalingV1CrossVersionObjectReference, bool)`

GetScaleTargetRefOk returns a tuple with the ScaleTargetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleTargetRef

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) SetScaleTargetRef(v IoK8sApiAutoscalingV1CrossVersionObjectReference)`

SetScaleTargetRef sets ScaleTargetRef field to given value.


### GetTargetCPUUtilizationPercentage

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) GetTargetCPUUtilizationPercentage() int32`

GetTargetCPUUtilizationPercentage returns the TargetCPUUtilizationPercentage field if non-nil, zero value otherwise.

### GetTargetCPUUtilizationPercentageOk

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) GetTargetCPUUtilizationPercentageOk() (*int32, bool)`

GetTargetCPUUtilizationPercentageOk returns a tuple with the TargetCPUUtilizationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCPUUtilizationPercentage

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) SetTargetCPUUtilizationPercentage(v int32)`

SetTargetCPUUtilizationPercentage sets TargetCPUUtilizationPercentage field to given value.

### HasTargetCPUUtilizationPercentage

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerSpec) HasTargetCPUUtilizationPercentage() bool`

HasTargetCPUUtilizationPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


