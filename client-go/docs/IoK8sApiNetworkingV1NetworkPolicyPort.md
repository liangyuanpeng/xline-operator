# IoK8sApiNetworkingV1NetworkPolicyPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndPort** | Pointer to **int32** | endPort indicates that the range of ports from port to endPort if set, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port. | [optional] 
**Port** | Pointer to **string** | IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number. | [optional] 
**Protocol** | Pointer to **string** | protocol represents the protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.  Possible enum values:  - &#x60;\&quot;SCTP\&quot;&#x60; is the SCTP protocol.  - &#x60;\&quot;TCP\&quot;&#x60; is the TCP protocol.  - &#x60;\&quot;UDP\&quot;&#x60; is the UDP protocol. | [optional] 

## Methods

### NewIoK8sApiNetworkingV1NetworkPolicyPort

`func NewIoK8sApiNetworkingV1NetworkPolicyPort() *IoK8sApiNetworkingV1NetworkPolicyPort`

NewIoK8sApiNetworkingV1NetworkPolicyPort instantiates a new IoK8sApiNetworkingV1NetworkPolicyPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1NetworkPolicyPortWithDefaults

`func NewIoK8sApiNetworkingV1NetworkPolicyPortWithDefaults() *IoK8sApiNetworkingV1NetworkPolicyPort`

NewIoK8sApiNetworkingV1NetworkPolicyPortWithDefaults instantiates a new IoK8sApiNetworkingV1NetworkPolicyPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndPort

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.

### HasEndPort

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) HasEndPort() bool`

HasEndPort returns a boolean if a field has been set.

### GetPort

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IoK8sApiNetworkingV1NetworkPolicyPort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


