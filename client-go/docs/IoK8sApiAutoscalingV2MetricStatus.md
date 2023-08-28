# IoK8sApiAutoscalingV2MetricStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerResource** | Pointer to [**IoK8sApiAutoscalingV2ContainerResourceMetricStatus**](IoK8sApiAutoscalingV2ContainerResourceMetricStatus.md) |  | [optional] 
**External** | Pointer to [**IoK8sApiAutoscalingV2ExternalMetricStatus**](IoK8sApiAutoscalingV2ExternalMetricStatus.md) |  | [optional] 
**Object** | Pointer to [**IoK8sApiAutoscalingV2ObjectMetricStatus**](IoK8sApiAutoscalingV2ObjectMetricStatus.md) |  | [optional] 
**Pods** | Pointer to [**IoK8sApiAutoscalingV2PodsMetricStatus**](IoK8sApiAutoscalingV2PodsMetricStatus.md) |  | [optional] 
**Resource** | Pointer to [**IoK8sApiAutoscalingV2ResourceMetricStatus**](IoK8sApiAutoscalingV2ResourceMetricStatus.md) |  | [optional] 
**Type** | **string** | type is the type of metric source.  It will be one of \&quot;ContainerResource\&quot;, \&quot;External\&quot;, \&quot;Object\&quot;, \&quot;Pods\&quot; or \&quot;Resource\&quot;, each corresponds to a matching field in the object. Note: \&quot;ContainerResource\&quot; type is available on when the feature-gate HPAContainerMetrics is enabled | 

## Methods

### NewIoK8sApiAutoscalingV2MetricStatus

`func NewIoK8sApiAutoscalingV2MetricStatus(type_ string, ) *IoK8sApiAutoscalingV2MetricStatus`

NewIoK8sApiAutoscalingV2MetricStatus instantiates a new IoK8sApiAutoscalingV2MetricStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2MetricStatusWithDefaults

`func NewIoK8sApiAutoscalingV2MetricStatusWithDefaults() *IoK8sApiAutoscalingV2MetricStatus`

NewIoK8sApiAutoscalingV2MetricStatusWithDefaults instantiates a new IoK8sApiAutoscalingV2MetricStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerResource

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetContainerResource() IoK8sApiAutoscalingV2ContainerResourceMetricStatus`

GetContainerResource returns the ContainerResource field if non-nil, zero value otherwise.

### GetContainerResourceOk

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetContainerResourceOk() (*IoK8sApiAutoscalingV2ContainerResourceMetricStatus, bool)`

GetContainerResourceOk returns a tuple with the ContainerResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerResource

`func (o *IoK8sApiAutoscalingV2MetricStatus) SetContainerResource(v IoK8sApiAutoscalingV2ContainerResourceMetricStatus)`

SetContainerResource sets ContainerResource field to given value.

### HasContainerResource

`func (o *IoK8sApiAutoscalingV2MetricStatus) HasContainerResource() bool`

HasContainerResource returns a boolean if a field has been set.

### GetExternal

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetExternal() IoK8sApiAutoscalingV2ExternalMetricStatus`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetExternalOk() (*IoK8sApiAutoscalingV2ExternalMetricStatus, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *IoK8sApiAutoscalingV2MetricStatus) SetExternal(v IoK8sApiAutoscalingV2ExternalMetricStatus)`

SetExternal sets External field to given value.

### HasExternal

`func (o *IoK8sApiAutoscalingV2MetricStatus) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetObject

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetObject() IoK8sApiAutoscalingV2ObjectMetricStatus`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetObjectOk() (*IoK8sApiAutoscalingV2ObjectMetricStatus, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IoK8sApiAutoscalingV2MetricStatus) SetObject(v IoK8sApiAutoscalingV2ObjectMetricStatus)`

SetObject sets Object field to given value.

### HasObject

`func (o *IoK8sApiAutoscalingV2MetricStatus) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPods

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetPods() IoK8sApiAutoscalingV2PodsMetricStatus`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetPodsOk() (*IoK8sApiAutoscalingV2PodsMetricStatus, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *IoK8sApiAutoscalingV2MetricStatus) SetPods(v IoK8sApiAutoscalingV2PodsMetricStatus)`

SetPods sets Pods field to given value.

### HasPods

`func (o *IoK8sApiAutoscalingV2MetricStatus) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetResource

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetResource() IoK8sApiAutoscalingV2ResourceMetricStatus`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetResourceOk() (*IoK8sApiAutoscalingV2ResourceMetricStatus, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IoK8sApiAutoscalingV2MetricStatus) SetResource(v IoK8sApiAutoscalingV2ResourceMetricStatus)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IoK8sApiAutoscalingV2MetricStatus) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiAutoscalingV2MetricStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiAutoscalingV2MetricStatus) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


