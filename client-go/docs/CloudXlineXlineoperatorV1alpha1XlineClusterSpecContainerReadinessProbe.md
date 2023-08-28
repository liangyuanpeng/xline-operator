# CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exec** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec.md) |  | [optional] 
**FailureThreshold** | Pointer to **int32** | Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1. | [optional] 
**Grpc** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc.md) |  | [optional] 
**HttpGet** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet.md) |  | [optional] 
**InitialDelaySeconds** | Pointer to **int32** | Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | [optional] 
**PeriodSeconds** | Pointer to **int32** | How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1. | [optional] 
**SuccessThreshold** | Pointer to **int32** | Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1. | [optional] 
**TcpSocket** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeTcpSocket**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeTcpSocket.md) |  | [optional] 
**TerminationGracePeriodSeconds** | Pointer to **int64** | Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod&#39;s terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset. | [optional] 
**TimeoutSeconds** | Pointer to **int32** | Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | [optional] 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbeWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbeWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbeWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExec

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetExec() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetExecOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) SetExec(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec)`

SetExec sets Exec field to given value.

### HasExec

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.

### GetGrpc

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetGrpc() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc`

GetGrpc returns the Grpc field if non-nil, zero value otherwise.

### GetGrpcOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetGrpcOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc, bool)`

GetGrpcOk returns a tuple with the Grpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpc

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) SetGrpc(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeGrpc)`

SetGrpc sets Grpc field to given value.

### HasGrpc

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) HasGrpc() bool`

HasGrpc returns a boolean if a field has been set.

### GetHttpGet

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetHttpGet() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet`

GetHttpGet returns the HttpGet field if non-nil, zero value otherwise.

### GetHttpGetOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetHttpGetOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet, bool)`

GetHttpGetOk returns a tuple with the HttpGet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpGet

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) SetHttpGet(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartHttpGet)`

SetHttpGet sets HttpGet field to given value.

### HasHttpGet

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) HasHttpGet() bool`

HasHttpGet returns a boolean if a field has been set.

### GetInitialDelaySeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetInitialDelaySeconds() int32`

GetInitialDelaySeconds returns the InitialDelaySeconds field if non-nil, zero value otherwise.

### GetInitialDelaySecondsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetInitialDelaySecondsOk() (*int32, bool)`

GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDelaySeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) SetInitialDelaySeconds(v int32)`

SetInitialDelaySeconds sets InitialDelaySeconds field to given value.

### HasInitialDelaySeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) HasInitialDelaySeconds() bool`

HasInitialDelaySeconds returns a boolean if a field has been set.

### GetPeriodSeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.

### HasPeriodSeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) HasPeriodSeconds() bool`

HasPeriodSeconds returns a boolean if a field has been set.

### GetSuccessThreshold

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.

### HasSuccessThreshold

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) HasSuccessThreshold() bool`

HasSuccessThreshold returns a boolean if a field has been set.

### GetTcpSocket

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetTcpSocket() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeTcpSocket`

GetTcpSocket returns the TcpSocket field if non-nil, zero value otherwise.

### GetTcpSocketOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetTcpSocketOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeTcpSocket, bool)`

GetTcpSocketOk returns a tuple with the TcpSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSocket

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) SetTcpSocket(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbeTcpSocket)`

SetTcpSocket sets TcpSocket field to given value.

### HasTcpSocket

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) HasTcpSocket() bool`

HasTcpSocket returns a boolean if a field has been set.

### GetTerminationGracePeriodSeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetTerminationGracePeriodSeconds() int64`

GetTerminationGracePeriodSeconds returns the TerminationGracePeriodSeconds field if non-nil, zero value otherwise.

### GetTerminationGracePeriodSecondsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetTerminationGracePeriodSecondsOk() (*int64, bool)`

GetTerminationGracePeriodSecondsOk returns a tuple with the TerminationGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationGracePeriodSeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) SetTerminationGracePeriodSeconds(v int64)`

SetTerminationGracePeriodSeconds sets TerminationGracePeriodSeconds field to given value.

### HasTerminationGracePeriodSeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) HasTerminationGracePeriodSeconds() bool`

HasTerminationGracePeriodSeconds returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


