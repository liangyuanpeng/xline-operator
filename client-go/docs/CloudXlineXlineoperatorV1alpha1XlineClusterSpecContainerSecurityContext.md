# CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPrivilegeEscalation** | Pointer to **bool** | AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows. | [optional] 
**Capabilities** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextCapabilities**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextCapabilities.md) |  | [optional] 
**Privileged** | Pointer to **bool** | Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**ProcMount** | Pointer to **string** | procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**ReadOnlyRootFilesystem** | Pointer to **bool** | Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**RunAsGroup** | Pointer to **int64** | The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**RunAsNonRoot** | Pointer to **bool** | Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. | [optional] 
**RunAsUser** | Pointer to **int64** | The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**SeLinuxOptions** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeLinuxOptions**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeLinuxOptions.md) |  | [optional] 
**SeccompProfile** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile.md) |  | [optional] 
**WindowsOptions** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextWindowsOptions**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextWindowsOptions.md) |  | [optional] 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPrivilegeEscalation

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetAllowPrivilegeEscalation() bool`

GetAllowPrivilegeEscalation returns the AllowPrivilegeEscalation field if non-nil, zero value otherwise.

### GetAllowPrivilegeEscalationOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetAllowPrivilegeEscalationOk() (*bool, bool)`

GetAllowPrivilegeEscalationOk returns a tuple with the AllowPrivilegeEscalation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrivilegeEscalation

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetAllowPrivilegeEscalation(v bool)`

SetAllowPrivilegeEscalation sets AllowPrivilegeEscalation field to given value.

### HasAllowPrivilegeEscalation

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasAllowPrivilegeEscalation() bool`

HasAllowPrivilegeEscalation returns a boolean if a field has been set.

### GetCapabilities

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetCapabilities() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetCapabilitiesOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetCapabilities(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetPrivileged

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetProcMount

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetProcMount() string`

GetProcMount returns the ProcMount field if non-nil, zero value otherwise.

### GetProcMountOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetProcMountOk() (*string, bool)`

GetProcMountOk returns a tuple with the ProcMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcMount

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetProcMount(v string)`

SetProcMount sets ProcMount field to given value.

### HasProcMount

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasProcMount() bool`

HasProcMount returns a boolean if a field has been set.

### GetReadOnlyRootFilesystem

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetReadOnlyRootFilesystem() bool`

GetReadOnlyRootFilesystem returns the ReadOnlyRootFilesystem field if non-nil, zero value otherwise.

### GetReadOnlyRootFilesystemOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetReadOnlyRootFilesystemOk() (*bool, bool)`

GetReadOnlyRootFilesystemOk returns a tuple with the ReadOnlyRootFilesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyRootFilesystem

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetReadOnlyRootFilesystem(v bool)`

SetReadOnlyRootFilesystem sets ReadOnlyRootFilesystem field to given value.

### HasReadOnlyRootFilesystem

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasReadOnlyRootFilesystem() bool`

HasReadOnlyRootFilesystem returns a boolean if a field has been set.

### GetRunAsGroup

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetRunAsGroup() int64`

GetRunAsGroup returns the RunAsGroup field if non-nil, zero value otherwise.

### GetRunAsGroupOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetRunAsGroupOk() (*int64, bool)`

GetRunAsGroupOk returns a tuple with the RunAsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsGroup

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetRunAsGroup(v int64)`

SetRunAsGroup sets RunAsGroup field to given value.

### HasRunAsGroup

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasRunAsGroup() bool`

HasRunAsGroup returns a boolean if a field has been set.

### GetRunAsNonRoot

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetRunAsNonRoot() bool`

GetRunAsNonRoot returns the RunAsNonRoot field if non-nil, zero value otherwise.

### GetRunAsNonRootOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetRunAsNonRootOk() (*bool, bool)`

GetRunAsNonRootOk returns a tuple with the RunAsNonRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsNonRoot

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetRunAsNonRoot(v bool)`

SetRunAsNonRoot sets RunAsNonRoot field to given value.

### HasRunAsNonRoot

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasRunAsNonRoot() bool`

HasRunAsNonRoot returns a boolean if a field has been set.

### GetRunAsUser

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetRunAsUser() int64`

GetRunAsUser returns the RunAsUser field if non-nil, zero value otherwise.

### GetRunAsUserOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetRunAsUserOk() (*int64, bool)`

GetRunAsUserOk returns a tuple with the RunAsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUser

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetRunAsUser(v int64)`

SetRunAsUser sets RunAsUser field to given value.

### HasRunAsUser

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasRunAsUser() bool`

HasRunAsUser returns a boolean if a field has been set.

### GetSeLinuxOptions

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetSeLinuxOptions() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeLinuxOptions`

GetSeLinuxOptions returns the SeLinuxOptions field if non-nil, zero value otherwise.

### GetSeLinuxOptionsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetSeLinuxOptionsOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeLinuxOptions, bool)`

GetSeLinuxOptionsOk returns a tuple with the SeLinuxOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeLinuxOptions

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetSeLinuxOptions(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeLinuxOptions)`

SetSeLinuxOptions sets SeLinuxOptions field to given value.

### HasSeLinuxOptions

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasSeLinuxOptions() bool`

HasSeLinuxOptions returns a boolean if a field has been set.

### GetSeccompProfile

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetSeccompProfile() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile`

GetSeccompProfile returns the SeccompProfile field if non-nil, zero value otherwise.

### GetSeccompProfileOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetSeccompProfileOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile, bool)`

GetSeccompProfileOk returns a tuple with the SeccompProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeccompProfile

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetSeccompProfile(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile)`

SetSeccompProfile sets SeccompProfile field to given value.

### HasSeccompProfile

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasSeccompProfile() bool`

HasSeccompProfile returns a boolean if a field has been set.

### GetWindowsOptions

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetWindowsOptions() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextWindowsOptions`

GetWindowsOptions returns the WindowsOptions field if non-nil, zero value otherwise.

### GetWindowsOptionsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) GetWindowsOptionsOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextWindowsOptions, bool)`

GetWindowsOptionsOk returns a tuple with the WindowsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsOptions

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) SetWindowsOptions(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextWindowsOptions)`

SetWindowsOptions sets WindowsOptions field to given value.

### HasWindowsOptions

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContext) HasWindowsOptions() bool`

HasWindowsOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


