# IoK8sApiAppsV1ReplicaSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | Pointer to **int32** | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) | [optional] 
**Replicas** | Pointer to **int32** | Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller | [optional] 
**Selector** | [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | 
**Template** | Pointer to [**IoK8sApiCoreV1PodTemplateSpec**](IoK8sApiCoreV1PodTemplateSpec.md) |  | [optional] 

## Methods

### NewIoK8sApiAppsV1ReplicaSetSpec

`func NewIoK8sApiAppsV1ReplicaSetSpec(selector IoK8sApimachineryPkgApisMetaV1LabelSelector, ) *IoK8sApiAppsV1ReplicaSetSpec`

NewIoK8sApiAppsV1ReplicaSetSpec instantiates a new IoK8sApiAppsV1ReplicaSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1ReplicaSetSpecWithDefaults

`func NewIoK8sApiAppsV1ReplicaSetSpecWithDefaults() *IoK8sApiAppsV1ReplicaSetSpec`

NewIoK8sApiAppsV1ReplicaSetSpecWithDefaults instantiates a new IoK8sApiAppsV1ReplicaSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinReadySeconds

`func (o *IoK8sApiAppsV1ReplicaSetSpec) GetMinReadySeconds() int32`

GetMinReadySeconds returns the MinReadySeconds field if non-nil, zero value otherwise.

### GetMinReadySecondsOk

`func (o *IoK8sApiAppsV1ReplicaSetSpec) GetMinReadySecondsOk() (*int32, bool)`

GetMinReadySecondsOk returns a tuple with the MinReadySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadySeconds

`func (o *IoK8sApiAppsV1ReplicaSetSpec) SetMinReadySeconds(v int32)`

SetMinReadySeconds sets MinReadySeconds field to given value.

### HasMinReadySeconds

`func (o *IoK8sApiAppsV1ReplicaSetSpec) HasMinReadySeconds() bool`

HasMinReadySeconds returns a boolean if a field has been set.

### GetReplicas

`func (o *IoK8sApiAppsV1ReplicaSetSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoK8sApiAppsV1ReplicaSetSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoK8sApiAppsV1ReplicaSetSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *IoK8sApiAppsV1ReplicaSetSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetSelector

`func (o *IoK8sApiAppsV1ReplicaSetSpec) GetSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiAppsV1ReplicaSetSpec) GetSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiAppsV1ReplicaSetSpec) SetSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetSelector sets Selector field to given value.


### GetTemplate

`func (o *IoK8sApiAppsV1ReplicaSetSpec) GetTemplate() IoK8sApiCoreV1PodTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoK8sApiAppsV1ReplicaSetSpec) GetTemplateOk() (*IoK8sApiCoreV1PodTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoK8sApiAppsV1ReplicaSetSpec) SetTemplate(v IoK8sApiCoreV1PodTemplateSpec)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IoK8sApiAppsV1ReplicaSetSpec) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


