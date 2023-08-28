# CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host name to connect to, defaults to the pod IP. You probably want to set \&quot;Host\&quot; in httpHeaders instead. | [optional] 
**HttpHeaders** | Pointer to [**[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGetHttpHeadersInner**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGetHttpHeadersInner.md) | Custom headers to set in the request. HTTP allows repeated headers. | [optional] 
**Path** | Pointer to **string** | Path to access on the HTTP server. | [optional] 
**Port** | **string** | Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. | 
**Scheme** | Pointer to **string** | Scheme to use for connecting to the host. Defaults to HTTP.   | [optional] 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet(port string, ) *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGetWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGetWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGetWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHttpHeaders

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) GetHttpHeaders() []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGetHttpHeadersInner`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) GetHttpHeadersOk() (*[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGetHttpHeadersInner, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) SetHttpHeaders(v []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGetHttpHeadersInner)`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.

### GetPath

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) SetPort(v string)`

SetPort sets Port field to given value.


### GetScheme

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet) HasScheme() bool`

HasScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


