# IoK8sApiNetworkingV1IngressLoadBalancerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingress** | Pointer to [**[]IoK8sApiNetworkingV1IngressLoadBalancerIngress**](IoK8sApiNetworkingV1IngressLoadBalancerIngress.md) | ingress is a list containing ingress points for the load-balancer. | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IngressLoadBalancerStatus

`func NewIoK8sApiNetworkingV1IngressLoadBalancerStatus() *IoK8sApiNetworkingV1IngressLoadBalancerStatus`

NewIoK8sApiNetworkingV1IngressLoadBalancerStatus instantiates a new IoK8sApiNetworkingV1IngressLoadBalancerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IngressLoadBalancerStatusWithDefaults

`func NewIoK8sApiNetworkingV1IngressLoadBalancerStatusWithDefaults() *IoK8sApiNetworkingV1IngressLoadBalancerStatus`

NewIoK8sApiNetworkingV1IngressLoadBalancerStatusWithDefaults instantiates a new IoK8sApiNetworkingV1IngressLoadBalancerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngress

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerStatus) GetIngress() []IoK8sApiNetworkingV1IngressLoadBalancerIngress`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerStatus) GetIngressOk() (*[]IoK8sApiNetworkingV1IngressLoadBalancerIngress, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerStatus) SetIngress(v []IoK8sApiNetworkingV1IngressLoadBalancerIngress)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerStatus) HasIngress() bool`

HasIngress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


