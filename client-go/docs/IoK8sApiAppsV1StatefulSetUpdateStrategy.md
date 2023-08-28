# IoK8sApiAppsV1StatefulSetUpdateStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollingUpdate** | Pointer to [**IoK8sApiAppsV1RollingUpdateStatefulSetStrategy**](IoK8sApiAppsV1RollingUpdateStatefulSetStrategy.md) |  | [optional] 
**Type** | Pointer to **string** | Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.  Possible enum values:  - &#x60;\&quot;OnDelete\&quot;&#x60; triggers the legacy behavior. Version tracking and ordered rolling restarts are disabled. Pods are recreated from the StatefulSetSpec when they are manually deleted. When a scale operation is performed with this strategy,specification version indicated by the StatefulSet&#39;s currentRevision.  - &#x60;\&quot;RollingUpdate\&quot;&#x60; indicates that update will be applied to all Pods in the StatefulSet with respect to the StatefulSet ordering constraints. When a scale operation is performed with this strategy, new Pods will be created from the specification version indicated by the StatefulSet&#39;s updateRevision. | [optional] 

## Methods

### NewIoK8sApiAppsV1StatefulSetUpdateStrategy

`func NewIoK8sApiAppsV1StatefulSetUpdateStrategy() *IoK8sApiAppsV1StatefulSetUpdateStrategy`

NewIoK8sApiAppsV1StatefulSetUpdateStrategy instantiates a new IoK8sApiAppsV1StatefulSetUpdateStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1StatefulSetUpdateStrategyWithDefaults

`func NewIoK8sApiAppsV1StatefulSetUpdateStrategyWithDefaults() *IoK8sApiAppsV1StatefulSetUpdateStrategy`

NewIoK8sApiAppsV1StatefulSetUpdateStrategyWithDefaults instantiates a new IoK8sApiAppsV1StatefulSetUpdateStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollingUpdate

`func (o *IoK8sApiAppsV1StatefulSetUpdateStrategy) GetRollingUpdate() IoK8sApiAppsV1RollingUpdateStatefulSetStrategy`

GetRollingUpdate returns the RollingUpdate field if non-nil, zero value otherwise.

### GetRollingUpdateOk

`func (o *IoK8sApiAppsV1StatefulSetUpdateStrategy) GetRollingUpdateOk() (*IoK8sApiAppsV1RollingUpdateStatefulSetStrategy, bool)`

GetRollingUpdateOk returns a tuple with the RollingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpdate

`func (o *IoK8sApiAppsV1StatefulSetUpdateStrategy) SetRollingUpdate(v IoK8sApiAppsV1RollingUpdateStatefulSetStrategy)`

SetRollingUpdate sets RollingUpdate field to given value.

### HasRollingUpdate

`func (o *IoK8sApiAppsV1StatefulSetUpdateStrategy) HasRollingUpdate() bool`

HasRollingUpdate returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiAppsV1StatefulSetUpdateStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiAppsV1StatefulSetUpdateStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiAppsV1StatefulSetUpdateStrategy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiAppsV1StatefulSetUpdateStrategy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


