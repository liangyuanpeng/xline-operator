# IoK8sApiCoreV1HTTPGetAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host name to connect to, defaults to the pod IP. You probably want to set \&quot;Host\&quot; in httpHeaders instead. | [optional] 
**HttpHeaders** | Pointer to [**[]IoK8sApiCoreV1HTTPHeader**](IoK8sApiCoreV1HTTPHeader.md) | Custom headers to set in the request. HTTP allows repeated headers. | [optional] 
**Path** | Pointer to **string** | Path to access on the HTTP server. | [optional] 
**Port** | **string** | IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number. | 
**Scheme** | Pointer to **string** | Scheme to use for connecting to the host. Defaults to HTTP.  Possible enum values:  - &#x60;\&quot;HTTP\&quot;&#x60; means that the scheme used will be http://  - &#x60;\&quot;HTTPS\&quot;&#x60; means that the scheme used will be https:// | [optional] 

## Methods

### NewIoK8sApiCoreV1HTTPGetAction

`func NewIoK8sApiCoreV1HTTPGetAction(port string, ) *IoK8sApiCoreV1HTTPGetAction`

NewIoK8sApiCoreV1HTTPGetAction instantiates a new IoK8sApiCoreV1HTTPGetAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1HTTPGetActionWithDefaults

`func NewIoK8sApiCoreV1HTTPGetActionWithDefaults() *IoK8sApiCoreV1HTTPGetAction`

NewIoK8sApiCoreV1HTTPGetActionWithDefaults instantiates a new IoK8sApiCoreV1HTTPGetAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *IoK8sApiCoreV1HTTPGetAction) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IoK8sApiCoreV1HTTPGetAction) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IoK8sApiCoreV1HTTPGetAction) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IoK8sApiCoreV1HTTPGetAction) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHttpHeaders

`func (o *IoK8sApiCoreV1HTTPGetAction) GetHttpHeaders() []IoK8sApiCoreV1HTTPHeader`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *IoK8sApiCoreV1HTTPGetAction) GetHttpHeadersOk() (*[]IoK8sApiCoreV1HTTPHeader, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *IoK8sApiCoreV1HTTPGetAction) SetHttpHeaders(v []IoK8sApiCoreV1HTTPHeader)`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *IoK8sApiCoreV1HTTPGetAction) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.

### GetPath

`func (o *IoK8sApiCoreV1HTTPGetAction) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoK8sApiCoreV1HTTPGetAction) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoK8sApiCoreV1HTTPGetAction) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IoK8sApiCoreV1HTTPGetAction) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *IoK8sApiCoreV1HTTPGetAction) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sApiCoreV1HTTPGetAction) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sApiCoreV1HTTPGetAction) SetPort(v string)`

SetPort sets Port field to given value.


### GetScheme

`func (o *IoK8sApiCoreV1HTTPGetAction) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *IoK8sApiCoreV1HTTPGetAction) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *IoK8sApiCoreV1HTTPGetAction) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *IoK8sApiCoreV1HTTPGetAction) HasScheme() bool`

HasScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


