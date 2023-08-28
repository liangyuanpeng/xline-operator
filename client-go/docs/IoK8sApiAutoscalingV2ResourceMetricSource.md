# IoK8sApiAutoscalingV2ResourceMetricSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is the name of the resource in question. | 
**Target** | [**IoK8sApiAutoscalingV2MetricTarget**](IoK8sApiAutoscalingV2MetricTarget.md) |  | 

## Methods

### NewIoK8sApiAutoscalingV2ResourceMetricSource

`func NewIoK8sApiAutoscalingV2ResourceMetricSource(name string, target IoK8sApiAutoscalingV2MetricTarget, ) *IoK8sApiAutoscalingV2ResourceMetricSource`

NewIoK8sApiAutoscalingV2ResourceMetricSource instantiates a new IoK8sApiAutoscalingV2ResourceMetricSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2ResourceMetricSourceWithDefaults

`func NewIoK8sApiAutoscalingV2ResourceMetricSourceWithDefaults() *IoK8sApiAutoscalingV2ResourceMetricSource`

NewIoK8sApiAutoscalingV2ResourceMetricSourceWithDefaults instantiates a new IoK8sApiAutoscalingV2ResourceMetricSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiAutoscalingV2ResourceMetricSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiAutoscalingV2ResourceMetricSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiAutoscalingV2ResourceMetricSource) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *IoK8sApiAutoscalingV2ResourceMetricSource) GetTarget() IoK8sApiAutoscalingV2MetricTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IoK8sApiAutoscalingV2ResourceMetricSource) GetTargetOk() (*IoK8sApiAutoscalingV2MetricTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IoK8sApiAutoscalingV2ResourceMetricSource) SetTarget(v IoK8sApiAutoscalingV2MetricTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


