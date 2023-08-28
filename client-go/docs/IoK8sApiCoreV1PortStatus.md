# IoK8sApiCoreV1PortStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use   CamelCase names - cloud provider specific error values must have names that comply with the   format foo.example.com/CamelCase. | [optional] 
**Port** | **int32** | Port is the port number of the service port of which status is recorded here | 
**Protocol** | **string** | Protocol is the protocol of the service port of which status is recorded here The supported values are: \&quot;TCP\&quot;, \&quot;UDP\&quot;, \&quot;SCTP\&quot;  Possible enum values:  - &#x60;\&quot;SCTP\&quot;&#x60; is the SCTP protocol.  - &#x60;\&quot;TCP\&quot;&#x60; is the TCP protocol.  - &#x60;\&quot;UDP\&quot;&#x60; is the UDP protocol. | 

## Methods

### NewIoK8sApiCoreV1PortStatus

`func NewIoK8sApiCoreV1PortStatus(port int32, protocol string, ) *IoK8sApiCoreV1PortStatus`

NewIoK8sApiCoreV1PortStatus instantiates a new IoK8sApiCoreV1PortStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PortStatusWithDefaults

`func NewIoK8sApiCoreV1PortStatusWithDefaults() *IoK8sApiCoreV1PortStatus`

NewIoK8sApiCoreV1PortStatusWithDefaults instantiates a new IoK8sApiCoreV1PortStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *IoK8sApiCoreV1PortStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IoK8sApiCoreV1PortStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IoK8sApiCoreV1PortStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *IoK8sApiCoreV1PortStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPort

`func (o *IoK8sApiCoreV1PortStatus) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sApiCoreV1PortStatus) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sApiCoreV1PortStatus) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *IoK8sApiCoreV1PortStatus) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IoK8sApiCoreV1PortStatus) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IoK8sApiCoreV1PortStatus) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


