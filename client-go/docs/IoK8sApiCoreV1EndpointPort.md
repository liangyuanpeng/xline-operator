# IoK8sApiCoreV1EndpointPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppProtocol** | Pointer to **string** | The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:  * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).  * Kubernetes-defined prefixed names:   * &#39;kubernetes.io/h2c&#39; - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540  * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol. | [optional] 
**Name** | Pointer to **string** | The name of this port.  This must match the &#39;name&#39; field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined. | [optional] 
**Port** | **int32** | The port number of the endpoint. | 
**Protocol** | Pointer to **string** | The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.  Possible enum values:  - &#x60;\&quot;SCTP\&quot;&#x60; is the SCTP protocol.  - &#x60;\&quot;TCP\&quot;&#x60; is the TCP protocol.  - &#x60;\&quot;UDP\&quot;&#x60; is the UDP protocol. | [optional] 

## Methods

### NewIoK8sApiCoreV1EndpointPort

`func NewIoK8sApiCoreV1EndpointPort(port int32, ) *IoK8sApiCoreV1EndpointPort`

NewIoK8sApiCoreV1EndpointPort instantiates a new IoK8sApiCoreV1EndpointPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1EndpointPortWithDefaults

`func NewIoK8sApiCoreV1EndpointPortWithDefaults() *IoK8sApiCoreV1EndpointPort`

NewIoK8sApiCoreV1EndpointPortWithDefaults instantiates a new IoK8sApiCoreV1EndpointPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppProtocol

`func (o *IoK8sApiCoreV1EndpointPort) GetAppProtocol() string`

GetAppProtocol returns the AppProtocol field if non-nil, zero value otherwise.

### GetAppProtocolOk

`func (o *IoK8sApiCoreV1EndpointPort) GetAppProtocolOk() (*string, bool)`

GetAppProtocolOk returns a tuple with the AppProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtocol

`func (o *IoK8sApiCoreV1EndpointPort) SetAppProtocol(v string)`

SetAppProtocol sets AppProtocol field to given value.

### HasAppProtocol

`func (o *IoK8sApiCoreV1EndpointPort) HasAppProtocol() bool`

HasAppProtocol returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiCoreV1EndpointPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1EndpointPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1EndpointPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiCoreV1EndpointPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *IoK8sApiCoreV1EndpointPort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sApiCoreV1EndpointPort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sApiCoreV1EndpointPort) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *IoK8sApiCoreV1EndpointPort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IoK8sApiCoreV1EndpointPort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IoK8sApiCoreV1EndpointPort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IoK8sApiCoreV1EndpointPort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


