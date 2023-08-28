# IoK8sApiCoreV1EndpointAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | The Hostname of this endpoint | [optional] 
**Ip** | **string** | The IP of this endpoint. May not be loopback (127.0.0.0/8 or ::1), link-local (169.254.0.0/16 or fe80::/10), or link-local multicast (224.0.0.0/24 or ff02::/16). | 
**NodeName** | Pointer to **string** | Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node. | [optional] 
**TargetRef** | Pointer to [**IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1EndpointAddress

`func NewIoK8sApiCoreV1EndpointAddress(ip string, ) *IoK8sApiCoreV1EndpointAddress`

NewIoK8sApiCoreV1EndpointAddress instantiates a new IoK8sApiCoreV1EndpointAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1EndpointAddressWithDefaults

`func NewIoK8sApiCoreV1EndpointAddressWithDefaults() *IoK8sApiCoreV1EndpointAddress`

NewIoK8sApiCoreV1EndpointAddressWithDefaults instantiates a new IoK8sApiCoreV1EndpointAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *IoK8sApiCoreV1EndpointAddress) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *IoK8sApiCoreV1EndpointAddress) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *IoK8sApiCoreV1EndpointAddress) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *IoK8sApiCoreV1EndpointAddress) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *IoK8sApiCoreV1EndpointAddress) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IoK8sApiCoreV1EndpointAddress) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IoK8sApiCoreV1EndpointAddress) SetIp(v string)`

SetIp sets Ip field to given value.


### GetNodeName

`func (o *IoK8sApiCoreV1EndpointAddress) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *IoK8sApiCoreV1EndpointAddress) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *IoK8sApiCoreV1EndpointAddress) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *IoK8sApiCoreV1EndpointAddress) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetTargetRef

`func (o *IoK8sApiCoreV1EndpointAddress) GetTargetRef() IoK8sApiCoreV1ObjectReference`

GetTargetRef returns the TargetRef field if non-nil, zero value otherwise.

### GetTargetRefOk

`func (o *IoK8sApiCoreV1EndpointAddress) GetTargetRefOk() (*IoK8sApiCoreV1ObjectReference, bool)`

GetTargetRefOk returns a tuple with the TargetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRef

`func (o *IoK8sApiCoreV1EndpointAddress) SetTargetRef(v IoK8sApiCoreV1ObjectReference)`

SetTargetRef sets TargetRef field to given value.

### HasTargetRef

`func (o *IoK8sApiCoreV1EndpointAddress) HasTargetRef() bool`

HasTargetRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


