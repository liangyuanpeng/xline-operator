# IoK8sApiNetworkingV1IngressClassSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to **string** | controller refers to the name of the controller that should handle this class. This allows for different \&quot;flavors\&quot; that are controlled by the same controller. For example, you may have different parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. \&quot;acme.io/ingress-controller\&quot;. This field is immutable. | [optional] 
**Parameters** | Pointer to [**IoK8sApiNetworkingV1IngressClassParametersReference**](IoK8sApiNetworkingV1IngressClassParametersReference.md) |  | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IngressClassSpec

`func NewIoK8sApiNetworkingV1IngressClassSpec() *IoK8sApiNetworkingV1IngressClassSpec`

NewIoK8sApiNetworkingV1IngressClassSpec instantiates a new IoK8sApiNetworkingV1IngressClassSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IngressClassSpecWithDefaults

`func NewIoK8sApiNetworkingV1IngressClassSpecWithDefaults() *IoK8sApiNetworkingV1IngressClassSpec`

NewIoK8sApiNetworkingV1IngressClassSpecWithDefaults instantiates a new IoK8sApiNetworkingV1IngressClassSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetController

`func (o *IoK8sApiNetworkingV1IngressClassSpec) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *IoK8sApiNetworkingV1IngressClassSpec) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *IoK8sApiNetworkingV1IngressClassSpec) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *IoK8sApiNetworkingV1IngressClassSpec) HasController() bool`

HasController returns a boolean if a field has been set.

### GetParameters

`func (o *IoK8sApiNetworkingV1IngressClassSpec) GetParameters() IoK8sApiNetworkingV1IngressClassParametersReference`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoK8sApiNetworkingV1IngressClassSpec) GetParametersOk() (*IoK8sApiNetworkingV1IngressClassParametersReference, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoK8sApiNetworkingV1IngressClassSpec) SetParameters(v IoK8sApiNetworkingV1IngressClassParametersReference)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoK8sApiNetworkingV1IngressClassSpec) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


