# IoK8sApiAppsV1DaemonSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | Pointer to **int32** | The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready). | [optional] 
**RevisionHistoryLimit** | Pointer to **int32** | The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10. | [optional] 
**Selector** | [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | 
**Template** | [**IoK8sApiCoreV1PodTemplateSpec**](IoK8sApiCoreV1PodTemplateSpec.md) |  | 
**UpdateStrategy** | Pointer to [**IoK8sApiAppsV1DaemonSetUpdateStrategy**](IoK8sApiAppsV1DaemonSetUpdateStrategy.md) |  | [optional] 

## Methods

### NewIoK8sApiAppsV1DaemonSetSpec

`func NewIoK8sApiAppsV1DaemonSetSpec(selector IoK8sApimachineryPkgApisMetaV1LabelSelector, template IoK8sApiCoreV1PodTemplateSpec, ) *IoK8sApiAppsV1DaemonSetSpec`

NewIoK8sApiAppsV1DaemonSetSpec instantiates a new IoK8sApiAppsV1DaemonSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1DaemonSetSpecWithDefaults

`func NewIoK8sApiAppsV1DaemonSetSpecWithDefaults() *IoK8sApiAppsV1DaemonSetSpec`

NewIoK8sApiAppsV1DaemonSetSpecWithDefaults instantiates a new IoK8sApiAppsV1DaemonSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinReadySeconds

`func (o *IoK8sApiAppsV1DaemonSetSpec) GetMinReadySeconds() int32`

GetMinReadySeconds returns the MinReadySeconds field if non-nil, zero value otherwise.

### GetMinReadySecondsOk

`func (o *IoK8sApiAppsV1DaemonSetSpec) GetMinReadySecondsOk() (*int32, bool)`

GetMinReadySecondsOk returns a tuple with the MinReadySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadySeconds

`func (o *IoK8sApiAppsV1DaemonSetSpec) SetMinReadySeconds(v int32)`

SetMinReadySeconds sets MinReadySeconds field to given value.

### HasMinReadySeconds

`func (o *IoK8sApiAppsV1DaemonSetSpec) HasMinReadySeconds() bool`

HasMinReadySeconds returns a boolean if a field has been set.

### GetRevisionHistoryLimit

`func (o *IoK8sApiAppsV1DaemonSetSpec) GetRevisionHistoryLimit() int32`

GetRevisionHistoryLimit returns the RevisionHistoryLimit field if non-nil, zero value otherwise.

### GetRevisionHistoryLimitOk

`func (o *IoK8sApiAppsV1DaemonSetSpec) GetRevisionHistoryLimitOk() (*int32, bool)`

GetRevisionHistoryLimitOk returns a tuple with the RevisionHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistoryLimit

`func (o *IoK8sApiAppsV1DaemonSetSpec) SetRevisionHistoryLimit(v int32)`

SetRevisionHistoryLimit sets RevisionHistoryLimit field to given value.

### HasRevisionHistoryLimit

`func (o *IoK8sApiAppsV1DaemonSetSpec) HasRevisionHistoryLimit() bool`

HasRevisionHistoryLimit returns a boolean if a field has been set.

### GetSelector

`func (o *IoK8sApiAppsV1DaemonSetSpec) GetSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiAppsV1DaemonSetSpec) GetSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiAppsV1DaemonSetSpec) SetSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetSelector sets Selector field to given value.


### GetTemplate

`func (o *IoK8sApiAppsV1DaemonSetSpec) GetTemplate() IoK8sApiCoreV1PodTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoK8sApiAppsV1DaemonSetSpec) GetTemplateOk() (*IoK8sApiCoreV1PodTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoK8sApiAppsV1DaemonSetSpec) SetTemplate(v IoK8sApiCoreV1PodTemplateSpec)`

SetTemplate sets Template field to given value.


### GetUpdateStrategy

`func (o *IoK8sApiAppsV1DaemonSetSpec) GetUpdateStrategy() IoK8sApiAppsV1DaemonSetUpdateStrategy`

GetUpdateStrategy returns the UpdateStrategy field if non-nil, zero value otherwise.

### GetUpdateStrategyOk

`func (o *IoK8sApiAppsV1DaemonSetSpec) GetUpdateStrategyOk() (*IoK8sApiAppsV1DaemonSetUpdateStrategy, bool)`

GetUpdateStrategyOk returns a tuple with the UpdateStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStrategy

`func (o *IoK8sApiAppsV1DaemonSetSpec) SetUpdateStrategy(v IoK8sApiAppsV1DaemonSetUpdateStrategy)`

SetUpdateStrategy sets UpdateStrategy field to given value.

### HasUpdateStrategy

`func (o *IoK8sApiAppsV1DaemonSetSpec) HasUpdateStrategy() bool`

HasUpdateStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


