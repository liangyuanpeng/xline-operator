# CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | Arguments to the entrypoint. The container image&#39;s CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container&#39;s environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. \&quot;$$(VAR_NAME)\&quot; will produce the string literal \&quot;$(VAR_NAME)\&quot;. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell | [optional] 
**Command** | Pointer to **[]string** | Entrypoint array. Not executed within a shell. The container image&#39;s ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container&#39;s environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. \&quot;$$(VAR_NAME)\&quot; will produce the string literal \&quot;$(VAR_NAME)\&quot;. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell | [optional] 
**Env** | Pointer to [**[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner.md) | List of environment variables to set in the container. Cannot be updated. | [optional] 
**EnvFrom** | Pointer to [**[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInner**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInner.md) | List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated. | [optional] 
**Image** | Pointer to **string** | Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets. | [optional] 
**ImagePullPolicy** | Pointer to **string** | Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images   | [optional] 
**Lifecycle** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecycle**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecycle.md) |  | [optional] 
**LivenessProbe** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbe**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbe.md) |  | [optional] 
**Name** | **string** | Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated. | 
**Ports** | Pointer to [**[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerPortsInner**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerPortsInner.md) | List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default \&quot;0.0.0.0\&quot; address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated. | [optional] 
**ReadinessProbe** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe.md) |  | [optional] 
**Resources** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources.md) |  | [optional] 
**SecurityContext** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext.md) |  | [optional] 
**StartupProbe** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerStartupProbe**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerStartupProbe.md) |  | [optional] 
**Stdin** | Pointer to **bool** | Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false. | [optional] 
**StdinOnce** | Pointer to **bool** | Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false | [optional] 
**TerminationMessagePath** | Pointer to **string** | Optional: Path at which the file to which the container&#39;s termination message will be written is mounted into the container&#39;s filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated. | [optional] 
**TerminationMessagePolicy** | Pointer to **string** | Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.   | [optional] 
**Tty** | Pointer to **bool** | Whether this container should allocate a TTY for itself, also requires &#39;stdin&#39; to be true. Default is false. | [optional] 
**VolumeDevices** | Pointer to [**[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerVolumeDevicesInner**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerVolumeDevicesInner.md) | volumeDevices is the list of block devices to be used by the container. | [optional] 
**VolumeMounts** | Pointer to [**[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerVolumeMountsInner**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerVolumeMountsInner.md) | Pod volumes to mount into the container&#39;s filesystem. Cannot be updated. | [optional] 
**WorkingDir** | Pointer to **string** | Container&#39;s working directory. If not specified, the container runtime&#39;s default will be used, which might be configured in the container image. Cannot be updated. | [optional] 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer(name string, ) *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCommand

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEnv

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetEnv() []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetEnvOk() (*[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetEnv(v []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvFrom

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetEnvFrom() []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInner`

GetEnvFrom returns the EnvFrom field if non-nil, zero value otherwise.

### GetEnvFromOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetEnvFromOk() (*[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInner, bool)`

GetEnvFromOk returns a tuple with the EnvFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvFrom

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetEnvFrom(v []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvFromInner)`

SetEnvFrom sets EnvFrom field to given value.

### HasEnvFrom

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasEnvFrom() bool`

HasEnvFrom returns a boolean if a field has been set.

### GetImage

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImagePullPolicy

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetImagePullPolicy() string`

GetImagePullPolicy returns the ImagePullPolicy field if non-nil, zero value otherwise.

### GetImagePullPolicyOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetImagePullPolicyOk() (*string, bool)`

GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullPolicy

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetImagePullPolicy(v string)`

SetImagePullPolicy sets ImagePullPolicy field to given value.

### HasImagePullPolicy

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasImagePullPolicy() bool`

HasImagePullPolicy returns a boolean if a field has been set.

### GetLifecycle

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetLifecycle() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetLifecycleOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetLifecycle(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLivenessProbe

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetLivenessProbe() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetLivenessProbeOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetLivenessProbe(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLivenessProbe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### GetName

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetName(v string)`

SetName sets Name field to given value.


### GetPorts

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetPorts() []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetPortsOk() (*[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetPorts(v []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetReadinessProbe() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetReadinessProbeOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetReadinessProbe(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerReadinessProbe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetResources

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetResources() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetResourcesOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetResources(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSecurityContext

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetSecurityContext() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext`

GetSecurityContext returns the SecurityContext field if non-nil, zero value otherwise.

### GetSecurityContextOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetSecurityContextOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext, bool)`

GetSecurityContextOk returns a tuple with the SecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContext

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetSecurityContext(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext)`

SetSecurityContext sets SecurityContext field to given value.

### HasSecurityContext

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasSecurityContext() bool`

HasSecurityContext returns a boolean if a field has been set.

### GetStartupProbe

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetStartupProbe() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerStartupProbe`

GetStartupProbe returns the StartupProbe field if non-nil, zero value otherwise.

### GetStartupProbeOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetStartupProbeOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerStartupProbe, bool)`

GetStartupProbeOk returns a tuple with the StartupProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupProbe

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetStartupProbe(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerStartupProbe)`

SetStartupProbe sets StartupProbe field to given value.

### HasStartupProbe

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasStartupProbe() bool`

HasStartupProbe returns a boolean if a field has been set.

### GetStdin

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetStdin() bool`

GetStdin returns the Stdin field if non-nil, zero value otherwise.

### GetStdinOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetStdinOk() (*bool, bool)`

GetStdinOk returns a tuple with the Stdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdin

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetStdin(v bool)`

SetStdin sets Stdin field to given value.

### HasStdin

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasStdin() bool`

HasStdin returns a boolean if a field has been set.

### GetStdinOnce

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetStdinOnce() bool`

GetStdinOnce returns the StdinOnce field if non-nil, zero value otherwise.

### GetStdinOnceOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetStdinOnceOk() (*bool, bool)`

GetStdinOnceOk returns a tuple with the StdinOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdinOnce

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetStdinOnce(v bool)`

SetStdinOnce sets StdinOnce field to given value.

### HasStdinOnce

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasStdinOnce() bool`

HasStdinOnce returns a boolean if a field has been set.

### GetTerminationMessagePath

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetTerminationMessagePath() string`

GetTerminationMessagePath returns the TerminationMessagePath field if non-nil, zero value otherwise.

### GetTerminationMessagePathOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetTerminationMessagePathOk() (*string, bool)`

GetTerminationMessagePathOk returns a tuple with the TerminationMessagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationMessagePath

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetTerminationMessagePath(v string)`

SetTerminationMessagePath sets TerminationMessagePath field to given value.

### HasTerminationMessagePath

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasTerminationMessagePath() bool`

HasTerminationMessagePath returns a boolean if a field has been set.

### GetTerminationMessagePolicy

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetTerminationMessagePolicy() string`

GetTerminationMessagePolicy returns the TerminationMessagePolicy field if non-nil, zero value otherwise.

### GetTerminationMessagePolicyOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetTerminationMessagePolicyOk() (*string, bool)`

GetTerminationMessagePolicyOk returns a tuple with the TerminationMessagePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationMessagePolicy

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetTerminationMessagePolicy(v string)`

SetTerminationMessagePolicy sets TerminationMessagePolicy field to given value.

### HasTerminationMessagePolicy

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasTerminationMessagePolicy() bool`

HasTerminationMessagePolicy returns a boolean if a field has been set.

### GetTty

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetTty() bool`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetTtyOk() (*bool, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetTty(v bool)`

SetTty sets Tty field to given value.

### HasTty

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetVolumeDevices

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetVolumeDevices() []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerVolumeDevicesInner`

GetVolumeDevices returns the VolumeDevices field if non-nil, zero value otherwise.

### GetVolumeDevicesOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetVolumeDevicesOk() (*[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerVolumeDevicesInner, bool)`

GetVolumeDevicesOk returns a tuple with the VolumeDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDevices

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetVolumeDevices(v []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerVolumeDevicesInner)`

SetVolumeDevices sets VolumeDevices field to given value.

### HasVolumeDevices

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasVolumeDevices() bool`

HasVolumeDevices returns a boolean if a field has been set.

### GetVolumeMounts

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetVolumeMounts() []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerVolumeMountsInner`

GetVolumeMounts returns the VolumeMounts field if non-nil, zero value otherwise.

### GetVolumeMountsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetVolumeMountsOk() (*[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerVolumeMountsInner, bool)`

GetVolumeMountsOk returns a tuple with the VolumeMounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMounts

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetVolumeMounts(v []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerVolumeMountsInner)`

SetVolumeMounts sets VolumeMounts field to given value.

### HasVolumeMounts

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasVolumeMounts() bool`

HasVolumeMounts returns a boolean if a field has been set.

### GetWorkingDir

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.

### HasWorkingDir

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer) HasWorkingDir() bool`

HasWorkingDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


