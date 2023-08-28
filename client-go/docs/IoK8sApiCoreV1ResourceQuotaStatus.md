# IoK8sApiCoreV1ResourceQuotaStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hard** | Pointer to **map[string]string** | Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/ | [optional] 
**Used** | Pointer to **map[string]string** | Used is the current observed total usage of the resource in the namespace. | [optional] 

## Methods

### NewIoK8sApiCoreV1ResourceQuotaStatus

`func NewIoK8sApiCoreV1ResourceQuotaStatus() *IoK8sApiCoreV1ResourceQuotaStatus`

NewIoK8sApiCoreV1ResourceQuotaStatus instantiates a new IoK8sApiCoreV1ResourceQuotaStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ResourceQuotaStatusWithDefaults

`func NewIoK8sApiCoreV1ResourceQuotaStatusWithDefaults() *IoK8sApiCoreV1ResourceQuotaStatus`

NewIoK8sApiCoreV1ResourceQuotaStatusWithDefaults instantiates a new IoK8sApiCoreV1ResourceQuotaStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHard

`func (o *IoK8sApiCoreV1ResourceQuotaStatus) GetHard() map[string]string`

GetHard returns the Hard field if non-nil, zero value otherwise.

### GetHardOk

`func (o *IoK8sApiCoreV1ResourceQuotaStatus) GetHardOk() (*map[string]string, bool)`

GetHardOk returns a tuple with the Hard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHard

`func (o *IoK8sApiCoreV1ResourceQuotaStatus) SetHard(v map[string]string)`

SetHard sets Hard field to given value.

### HasHard

`func (o *IoK8sApiCoreV1ResourceQuotaStatus) HasHard() bool`

HasHard returns a boolean if a field has been set.

### GetUsed

`func (o *IoK8sApiCoreV1ResourceQuotaStatus) GetUsed() map[string]string`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *IoK8sApiCoreV1ResourceQuotaStatus) GetUsedOk() (*map[string]string, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *IoK8sApiCoreV1ResourceQuotaStatus) SetUsed(v map[string]string)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *IoK8sApiCoreV1ResourceQuotaStatus) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


