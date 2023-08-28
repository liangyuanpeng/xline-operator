# IoK8sApiPolicyV1PodDisruptionBudgetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IoK8sApimachineryPkgApisMetaV1Condition**](IoK8sApimachineryPkgApisMetaV1Condition.md) | Conditions contain conditions for PDB. The disruption controller sets the DisruptionAllowed condition. The following are known values for the reason field (additional reasons could be added in the future): - SyncFailed: The controller encountered an error and wasn&#39;t able to compute               the number of allowed disruptions. Therefore no disruptions are               allowed and the status of the condition will be False. - InsufficientPods: The number of pods are either at or below the number                     required by the PodDisruptionBudget. No disruptions are                     allowed and the status of the condition will be False. - SufficientPods: There are more pods than required by the PodDisruptionBudget.                   The condition will be True, and the number of allowed                   disruptions are provided by the disruptionsAllowed property. | [optional] 
**CurrentHealthy** | **int32** | current number of healthy pods | 
**DesiredHealthy** | **int32** | minimum desired number of healthy pods | 
**DisruptedPods** | Pointer to [**map[string]time.Time**](time.Time.md) | DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn&#39;t occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions. | [optional] 
**DisruptionsAllowed** | **int32** | Number of pod disruptions that are currently allowed. | 
**ExpectedPods** | **int32** | total number of pods counted by this disruption budget | 
**ObservedGeneration** | Pointer to **int64** | Most recent generation observed when updating this PDB status. DisruptionsAllowed and other status information is valid only if observedGeneration equals to PDB&#39;s object generation. | [optional] 

## Methods

### NewIoK8sApiPolicyV1PodDisruptionBudgetStatus

`func NewIoK8sApiPolicyV1PodDisruptionBudgetStatus(currentHealthy int32, desiredHealthy int32, disruptionsAllowed int32, expectedPods int32, ) *IoK8sApiPolicyV1PodDisruptionBudgetStatus`

NewIoK8sApiPolicyV1PodDisruptionBudgetStatus instantiates a new IoK8sApiPolicyV1PodDisruptionBudgetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiPolicyV1PodDisruptionBudgetStatusWithDefaults

`func NewIoK8sApiPolicyV1PodDisruptionBudgetStatusWithDefaults() *IoK8sApiPolicyV1PodDisruptionBudgetStatus`

NewIoK8sApiPolicyV1PodDisruptionBudgetStatusWithDefaults instantiates a new IoK8sApiPolicyV1PodDisruptionBudgetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetConditions() []IoK8sApimachineryPkgApisMetaV1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetConditionsOk() (*[]IoK8sApimachineryPkgApisMetaV1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) SetConditions(v []IoK8sApimachineryPkgApisMetaV1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCurrentHealthy

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetCurrentHealthy() int32`

GetCurrentHealthy returns the CurrentHealthy field if non-nil, zero value otherwise.

### GetCurrentHealthyOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetCurrentHealthyOk() (*int32, bool)`

GetCurrentHealthyOk returns a tuple with the CurrentHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentHealthy

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) SetCurrentHealthy(v int32)`

SetCurrentHealthy sets CurrentHealthy field to given value.


### GetDesiredHealthy

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetDesiredHealthy() int32`

GetDesiredHealthy returns the DesiredHealthy field if non-nil, zero value otherwise.

### GetDesiredHealthyOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetDesiredHealthyOk() (*int32, bool)`

GetDesiredHealthyOk returns a tuple with the DesiredHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredHealthy

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) SetDesiredHealthy(v int32)`

SetDesiredHealthy sets DesiredHealthy field to given value.


### GetDisruptedPods

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetDisruptedPods() map[string]time.Time`

GetDisruptedPods returns the DisruptedPods field if non-nil, zero value otherwise.

### GetDisruptedPodsOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetDisruptedPodsOk() (*map[string]time.Time, bool)`

GetDisruptedPodsOk returns a tuple with the DisruptedPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptedPods

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) SetDisruptedPods(v map[string]time.Time)`

SetDisruptedPods sets DisruptedPods field to given value.

### HasDisruptedPods

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) HasDisruptedPods() bool`

HasDisruptedPods returns a boolean if a field has been set.

### GetDisruptionsAllowed

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetDisruptionsAllowed() int32`

GetDisruptionsAllowed returns the DisruptionsAllowed field if non-nil, zero value otherwise.

### GetDisruptionsAllowedOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetDisruptionsAllowedOk() (*int32, bool)`

GetDisruptionsAllowedOk returns a tuple with the DisruptionsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptionsAllowed

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) SetDisruptionsAllowed(v int32)`

SetDisruptionsAllowed sets DisruptionsAllowed field to given value.


### GetExpectedPods

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetExpectedPods() int32`

GetExpectedPods returns the ExpectedPods field if non-nil, zero value otherwise.

### GetExpectedPodsOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetExpectedPodsOk() (*int32, bool)`

GetExpectedPodsOk returns a tuple with the ExpectedPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedPods

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) SetExpectedPods(v int32)`

SetExpectedPods sets ExpectedPods field to given value.


### GetObservedGeneration

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


