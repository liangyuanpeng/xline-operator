# IoK8sApiCoreV1LimitRangeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | [**[]IoK8sApiCoreV1LimitRangeItem**](IoK8sApiCoreV1LimitRangeItem.md) | Limits is the list of LimitRangeItem objects that are enforced. | 

## Methods

### NewIoK8sApiCoreV1LimitRangeSpec

`func NewIoK8sApiCoreV1LimitRangeSpec(limits []IoK8sApiCoreV1LimitRangeItem, ) *IoK8sApiCoreV1LimitRangeSpec`

NewIoK8sApiCoreV1LimitRangeSpec instantiates a new IoK8sApiCoreV1LimitRangeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1LimitRangeSpecWithDefaults

`func NewIoK8sApiCoreV1LimitRangeSpecWithDefaults() *IoK8sApiCoreV1LimitRangeSpec`

NewIoK8sApiCoreV1LimitRangeSpecWithDefaults instantiates a new IoK8sApiCoreV1LimitRangeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *IoK8sApiCoreV1LimitRangeSpec) GetLimits() []IoK8sApiCoreV1LimitRangeItem`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *IoK8sApiCoreV1LimitRangeSpec) GetLimitsOk() (*[]IoK8sApiCoreV1LimitRangeItem, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *IoK8sApiCoreV1LimitRangeSpec) SetLimits(v []IoK8sApiCoreV1LimitRangeItem)`

SetLimits sets Limits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


