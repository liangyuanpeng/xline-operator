# IoK8sApiNetworkingV1IPBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | cidr is a string representing the IPBlock Valid examples are \&quot;192.168.1.0/24\&quot; or \&quot;2001:db8::/64\&quot; | 
**Except** | Pointer to **[]string** | except is a slice of CIDRs that should not be included within an IPBlock Valid examples are \&quot;192.168.1.0/24\&quot; or \&quot;2001:db8::/64\&quot; Except values will be rejected if they are outside the cidr range | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IPBlock

`func NewIoK8sApiNetworkingV1IPBlock(cidr string, ) *IoK8sApiNetworkingV1IPBlock`

NewIoK8sApiNetworkingV1IPBlock instantiates a new IoK8sApiNetworkingV1IPBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IPBlockWithDefaults

`func NewIoK8sApiNetworkingV1IPBlockWithDefaults() *IoK8sApiNetworkingV1IPBlock`

NewIoK8sApiNetworkingV1IPBlockWithDefaults instantiates a new IoK8sApiNetworkingV1IPBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *IoK8sApiNetworkingV1IPBlock) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IoK8sApiNetworkingV1IPBlock) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IoK8sApiNetworkingV1IPBlock) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetExcept

`func (o *IoK8sApiNetworkingV1IPBlock) GetExcept() []string`

GetExcept returns the Except field if non-nil, zero value otherwise.

### GetExceptOk

`func (o *IoK8sApiNetworkingV1IPBlock) GetExceptOk() (*[]string, bool)`

GetExceptOk returns a tuple with the Except field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcept

`func (o *IoK8sApiNetworkingV1IPBlock) SetExcept(v []string)`

SetExcept sets Except field to given value.

### HasExcept

`func (o *IoK8sApiNetworkingV1IPBlock) HasExcept() bool`

HasExcept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


