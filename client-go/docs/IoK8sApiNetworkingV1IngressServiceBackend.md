# IoK8sApiNetworkingV1IngressServiceBackend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is the referenced service. The service must exist in the same namespace as the Ingress object. | 
**Port** | Pointer to [**IoK8sApiNetworkingV1ServiceBackendPort**](IoK8sApiNetworkingV1ServiceBackendPort.md) |  | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IngressServiceBackend

`func NewIoK8sApiNetworkingV1IngressServiceBackend(name string, ) *IoK8sApiNetworkingV1IngressServiceBackend`

NewIoK8sApiNetworkingV1IngressServiceBackend instantiates a new IoK8sApiNetworkingV1IngressServiceBackend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IngressServiceBackendWithDefaults

`func NewIoK8sApiNetworkingV1IngressServiceBackendWithDefaults() *IoK8sApiNetworkingV1IngressServiceBackend`

NewIoK8sApiNetworkingV1IngressServiceBackendWithDefaults instantiates a new IoK8sApiNetworkingV1IngressServiceBackend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiNetworkingV1IngressServiceBackend) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiNetworkingV1IngressServiceBackend) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiNetworkingV1IngressServiceBackend) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *IoK8sApiNetworkingV1IngressServiceBackend) GetPort() IoK8sApiNetworkingV1ServiceBackendPort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sApiNetworkingV1IngressServiceBackend) GetPortOk() (*IoK8sApiNetworkingV1ServiceBackendPort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sApiNetworkingV1IngressServiceBackend) SetPort(v IoK8sApiNetworkingV1ServiceBackendPort)`

SetPort sets Port field to given value.

### HasPort

`func (o *IoK8sApiNetworkingV1IngressServiceBackend) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


