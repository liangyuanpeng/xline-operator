# IoK8sApiCoreV1TCPSocketAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Optional: Host name to connect to, defaults to the pod IP. | [optional] 
**Port** | **string** | IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number. | 

## Methods

### NewIoK8sApiCoreV1TCPSocketAction

`func NewIoK8sApiCoreV1TCPSocketAction(port string, ) *IoK8sApiCoreV1TCPSocketAction`

NewIoK8sApiCoreV1TCPSocketAction instantiates a new IoK8sApiCoreV1TCPSocketAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1TCPSocketActionWithDefaults

`func NewIoK8sApiCoreV1TCPSocketActionWithDefaults() *IoK8sApiCoreV1TCPSocketAction`

NewIoK8sApiCoreV1TCPSocketActionWithDefaults instantiates a new IoK8sApiCoreV1TCPSocketAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *IoK8sApiCoreV1TCPSocketAction) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IoK8sApiCoreV1TCPSocketAction) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IoK8sApiCoreV1TCPSocketAction) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IoK8sApiCoreV1TCPSocketAction) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *IoK8sApiCoreV1TCPSocketAction) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sApiCoreV1TCPSocketAction) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sApiCoreV1TCPSocketAction) SetPort(v string)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


