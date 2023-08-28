# IoK8sApiCoreV1GRPCAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | Port number of the gRPC service. Number must be in the range 1 to 65535. | 
**Service** | Pointer to **string** | Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC. | [optional] 

## Methods

### NewIoK8sApiCoreV1GRPCAction

`func NewIoK8sApiCoreV1GRPCAction(port int32, ) *IoK8sApiCoreV1GRPCAction`

NewIoK8sApiCoreV1GRPCAction instantiates a new IoK8sApiCoreV1GRPCAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1GRPCActionWithDefaults

`func NewIoK8sApiCoreV1GRPCActionWithDefaults() *IoK8sApiCoreV1GRPCAction`

NewIoK8sApiCoreV1GRPCActionWithDefaults instantiates a new IoK8sApiCoreV1GRPCAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *IoK8sApiCoreV1GRPCAction) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sApiCoreV1GRPCAction) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sApiCoreV1GRPCAction) SetPort(v int32)`

SetPort sets Port field to given value.


### GetService

`func (o *IoK8sApiCoreV1GRPCAction) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IoK8sApiCoreV1GRPCAction) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IoK8sApiCoreV1GRPCAction) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *IoK8sApiCoreV1GRPCAction) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


