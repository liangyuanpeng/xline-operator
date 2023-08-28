# IoK8sApiCoreV1Probe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exec** | Pointer to [**IoK8sApiCoreV1ExecAction**](IoK8sApiCoreV1ExecAction.md) |  | [optional] 
**FailureThreshold** | Pointer to **int32** | Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1. | [optional] 
**Grpc** | Pointer to [**IoK8sApiCoreV1GRPCAction**](IoK8sApiCoreV1GRPCAction.md) |  | [optional] 
**HttpGet** | Pointer to [**IoK8sApiCoreV1HTTPGetAction**](IoK8sApiCoreV1HTTPGetAction.md) |  | [optional] 
**InitialDelaySeconds** | Pointer to **int32** | Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | [optional] 
**PeriodSeconds** | Pointer to **int32** | How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1. | [optional] 
**SuccessThreshold** | Pointer to **int32** | Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1. | [optional] 
**TcpSocket** | Pointer to [**IoK8sApiCoreV1TCPSocketAction**](IoK8sApiCoreV1TCPSocketAction.md) |  | [optional] 
**TerminationGracePeriodSeconds** | Pointer to **int64** | Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod&#39;s terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset. | [optional] 
**TimeoutSeconds** | Pointer to **int32** | Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | [optional] 

## Methods

### NewIoK8sApiCoreV1Probe

`func NewIoK8sApiCoreV1Probe() *IoK8sApiCoreV1Probe`

NewIoK8sApiCoreV1Probe instantiates a new IoK8sApiCoreV1Probe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ProbeWithDefaults

`func NewIoK8sApiCoreV1ProbeWithDefaults() *IoK8sApiCoreV1Probe`

NewIoK8sApiCoreV1ProbeWithDefaults instantiates a new IoK8sApiCoreV1Probe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExec

`func (o *IoK8sApiCoreV1Probe) GetExec() IoK8sApiCoreV1ExecAction`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *IoK8sApiCoreV1Probe) GetExecOk() (*IoK8sApiCoreV1ExecAction, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *IoK8sApiCoreV1Probe) SetExec(v IoK8sApiCoreV1ExecAction)`

SetExec sets Exec field to given value.

### HasExec

`func (o *IoK8sApiCoreV1Probe) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *IoK8sApiCoreV1Probe) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *IoK8sApiCoreV1Probe) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *IoK8sApiCoreV1Probe) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *IoK8sApiCoreV1Probe) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.

### GetGrpc

`func (o *IoK8sApiCoreV1Probe) GetGrpc() IoK8sApiCoreV1GRPCAction`

GetGrpc returns the Grpc field if non-nil, zero value otherwise.

### GetGrpcOk

`func (o *IoK8sApiCoreV1Probe) GetGrpcOk() (*IoK8sApiCoreV1GRPCAction, bool)`

GetGrpcOk returns a tuple with the Grpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpc

`func (o *IoK8sApiCoreV1Probe) SetGrpc(v IoK8sApiCoreV1GRPCAction)`

SetGrpc sets Grpc field to given value.

### HasGrpc

`func (o *IoK8sApiCoreV1Probe) HasGrpc() bool`

HasGrpc returns a boolean if a field has been set.

### GetHttpGet

`func (o *IoK8sApiCoreV1Probe) GetHttpGet() IoK8sApiCoreV1HTTPGetAction`

GetHttpGet returns the HttpGet field if non-nil, zero value otherwise.

### GetHttpGetOk

`func (o *IoK8sApiCoreV1Probe) GetHttpGetOk() (*IoK8sApiCoreV1HTTPGetAction, bool)`

GetHttpGetOk returns a tuple with the HttpGet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpGet

`func (o *IoK8sApiCoreV1Probe) SetHttpGet(v IoK8sApiCoreV1HTTPGetAction)`

SetHttpGet sets HttpGet field to given value.

### HasHttpGet

`func (o *IoK8sApiCoreV1Probe) HasHttpGet() bool`

HasHttpGet returns a boolean if a field has been set.

### GetInitialDelaySeconds

`func (o *IoK8sApiCoreV1Probe) GetInitialDelaySeconds() int32`

GetInitialDelaySeconds returns the InitialDelaySeconds field if non-nil, zero value otherwise.

### GetInitialDelaySecondsOk

`func (o *IoK8sApiCoreV1Probe) GetInitialDelaySecondsOk() (*int32, bool)`

GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDelaySeconds

`func (o *IoK8sApiCoreV1Probe) SetInitialDelaySeconds(v int32)`

SetInitialDelaySeconds sets InitialDelaySeconds field to given value.

### HasInitialDelaySeconds

`func (o *IoK8sApiCoreV1Probe) HasInitialDelaySeconds() bool`

HasInitialDelaySeconds returns a boolean if a field has been set.

### GetPeriodSeconds

`func (o *IoK8sApiCoreV1Probe) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *IoK8sApiCoreV1Probe) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *IoK8sApiCoreV1Probe) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.

### HasPeriodSeconds

`func (o *IoK8sApiCoreV1Probe) HasPeriodSeconds() bool`

HasPeriodSeconds returns a boolean if a field has been set.

### GetSuccessThreshold

`func (o *IoK8sApiCoreV1Probe) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *IoK8sApiCoreV1Probe) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *IoK8sApiCoreV1Probe) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.

### HasSuccessThreshold

`func (o *IoK8sApiCoreV1Probe) HasSuccessThreshold() bool`

HasSuccessThreshold returns a boolean if a field has been set.

### GetTcpSocket

`func (o *IoK8sApiCoreV1Probe) GetTcpSocket() IoK8sApiCoreV1TCPSocketAction`

GetTcpSocket returns the TcpSocket field if non-nil, zero value otherwise.

### GetTcpSocketOk

`func (o *IoK8sApiCoreV1Probe) GetTcpSocketOk() (*IoK8sApiCoreV1TCPSocketAction, bool)`

GetTcpSocketOk returns a tuple with the TcpSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSocket

`func (o *IoK8sApiCoreV1Probe) SetTcpSocket(v IoK8sApiCoreV1TCPSocketAction)`

SetTcpSocket sets TcpSocket field to given value.

### HasTcpSocket

`func (o *IoK8sApiCoreV1Probe) HasTcpSocket() bool`

HasTcpSocket returns a boolean if a field has been set.

### GetTerminationGracePeriodSeconds

`func (o *IoK8sApiCoreV1Probe) GetTerminationGracePeriodSeconds() int64`

GetTerminationGracePeriodSeconds returns the TerminationGracePeriodSeconds field if non-nil, zero value otherwise.

### GetTerminationGracePeriodSecondsOk

`func (o *IoK8sApiCoreV1Probe) GetTerminationGracePeriodSecondsOk() (*int64, bool)`

GetTerminationGracePeriodSecondsOk returns a tuple with the TerminationGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationGracePeriodSeconds

`func (o *IoK8sApiCoreV1Probe) SetTerminationGracePeriodSeconds(v int64)`

SetTerminationGracePeriodSeconds sets TerminationGracePeriodSeconds field to given value.

### HasTerminationGracePeriodSeconds

`func (o *IoK8sApiCoreV1Probe) HasTerminationGracePeriodSeconds() bool`

HasTerminationGracePeriodSeconds returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *IoK8sApiCoreV1Probe) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *IoK8sApiCoreV1Probe) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *IoK8sApiCoreV1Probe) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *IoK8sApiCoreV1Probe) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


