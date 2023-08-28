# IoK8sApiCoreV1ReplicationControllerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | Pointer to **int32** | The number of available replicas (ready for at least minReadySeconds) for this replication controller. | [optional] 
**Conditions** | Pointer to [**[]IoK8sApiCoreV1ReplicationControllerCondition**](IoK8sApiCoreV1ReplicationControllerCondition.md) | Represents the latest available observations of a replication controller&#39;s current state. | [optional] 
**FullyLabeledReplicas** | Pointer to **int32** | The number of pods that have labels matching the labels of the pod template of the replication controller. | [optional] 
**ObservedGeneration** | Pointer to **int64** | ObservedGeneration reflects the generation of the most recently observed replication controller. | [optional] 
**ReadyReplicas** | Pointer to **int32** | The number of ready replicas for this replication controller. | [optional] 
**Replicas** | **int32** | Replicas is the most recently observed number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller | 

## Methods

### NewIoK8sApiCoreV1ReplicationControllerStatus

`func NewIoK8sApiCoreV1ReplicationControllerStatus(replicas int32, ) *IoK8sApiCoreV1ReplicationControllerStatus`

NewIoK8sApiCoreV1ReplicationControllerStatus instantiates a new IoK8sApiCoreV1ReplicationControllerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ReplicationControllerStatusWithDefaults

`func NewIoK8sApiCoreV1ReplicationControllerStatusWithDefaults() *IoK8sApiCoreV1ReplicationControllerStatus`

NewIoK8sApiCoreV1ReplicationControllerStatusWithDefaults instantiates a new IoK8sApiCoreV1ReplicationControllerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetConditions

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetConditions() []IoK8sApiCoreV1ReplicationControllerCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetConditionsOk() (*[]IoK8sApiCoreV1ReplicationControllerCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) SetConditions(v []IoK8sApiCoreV1ReplicationControllerCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetFullyLabeledReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetFullyLabeledReplicas() int32`

GetFullyLabeledReplicas returns the FullyLabeledReplicas field if non-nil, zero value otherwise.

### GetFullyLabeledReplicasOk

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetFullyLabeledReplicasOk() (*int32, bool)`

GetFullyLabeledReplicasOk returns a tuple with the FullyLabeledReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyLabeledReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) SetFullyLabeledReplicas(v int32)`

SetFullyLabeledReplicas sets FullyLabeledReplicas field to given value.

### HasFullyLabeledReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) HasFullyLabeledReplicas() bool`

HasFullyLabeledReplicas returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoK8sApiCoreV1ReplicationControllerStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


