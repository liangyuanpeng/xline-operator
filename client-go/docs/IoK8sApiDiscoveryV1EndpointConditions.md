# IoK8sApiDiscoveryV1EndpointConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ready** | Pointer to **bool** | ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready. For compatibility reasons, ready should never be \&quot;true\&quot; for terminating endpoints, except when the normal readiness behavior is being explicitly overridden, for example when the associated Service has set the publishNotReadyAddresses flag. | [optional] 
**Serving** | Pointer to **bool** | serving is identical to ready except that it is set regardless of the terminating state of endpoints. This condition should be set to true for a ready endpoint that is terminating. If nil, consumers should defer to the ready condition. | [optional] 
**Terminating** | Pointer to **bool** | terminating indicates that this endpoint is terminating. A nil value indicates an unknown state. Consumers should interpret this unknown state to mean that the endpoint is not terminating. | [optional] 

## Methods

### NewIoK8sApiDiscoveryV1EndpointConditions

`func NewIoK8sApiDiscoveryV1EndpointConditions() *IoK8sApiDiscoveryV1EndpointConditions`

NewIoK8sApiDiscoveryV1EndpointConditions instantiates a new IoK8sApiDiscoveryV1EndpointConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiDiscoveryV1EndpointConditionsWithDefaults

`func NewIoK8sApiDiscoveryV1EndpointConditionsWithDefaults() *IoK8sApiDiscoveryV1EndpointConditions`

NewIoK8sApiDiscoveryV1EndpointConditionsWithDefaults instantiates a new IoK8sApiDiscoveryV1EndpointConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReady

`func (o *IoK8sApiDiscoveryV1EndpointConditions) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *IoK8sApiDiscoveryV1EndpointConditions) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *IoK8sApiDiscoveryV1EndpointConditions) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *IoK8sApiDiscoveryV1EndpointConditions) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetServing

`func (o *IoK8sApiDiscoveryV1EndpointConditions) GetServing() bool`

GetServing returns the Serving field if non-nil, zero value otherwise.

### GetServingOk

`func (o *IoK8sApiDiscoveryV1EndpointConditions) GetServingOk() (*bool, bool)`

GetServingOk returns a tuple with the Serving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServing

`func (o *IoK8sApiDiscoveryV1EndpointConditions) SetServing(v bool)`

SetServing sets Serving field to given value.

### HasServing

`func (o *IoK8sApiDiscoveryV1EndpointConditions) HasServing() bool`

HasServing returns a boolean if a field has been set.

### GetTerminating

`func (o *IoK8sApiDiscoveryV1EndpointConditions) GetTerminating() bool`

GetTerminating returns the Terminating field if non-nil, zero value otherwise.

### GetTerminatingOk

`func (o *IoK8sApiDiscoveryV1EndpointConditions) GetTerminatingOk() (*bool, bool)`

GetTerminatingOk returns a tuple with the Terminating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminating

`func (o *IoK8sApiDiscoveryV1EndpointConditions) SetTerminating(v bool)`

SetTerminating sets Terminating field to given value.

### HasTerminating

`func (o *IoK8sApiDiscoveryV1EndpointConditions) HasTerminating() bool`

HasTerminating returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


