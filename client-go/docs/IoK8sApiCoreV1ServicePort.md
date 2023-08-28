# IoK8sApiCoreV1ServicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppProtocol** | Pointer to **string** | The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. | [optional] 
**Name** | Pointer to **string** | The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the &#39;name&#39; field in the EndpointPort. Optional if only one ServicePort is defined on this service. | [optional] 
**NodePort** | Pointer to **int32** | The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport | [optional] 
**Port** | **int32** | The port that will be exposed by this service. | 
**Protocol** | Pointer to **string** | The IP protocol for this port. Supports \&quot;TCP\&quot;, \&quot;UDP\&quot;, and \&quot;SCTP\&quot;. Default is TCP.  Possible enum values:  - &#x60;\&quot;SCTP\&quot;&#x60; is the SCTP protocol.  - &#x60;\&quot;TCP\&quot;&#x60; is the TCP protocol.  - &#x60;\&quot;UDP\&quot;&#x60; is the UDP protocol. | [optional] 
**TargetPort** | Pointer to **string** | IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number. | [optional] 

## Methods

### NewIoK8sApiCoreV1ServicePort

`func NewIoK8sApiCoreV1ServicePort(port int32, ) *IoK8sApiCoreV1ServicePort`

NewIoK8sApiCoreV1ServicePort instantiates a new IoK8sApiCoreV1ServicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ServicePortWithDefaults

`func NewIoK8sApiCoreV1ServicePortWithDefaults() *IoK8sApiCoreV1ServicePort`

NewIoK8sApiCoreV1ServicePortWithDefaults instantiates a new IoK8sApiCoreV1ServicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppProtocol

`func (o *IoK8sApiCoreV1ServicePort) GetAppProtocol() string`

GetAppProtocol returns the AppProtocol field if non-nil, zero value otherwise.

### GetAppProtocolOk

`func (o *IoK8sApiCoreV1ServicePort) GetAppProtocolOk() (*string, bool)`

GetAppProtocolOk returns a tuple with the AppProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtocol

`func (o *IoK8sApiCoreV1ServicePort) SetAppProtocol(v string)`

SetAppProtocol sets AppProtocol field to given value.

### HasAppProtocol

`func (o *IoK8sApiCoreV1ServicePort) HasAppProtocol() bool`

HasAppProtocol returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiCoreV1ServicePort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1ServicePort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1ServicePort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiCoreV1ServicePort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodePort

`func (o *IoK8sApiCoreV1ServicePort) GetNodePort() int32`

GetNodePort returns the NodePort field if non-nil, zero value otherwise.

### GetNodePortOk

`func (o *IoK8sApiCoreV1ServicePort) GetNodePortOk() (*int32, bool)`

GetNodePortOk returns a tuple with the NodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePort

`func (o *IoK8sApiCoreV1ServicePort) SetNodePort(v int32)`

SetNodePort sets NodePort field to given value.

### HasNodePort

`func (o *IoK8sApiCoreV1ServicePort) HasNodePort() bool`

HasNodePort returns a boolean if a field has been set.

### GetPort

`func (o *IoK8sApiCoreV1ServicePort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sApiCoreV1ServicePort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sApiCoreV1ServicePort) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *IoK8sApiCoreV1ServicePort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IoK8sApiCoreV1ServicePort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IoK8sApiCoreV1ServicePort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IoK8sApiCoreV1ServicePort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTargetPort

`func (o *IoK8sApiCoreV1ServicePort) GetTargetPort() string`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *IoK8sApiCoreV1ServicePort) GetTargetPortOk() (*string, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *IoK8sApiCoreV1ServicePort) SetTargetPort(v string)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *IoK8sApiCoreV1ServicePort) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


