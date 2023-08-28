# IoK8sApiNetworkingV1IngressStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | Pointer to [**IoK8sApiNetworkingV1IngressLoadBalancerStatus**](IoK8sApiNetworkingV1IngressLoadBalancerStatus.md) |  | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IngressStatus

`func NewIoK8sApiNetworkingV1IngressStatus() *IoK8sApiNetworkingV1IngressStatus`

NewIoK8sApiNetworkingV1IngressStatus instantiates a new IoK8sApiNetworkingV1IngressStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IngressStatusWithDefaults

`func NewIoK8sApiNetworkingV1IngressStatusWithDefaults() *IoK8sApiNetworkingV1IngressStatus`

NewIoK8sApiNetworkingV1IngressStatusWithDefaults instantiates a new IoK8sApiNetworkingV1IngressStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancer

`func (o *IoK8sApiNetworkingV1IngressStatus) GetLoadBalancer() IoK8sApiNetworkingV1IngressLoadBalancerStatus`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *IoK8sApiNetworkingV1IngressStatus) GetLoadBalancerOk() (*IoK8sApiNetworkingV1IngressLoadBalancerStatus, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *IoK8sApiNetworkingV1IngressStatus) SetLoadBalancer(v IoK8sApiNetworkingV1IngressLoadBalancerStatus)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *IoK8sApiNetworkingV1IngressStatus) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


