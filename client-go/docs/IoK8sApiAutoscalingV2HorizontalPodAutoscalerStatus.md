# IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition**](IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition.md) | conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met. | [optional] 
**CurrentMetrics** | Pointer to [**[]IoK8sApiAutoscalingV2MetricStatus**](IoK8sApiAutoscalingV2MetricStatus.md) | currentMetrics is the last read state of the metrics used by this autoscaler. | [optional] 
**CurrentReplicas** | Pointer to **int32** | currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler. | [optional] 
**DesiredReplicas** | **int32** | desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler. | 
**LastScaleTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**ObservedGeneration** | Pointer to **int64** | observedGeneration is the most recent generation observed by this autoscaler. | [optional] 

## Methods

### NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus

`func NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus(desiredReplicas int32, ) *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus`

NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus instantiates a new IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerStatusWithDefaults

`func NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerStatusWithDefaults() *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus`

NewIoK8sApiAutoscalingV2HorizontalPodAutoscalerStatusWithDefaults instantiates a new IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetConditions() []IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetConditionsOk() (*[]IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) SetConditions(v []IoK8sApiAutoscalingV2HorizontalPodAutoscalerCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCurrentMetrics

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetCurrentMetrics() []IoK8sApiAutoscalingV2MetricStatus`

GetCurrentMetrics returns the CurrentMetrics field if non-nil, zero value otherwise.

### GetCurrentMetricsOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetCurrentMetricsOk() (*[]IoK8sApiAutoscalingV2MetricStatus, bool)`

GetCurrentMetricsOk returns a tuple with the CurrentMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMetrics

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) SetCurrentMetrics(v []IoK8sApiAutoscalingV2MetricStatus)`

SetCurrentMetrics sets CurrentMetrics field to given value.

### HasCurrentMetrics

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) HasCurrentMetrics() bool`

HasCurrentMetrics returns a boolean if a field has been set.

### GetCurrentReplicas

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetCurrentReplicas() int32`

GetCurrentReplicas returns the CurrentReplicas field if non-nil, zero value otherwise.

### GetCurrentReplicasOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetCurrentReplicasOk() (*int32, bool)`

GetCurrentReplicasOk returns a tuple with the CurrentReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReplicas

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) SetCurrentReplicas(v int32)`

SetCurrentReplicas sets CurrentReplicas field to given value.

### HasCurrentReplicas

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) HasCurrentReplicas() bool`

HasCurrentReplicas returns a boolean if a field has been set.

### GetDesiredReplicas

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetDesiredReplicas() int32`

GetDesiredReplicas returns the DesiredReplicas field if non-nil, zero value otherwise.

### GetDesiredReplicasOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetDesiredReplicasOk() (*int32, bool)`

GetDesiredReplicasOk returns a tuple with the DesiredReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredReplicas

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) SetDesiredReplicas(v int32)`

SetDesiredReplicas sets DesiredReplicas field to given value.


### GetLastScaleTime

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetLastScaleTime() time.Time`

GetLastScaleTime returns the LastScaleTime field if non-nil, zero value otherwise.

### GetLastScaleTimeOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetLastScaleTimeOk() (*time.Time, bool)`

GetLastScaleTimeOk returns a tuple with the LastScaleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScaleTime

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) SetLastScaleTime(v time.Time)`

SetLastScaleTime sets LastScaleTime field to given value.

### HasLastScaleTime

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) HasLastScaleTime() bool`

HasLastScaleTime returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *IoK8sApiAutoscalingV2HorizontalPodAutoscalerStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


