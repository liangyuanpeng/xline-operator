# IoK8sApiAutoscalingV2PodsMetricSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | [**IoK8sApiAutoscalingV2MetricIdentifier**](IoK8sApiAutoscalingV2MetricIdentifier.md) |  | 
**Target** | [**IoK8sApiAutoscalingV2MetricTarget**](IoK8sApiAutoscalingV2MetricTarget.md) |  | 

## Methods

### NewIoK8sApiAutoscalingV2PodsMetricSource

`func NewIoK8sApiAutoscalingV2PodsMetricSource(metric IoK8sApiAutoscalingV2MetricIdentifier, target IoK8sApiAutoscalingV2MetricTarget, ) *IoK8sApiAutoscalingV2PodsMetricSource`

NewIoK8sApiAutoscalingV2PodsMetricSource instantiates a new IoK8sApiAutoscalingV2PodsMetricSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2PodsMetricSourceWithDefaults

`func NewIoK8sApiAutoscalingV2PodsMetricSourceWithDefaults() *IoK8sApiAutoscalingV2PodsMetricSource`

NewIoK8sApiAutoscalingV2PodsMetricSourceWithDefaults instantiates a new IoK8sApiAutoscalingV2PodsMetricSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *IoK8sApiAutoscalingV2PodsMetricSource) GetMetric() IoK8sApiAutoscalingV2MetricIdentifier`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *IoK8sApiAutoscalingV2PodsMetricSource) GetMetricOk() (*IoK8sApiAutoscalingV2MetricIdentifier, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *IoK8sApiAutoscalingV2PodsMetricSource) SetMetric(v IoK8sApiAutoscalingV2MetricIdentifier)`

SetMetric sets Metric field to given value.


### GetTarget

`func (o *IoK8sApiAutoscalingV2PodsMetricSource) GetTarget() IoK8sApiAutoscalingV2MetricTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IoK8sApiAutoscalingV2PodsMetricSource) GetTargetOk() (*IoK8sApiAutoscalingV2MetricTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IoK8sApiAutoscalingV2PodsMetricSource) SetTarget(v IoK8sApiAutoscalingV2MetricTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


