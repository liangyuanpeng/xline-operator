# IoK8sApiAutoscalingV2ObjectMetricStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | [**IoK8sApiAutoscalingV2MetricValueStatus**](IoK8sApiAutoscalingV2MetricValueStatus.md) |  | 
**DescribedObject** | [**IoK8sApiAutoscalingV2CrossVersionObjectReference**](IoK8sApiAutoscalingV2CrossVersionObjectReference.md) |  | 
**Metric** | [**IoK8sApiAutoscalingV2MetricIdentifier**](IoK8sApiAutoscalingV2MetricIdentifier.md) |  | 

## Methods

### NewIoK8sApiAutoscalingV2ObjectMetricStatus

`func NewIoK8sApiAutoscalingV2ObjectMetricStatus(current IoK8sApiAutoscalingV2MetricValueStatus, describedObject IoK8sApiAutoscalingV2CrossVersionObjectReference, metric IoK8sApiAutoscalingV2MetricIdentifier, ) *IoK8sApiAutoscalingV2ObjectMetricStatus`

NewIoK8sApiAutoscalingV2ObjectMetricStatus instantiates a new IoK8sApiAutoscalingV2ObjectMetricStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2ObjectMetricStatusWithDefaults

`func NewIoK8sApiAutoscalingV2ObjectMetricStatusWithDefaults() *IoK8sApiAutoscalingV2ObjectMetricStatus`

NewIoK8sApiAutoscalingV2ObjectMetricStatusWithDefaults instantiates a new IoK8sApiAutoscalingV2ObjectMetricStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *IoK8sApiAutoscalingV2ObjectMetricStatus) GetCurrent() IoK8sApiAutoscalingV2MetricValueStatus`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *IoK8sApiAutoscalingV2ObjectMetricStatus) GetCurrentOk() (*IoK8sApiAutoscalingV2MetricValueStatus, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *IoK8sApiAutoscalingV2ObjectMetricStatus) SetCurrent(v IoK8sApiAutoscalingV2MetricValueStatus)`

SetCurrent sets Current field to given value.


### GetDescribedObject

`func (o *IoK8sApiAutoscalingV2ObjectMetricStatus) GetDescribedObject() IoK8sApiAutoscalingV2CrossVersionObjectReference`

GetDescribedObject returns the DescribedObject field if non-nil, zero value otherwise.

### GetDescribedObjectOk

`func (o *IoK8sApiAutoscalingV2ObjectMetricStatus) GetDescribedObjectOk() (*IoK8sApiAutoscalingV2CrossVersionObjectReference, bool)`

GetDescribedObjectOk returns a tuple with the DescribedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribedObject

`func (o *IoK8sApiAutoscalingV2ObjectMetricStatus) SetDescribedObject(v IoK8sApiAutoscalingV2CrossVersionObjectReference)`

SetDescribedObject sets DescribedObject field to given value.


### GetMetric

`func (o *IoK8sApiAutoscalingV2ObjectMetricStatus) GetMetric() IoK8sApiAutoscalingV2MetricIdentifier`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *IoK8sApiAutoscalingV2ObjectMetricStatus) GetMetricOk() (*IoK8sApiAutoscalingV2MetricIdentifier, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *IoK8sApiAutoscalingV2ObjectMetricStatus) SetMetric(v IoK8sApiAutoscalingV2MetricIdentifier)`

SetMetric sets Metric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


