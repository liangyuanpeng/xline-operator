# CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | Port number of the gRPC service. Number must be in the range 1 to 65535. | 
**Service** | Pointer to **string** | Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC. | [optional] 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc(port int32, ) *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpcWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpcWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpcWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc) SetPort(v int32)`

SetPort sets Port field to given value.


### GetService

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


