# IoK8sApiCoreV1NodeAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The node address. | 
**Type** | **string** | Node address type, one of Hostname, ExternalIP or InternalIP. | 

## Methods

### NewIoK8sApiCoreV1NodeAddress

`func NewIoK8sApiCoreV1NodeAddress(address string, type_ string, ) *IoK8sApiCoreV1NodeAddress`

NewIoK8sApiCoreV1NodeAddress instantiates a new IoK8sApiCoreV1NodeAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1NodeAddressWithDefaults

`func NewIoK8sApiCoreV1NodeAddressWithDefaults() *IoK8sApiCoreV1NodeAddress`

NewIoK8sApiCoreV1NodeAddressWithDefaults instantiates a new IoK8sApiCoreV1NodeAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IoK8sApiCoreV1NodeAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IoK8sApiCoreV1NodeAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IoK8sApiCoreV1NodeAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetType

`func (o *IoK8sApiCoreV1NodeAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiCoreV1NodeAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiCoreV1NodeAddress) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


