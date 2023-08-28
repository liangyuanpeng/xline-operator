# IoK8sApiAutoscalingV2MetricSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerResource** | Pointer to [**IoK8sApiAutoscalingV2ContainerResourceMetricSource**](IoK8sApiAutoscalingV2ContainerResourceMetricSource.md) |  | [optional] 
**External** | Pointer to [**IoK8sApiAutoscalingV2ExternalMetricSource**](IoK8sApiAutoscalingV2ExternalMetricSource.md) |  | [optional] 
**Object** | Pointer to [**IoK8sApiAutoscalingV2ObjectMetricSource**](IoK8sApiAutoscalingV2ObjectMetricSource.md) |  | [optional] 
**Pods** | Pointer to [**IoK8sApiAutoscalingV2PodsMetricSource**](IoK8sApiAutoscalingV2PodsMetricSource.md) |  | [optional] 
**Resource** | Pointer to [**IoK8sApiAutoscalingV2ResourceMetricSource**](IoK8sApiAutoscalingV2ResourceMetricSource.md) |  | [optional] 
**Type** | **string** | type is the type of metric source.  It should be one of \&quot;ContainerResource\&quot;, \&quot;External\&quot;, \&quot;Object\&quot;, \&quot;Pods\&quot; or \&quot;Resource\&quot;, each mapping to a matching field in the object. Note: \&quot;ContainerResource\&quot; type is available on when the feature-gate HPAContainerMetrics is enabled | 

## Methods

### NewIoK8sApiAutoscalingV2MetricSpec

`func NewIoK8sApiAutoscalingV2MetricSpec(type_ string, ) *IoK8sApiAutoscalingV2MetricSpec`

NewIoK8sApiAutoscalingV2MetricSpec instantiates a new IoK8sApiAutoscalingV2MetricSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2MetricSpecWithDefaults

`func NewIoK8sApiAutoscalingV2MetricSpecWithDefaults() *IoK8sApiAutoscalingV2MetricSpec`

NewIoK8sApiAutoscalingV2MetricSpecWithDefaults instantiates a new IoK8sApiAutoscalingV2MetricSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerResource

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetContainerResource() IoK8sApiAutoscalingV2ContainerResourceMetricSource`

GetContainerResource returns the ContainerResource field if non-nil, zero value otherwise.

### GetContainerResourceOk

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetContainerResourceOk() (*IoK8sApiAutoscalingV2ContainerResourceMetricSource, bool)`

GetContainerResourceOk returns a tuple with the ContainerResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerResource

`func (o *IoK8sApiAutoscalingV2MetricSpec) SetContainerResource(v IoK8sApiAutoscalingV2ContainerResourceMetricSource)`

SetContainerResource sets ContainerResource field to given value.

### HasContainerResource

`func (o *IoK8sApiAutoscalingV2MetricSpec) HasContainerResource() bool`

HasContainerResource returns a boolean if a field has been set.

### GetExternal

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetExternal() IoK8sApiAutoscalingV2ExternalMetricSource`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetExternalOk() (*IoK8sApiAutoscalingV2ExternalMetricSource, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *IoK8sApiAutoscalingV2MetricSpec) SetExternal(v IoK8sApiAutoscalingV2ExternalMetricSource)`

SetExternal sets External field to given value.

### HasExternal

`func (o *IoK8sApiAutoscalingV2MetricSpec) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetObject

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetObject() IoK8sApiAutoscalingV2ObjectMetricSource`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetObjectOk() (*IoK8sApiAutoscalingV2ObjectMetricSource, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IoK8sApiAutoscalingV2MetricSpec) SetObject(v IoK8sApiAutoscalingV2ObjectMetricSource)`

SetObject sets Object field to given value.

### HasObject

`func (o *IoK8sApiAutoscalingV2MetricSpec) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPods

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetPods() IoK8sApiAutoscalingV2PodsMetricSource`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetPodsOk() (*IoK8sApiAutoscalingV2PodsMetricSource, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *IoK8sApiAutoscalingV2MetricSpec) SetPods(v IoK8sApiAutoscalingV2PodsMetricSource)`

SetPods sets Pods field to given value.

### HasPods

`func (o *IoK8sApiAutoscalingV2MetricSpec) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetResource

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetResource() IoK8sApiAutoscalingV2ResourceMetricSource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetResourceOk() (*IoK8sApiAutoscalingV2ResourceMetricSource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IoK8sApiAutoscalingV2MetricSpec) SetResource(v IoK8sApiAutoscalingV2ResourceMetricSource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IoK8sApiAutoscalingV2MetricSpec) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiAutoscalingV2MetricSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiAutoscalingV2MetricSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


