# IoK8sApiCoreV1LoadBalancerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingress** | Pointer to [**[]IoK8sApiCoreV1LoadBalancerIngress**](IoK8sApiCoreV1LoadBalancerIngress.md) | Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points. | [optional] 

## Methods

### NewIoK8sApiCoreV1LoadBalancerStatus

`func NewIoK8sApiCoreV1LoadBalancerStatus() *IoK8sApiCoreV1LoadBalancerStatus`

NewIoK8sApiCoreV1LoadBalancerStatus instantiates a new IoK8sApiCoreV1LoadBalancerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1LoadBalancerStatusWithDefaults

`func NewIoK8sApiCoreV1LoadBalancerStatusWithDefaults() *IoK8sApiCoreV1LoadBalancerStatus`

NewIoK8sApiCoreV1LoadBalancerStatusWithDefaults instantiates a new IoK8sApiCoreV1LoadBalancerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngress

`func (o *IoK8sApiCoreV1LoadBalancerStatus) GetIngress() []IoK8sApiCoreV1LoadBalancerIngress`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *IoK8sApiCoreV1LoadBalancerStatus) GetIngressOk() (*[]IoK8sApiCoreV1LoadBalancerIngress, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *IoK8sApiCoreV1LoadBalancerStatus) SetIngress(v []IoK8sApiCoreV1LoadBalancerIngress)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *IoK8sApiCoreV1LoadBalancerStatus) HasIngress() bool`

HasIngress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


