# IoK8sApiAutoscalingV2ObjectMetricSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DescribedObject** | [**IoK8sApiAutoscalingV2CrossVersionObjectReference**](IoK8sApiAutoscalingV2CrossVersionObjectReference.md) |  | 
**Metric** | [**IoK8sApiAutoscalingV2MetricIdentifier**](IoK8sApiAutoscalingV2MetricIdentifier.md) |  | 
**Target** | [**IoK8sApiAutoscalingV2MetricTarget**](IoK8sApiAutoscalingV2MetricTarget.md) |  | 

## Methods

### NewIoK8sApiAutoscalingV2ObjectMetricSource

`func NewIoK8sApiAutoscalingV2ObjectMetricSource(describedObject IoK8sApiAutoscalingV2CrossVersionObjectReference, metric IoK8sApiAutoscalingV2MetricIdentifier, target IoK8sApiAutoscalingV2MetricTarget, ) *IoK8sApiAutoscalingV2ObjectMetricSource`

NewIoK8sApiAutoscalingV2ObjectMetricSource instantiates a new IoK8sApiAutoscalingV2ObjectMetricSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2ObjectMetricSourceWithDefaults

`func NewIoK8sApiAutoscalingV2ObjectMetricSourceWithDefaults() *IoK8sApiAutoscalingV2ObjectMetricSource`

NewIoK8sApiAutoscalingV2ObjectMetricSourceWithDefaults instantiates a new IoK8sApiAutoscalingV2ObjectMetricSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescribedObject

`func (o *IoK8sApiAutoscalingV2ObjectMetricSource) GetDescribedObject() IoK8sApiAutoscalingV2CrossVersionObjectReference`

GetDescribedObject returns the DescribedObject field if non-nil, zero value otherwise.

### GetDescribedObjectOk

`func (o *IoK8sApiAutoscalingV2ObjectMetricSource) GetDescribedObjectOk() (*IoK8sApiAutoscalingV2CrossVersionObjectReference, bool)`

GetDescribedObjectOk returns a tuple with the DescribedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribedObject

`func (o *IoK8sApiAutoscalingV2ObjectMetricSource) SetDescribedObject(v IoK8sApiAutoscalingV2CrossVersionObjectReference)`

SetDescribedObject sets DescribedObject field to given value.


### GetMetric

`func (o *IoK8sApiAutoscalingV2ObjectMetricSource) GetMetric() IoK8sApiAutoscalingV2MetricIdentifier`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *IoK8sApiAutoscalingV2ObjectMetricSource) GetMetricOk() (*IoK8sApiAutoscalingV2MetricIdentifier, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *IoK8sApiAutoscalingV2ObjectMetricSource) SetMetric(v IoK8sApiAutoscalingV2MetricIdentifier)`

SetMetric sets Metric field to given value.


### GetTarget

`func (o *IoK8sApiAutoscalingV2ObjectMetricSource) GetTarget() IoK8sApiAutoscalingV2MetricTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IoK8sApiAutoscalingV2ObjectMetricSource) GetTargetOk() (*IoK8sApiAutoscalingV2MetricTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IoK8sApiAutoscalingV2ObjectMetricSource) SetTarget(v IoK8sApiAutoscalingV2MetricTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


