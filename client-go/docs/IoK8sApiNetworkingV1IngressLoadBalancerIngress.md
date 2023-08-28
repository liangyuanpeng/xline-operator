# IoK8sApiNetworkingV1IngressLoadBalancerIngress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | hostname is set for load-balancer ingress points that are DNS based. | [optional] 
**Ip** | Pointer to **string** | ip is set for load-balancer ingress points that are IP based. | [optional] 
**Ports** | Pointer to [**[]IoK8sApiNetworkingV1IngressPortStatus**](IoK8sApiNetworkingV1IngressPortStatus.md) | ports provides information about the ports exposed by this LoadBalancer. | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IngressLoadBalancerIngress

`func NewIoK8sApiNetworkingV1IngressLoadBalancerIngress() *IoK8sApiNetworkingV1IngressLoadBalancerIngress`

NewIoK8sApiNetworkingV1IngressLoadBalancerIngress instantiates a new IoK8sApiNetworkingV1IngressLoadBalancerIngress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IngressLoadBalancerIngressWithDefaults

`func NewIoK8sApiNetworkingV1IngressLoadBalancerIngressWithDefaults() *IoK8sApiNetworkingV1IngressLoadBalancerIngress`

NewIoK8sApiNetworkingV1IngressLoadBalancerIngressWithDefaults instantiates a new IoK8sApiNetworkingV1IngressLoadBalancerIngress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPorts

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) GetPorts() []IoK8sApiNetworkingV1IngressPortStatus`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) GetPortsOk() (*[]IoK8sApiNetworkingV1IngressPortStatus, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) SetPorts(v []IoK8sApiNetworkingV1IngressPortStatus)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *IoK8sApiNetworkingV1IngressLoadBalancerIngress) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


