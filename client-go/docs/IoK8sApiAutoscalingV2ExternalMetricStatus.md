# IoK8sApiAutoscalingV2ExternalMetricStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | [**IoK8sApiAutoscalingV2MetricValueStatus**](IoK8sApiAutoscalingV2MetricValueStatus.md) |  | 
**Metric** | [**IoK8sApiAutoscalingV2MetricIdentifier**](IoK8sApiAutoscalingV2MetricIdentifier.md) |  | 

## Methods

### NewIoK8sApiAutoscalingV2ExternalMetricStatus

`func NewIoK8sApiAutoscalingV2ExternalMetricStatus(current IoK8sApiAutoscalingV2MetricValueStatus, metric IoK8sApiAutoscalingV2MetricIdentifier, ) *IoK8sApiAutoscalingV2ExternalMetricStatus`

NewIoK8sApiAutoscalingV2ExternalMetricStatus instantiates a new IoK8sApiAutoscalingV2ExternalMetricStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2ExternalMetricStatusWithDefaults

`func NewIoK8sApiAutoscalingV2ExternalMetricStatusWithDefaults() *IoK8sApiAutoscalingV2ExternalMetricStatus`

NewIoK8sApiAutoscalingV2ExternalMetricStatusWithDefaults instantiates a new IoK8sApiAutoscalingV2ExternalMetricStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *IoK8sApiAutoscalingV2ExternalMetricStatus) GetCurrent() IoK8sApiAutoscalingV2MetricValueStatus`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *IoK8sApiAutoscalingV2ExternalMetricStatus) GetCurrentOk() (*IoK8sApiAutoscalingV2MetricValueStatus, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *IoK8sApiAutoscalingV2ExternalMetricStatus) SetCurrent(v IoK8sApiAutoscalingV2MetricValueStatus)`

SetCurrent sets Current field to given value.


### GetMetric

`func (o *IoK8sApiAutoscalingV2ExternalMetricStatus) GetMetric() IoK8sApiAutoscalingV2MetricIdentifier`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *IoK8sApiAutoscalingV2ExternalMetricStatus) GetMetricOk() (*IoK8sApiAutoscalingV2MetricIdentifier, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *IoK8sApiAutoscalingV2ExternalMetricStatus) SetMetric(v IoK8sApiAutoscalingV2MetricIdentifier)`

SetMetric sets Metric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


