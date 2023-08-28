# IoK8sApiAppsV1StatefulSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | Pointer to **int32** | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) | [optional] 
**Ordinals** | Pointer to [**IoK8sApiAppsV1StatefulSetOrdinals**](IoK8sApiAppsV1StatefulSetOrdinals.md) |  | [optional] 
**PersistentVolumeClaimRetentionPolicy** | Pointer to [**IoK8sApiAppsV1StatefulSetPersistentVolumeClaimRetentionPolicy**](IoK8sApiAppsV1StatefulSetPersistentVolumeClaimRetentionPolicy.md) |  | [optional] 
**PodManagementPolicy** | Pointer to **string** | podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is &#x60;OrderedReady&#x60;, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is &#x60;Parallel&#x60; which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.  Possible enum values:  - &#x60;\&quot;OrderedReady\&quot;&#x60; will create pods in strictly increasing order on scale up and strictly decreasing order on scale down, progressing only when the previous pod is ready or terminated. At most one pod will be changed at any time.  - &#x60;\&quot;Parallel\&quot;&#x60; will create and delete pods as soon as the stateful set replica count is changed, and will not wait for pods to be ready or complete termination. | [optional] 
**Replicas** | Pointer to **int32** | replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1. | [optional] 
**RevisionHistoryLimit** | Pointer to **int32** | revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet&#39;s revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10. | [optional] 
**Selector** | [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | 
**ServiceName** | **string** | serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where \&quot;pod-specific-string\&quot; is managed by the StatefulSet controller. | 
**Template** | [**IoK8sApiCoreV1PodTemplateSpec**](IoK8sApiCoreV1PodTemplateSpec.md) |  | 
**UpdateStrategy** | Pointer to [**IoK8sApiAppsV1StatefulSetUpdateStrategy**](IoK8sApiAppsV1StatefulSetUpdateStrategy.md) |  | [optional] 
**VolumeClaimTemplates** | Pointer to [**[]IoK8sApiCoreV1PersistentVolumeClaim**](IoK8sApiCoreV1PersistentVolumeClaim.md) | volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name. | [optional] 

## Methods

### NewIoK8sApiAppsV1StatefulSetSpec

`func NewIoK8sApiAppsV1StatefulSetSpec(selector IoK8sApimachineryPkgApisMetaV1LabelSelector, serviceName string, template IoK8sApiCoreV1PodTemplateSpec, ) *IoK8sApiAppsV1StatefulSetSpec`

NewIoK8sApiAppsV1StatefulSetSpec instantiates a new IoK8sApiAppsV1StatefulSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1StatefulSetSpecWithDefaults

`func NewIoK8sApiAppsV1StatefulSetSpecWithDefaults() *IoK8sApiAppsV1StatefulSetSpec`

NewIoK8sApiAppsV1StatefulSetSpecWithDefaults instantiates a new IoK8sApiAppsV1StatefulSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinReadySeconds

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetMinReadySeconds() int32`

GetMinReadySeconds returns the MinReadySeconds field if non-nil, zero value otherwise.

### GetMinReadySecondsOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetMinReadySecondsOk() (*int32, bool)`

GetMinReadySecondsOk returns a tuple with the MinReadySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadySeconds

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetMinReadySeconds(v int32)`

SetMinReadySeconds sets MinReadySeconds field to given value.

### HasMinReadySeconds

`func (o *IoK8sApiAppsV1StatefulSetSpec) HasMinReadySeconds() bool`

HasMinReadySeconds returns a boolean if a field has been set.

### GetOrdinals

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetOrdinals() IoK8sApiAppsV1StatefulSetOrdinals`

GetOrdinals returns the Ordinals field if non-nil, zero value otherwise.

### GetOrdinalsOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetOrdinalsOk() (*IoK8sApiAppsV1StatefulSetOrdinals, bool)`

GetOrdinalsOk returns a tuple with the Ordinals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinals

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetOrdinals(v IoK8sApiAppsV1StatefulSetOrdinals)`

SetOrdinals sets Ordinals field to given value.

### HasOrdinals

`func (o *IoK8sApiAppsV1StatefulSetSpec) HasOrdinals() bool`

HasOrdinals returns a boolean if a field has been set.

### GetPersistentVolumeClaimRetentionPolicy

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetPersistentVolumeClaimRetentionPolicy() IoK8sApiAppsV1StatefulSetPersistentVolumeClaimRetentionPolicy`

GetPersistentVolumeClaimRetentionPolicy returns the PersistentVolumeClaimRetentionPolicy field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimRetentionPolicyOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetPersistentVolumeClaimRetentionPolicyOk() (*IoK8sApiAppsV1StatefulSetPersistentVolumeClaimRetentionPolicy, bool)`

GetPersistentVolumeClaimRetentionPolicyOk returns a tuple with the PersistentVolumeClaimRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaimRetentionPolicy

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetPersistentVolumeClaimRetentionPolicy(v IoK8sApiAppsV1StatefulSetPersistentVolumeClaimRetentionPolicy)`

SetPersistentVolumeClaimRetentionPolicy sets PersistentVolumeClaimRetentionPolicy field to given value.

### HasPersistentVolumeClaimRetentionPolicy

`func (o *IoK8sApiAppsV1StatefulSetSpec) HasPersistentVolumeClaimRetentionPolicy() bool`

HasPersistentVolumeClaimRetentionPolicy returns a boolean if a field has been set.

### GetPodManagementPolicy

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetPodManagementPolicy() string`

GetPodManagementPolicy returns the PodManagementPolicy field if non-nil, zero value otherwise.

### GetPodManagementPolicyOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetPodManagementPolicyOk() (*string, bool)`

GetPodManagementPolicyOk returns a tuple with the PodManagementPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodManagementPolicy

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetPodManagementPolicy(v string)`

SetPodManagementPolicy sets PodManagementPolicy field to given value.

### HasPodManagementPolicy

`func (o *IoK8sApiAppsV1StatefulSetSpec) HasPodManagementPolicy() bool`

HasPodManagementPolicy returns a boolean if a field has been set.

### GetReplicas

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *IoK8sApiAppsV1StatefulSetSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRevisionHistoryLimit

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetRevisionHistoryLimit() int32`

GetRevisionHistoryLimit returns the RevisionHistoryLimit field if non-nil, zero value otherwise.

### GetRevisionHistoryLimitOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetRevisionHistoryLimitOk() (*int32, bool)`

GetRevisionHistoryLimitOk returns a tuple with the RevisionHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistoryLimit

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetRevisionHistoryLimit(v int32)`

SetRevisionHistoryLimit sets RevisionHistoryLimit field to given value.

### HasRevisionHistoryLimit

`func (o *IoK8sApiAppsV1StatefulSetSpec) HasRevisionHistoryLimit() bool`

HasRevisionHistoryLimit returns a boolean if a field has been set.

### GetSelector

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetSelector sets Selector field to given value.


### GetServiceName

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetTemplate

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetTemplate() IoK8sApiCoreV1PodTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetTemplateOk() (*IoK8sApiCoreV1PodTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetTemplate(v IoK8sApiCoreV1PodTemplateSpec)`

SetTemplate sets Template field to given value.


### GetUpdateStrategy

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetUpdateStrategy() IoK8sApiAppsV1StatefulSetUpdateStrategy`

GetUpdateStrategy returns the UpdateStrategy field if non-nil, zero value otherwise.

### GetUpdateStrategyOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetUpdateStrategyOk() (*IoK8sApiAppsV1StatefulSetUpdateStrategy, bool)`

GetUpdateStrategyOk returns a tuple with the UpdateStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStrategy

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetUpdateStrategy(v IoK8sApiAppsV1StatefulSetUpdateStrategy)`

SetUpdateStrategy sets UpdateStrategy field to given value.

### HasUpdateStrategy

`func (o *IoK8sApiAppsV1StatefulSetSpec) HasUpdateStrategy() bool`

HasUpdateStrategy returns a boolean if a field has been set.

### GetVolumeClaimTemplates

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetVolumeClaimTemplates() []IoK8sApiCoreV1PersistentVolumeClaim`

GetVolumeClaimTemplates returns the VolumeClaimTemplates field if non-nil, zero value otherwise.

### GetVolumeClaimTemplatesOk

`func (o *IoK8sApiAppsV1StatefulSetSpec) GetVolumeClaimTemplatesOk() (*[]IoK8sApiCoreV1PersistentVolumeClaim, bool)`

GetVolumeClaimTemplatesOk returns a tuple with the VolumeClaimTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClaimTemplates

`func (o *IoK8sApiAppsV1StatefulSetSpec) SetVolumeClaimTemplates(v []IoK8sApiCoreV1PersistentVolumeClaim)`

SetVolumeClaimTemplates sets VolumeClaimTemplates field to given value.

### HasVolumeClaimTemplates

`func (o *IoK8sApiAppsV1StatefulSetSpec) HasVolumeClaimTemplates() bool`

HasVolumeClaimTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


