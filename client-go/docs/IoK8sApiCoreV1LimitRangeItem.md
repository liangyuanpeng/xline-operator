# IoK8sApiCoreV1LimitRangeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **map[string]string** | Default resource requirement limit value by resource name if resource limit is omitted. | [optional] 
**DefaultRequest** | Pointer to **map[string]string** | DefaultRequest is the default resource requirement request value by resource name if resource request is omitted. | [optional] 
**Max** | Pointer to **map[string]string** | Max usage constraints on this kind by resource name. | [optional] 
**MaxLimitRequestRatio** | Pointer to **map[string]string** | MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource. | [optional] 
**Min** | Pointer to **map[string]string** | Min usage constraints on this kind by resource name. | [optional] 
**Type** | **string** | Type of resource that this limit applies to. | 

## Methods

### NewIoK8sApiCoreV1LimitRangeItem

`func NewIoK8sApiCoreV1LimitRangeItem(type_ string, ) *IoK8sApiCoreV1LimitRangeItem`

NewIoK8sApiCoreV1LimitRangeItem instantiates a new IoK8sApiCoreV1LimitRangeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1LimitRangeItemWithDefaults

`func NewIoK8sApiCoreV1LimitRangeItemWithDefaults() *IoK8sApiCoreV1LimitRangeItem`

NewIoK8sApiCoreV1LimitRangeItemWithDefaults instantiates a new IoK8sApiCoreV1LimitRangeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *IoK8sApiCoreV1LimitRangeItem) GetDefault() map[string]string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *IoK8sApiCoreV1LimitRangeItem) GetDefaultOk() (*map[string]string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *IoK8sApiCoreV1LimitRangeItem) SetDefault(v map[string]string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *IoK8sApiCoreV1LimitRangeItem) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDefaultRequest

`func (o *IoK8sApiCoreV1LimitRangeItem) GetDefaultRequest() map[string]string`

GetDefaultRequest returns the DefaultRequest field if non-nil, zero value otherwise.

### GetDefaultRequestOk

`func (o *IoK8sApiCoreV1LimitRangeItem) GetDefaultRequestOk() (*map[string]string, bool)`

GetDefaultRequestOk returns a tuple with the DefaultRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRequest

`func (o *IoK8sApiCoreV1LimitRangeItem) SetDefaultRequest(v map[string]string)`

SetDefaultRequest sets DefaultRequest field to given value.

### HasDefaultRequest

`func (o *IoK8sApiCoreV1LimitRangeItem) HasDefaultRequest() bool`

HasDefaultRequest returns a boolean if a field has been set.

### GetMax

`func (o *IoK8sApiCoreV1LimitRangeItem) GetMax() map[string]string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *IoK8sApiCoreV1LimitRangeItem) GetMaxOk() (*map[string]string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *IoK8sApiCoreV1LimitRangeItem) SetMax(v map[string]string)`

SetMax sets Max field to given value.

### HasMax

`func (o *IoK8sApiCoreV1LimitRangeItem) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMaxLimitRequestRatio

`func (o *IoK8sApiCoreV1LimitRangeItem) GetMaxLimitRequestRatio() map[string]string`

GetMaxLimitRequestRatio returns the MaxLimitRequestRatio field if non-nil, zero value otherwise.

### GetMaxLimitRequestRatioOk

`func (o *IoK8sApiCoreV1LimitRangeItem) GetMaxLimitRequestRatioOk() (*map[string]string, bool)`

GetMaxLimitRequestRatioOk returns a tuple with the MaxLimitRequestRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimitRequestRatio

`func (o *IoK8sApiCoreV1LimitRangeItem) SetMaxLimitRequestRatio(v map[string]string)`

SetMaxLimitRequestRatio sets MaxLimitRequestRatio field to given value.

### HasMaxLimitRequestRatio

`func (o *IoK8sApiCoreV1LimitRangeItem) HasMaxLimitRequestRatio() bool`

HasMaxLimitRequestRatio returns a boolean if a field has been set.

### GetMin

`func (o *IoK8sApiCoreV1LimitRangeItem) GetMin() map[string]string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *IoK8sApiCoreV1LimitRangeItem) GetMinOk() (*map[string]string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *IoK8sApiCoreV1LimitRangeItem) SetMin(v map[string]string)`

SetMin sets Min field to given value.

### HasMin

`func (o *IoK8sApiCoreV1LimitRangeItem) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiCoreV1LimitRangeItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiCoreV1LimitRangeItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiCoreV1LimitRangeItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


