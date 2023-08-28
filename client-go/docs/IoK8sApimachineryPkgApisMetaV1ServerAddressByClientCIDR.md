# IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCIDR** | **string** | The CIDR with which clients can match their IP to figure out the server address that they should use. | 
**ServerAddress** | **string** | Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port. | 

## Methods

### NewIoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR

`func NewIoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR(clientCIDR string, serverAddress string, ) *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR`

NewIoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR instantiates a new IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDRWithDefaults

`func NewIoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDRWithDefaults() *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR`

NewIoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDRWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCIDR

`func (o *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) GetClientCIDR() string`

GetClientCIDR returns the ClientCIDR field if non-nil, zero value otherwise.

### GetClientCIDROk

`func (o *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) GetClientCIDROk() (*string, bool)`

GetClientCIDROk returns a tuple with the ClientCIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCIDR

`func (o *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) SetClientCIDR(v string)`

SetClientCIDR sets ClientCIDR field to given value.


### GetServerAddress

`func (o *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


