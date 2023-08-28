# IoK8sApiAppsV1DeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | Pointer to **int32** | Total number of available pods (ready for at least minReadySeconds) targeted by this deployment. | [optional] 
**CollisionCount** | Pointer to **int32** | Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet. | [optional] 
**Conditions** | Pointer to [**[]IoK8sApiAppsV1DeploymentCondition**](IoK8sApiAppsV1DeploymentCondition.md) | Represents the latest available observations of a deployment&#39;s current state. | [optional] 
**ObservedGeneration** | Pointer to **int64** | The generation observed by the deployment controller. | [optional] 
**ReadyReplicas** | Pointer to **int32** | readyReplicas is the number of pods targeted by this Deployment with a Ready Condition. | [optional] 
**Replicas** | Pointer to **int32** | Total number of non-terminated pods targeted by this deployment (their labels match the selector). | [optional] 
**UnavailableReplicas** | Pointer to **int32** | Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created. | [optional] 
**UpdatedReplicas** | Pointer to **int32** | Total number of non-terminated pods targeted by this deployment that have the desired template spec. | [optional] 

## Methods

### NewIoK8sApiAppsV1DeploymentStatus

`func NewIoK8sApiAppsV1DeploymentStatus() *IoK8sApiAppsV1DeploymentStatus`

NewIoK8sApiAppsV1DeploymentStatus instantiates a new IoK8sApiAppsV1DeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1DeploymentStatusWithDefaults

`func NewIoK8sApiAppsV1DeploymentStatusWithDefaults() *IoK8sApiAppsV1DeploymentStatus`

NewIoK8sApiAppsV1DeploymentStatusWithDefaults instantiates a new IoK8sApiAppsV1DeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *IoK8sApiAppsV1DeploymentStatus) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetCollisionCount

`func (o *IoK8sApiAppsV1DeploymentStatus) GetCollisionCount() int32`

GetCollisionCount returns the CollisionCount field if non-nil, zero value otherwise.

### GetCollisionCountOk

`func (o *IoK8sApiAppsV1DeploymentStatus) GetCollisionCountOk() (*int32, bool)`

GetCollisionCountOk returns a tuple with the CollisionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollisionCount

`func (o *IoK8sApiAppsV1DeploymentStatus) SetCollisionCount(v int32)`

SetCollisionCount sets CollisionCount field to given value.

### HasCollisionCount

`func (o *IoK8sApiAppsV1DeploymentStatus) HasCollisionCount() bool`

HasCollisionCount returns a boolean if a field has been set.

### GetConditions

`func (o *IoK8sApiAppsV1DeploymentStatus) GetConditions() []IoK8sApiAppsV1DeploymentCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiAppsV1DeploymentStatus) GetConditionsOk() (*[]IoK8sApiAppsV1DeploymentCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiAppsV1DeploymentStatus) SetConditions(v []IoK8sApiAppsV1DeploymentCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiAppsV1DeploymentStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *IoK8sApiAppsV1DeploymentStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *IoK8sApiAppsV1DeploymentStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *IoK8sApiAppsV1DeploymentStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *IoK8sApiAppsV1DeploymentStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *IoK8sApiAppsV1DeploymentStatus) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoK8sApiAppsV1DeploymentStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetUnavailableReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) GetUnavailableReplicas() int32`

GetUnavailableReplicas returns the UnavailableReplicas field if non-nil, zero value otherwise.

### GetUnavailableReplicasOk

`func (o *IoK8sApiAppsV1DeploymentStatus) GetUnavailableReplicasOk() (*int32, bool)`

GetUnavailableReplicasOk returns a tuple with the UnavailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) SetUnavailableReplicas(v int32)`

SetUnavailableReplicas sets UnavailableReplicas field to given value.

### HasUnavailableReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) HasUnavailableReplicas() bool`

HasUnavailableReplicas returns a boolean if a field has been set.

### GetUpdatedReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) GetUpdatedReplicas() int32`

GetUpdatedReplicas returns the UpdatedReplicas field if non-nil, zero value otherwise.

### GetUpdatedReplicasOk

`func (o *IoK8sApiAppsV1DeploymentStatus) GetUpdatedReplicasOk() (*int32, bool)`

GetUpdatedReplicasOk returns a tuple with the UpdatedReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) SetUpdatedReplicas(v int32)`

SetUpdatedReplicas sets UpdatedReplicas field to given value.

### HasUpdatedReplicas

`func (o *IoK8sApiAppsV1DeploymentStatus) HasUpdatedReplicas() bool`

HasUpdatedReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


