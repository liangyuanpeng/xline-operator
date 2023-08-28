# IoK8sApiCoreV1Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | Arguments to the entrypoint. The container image&#39;s CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container&#39;s environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. \&quot;$$(VAR_NAME)\&quot; will produce the string literal \&quot;$(VAR_NAME)\&quot;. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell | [optional] 
**Command** | Pointer to **[]string** | Entrypoint array. Not executed within a shell. The container image&#39;s ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container&#39;s environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. \&quot;$$(VAR_NAME)\&quot; will produce the string literal \&quot;$(VAR_NAME)\&quot;. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell | [optional] 
**Env** | Pointer to [**[]IoK8sApiCoreV1EnvVar**](IoK8sApiCoreV1EnvVar.md) | List of environment variables to set in the container. Cannot be updated. | [optional] 
**EnvFrom** | Pointer to [**[]IoK8sApiCoreV1EnvFromSource**](IoK8sApiCoreV1EnvFromSource.md) | List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated. | [optional] 
**Image** | Pointer to **string** | Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets. | [optional] 
**ImagePullPolicy** | Pointer to **string** | Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images  Possible enum values:  - &#x60;\&quot;Always\&quot;&#x60; means that kubelet always attempts to pull the latest image. Container will fail If the pull fails.  - &#x60;\&quot;IfNotPresent\&quot;&#x60; means that kubelet pulls if the image isn&#39;t present on disk. Container will fail if the image isn&#39;t present and the pull fails.  - &#x60;\&quot;Never\&quot;&#x60; means that kubelet never pulls an image, but only uses a local image. Container will fail if the image isn&#39;t present | [optional] 
**Lifecycle** | Pointer to [**IoK8sApiCoreV1Lifecycle**](IoK8sApiCoreV1Lifecycle.md) |  | [optional] 
**LivenessProbe** | Pointer to [**IoK8sApiCoreV1Probe**](IoK8sApiCoreV1Probe.md) |  | [optional] 
**Name** | **string** | Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated. | 
**Ports** | Pointer to [**[]IoK8sApiCoreV1ContainerPort**](IoK8sApiCoreV1ContainerPort.md) | List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default \&quot;0.0.0.0\&quot; address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated. | [optional] 
**ReadinessProbe** | Pointer to [**IoK8sApiCoreV1Probe**](IoK8sApiCoreV1Probe.md) |  | [optional] 
**ResizePolicy** | Pointer to [**[]IoK8sApiCoreV1ContainerResizePolicy**](IoK8sApiCoreV1ContainerResizePolicy.md) | Resources resize policy for the container. | [optional] 
**Resources** | Pointer to [**IoK8sApiCoreV1ResourceRequirements**](IoK8sApiCoreV1ResourceRequirements.md) |  | [optional] 
**SecurityContext** | Pointer to [**IoK8sApiCoreV1SecurityContext**](IoK8sApiCoreV1SecurityContext.md) |  | [optional] 
**StartupProbe** | Pointer to [**IoK8sApiCoreV1Probe**](IoK8sApiCoreV1Probe.md) |  | [optional] 
**Stdin** | Pointer to **bool** | Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false. | [optional] 
**StdinOnce** | Pointer to **bool** | Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false | [optional] 
**TerminationMessagePath** | Pointer to **string** | Optional: Path at which the file to which the container&#39;s termination message will be written is mounted into the container&#39;s filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated. | [optional] 
**TerminationMessagePolicy** | Pointer to **string** | Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.  Possible enum values:  - &#x60;\&quot;FallbackToLogsOnError\&quot;&#x60; will read the most recent contents of the container logs for the container status message when the container exits with an error and the terminationMessagePath has no contents.  - &#x60;\&quot;File\&quot;&#x60; is the default behavior and will set the container status message to the contents of the container&#39;s terminationMessagePath when the container exits. | [optional] 
**Tty** | Pointer to **bool** | Whether this container should allocate a TTY for itself, also requires &#39;stdin&#39; to be true. Default is false. | [optional] 
**VolumeDevices** | Pointer to [**[]IoK8sApiCoreV1VolumeDevice**](IoK8sApiCoreV1VolumeDevice.md) | volumeDevices is the list of block devices to be used by the container. | [optional] 
**VolumeMounts** | Pointer to [**[]IoK8sApiCoreV1VolumeMount**](IoK8sApiCoreV1VolumeMount.md) | Pod volumes to mount into the container&#39;s filesystem. Cannot be updated. | [optional] 
**WorkingDir** | Pointer to **string** | Container&#39;s working directory. If not specified, the container runtime&#39;s default will be used, which might be configured in the container image. Cannot be updated. | [optional] 

## Methods

### NewIoK8sApiCoreV1Container

`func NewIoK8sApiCoreV1Container(name string, ) *IoK8sApiCoreV1Container`

NewIoK8sApiCoreV1Container instantiates a new IoK8sApiCoreV1Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ContainerWithDefaults

`func NewIoK8sApiCoreV1ContainerWithDefaults() *IoK8sApiCoreV1Container`

NewIoK8sApiCoreV1ContainerWithDefaults instantiates a new IoK8sApiCoreV1Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *IoK8sApiCoreV1Container) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *IoK8sApiCoreV1Container) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *IoK8sApiCoreV1Container) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *IoK8sApiCoreV1Container) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCommand

`func (o *IoK8sApiCoreV1Container) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *IoK8sApiCoreV1Container) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *IoK8sApiCoreV1Container) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *IoK8sApiCoreV1Container) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEnv

`func (o *IoK8sApiCoreV1Container) GetEnv() []IoK8sApiCoreV1EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *IoK8sApiCoreV1Container) GetEnvOk() (*[]IoK8sApiCoreV1EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *IoK8sApiCoreV1Container) SetEnv(v []IoK8sApiCoreV1EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *IoK8sApiCoreV1Container) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEnvFrom

`func (o *IoK8sApiCoreV1Container) GetEnvFrom() []IoK8sApiCoreV1EnvFromSource`

GetEnvFrom returns the EnvFrom field if non-nil, zero value otherwise.

### GetEnvFromOk

`func (o *IoK8sApiCoreV1Container) GetEnvFromOk() (*[]IoK8sApiCoreV1EnvFromSource, bool)`

GetEnvFromOk returns a tuple with the EnvFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvFrom

`func (o *IoK8sApiCoreV1Container) SetEnvFrom(v []IoK8sApiCoreV1EnvFromSource)`

SetEnvFrom sets EnvFrom field to given value.

### HasEnvFrom

`func (o *IoK8sApiCoreV1Container) HasEnvFrom() bool`

HasEnvFrom returns a boolean if a field has been set.

### GetImage

`func (o *IoK8sApiCoreV1Container) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *IoK8sApiCoreV1Container) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *IoK8sApiCoreV1Container) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *IoK8sApiCoreV1Container) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImagePullPolicy

`func (o *IoK8sApiCoreV1Container) GetImagePullPolicy() string`

GetImagePullPolicy returns the ImagePullPolicy field if non-nil, zero value otherwise.

### GetImagePullPolicyOk

`func (o *IoK8sApiCoreV1Container) GetImagePullPolicyOk() (*string, bool)`

GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullPolicy

`func (o *IoK8sApiCoreV1Container) SetImagePullPolicy(v string)`

SetImagePullPolicy sets ImagePullPolicy field to given value.

### HasImagePullPolicy

`func (o *IoK8sApiCoreV1Container) HasImagePullPolicy() bool`

HasImagePullPolicy returns a boolean if a field has been set.

### GetLifecycle

`func (o *IoK8sApiCoreV1Container) GetLifecycle() IoK8sApiCoreV1Lifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *IoK8sApiCoreV1Container) GetLifecycleOk() (*IoK8sApiCoreV1Lifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *IoK8sApiCoreV1Container) SetLifecycle(v IoK8sApiCoreV1Lifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *IoK8sApiCoreV1Container) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLivenessProbe

`func (o *IoK8sApiCoreV1Container) GetLivenessProbe() IoK8sApiCoreV1Probe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *IoK8sApiCoreV1Container) GetLivenessProbeOk() (*IoK8sApiCoreV1Probe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *IoK8sApiCoreV1Container) SetLivenessProbe(v IoK8sApiCoreV1Probe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *IoK8sApiCoreV1Container) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiCoreV1Container) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1Container) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1Container) SetName(v string)`

SetName sets Name field to given value.


### GetPorts

`func (o *IoK8sApiCoreV1Container) GetPorts() []IoK8sApiCoreV1ContainerPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *IoK8sApiCoreV1Container) GetPortsOk() (*[]IoK8sApiCoreV1ContainerPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *IoK8sApiCoreV1Container) SetPorts(v []IoK8sApiCoreV1ContainerPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *IoK8sApiCoreV1Container) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *IoK8sApiCoreV1Container) GetReadinessProbe() IoK8sApiCoreV1Probe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *IoK8sApiCoreV1Container) GetReadinessProbeOk() (*IoK8sApiCoreV1Probe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *IoK8sApiCoreV1Container) SetReadinessProbe(v IoK8sApiCoreV1Probe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *IoK8sApiCoreV1Container) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetResizePolicy

`func (o *IoK8sApiCoreV1Container) GetResizePolicy() []IoK8sApiCoreV1ContainerResizePolicy`

GetResizePolicy returns the ResizePolicy field if non-nil, zero value otherwise.

### GetResizePolicyOk

`func (o *IoK8sApiCoreV1Container) GetResizePolicyOk() (*[]IoK8sApiCoreV1ContainerResizePolicy, bool)`

GetResizePolicyOk returns a tuple with the ResizePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizePolicy

`func (o *IoK8sApiCoreV1Container) SetResizePolicy(v []IoK8sApiCoreV1ContainerResizePolicy)`

SetResizePolicy sets ResizePolicy field to given value.

### HasResizePolicy

`func (o *IoK8sApiCoreV1Container) HasResizePolicy() bool`

HasResizePolicy returns a boolean if a field has been set.

### GetResources

`func (o *IoK8sApiCoreV1Container) GetResources() IoK8sApiCoreV1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IoK8sApiCoreV1Container) GetResourcesOk() (*IoK8sApiCoreV1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IoK8sApiCoreV1Container) SetResources(v IoK8sApiCoreV1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *IoK8sApiCoreV1Container) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSecurityContext

`func (o *IoK8sApiCoreV1Container) GetSecurityContext() IoK8sApiCoreV1SecurityContext`

GetSecurityContext returns the SecurityContext field if non-nil, zero value otherwise.

### GetSecurityContextOk

`func (o *IoK8sApiCoreV1Container) GetSecurityContextOk() (*IoK8sApiCoreV1SecurityContext, bool)`

GetSecurityContextOk returns a tuple with the SecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContext

`func (o *IoK8sApiCoreV1Container) SetSecurityContext(v IoK8sApiCoreV1SecurityContext)`

SetSecurityContext sets SecurityContext field to given value.

### HasSecurityContext

`func (o *IoK8sApiCoreV1Container) HasSecurityContext() bool`

HasSecurityContext returns a boolean if a field has been set.

### GetStartupProbe

`func (o *IoK8sApiCoreV1Container) GetStartupProbe() IoK8sApiCoreV1Probe`

GetStartupProbe returns the StartupProbe field if non-nil, zero value otherwise.

### GetStartupProbeOk

`func (o *IoK8sApiCoreV1Container) GetStartupProbeOk() (*IoK8sApiCoreV1Probe, bool)`

GetStartupProbeOk returns a tuple with the StartupProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupProbe

`func (o *IoK8sApiCoreV1Container) SetStartupProbe(v IoK8sApiCoreV1Probe)`

SetStartupProbe sets StartupProbe field to given value.

### HasStartupProbe

`func (o *IoK8sApiCoreV1Container) HasStartupProbe() bool`

HasStartupProbe returns a boolean if a field has been set.

### GetStdin

`func (o *IoK8sApiCoreV1Container) GetStdin() bool`

GetStdin returns the Stdin field if non-nil, zero value otherwise.

### GetStdinOk

`func (o *IoK8sApiCoreV1Container) GetStdinOk() (*bool, bool)`

GetStdinOk returns a tuple with the Stdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdin

`func (o *IoK8sApiCoreV1Container) SetStdin(v bool)`

SetStdin sets Stdin field to given value.

### HasStdin

`func (o *IoK8sApiCoreV1Container) HasStdin() bool`

HasStdin returns a boolean if a field has been set.

### GetStdinOnce

`func (o *IoK8sApiCoreV1Container) GetStdinOnce() bool`

GetStdinOnce returns the StdinOnce field if non-nil, zero value otherwise.

### GetStdinOnceOk

`func (o *IoK8sApiCoreV1Container) GetStdinOnceOk() (*bool, bool)`

GetStdinOnceOk returns a tuple with the StdinOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdinOnce

`func (o *IoK8sApiCoreV1Container) SetStdinOnce(v bool)`

SetStdinOnce sets StdinOnce field to given value.

### HasStdinOnce

`func (o *IoK8sApiCoreV1Container) HasStdinOnce() bool`

HasStdinOnce returns a boolean if a field has been set.

### GetTerminationMessagePath

`func (o *IoK8sApiCoreV1Container) GetTerminationMessagePath() string`

GetTerminationMessagePath returns the TerminationMessagePath field if non-nil, zero value otherwise.

### GetTerminationMessagePathOk

`func (o *IoK8sApiCoreV1Container) GetTerminationMessagePathOk() (*string, bool)`

GetTerminationMessagePathOk returns a tuple with the TerminationMessagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationMessagePath

`func (o *IoK8sApiCoreV1Container) SetTerminationMessagePath(v string)`

SetTerminationMessagePath sets TerminationMessagePath field to given value.

### HasTerminationMessagePath

`func (o *IoK8sApiCoreV1Container) HasTerminationMessagePath() bool`

HasTerminationMessagePath returns a boolean if a field has been set.

### GetTerminationMessagePolicy

`func (o *IoK8sApiCoreV1Container) GetTerminationMessagePolicy() string`

GetTerminationMessagePolicy returns the TerminationMessagePolicy field if non-nil, zero value otherwise.

### GetTerminationMessagePolicyOk

`func (o *IoK8sApiCoreV1Container) GetTerminationMessagePolicyOk() (*string, bool)`

GetTerminationMessagePolicyOk returns a tuple with the TerminationMessagePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationMessagePolicy

`func (o *IoK8sApiCoreV1Container) SetTerminationMessagePolicy(v string)`

SetTerminationMessagePolicy sets TerminationMessagePolicy field to given value.

### HasTerminationMessagePolicy

`func (o *IoK8sApiCoreV1Container) HasTerminationMessagePolicy() bool`

HasTerminationMessagePolicy returns a boolean if a field has been set.

### GetTty

`func (o *IoK8sApiCoreV1Container) GetTty() bool`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *IoK8sApiCoreV1Container) GetTtyOk() (*bool, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *IoK8sApiCoreV1Container) SetTty(v bool)`

SetTty sets Tty field to given value.

### HasTty

`func (o *IoK8sApiCoreV1Container) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetVolumeDevices

`func (o *IoK8sApiCoreV1Container) GetVolumeDevices() []IoK8sApiCoreV1VolumeDevice`

GetVolumeDevices returns the VolumeDevices field if non-nil, zero value otherwise.

### GetVolumeDevicesOk

`func (o *IoK8sApiCoreV1Container) GetVolumeDevicesOk() (*[]IoK8sApiCoreV1VolumeDevice, bool)`

GetVolumeDevicesOk returns a tuple with the VolumeDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDevices

`func (o *IoK8sApiCoreV1Container) SetVolumeDevices(v []IoK8sApiCoreV1VolumeDevice)`

SetVolumeDevices sets VolumeDevices field to given value.

### HasVolumeDevices

`func (o *IoK8sApiCoreV1Container) HasVolumeDevices() bool`

HasVolumeDevices returns a boolean if a field has been set.

### GetVolumeMounts

`func (o *IoK8sApiCoreV1Container) GetVolumeMounts() []IoK8sApiCoreV1VolumeMount`

GetVolumeMounts returns the VolumeMounts field if non-nil, zero value otherwise.

### GetVolumeMountsOk

`func (o *IoK8sApiCoreV1Container) GetVolumeMountsOk() (*[]IoK8sApiCoreV1VolumeMount, bool)`

GetVolumeMountsOk returns a tuple with the VolumeMounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMounts

`func (o *IoK8sApiCoreV1Container) SetVolumeMounts(v []IoK8sApiCoreV1VolumeMount)`

SetVolumeMounts sets VolumeMounts field to given value.

### HasVolumeMounts

`func (o *IoK8sApiCoreV1Container) HasVolumeMounts() bool`

HasVolumeMounts returns a boolean if a field has been set.

### GetWorkingDir

`func (o *IoK8sApiCoreV1Container) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *IoK8sApiCoreV1Container) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *IoK8sApiCoreV1Container) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.

### HasWorkingDir

`func (o *IoK8sApiCoreV1Container) HasWorkingDir() bool`

HasWorkingDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


