# IoK8sApiCoreV1SecurityContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPrivilegeEscalation** | Pointer to **bool** | AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows. | [optional] 
**Capabilities** | Pointer to [**IoK8sApiCoreV1Capabilities**](IoK8sApiCoreV1Capabilities.md) |  | [optional] 
**Privileged** | Pointer to **bool** | Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**ProcMount** | Pointer to **string** | procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.  Possible enum values:  - &#x60;\&quot;Default\&quot;&#x60; uses the container runtime defaults for readonly and masked paths for /proc. Most container runtimes mask certain paths in /proc to avoid accidental security exposure of special devices or information.  - &#x60;\&quot;Unmasked\&quot;&#x60; bypasses the default masking behavior of the container runtime and ensures the newly created /proc the container stays in tact with no modifications. | [optional] 
**ReadOnlyRootFilesystem** | Pointer to **bool** | Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**RunAsGroup** | Pointer to **int64** | The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**RunAsNonRoot** | Pointer to **bool** | Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. | [optional] 
**RunAsUser** | Pointer to **int64** | The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**SeLinuxOptions** | Pointer to [**IoK8sApiCoreV1SELinuxOptions**](IoK8sApiCoreV1SELinuxOptions.md) |  | [optional] 
**SeccompProfile** | Pointer to [**IoK8sApiCoreV1SeccompProfile**](IoK8sApiCoreV1SeccompProfile.md) |  | [optional] 
**WindowsOptions** | Pointer to [**IoK8sApiCoreV1WindowsSecurityContextOptions**](IoK8sApiCoreV1WindowsSecurityContextOptions.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1SecurityContext

`func NewIoK8sApiCoreV1SecurityContext() *IoK8sApiCoreV1SecurityContext`

NewIoK8sApiCoreV1SecurityContext instantiates a new IoK8sApiCoreV1SecurityContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1SecurityContextWithDefaults

`func NewIoK8sApiCoreV1SecurityContextWithDefaults() *IoK8sApiCoreV1SecurityContext`

NewIoK8sApiCoreV1SecurityContextWithDefaults instantiates a new IoK8sApiCoreV1SecurityContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPrivilegeEscalation

`func (o *IoK8sApiCoreV1SecurityContext) GetAllowPrivilegeEscalation() bool`

GetAllowPrivilegeEscalation returns the AllowPrivilegeEscalation field if non-nil, zero value otherwise.

### GetAllowPrivilegeEscalationOk

`func (o *IoK8sApiCoreV1SecurityContext) GetAllowPrivilegeEscalationOk() (*bool, bool)`

GetAllowPrivilegeEscalationOk returns a tuple with the AllowPrivilegeEscalation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrivilegeEscalation

`func (o *IoK8sApiCoreV1SecurityContext) SetAllowPrivilegeEscalation(v bool)`

SetAllowPrivilegeEscalation sets AllowPrivilegeEscalation field to given value.

### HasAllowPrivilegeEscalation

`func (o *IoK8sApiCoreV1SecurityContext) HasAllowPrivilegeEscalation() bool`

HasAllowPrivilegeEscalation returns a boolean if a field has been set.

### GetCapabilities

`func (o *IoK8sApiCoreV1SecurityContext) GetCapabilities() IoK8sApiCoreV1Capabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *IoK8sApiCoreV1SecurityContext) GetCapabilitiesOk() (*IoK8sApiCoreV1Capabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *IoK8sApiCoreV1SecurityContext) SetCapabilities(v IoK8sApiCoreV1Capabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *IoK8sApiCoreV1SecurityContext) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetPrivileged

`func (o *IoK8sApiCoreV1SecurityContext) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *IoK8sApiCoreV1SecurityContext) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *IoK8sApiCoreV1SecurityContext) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *IoK8sApiCoreV1SecurityContext) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetProcMount

`func (o *IoK8sApiCoreV1SecurityContext) GetProcMount() string`

GetProcMount returns the ProcMount field if non-nil, zero value otherwise.

### GetProcMountOk

`func (o *IoK8sApiCoreV1SecurityContext) GetProcMountOk() (*string, bool)`

GetProcMountOk returns a tuple with the ProcMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcMount

`func (o *IoK8sApiCoreV1SecurityContext) SetProcMount(v string)`

SetProcMount sets ProcMount field to given value.

### HasProcMount

`func (o *IoK8sApiCoreV1SecurityContext) HasProcMount() bool`

HasProcMount returns a boolean if a field has been set.

### GetReadOnlyRootFilesystem

`func (o *IoK8sApiCoreV1SecurityContext) GetReadOnlyRootFilesystem() bool`

GetReadOnlyRootFilesystem returns the ReadOnlyRootFilesystem field if non-nil, zero value otherwise.

### GetReadOnlyRootFilesystemOk

`func (o *IoK8sApiCoreV1SecurityContext) GetReadOnlyRootFilesystemOk() (*bool, bool)`

GetReadOnlyRootFilesystemOk returns a tuple with the ReadOnlyRootFilesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyRootFilesystem

`func (o *IoK8sApiCoreV1SecurityContext) SetReadOnlyRootFilesystem(v bool)`

SetReadOnlyRootFilesystem sets ReadOnlyRootFilesystem field to given value.

### HasReadOnlyRootFilesystem

`func (o *IoK8sApiCoreV1SecurityContext) HasReadOnlyRootFilesystem() bool`

HasReadOnlyRootFilesystem returns a boolean if a field has been set.

### GetRunAsGroup

`func (o *IoK8sApiCoreV1SecurityContext) GetRunAsGroup() int64`

GetRunAsGroup returns the RunAsGroup field if non-nil, zero value otherwise.

### GetRunAsGroupOk

`func (o *IoK8sApiCoreV1SecurityContext) GetRunAsGroupOk() (*int64, bool)`

GetRunAsGroupOk returns a tuple with the RunAsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsGroup

`func (o *IoK8sApiCoreV1SecurityContext) SetRunAsGroup(v int64)`

SetRunAsGroup sets RunAsGroup field to given value.

### HasRunAsGroup

`func (o *IoK8sApiCoreV1SecurityContext) HasRunAsGroup() bool`

HasRunAsGroup returns a boolean if a field has been set.

### GetRunAsNonRoot

`func (o *IoK8sApiCoreV1SecurityContext) GetRunAsNonRoot() bool`

GetRunAsNonRoot returns the RunAsNonRoot field if non-nil, zero value otherwise.

### GetRunAsNonRootOk

`func (o *IoK8sApiCoreV1SecurityContext) GetRunAsNonRootOk() (*bool, bool)`

GetRunAsNonRootOk returns a tuple with the RunAsNonRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsNonRoot

`func (o *IoK8sApiCoreV1SecurityContext) SetRunAsNonRoot(v bool)`

SetRunAsNonRoot sets RunAsNonRoot field to given value.

### HasRunAsNonRoot

`func (o *IoK8sApiCoreV1SecurityContext) HasRunAsNonRoot() bool`

HasRunAsNonRoot returns a boolean if a field has been set.

### GetRunAsUser

`func (o *IoK8sApiCoreV1SecurityContext) GetRunAsUser() int64`

GetRunAsUser returns the RunAsUser field if non-nil, zero value otherwise.

### GetRunAsUserOk

`func (o *IoK8sApiCoreV1SecurityContext) GetRunAsUserOk() (*int64, bool)`

GetRunAsUserOk returns a tuple with the RunAsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUser

`func (o *IoK8sApiCoreV1SecurityContext) SetRunAsUser(v int64)`

SetRunAsUser sets RunAsUser field to given value.

### HasRunAsUser

`func (o *IoK8sApiCoreV1SecurityContext) HasRunAsUser() bool`

HasRunAsUser returns a boolean if a field has been set.

### GetSeLinuxOptions

`func (o *IoK8sApiCoreV1SecurityContext) GetSeLinuxOptions() IoK8sApiCoreV1SELinuxOptions`

GetSeLinuxOptions returns the SeLinuxOptions field if non-nil, zero value otherwise.

### GetSeLinuxOptionsOk

`func (o *IoK8sApiCoreV1SecurityContext) GetSeLinuxOptionsOk() (*IoK8sApiCoreV1SELinuxOptions, bool)`

GetSeLinuxOptionsOk returns a tuple with the SeLinuxOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeLinuxOptions

`func (o *IoK8sApiCoreV1SecurityContext) SetSeLinuxOptions(v IoK8sApiCoreV1SELinuxOptions)`

SetSeLinuxOptions sets SeLinuxOptions field to given value.

### HasSeLinuxOptions

`func (o *IoK8sApiCoreV1SecurityContext) HasSeLinuxOptions() bool`

HasSeLinuxOptions returns a boolean if a field has been set.

### GetSeccompProfile

`func (o *IoK8sApiCoreV1SecurityContext) GetSeccompProfile() IoK8sApiCoreV1SeccompProfile`

GetSeccompProfile returns the SeccompProfile field if non-nil, zero value otherwise.

### GetSeccompProfileOk

`func (o *IoK8sApiCoreV1SecurityContext) GetSeccompProfileOk() (*IoK8sApiCoreV1SeccompProfile, bool)`

GetSeccompProfileOk returns a tuple with the SeccompProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeccompProfile

`func (o *IoK8sApiCoreV1SecurityContext) SetSeccompProfile(v IoK8sApiCoreV1SeccompProfile)`

SetSeccompProfile sets SeccompProfile field to given value.

### HasSeccompProfile

`func (o *IoK8sApiCoreV1SecurityContext) HasSeccompProfile() bool`

HasSeccompProfile returns a boolean if a field has been set.

### GetWindowsOptions

`func (o *IoK8sApiCoreV1SecurityContext) GetWindowsOptions() IoK8sApiCoreV1WindowsSecurityContextOptions`

GetWindowsOptions returns the WindowsOptions field if non-nil, zero value otherwise.

### GetWindowsOptionsOk

`func (o *IoK8sApiCoreV1SecurityContext) GetWindowsOptionsOk() (*IoK8sApiCoreV1WindowsSecurityContextOptions, bool)`

GetWindowsOptionsOk returns a tuple with the WindowsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsOptions

`func (o *IoK8sApiCoreV1SecurityContext) SetWindowsOptions(v IoK8sApiCoreV1WindowsSecurityContextOptions)`

SetWindowsOptions sets WindowsOptions field to given value.

### HasWindowsOptions

`func (o *IoK8sApiCoreV1SecurityContext) HasWindowsOptions() bool`

HasWindowsOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


