# IoK8sApiAppsV1ReplicaSetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | Pointer to **int32** | The number of available replicas (ready for at least minReadySeconds) for this replica set. | [optional] 
**Conditions** | Pointer to [**[]IoK8sApiAppsV1ReplicaSetCondition**](IoK8sApiAppsV1ReplicaSetCondition.md) | Represents the latest available observations of a replica set&#39;s current state. | [optional] 
**FullyLabeledReplicas** | Pointer to **int32** | The number of pods that have labels matching the labels of the pod template of the replicaset. | [optional] 
**ObservedGeneration** | Pointer to **int64** | ObservedGeneration reflects the generation of the most recently observed ReplicaSet. | [optional] 
**ReadyReplicas** | Pointer to **int32** | readyReplicas is the number of pods targeted by this ReplicaSet with a Ready Condition. | [optional] 
**Replicas** | **int32** | Replicas is the most recently observed number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller | 

## Methods

### NewIoK8sApiAppsV1ReplicaSetStatus

`func NewIoK8sApiAppsV1ReplicaSetStatus(replicas int32, ) *IoK8sApiAppsV1ReplicaSetStatus`

NewIoK8sApiAppsV1ReplicaSetStatus instantiates a new IoK8sApiAppsV1ReplicaSetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1ReplicaSetStatusWithDefaults

`func NewIoK8sApiAppsV1ReplicaSetStatusWithDefaults() *IoK8sApiAppsV1ReplicaSetStatus`

NewIoK8sApiAppsV1ReplicaSetStatusWithDefaults instantiates a new IoK8sApiAppsV1ReplicaSetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetConditions

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetConditions() []IoK8sApiAppsV1ReplicaSetCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetConditionsOk() (*[]IoK8sApiAppsV1ReplicaSetCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiAppsV1ReplicaSetStatus) SetConditions(v []IoK8sApiAppsV1ReplicaSetCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiAppsV1ReplicaSetStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetFullyLabeledReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetFullyLabeledReplicas() int32`

GetFullyLabeledReplicas returns the FullyLabeledReplicas field if non-nil, zero value otherwise.

### GetFullyLabeledReplicasOk

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetFullyLabeledReplicasOk() (*int32, bool)`

GetFullyLabeledReplicasOk returns a tuple with the FullyLabeledReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyLabeledReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) SetFullyLabeledReplicas(v int32)`

SetFullyLabeledReplicas sets FullyLabeledReplicas field to given value.

### HasFullyLabeledReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) HasFullyLabeledReplicas() bool`

HasFullyLabeledReplicas returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *IoK8sApiAppsV1ReplicaSetStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *IoK8sApiAppsV1ReplicaSetStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoK8sApiAppsV1ReplicaSetStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoK8sApiAppsV1ReplicaSetStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


