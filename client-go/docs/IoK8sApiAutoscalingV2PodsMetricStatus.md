# IoK8sApiAutoscalingV2PodsMetricStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | [**IoK8sApiAutoscalingV2MetricValueStatus**](IoK8sApiAutoscalingV2MetricValueStatus.md) |  | 
**Metric** | [**IoK8sApiAutoscalingV2MetricIdentifier**](IoK8sApiAutoscalingV2MetricIdentifier.md) |  | 

## Methods

### NewIoK8sApiAutoscalingV2PodsMetricStatus

`func NewIoK8sApiAutoscalingV2PodsMetricStatus(current IoK8sApiAutoscalingV2MetricValueStatus, metric IoK8sApiAutoscalingV2MetricIdentifier, ) *IoK8sApiAutoscalingV2PodsMetricStatus`

NewIoK8sApiAutoscalingV2PodsMetricStatus instantiates a new IoK8sApiAutoscalingV2PodsMetricStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2PodsMetricStatusWithDefaults

`func NewIoK8sApiAutoscalingV2PodsMetricStatusWithDefaults() *IoK8sApiAutoscalingV2PodsMetricStatus`

NewIoK8sApiAutoscalingV2PodsMetricStatusWithDefaults instantiates a new IoK8sApiAutoscalingV2PodsMetricStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *IoK8sApiAutoscalingV2PodsMetricStatus) GetCurrent() IoK8sApiAutoscalingV2MetricValueStatus`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *IoK8sApiAutoscalingV2PodsMetricStatus) GetCurrentOk() (*IoK8sApiAutoscalingV2MetricValueStatus, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *IoK8sApiAutoscalingV2PodsMetricStatus) SetCurrent(v IoK8sApiAutoscalingV2MetricValueStatus)`

SetCurrent sets Current field to given value.


### GetMetric

`func (o *IoK8sApiAutoscalingV2PodsMetricStatus) GetMetric() IoK8sApiAutoscalingV2MetricIdentifier`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *IoK8sApiAutoscalingV2PodsMetricStatus) GetMetricOk() (*IoK8sApiAutoscalingV2MetricIdentifier, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *IoK8sApiAutoscalingV2PodsMetricStatus) SetMetric(v IoK8sApiAutoscalingV2MetricIdentifier)`

SetMetric sets Metric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


