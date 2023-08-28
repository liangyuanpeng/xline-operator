# IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentCPUUtilizationPercentage** | Pointer to **int32** | currentCPUUtilizationPercentage is the current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU. | [optional] 
**CurrentReplicas** | **int32** | currentReplicas is the current number of replicas of pods managed by this autoscaler. | 
**DesiredReplicas** | **int32** | desiredReplicas is the  desired number of replicas of pods managed by this autoscaler. | 
**LastScaleTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**ObservedGeneration** | Pointer to **int64** | observedGeneration is the most recent generation observed by this autoscaler. | [optional] 

## Methods

### NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus

`func NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus(currentReplicas int32, desiredReplicas int32, ) *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus`

NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus instantiates a new IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerStatusWithDefaults

`func NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerStatusWithDefaults() *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus`

NewIoK8sApiAutoscalingV1HorizontalPodAutoscalerStatusWithDefaults instantiates a new IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentCPUUtilizationPercentage

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) GetCurrentCPUUtilizationPercentage() int32`

GetCurrentCPUUtilizationPercentage returns the CurrentCPUUtilizationPercentage field if non-nil, zero value otherwise.

### GetCurrentCPUUtilizationPercentageOk

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) GetCurrentCPUUtilizationPercentageOk() (*int32, bool)`

GetCurrentCPUUtilizationPercentageOk returns a tuple with the CurrentCPUUtilizationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCPUUtilizationPercentage

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) SetCurrentCPUUtilizationPercentage(v int32)`

SetCurrentCPUUtilizationPercentage sets CurrentCPUUtilizationPercentage field to given value.

### HasCurrentCPUUtilizationPercentage

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) HasCurrentCPUUtilizationPercentage() bool`

HasCurrentCPUUtilizationPercentage returns a boolean if a field has been set.

### GetCurrentReplicas

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) GetCurrentReplicas() int32`

GetCurrentReplicas returns the CurrentReplicas field if non-nil, zero value otherwise.

### GetCurrentReplicasOk

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) GetCurrentReplicasOk() (*int32, bool)`

GetCurrentReplicasOk returns a tuple with the CurrentReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReplicas

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) SetCurrentReplicas(v int32)`

SetCurrentReplicas sets CurrentReplicas field to given value.


### GetDesiredReplicas

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) GetDesiredReplicas() int32`

GetDesiredReplicas returns the DesiredReplicas field if non-nil, zero value otherwise.

### GetDesiredReplicasOk

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) GetDesiredReplicasOk() (*int32, bool)`

GetDesiredReplicasOk returns a tuple with the DesiredReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredReplicas

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) SetDesiredReplicas(v int32)`

SetDesiredReplicas sets DesiredReplicas field to given value.


### GetLastScaleTime

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) GetLastScaleTime() time.Time`

GetLastScaleTime returns the LastScaleTime field if non-nil, zero value otherwise.

### GetLastScaleTimeOk

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) GetLastScaleTimeOk() (*time.Time, bool)`

GetLastScaleTimeOk returns a tuple with the LastScaleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScaleTime

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) SetLastScaleTime(v time.Time)`

SetLastScaleTime sets LastScaleTime field to given value.

### HasLastScaleTime

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) HasLastScaleTime() bool`

HasLastScaleTime returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *IoK8sApiAutoscalingV1HorizontalPodAutoscalerStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


