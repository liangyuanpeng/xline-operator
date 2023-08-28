# IoK8sApiPolicyV1PodDisruptionBudgetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxUnavailable** | Pointer to **string** | IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number. | [optional] 
**MinAvailable** | Pointer to **string** | IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number. | [optional] 
**Selector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**UnhealthyPodEvictionPolicy** | Pointer to **string** | UnhealthyPodEvictionPolicy defines the criteria for when unhealthy pods should be considered for eviction. Current implementation considers healthy pods, as pods that have status.conditions item with type&#x3D;\&quot;Ready\&quot;,status&#x3D;\&quot;True\&quot;.  Valid policies are IfHealthyBudget and AlwaysAllow. If no policy is specified, the default behavior will be used, which corresponds to the IfHealthyBudget policy.  IfHealthyBudget policy means that running pods (status.phase&#x3D;\&quot;Running\&quot;), but not yet healthy can be evicted only if the guarded application is not disrupted (status.currentHealthy is at least equal to status.desiredHealthy). Healthy pods will be subject to the PDB for eviction.  AlwaysAllow policy means that all running pods (status.phase&#x3D;\&quot;Running\&quot;), but not yet healthy are considered disrupted and can be evicted regardless of whether the criteria in a PDB is met. This means perspective running pods of a disrupted application might not get a chance to become healthy. Healthy pods will be subject to the PDB for eviction.  Additional policies may be added in the future. Clients making eviction decisions should disallow eviction of unhealthy pods if they encounter an unrecognized policy in this field.  This field is beta-level. The eviction API uses this field when the feature gate PDBUnhealthyPodEvictionPolicy is enabled (enabled by default).  Possible enum values:  - &#x60;\&quot;AlwaysAllow\&quot;&#x60; policy means that all running pods (status.phase&#x3D;\&quot;Running\&quot;), but not yet healthy are considered disrupted and can be evicted regardless of whether the criteria in a PDB is met. This means perspective running pods of a disrupted application might not get a chance to become healthy. Healthy pods will be subject to the PDB for eviction.  - &#x60;\&quot;IfHealthyBudget\&quot;&#x60; policy means that running pods (status.phase&#x3D;\&quot;Running\&quot;), but not yet healthy can be evicted only if the guarded application is not disrupted (status.currentHealthy is at least equal to status.desiredHealthy). Healthy pods will be subject to the PDB for eviction. | [optional] 

## Methods

### NewIoK8sApiPolicyV1PodDisruptionBudgetSpec

`func NewIoK8sApiPolicyV1PodDisruptionBudgetSpec() *IoK8sApiPolicyV1PodDisruptionBudgetSpec`

NewIoK8sApiPolicyV1PodDisruptionBudgetSpec instantiates a new IoK8sApiPolicyV1PodDisruptionBudgetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiPolicyV1PodDisruptionBudgetSpecWithDefaults

`func NewIoK8sApiPolicyV1PodDisruptionBudgetSpecWithDefaults() *IoK8sApiPolicyV1PodDisruptionBudgetSpec`

NewIoK8sApiPolicyV1PodDisruptionBudgetSpecWithDefaults instantiates a new IoK8sApiPolicyV1PodDisruptionBudgetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxUnavailable

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.

### GetMinAvailable

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) GetMinAvailable() string`

GetMinAvailable returns the MinAvailable field if non-nil, zero value otherwise.

### GetMinAvailableOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) GetMinAvailableOk() (*string, bool)`

GetMinAvailableOk returns a tuple with the MinAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAvailable

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) SetMinAvailable(v string)`

SetMinAvailable sets MinAvailable field to given value.

### HasMinAvailable

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) HasMinAvailable() bool`

HasMinAvailable returns a boolean if a field has been set.

### GetSelector

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) GetSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) GetSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) SetSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetUnhealthyPodEvictionPolicy

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) GetUnhealthyPodEvictionPolicy() string`

GetUnhealthyPodEvictionPolicy returns the UnhealthyPodEvictionPolicy field if non-nil, zero value otherwise.

### GetUnhealthyPodEvictionPolicyOk

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) GetUnhealthyPodEvictionPolicyOk() (*string, bool)`

GetUnhealthyPodEvictionPolicyOk returns a tuple with the UnhealthyPodEvictionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyPodEvictionPolicy

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) SetUnhealthyPodEvictionPolicy(v string)`

SetUnhealthyPodEvictionPolicy sets UnhealthyPodEvictionPolicy field to given value.

### HasUnhealthyPodEvictionPolicy

`func (o *IoK8sApiPolicyV1PodDisruptionBudgetSpec) HasUnhealthyPodEvictionPolicy() bool`

HasUnhealthyPodEvictionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


