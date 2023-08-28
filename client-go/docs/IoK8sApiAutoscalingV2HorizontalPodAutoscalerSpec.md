# IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | Pointer to [**IoK8sApiAutoscalingV2HorizontalPodAutoscalerBehavior**](IoK8sApiAutoscalingV2HorizontalPodAutoscalerBehavior.md) |  | [optional] 
**MaxReplicas** | **int32** | maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas. | 
**Metrics** | Pointer to [**[]IoK8sApiAutoscalingV2MetricSpec**](IoK8sApiAutoscalingV2MetricSpec.md) | metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used).  The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods.  Ergo, metrics used must decrease as the pod count is increased, and vice-versa.  See the individual metric source types for more information about how each type of metric must respond. If not set, the default metric will be set to 80% average CPU utilization. | [optional] 
**MinReplicas** | Pointer to **int32** | minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available. | [optional] 
**ScaleTargetRef** | [**IoK8sApiAutoscalingV2CrossVersionObjectReference**](IoK8sApiAutoscalingV2CrossVersionObjectReference.md) |  | 

## Methods

### NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec

`func NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec(maxReplicas int32, scaleTargetRef IoK8sApiAutoscalingV2CrossVersionObjectReference, ) *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec`

NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec instantiates a new IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerSpecWithDefaults

`func NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerSpecWithDefaults() *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec`

NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerSpecWithDefaults instantiates a new IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) GetBehavior() IoK8sApiAutoscalingV2HorizontalPodAutoscalerBehavior`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) GetBehaviorOk() (*IoK8sApiAutoscalingV2HorizontalPodAutoscalerBehavior, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) SetBehavior(v IoK8sApiAutoscalingV2HorizontalPodAutoscalerBehavior)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.


### GetMetrics

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) GetMetrics() []IoK8sApiAutoscalingV2MetricSpec`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) GetMetricsOk() (*[]IoK8sApiAutoscalingV2MetricSpec, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) SetMetrics(v []IoK8sApiAutoscalingV2MetricSpec)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMinReplicas

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetScaleTargetRef

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) GetScaleTargetRef() IoK8sApiAutoscalingV2CrossVersionObjectReference`

GetScaleTargetRef returns the ScaleTargetRef field if non-nil, zero value otherwise.

### GetScaleTargetRefOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) GetScaleTargetRefOk() (*IoK8sApiAutoscalingV2CrossVersionObjectReference, bool)`

GetScaleTargetRefOk returns a tuple with the ScaleTargetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleTargetRef

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerSpec) SetScaleTargetRef(v IoK8sApiAutoscalingV2CrossVersionObjectReference)`

SetScaleTargetRef sets ScaleTargetRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


