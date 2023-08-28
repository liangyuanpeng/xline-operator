# IoK8sApiNetworkingV1IngressBackend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to [**IoK8sApiCoreV1TypedLocalObjectReference**](IoK8sApiCoreV1TypedLocalObjectReference.md) |  | [optional] 
**Service** | Pointer to [**IoK8sApiNetworkingV1IngressServiceBackend**](IoK8sApiNetworkingV1IngressServiceBackend.md) |  | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IngressBackend

`func NewIoK8sApiNetworkingV1IngressBackend() *IoK8sApiNetworkingV1IngressBackend`

NewIoK8sApiNetworkingV1IngressBackend instantiates a new IoK8sApiNetworkingV1IngressBackend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IngressBackendWithDefaults

`func NewIoK8sApiNetworkingV1IngressBackendWithDefaults() *IoK8sApiNetworkingV1IngressBackend`

NewIoK8sApiNetworkingV1IngressBackendWithDefaults instantiates a new IoK8sApiNetworkingV1IngressBackend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *IoK8sApiNetworkingV1IngressBackend) GetResource() IoK8sApiCoreV1TypedLocalObjectReference`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IoK8sApiNetworkingV1IngressBackend) GetResourceOk() (*IoK8sApiCoreV1TypedLocalObjectReference, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IoK8sApiNetworkingV1IngressBackend) SetResource(v IoK8sApiCoreV1TypedLocalObjectReference)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IoK8sApiNetworkingV1IngressBackend) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetService

`func (o *IoK8sApiNetworkingV1IngressBackend) GetService() IoK8sApiNetworkingV1IngressServiceBackend`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IoK8sApiNetworkingV1IngressBackend) GetServiceOk() (*IoK8sApiNetworkingV1IngressServiceBackend, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IoK8sApiNetworkingV1IngressBackend) SetService(v IoK8sApiNetworkingV1IngressServiceBackend)`

SetService sets Service field to given value.

### HasService

`func (o *IoK8sApiNetworkingV1IngressBackend) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


