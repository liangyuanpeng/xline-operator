# IoK8sApiAutoscalingV2ContainerResourceMetricSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | **string** | container is the name of the container in the pods of the scaling target | 
**Name** | **string** | name is the name of the resource in question. | 
**Target** | [**IoK8sApiAutoscalingV2MetricTarget**](IoK8sApiAutoscalingV2MetricTarget.md) |  | 

## Methods

### NewIoK8sApiAutoscalingV2ContainerResourceMetricSource

`func NewIoK8sApiAutoscalingV2ContainerResourceMetricSource(container string, name string, target IoK8sApiAutoscalingV2MetricTarget, ) *IoK8sApiAutoscalingV2ContainerResourceMetricSource`

NewIoK8sApiAutoscalingV2ContainerResourceMetricSource instantiates a new IoK8sApiAutoscalingV2ContainerResourceMetricSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2ContainerResourceMetricSourceWithDefaults

`func NewIoK8sApiAutoscalingV2ContainerResourceMetricSourceWithDefaults() *IoK8sApiAutoscalingV2ContainerResourceMetricSource`

NewIoK8sApiAutoscalingV2ContainerResourceMetricSourceWithDefaults instantiates a new IoK8sApiAutoscalingV2ContainerResourceMetricSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *IoK8sApiAutoscalingV2ContainerResourceMetricSource) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *IoK8sApiAutoscalingV2ContainerResourceMetricSource) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *IoK8sApiAutoscalingV2ContainerResourceMetricSource) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetName

`func (o *IoK8sApiAutoscalingV2ContainerResourceMetricSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiAutoscalingV2ContainerResourceMetricSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiAutoscalingV2ContainerResourceMetricSource) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *IoK8sApiAutoscalingV2ContainerResourceMetricSource) GetTarget() IoK8sApiAutoscalingV2MetricTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IoK8sApiAutoscalingV2ContainerResourceMetricSource) GetTargetOk() (*IoK8sApiAutoscalingV2MetricTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IoK8sApiAutoscalingV2ContainerResourceMetricSource) SetTarget(v IoK8sApiAutoscalingV2MetricTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


