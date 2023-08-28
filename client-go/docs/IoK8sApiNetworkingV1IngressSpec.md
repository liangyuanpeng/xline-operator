# IoK8sApiNetworkingV1IngressSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultBackend** | Pointer to [**IoK8sApiNetworkingV1IngressBackend**](IoK8sApiNetworkingV1IngressBackend.md) |  | [optional] 
**IngressClassName** | Pointer to **string** | ingressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -&gt; IngressClass -&gt; Ingress resource). Although the &#x60;kubernetes.io/ingress.class&#x60; annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present. | [optional] 
**Rules** | Pointer to [**[]IoK8sApiNetworkingV1IngressRule**](IoK8sApiNetworkingV1IngressRule.md) | rules is a list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend. | [optional] 
**Tls** | Pointer to [**[]IoK8sApiNetworkingV1IngressTLS**](IoK8sApiNetworkingV1IngressTLS.md) | tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI. | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IngressSpec

`func NewIoK8sApiNetworkingV1IngressSpec() *IoK8sApiNetworkingV1IngressSpec`

NewIoK8sApiNetworkingV1IngressSpec instantiates a new IoK8sApiNetworkingV1IngressSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IngressSpecWithDefaults

`func NewIoK8sApiNetworkingV1IngressSpecWithDefaults() *IoK8sApiNetworkingV1IngressSpec`

NewIoK8sApiNetworkingV1IngressSpecWithDefaults instantiates a new IoK8sApiNetworkingV1IngressSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultBackend

`func (o *IoK8sApiNetworkingV1IngressSpec) GetDefaultBackend() IoK8sApiNetworkingV1IngressBackend`

GetDefaultBackend returns the DefaultBackend field if non-nil, zero value otherwise.

### GetDefaultBackendOk

`func (o *IoK8sApiNetworkingV1IngressSpec) GetDefaultBackendOk() (*IoK8sApiNetworkingV1IngressBackend, bool)`

GetDefaultBackendOk returns a tuple with the DefaultBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackend

`func (o *IoK8sApiNetworkingV1IngressSpec) SetDefaultBackend(v IoK8sApiNetworkingV1IngressBackend)`

SetDefaultBackend sets DefaultBackend field to given value.

### HasDefaultBackend

`func (o *IoK8sApiNetworkingV1IngressSpec) HasDefaultBackend() bool`

HasDefaultBackend returns a boolean if a field has been set.

### GetIngressClassName

`func (o *IoK8sApiNetworkingV1IngressSpec) GetIngressClassName() string`

GetIngressClassName returns the IngressClassName field if non-nil, zero value otherwise.

### GetIngressClassNameOk

`func (o *IoK8sApiNetworkingV1IngressSpec) GetIngressClassNameOk() (*string, bool)`

GetIngressClassNameOk returns a tuple with the IngressClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressClassName

`func (o *IoK8sApiNetworkingV1IngressSpec) SetIngressClassName(v string)`

SetIngressClassName sets IngressClassName field to given value.

### HasIngressClassName

`func (o *IoK8sApiNetworkingV1IngressSpec) HasIngressClassName() bool`

HasIngressClassName returns a boolean if a field has been set.

### GetRules

`func (o *IoK8sApiNetworkingV1IngressSpec) GetRules() []IoK8sApiNetworkingV1IngressRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *IoK8sApiNetworkingV1IngressSpec) GetRulesOk() (*[]IoK8sApiNetworkingV1IngressRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *IoK8sApiNetworkingV1IngressSpec) SetRules(v []IoK8sApiNetworkingV1IngressRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *IoK8sApiNetworkingV1IngressSpec) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTls

`func (o *IoK8sApiNetworkingV1IngressSpec) GetTls() []IoK8sApiNetworkingV1IngressTLS`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoK8sApiNetworkingV1IngressSpec) GetTlsOk() (*[]IoK8sApiNetworkingV1IngressTLS, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoK8sApiNetworkingV1IngressSpec) SetTls(v []IoK8sApiNetworkingV1IngressTLS)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoK8sApiNetworkingV1IngressSpec) HasTls() bool`

HasTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


