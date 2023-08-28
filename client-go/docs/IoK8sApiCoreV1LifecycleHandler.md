# IoK8sApiCoreV1LifecycleHandler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exec** | Pointer to [**IoK8sApiCoreV1ExecAction**](IoK8sApiCoreV1ExecAction.md) |  | [optional] 
**HttpGet** | Pointer to [**IoK8sApiCoreV1HTTPGetAction**](IoK8sApiCoreV1HTTPGetAction.md) |  | [optional] 
**TcpSocket** | Pointer to [**IoK8sApiCoreV1TCPSocketAction**](IoK8sApiCoreV1TCPSocketAction.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1LifecycleHandler

`func NewIoK8sApiCoreV1LifecycleHandler() *IoK8sApiCoreV1LifecycleHandler`

NewIoK8sApiCoreV1LifecycleHandler instantiates a new IoK8sApiCoreV1LifecycleHandler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1LifecycleHandlerWithDefaults

`func NewIoK8sApiCoreV1LifecycleHandlerWithDefaults() *IoK8sApiCoreV1LifecycleHandler`

NewIoK8sApiCoreV1LifecycleHandlerWithDefaults instantiates a new IoK8sApiCoreV1LifecycleHandler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExec

`func (o *IoK8sApiCoreV1LifecycleHandler) GetExec() IoK8sApiCoreV1ExecAction`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *IoK8sApiCoreV1LifecycleHandler) GetExecOk() (*IoK8sApiCoreV1ExecAction, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *IoK8sApiCoreV1LifecycleHandler) SetExec(v IoK8sApiCoreV1ExecAction)`

SetExec sets Exec field to given value.

### HasExec

`func (o *IoK8sApiCoreV1LifecycleHandler) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetHttpGet

`func (o *IoK8sApiCoreV1LifecycleHandler) GetHttpGet() IoK8sApiCoreV1HTTPGetAction`

GetHttpGet returns the HttpGet field if non-nil, zero value otherwise.

### GetHttpGetOk

`func (o *IoK8sApiCoreV1LifecycleHandler) GetHttpGetOk() (*IoK8sApiCoreV1HTTPGetAction, bool)`

GetHttpGetOk returns a tuple with the HttpGet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpGet

`func (o *IoK8sApiCoreV1LifecycleHandler) SetHttpGet(v IoK8sApiCoreV1HTTPGetAction)`

SetHttpGet sets HttpGet field to given value.

### HasHttpGet

`func (o *IoK8sApiCoreV1LifecycleHandler) HasHttpGet() bool`

HasHttpGet returns a boolean if a field has been set.

### GetTcpSocket

`func (o *IoK8sApiCoreV1LifecycleHandler) GetTcpSocket() IoK8sApiCoreV1TCPSocketAction`

GetTcpSocket returns the TcpSocket field if non-nil, zero value otherwise.

### GetTcpSocketOk

`func (o *IoK8sApiCoreV1LifecycleHandler) GetTcpSocketOk() (*IoK8sApiCoreV1TCPSocketAction, bool)`

GetTcpSocketOk returns a tuple with the TcpSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSocket

`func (o *IoK8sApiCoreV1LifecycleHandler) SetTcpSocket(v IoK8sApiCoreV1TCPSocketAction)`

SetTcpSocket sets TcpSocket field to given value.

### HasTcpSocket

`func (o *IoK8sApiCoreV1LifecycleHandler) HasTcpSocket() bool`

HasTcpSocket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


