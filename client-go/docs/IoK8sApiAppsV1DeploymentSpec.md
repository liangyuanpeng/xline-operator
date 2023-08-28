# IoK8sApiAppsV1DeploymentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | Pointer to **int32** | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) | [optional] 
**Paused** | Pointer to **bool** | Indicates that the deployment is paused. | [optional] 
**ProgressDeadlineSeconds** | Pointer to **int32** | The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s. | [optional] 
**Replicas** | Pointer to **int32** | Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1. | [optional] 
**RevisionHistoryLimit** | Pointer to **int32** | The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10. | [optional] 
**Selector** | [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | 
**Strategy** | Pointer to [**IoK8sApiAppsV1DeploymentStrategy**](IoK8sApiAppsV1DeploymentStrategy.md) |  | [optional] 
**Template** | [**IoK8sApiCoreV1PodTemplateSpec**](IoK8sApiCoreV1PodTemplateSpec.md) |  | 

## Methods

### NewIoK8sApiAppsV1DeploymentSpec

`func NewIoK8sApiAppsV1DeploymentSpec(selector IoK8sApimachineryPkgApisMetaV1LabelSelector, template IoK8sApiCoreV1PodTemplateSpec, ) *IoK8sApiAppsV1DeploymentSpec`

NewIoK8sApiAppsV1DeploymentSpec instantiates a new IoK8sApiAppsV1DeploymentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1DeploymentSpecWithDefaults

`func NewIoK8sApiAppsV1DeploymentSpecWithDefaults() *IoK8sApiAppsV1DeploymentSpec`

NewIoK8sApiAppsV1DeploymentSpecWithDefaults instantiates a new IoK8sApiAppsV1DeploymentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinReadySeconds

`func (o *IoK8sApiAppsV1DeploymentSpec) GetMinReadySeconds() int32`

GetMinReadySeconds returns the MinReadySeconds field if non-nil, zero value otherwise.

### GetMinReadySecondsOk

`func (o *IoK8sApiAppsV1DeploymentSpec) GetMinReadySecondsOk() (*int32, bool)`

GetMinReadySecondsOk returns a tuple with the MinReadySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadySeconds

`func (o *IoK8sApiAppsV1DeploymentSpec) SetMinReadySeconds(v int32)`

SetMinReadySeconds sets MinReadySeconds field to given value.

### HasMinReadySeconds

`func (o *IoK8sApiAppsV1DeploymentSpec) HasMinReadySeconds() bool`

HasMinReadySeconds returns a boolean if a field has been set.

### GetPaused

`func (o *IoK8sApiAppsV1DeploymentSpec) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *IoK8sApiAppsV1DeploymentSpec) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *IoK8sApiAppsV1DeploymentSpec) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *IoK8sApiAppsV1DeploymentSpec) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetProgressDeadlineSeconds

`func (o *IoK8sApiAppsV1DeploymentSpec) GetProgressDeadlineSeconds() int32`

GetProgressDeadlineSeconds returns the ProgressDeadlineSeconds field if non-nil, zero value otherwise.

### GetProgressDeadlineSecondsOk

`func (o *IoK8sApiAppsV1DeploymentSpec) GetProgressDeadlineSecondsOk() (*int32, bool)`

GetProgressDeadlineSecondsOk returns a tuple with the ProgressDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressDeadlineSeconds

`func (o *IoK8sApiAppsV1DeploymentSpec) SetProgressDeadlineSeconds(v int32)`

SetProgressDeadlineSeconds sets ProgressDeadlineSeconds field to given value.

### HasProgressDeadlineSeconds

`func (o *IoK8sApiAppsV1DeploymentSpec) HasProgressDeadlineSeconds() bool`

HasProgressDeadlineSeconds returns a boolean if a field has been set.

### GetReplicas

`func (o *IoK8sApiAppsV1DeploymentSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoK8sApiAppsV1DeploymentSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoK8sApiAppsV1DeploymentSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *IoK8sApiAppsV1DeploymentSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRevisionHistoryLimit

`func (o *IoK8sApiAppsV1DeploymentSpec) GetRevisionHistoryLimit() int32`

GetRevisionHistoryLimit returns the RevisionHistoryLimit field if non-nil, zero value otherwise.

### GetRevisionHistoryLimitOk

`func (o *IoK8sApiAppsV1DeploymentSpec) GetRevisionHistoryLimitOk() (*int32, bool)`

GetRevisionHistoryLimitOk returns a tuple with the RevisionHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistoryLimit

`func (o *IoK8sApiAppsV1DeploymentSpec) SetRevisionHistoryLimit(v int32)`

SetRevisionHistoryLimit sets RevisionHistoryLimit field to given value.

### HasRevisionHistoryLimit

`func (o *IoK8sApiAppsV1DeploymentSpec) HasRevisionHistoryLimit() bool`

HasRevisionHistoryLimit returns a boolean if a field has been set.

### GetSelector

`func (o *IoK8sApiAppsV1DeploymentSpec) GetSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiAppsV1DeploymentSpec) GetSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiAppsV1DeploymentSpec) SetSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetSelector sets Selector field to given value.


### GetStrategy

`func (o *IoK8sApiAppsV1DeploymentSpec) GetStrategy() IoK8sApiAppsV1DeploymentStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *IoK8sApiAppsV1DeploymentSpec) GetStrategyOk() (*IoK8sApiAppsV1DeploymentStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *IoK8sApiAppsV1DeploymentSpec) SetStrategy(v IoK8sApiAppsV1DeploymentStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *IoK8sApiAppsV1DeploymentSpec) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetTemplate

`func (o *IoK8sApiAppsV1DeploymentSpec) GetTemplate() IoK8sApiCoreV1PodTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoK8sApiAppsV1DeploymentSpec) GetTemplateOk() (*IoK8sApiCoreV1PodTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoK8sApiAppsV1DeploymentSpec) SetTemplate(v IoK8sApiCoreV1PodTemplateSpec)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


