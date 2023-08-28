# IoK8sApiCoreV1ResourceQuotaSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hard** | Pointer to **map[string]string** | hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/ | [optional] 
**ScopeSelector** | Pointer to [**IoK8sApiCoreV1ScopeSelector**](IoK8sApiCoreV1ScopeSelector.md) |  | [optional] 
**Scopes** | Pointer to **[]string** | A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects. | [optional] 

## Methods

### NewIoK8sApiCoreV1ResourceQuotaSpec

`func NewIoK8sApiCoreV1ResourceQuotaSpec() *IoK8sApiCoreV1ResourceQuotaSpec`

NewIoK8sApiCoreV1ResourceQuotaSpec instantiates a new IoK8sApiCoreV1ResourceQuotaSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ResourceQuotaSpecWithDefaults

`func NewIoK8sApiCoreV1ResourceQuotaSpecWithDefaults() *IoK8sApiCoreV1ResourceQuotaSpec`

NewIoK8sApiCoreV1ResourceQuotaSpecWithDefaults instantiates a new IoK8sApiCoreV1ResourceQuotaSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHard

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) GetHard() map[string]string`

GetHard returns the Hard field if non-nil, zero value otherwise.

### GetHardOk

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) GetHardOk() (*map[string]string, bool)`

GetHardOk returns a tuple with the Hard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHard

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) SetHard(v map[string]string)`

SetHard sets Hard field to given value.

### HasHard

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) HasHard() bool`

HasHard returns a boolean if a field has been set.

### GetScopeSelector

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) GetScopeSelector() IoK8sApiCoreV1ScopeSelector`

GetScopeSelector returns the ScopeSelector field if non-nil, zero value otherwise.

### GetScopeSelectorOk

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) GetScopeSelectorOk() (*IoK8sApiCoreV1ScopeSelector, bool)`

GetScopeSelectorOk returns a tuple with the ScopeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSelector

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) SetScopeSelector(v IoK8sApiCoreV1ScopeSelector)`

SetScopeSelector sets ScopeSelector field to given value.

### HasScopeSelector

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) HasScopeSelector() bool`

HasScopeSelector returns a boolean if a field has been set.

### GetScopes

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IoK8sApiCoreV1ResourceQuotaSpec) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


