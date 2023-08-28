# IoK8sApiAutoscalingV2ResourceMetricStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | [**IoK8sApiAutoscalingV2MetricValueStatus**](IoK8sApiAutoscalingV2MetricValueStatus.md) |  | 
**Name** | **string** | name is the name of the resource in question. | 

## Methods

### NewIoK8sApiAutoscalingV2ResourceMetricStatus

`func NewIoK8sApiAutoscalingV2ResourceMetricStatus(current IoK8sApiAutoscalingV2MetricValueStatus, name string, ) *IoK8sApiAutoscalingV2ResourceMetricStatus`

NewIoK8sApiAutoscalingV2ResourceMetricStatus instantiates a new IoK8sApiAutoscalingV2ResourceMetricStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2ResourceMetricStatusWithDefaults

`func NewIoK8sApiAutoscalingV2ResourceMetricStatusWithDefaults() *IoK8sApiAutoscalingV2ResourceMetricStatus`

NewIoK8sApiAutoscalingV2ResourceMetricStatusWithDefaults instantiates a new IoK8sApiAutoscalingV2ResourceMetricStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *IoK8sApiAutoscalingV2ResourceMetricStatus) GetCurrent() IoK8sApiAutoscalingV2MetricValueStatus`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *IoK8sApiAutoscalingV2ResourceMetricStatus) GetCurrentOk() (*IoK8sApiAutoscalingV2MetricValueStatus, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *IoK8sApiAutoscalingV2ResourceMetricStatus) SetCurrent(v IoK8sApiAutoscalingV2MetricValueStatus)`

SetCurrent sets Current field to given value.


### GetName

`func (o *IoK8sApiAutoscalingV2ResourceMetricStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiAutoscalingV2ResourceMetricStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiAutoscalingV2ResourceMetricStatus) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


