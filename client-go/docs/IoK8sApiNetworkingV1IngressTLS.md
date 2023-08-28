# IoK8sApiNetworkingV1IngressTLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** | hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified. | [optional] 
**SecretName** | Pointer to **string** | secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \&quot;Host\&quot; header field used by an IngressRule, the SNI host is used for termination and value of the \&quot;Host\&quot; header is used for routing. | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IngressTLS

`func NewIoK8sApiNetworkingV1IngressTLS() *IoK8sApiNetworkingV1IngressTLS`

NewIoK8sApiNetworkingV1IngressTLS instantiates a new IoK8sApiNetworkingV1IngressTLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IngressTLSWithDefaults

`func NewIoK8sApiNetworkingV1IngressTLSWithDefaults() *IoK8sApiNetworkingV1IngressTLS`

NewIoK8sApiNetworkingV1IngressTLSWithDefaults instantiates a new IoK8sApiNetworkingV1IngressTLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *IoK8sApiNetworkingV1IngressTLS) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *IoK8sApiNetworkingV1IngressTLS) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *IoK8sApiNetworkingV1IngressTLS) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *IoK8sApiNetworkingV1IngressTLS) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetSecretName

`func (o *IoK8sApiNetworkingV1IngressTLS) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *IoK8sApiNetworkingV1IngressTLS) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *IoK8sApiNetworkingV1IngressTLS) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.

### HasSecretName

`func (o *IoK8sApiNetworkingV1IngressTLS) HasSecretName() bool`

HasSecretName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


