# IoK8sApiCoreV1ReplicationControllerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | Pointer to **int32** | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) | [optional] 
**Replicas** | Pointer to **int32** | Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller | [optional] 
**Selector** | Pointer to **map[string]string** | Selector is a label query over pods that should match the Replicas count. If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors | [optional] 
**Template** | Pointer to [**IoK8sApiCoreV1PodTemplateSpec**](IoK8sApiCoreV1PodTemplateSpec.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1ReplicationControllerSpec

`func NewIoK8sApiCoreV1ReplicationControllerSpec() *IoK8sApiCoreV1ReplicationControllerSpec`

NewIoK8sApiCoreV1ReplicationControllerSpec instantiates a new IoK8sApiCoreV1ReplicationControllerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ReplicationControllerSpecWithDefaults

`func NewIoK8sApiCoreV1ReplicationControllerSpecWithDefaults() *IoK8sApiCoreV1ReplicationControllerSpec`

NewIoK8sApiCoreV1ReplicationControllerSpecWithDefaults instantiates a new IoK8sApiCoreV1ReplicationControllerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinReadySeconds

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) GetMinReadySeconds() int32`

GetMinReadySeconds returns the MinReadySeconds field if non-nil, zero value otherwise.

### GetMinReadySecondsOk

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) GetMinReadySecondsOk() (*int32, bool)`

GetMinReadySecondsOk returns a tuple with the MinReadySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadySeconds

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) SetMinReadySeconds(v int32)`

SetMinReadySeconds sets MinReadySeconds field to given value.

### HasMinReadySeconds

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) HasMinReadySeconds() bool`

HasMinReadySeconds returns a boolean if a field has been set.

### GetReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetSelector

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) GetSelector() map[string]string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) GetSelectorOk() (*map[string]string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) SetSelector(v map[string]string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetTemplate

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) GetTemplate() IoK8sApiCoreV1PodTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) GetTemplateOk() (*IoK8sApiCoreV1PodTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) SetTemplate(v IoK8sApiCoreV1PodTemplateSpec)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IoK8sApiCoreV1ReplicationControllerSpec) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


