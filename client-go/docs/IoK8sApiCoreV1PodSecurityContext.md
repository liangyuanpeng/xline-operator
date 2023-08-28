# IoK8sApiCoreV1PodSecurityContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsGroup** | Pointer to **int64** | A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR&#39;d with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**FsGroupChangePolicy** | Pointer to **string** | fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are \&quot;OnRootMismatch\&quot; and \&quot;Always\&quot;. If not specified, \&quot;Always\&quot; is used. Note that this field cannot be set when spec.os.name is windows.  Possible enum values:  - &#x60;\&quot;Always\&quot;&#x60; indicates that volume&#39;s ownership and permissions should always be changed whenever volume is mounted inside a Pod. This the default behavior.  - &#x60;\&quot;OnRootMismatch\&quot;&#x60; indicates that volume&#39;s ownership and permissions will be changed only when permission and ownership of root directory does not match with expected permissions on the volume. This can help shorten the time it takes to change ownership and permissions of a volume. | [optional] 
**RunAsGroup** | Pointer to **int64** | The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**RunAsNonRoot** | Pointer to **bool** | Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. | [optional] 
**RunAsUser** | Pointer to **int64** | The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**SeLinuxOptions** | Pointer to [**IoK8sApiCoreV1SELinuxOptions**](IoK8sApiCoreV1SELinuxOptions.md) |  | [optional] 
**SeccompProfile** | Pointer to [**IoK8sApiCoreV1SeccompProfile**](IoK8sApiCoreV1SeccompProfile.md) |  | [optional] 
**SupplementalGroups** | Pointer to **[]int64** | A list of groups applied to the first process run in each container, in addition to the container&#39;s primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**Sysctls** | Pointer to [**[]IoK8sApiCoreV1Sysctl**](IoK8sApiCoreV1Sysctl.md) | Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**WindowsOptions** | Pointer to [**IoK8sApiCoreV1WindowsSecurityContextOptions**](IoK8sApiCoreV1WindowsSecurityContextOptions.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1PodSecurityContext

`func NewIoK8sApiCoreV1PodSecurityContext() *IoK8sApiCoreV1PodSecurityContext`

NewIoK8sApiCoreV1PodSecurityContext instantiates a new IoK8sApiCoreV1PodSecurityContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PodSecurityContextWithDefaults

`func NewIoK8sApiCoreV1PodSecurityContextWithDefaults() *IoK8sApiCoreV1PodSecurityContext`

NewIoK8sApiCoreV1PodSecurityContextWithDefaults instantiates a new IoK8sApiCoreV1PodSecurityContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsGroup

`func (o *IoK8sApiCoreV1PodSecurityContext) GetFsGroup() int64`

GetFsGroup returns the FsGroup field if non-nil, zero value otherwise.

### GetFsGroupOk

`func (o *IoK8sApiCoreV1PodSecurityContext) GetFsGroupOk() (*int64, bool)`

GetFsGroupOk returns a tuple with the FsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsGroup

`func (o *IoK8sApiCoreV1PodSecurityContext) SetFsGroup(v int64)`

SetFsGroup sets FsGroup field to given value.

### HasFsGroup

`func (o *IoK8sApiCoreV1PodSecurityContext) HasFsGroup() bool`

HasFsGroup returns a boolean if a field has been set.

### GetFsGroupChangePolicy

`func (o *IoK8sApiCoreV1PodSecurityContext) GetFsGroupChangePolicy() string`

GetFsGroupChangePolicy returns the FsGroupChangePolicy field if non-nil, zero value otherwise.

### GetFsGroupChangePolicyOk

`func (o *IoK8sApiCoreV1PodSecurityContext) GetFsGroupChangePolicyOk() (*string, bool)`

GetFsGroupChangePolicyOk returns a tuple with the FsGroupChangePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsGroupChangePolicy

`func (o *IoK8sApiCoreV1PodSecurityContext) SetFsGroupChangePolicy(v string)`

SetFsGroupChangePolicy sets FsGroupChangePolicy field to given value.

### HasFsGroupChangePolicy

`func (o *IoK8sApiCoreV1PodSecurityContext) HasFsGroupChangePolicy() bool`

HasFsGroupChangePolicy returns a boolean if a field has been set.

### GetRunAsGroup

`func (o *IoK8sApiCoreV1PodSecurityContext) GetRunAsGroup() int64`

GetRunAsGroup returns the RunAsGroup field if non-nil, zero value otherwise.

### GetRunAsGroupOk

`func (o *IoK8sApiCoreV1PodSecurityContext) GetRunAsGroupOk() (*int64, bool)`

GetRunAsGroupOk returns a tuple with the RunAsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsGroup

`func (o *IoK8sApiCoreV1PodSecurityContext) SetRunAsGroup(v int64)`

SetRunAsGroup sets RunAsGroup field to given value.

### HasRunAsGroup

`func (o *IoK8sApiCoreV1PodSecurityContext) HasRunAsGroup() bool`

HasRunAsGroup returns a boolean if a field has been set.

### GetRunAsNonRoot

`func (o *IoK8sApiCoreV1PodSecurityContext) GetRunAsNonRoot() bool`

GetRunAsNonRoot returns the RunAsNonRoot field if non-nil, zero value otherwise.

### GetRunAsNonRootOk

`func (o *IoK8sApiCoreV1PodSecurityContext) GetRunAsNonRootOk() (*bool, bool)`

GetRunAsNonRootOk returns a tuple with the RunAsNonRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsNonRoot

`func (o *IoK8sApiCoreV1PodSecurityContext) SetRunAsNonRoot(v bool)`

SetRunAsNonRoot sets RunAsNonRoot field to given value.

### HasRunAsNonRoot

`func (o *IoK8sApiCoreV1PodSecurityContext) HasRunAsNonRoot() bool`

HasRunAsNonRoot returns a boolean if a field has been set.

### GetRunAsUser

`func (o *IoK8sApiCoreV1PodSecurityContext) GetRunAsUser() int64`

GetRunAsUser returns the RunAsUser field if non-nil, zero value otherwise.

### GetRunAsUserOk

`func (o *IoK8sApiCoreV1PodSecurityContext) GetRunAsUserOk() (*int64, bool)`

GetRunAsUserOk returns a tuple with the RunAsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUser

`func (o *IoK8sApiCoreV1PodSecurityContext) SetRunAsUser(v int64)`

SetRunAsUser sets RunAsUser field to given value.

### HasRunAsUser

`func (o *IoK8sApiCoreV1PodSecurityContext) HasRunAsUser() bool`

HasRunAsUser returns a boolean if a field has been set.

### GetSeLinuxOptions

`func (o *IoK8sApiCoreV1PodSecurityContext) GetSeLinuxOptions() IoK8sApiCoreV1SELinuxOptions`

GetSeLinuxOptions returns the SeLinuxOptions field if non-nil, zero value otherwise.

### GetSeLinuxOptionsOk

`func (o *IoK8sApiCoreV1PodSecurityContext) GetSeLinuxOptionsOk() (*IoK8sApiCoreV1SELinuxOptions, bool)`

GetSeLinuxOptionsOk returns a tuple with the SeLinuxOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeLinuxOptions

`func (o *IoK8sApiCoreV1PodSecurityContext) SetSeLinuxOptions(v IoK8sApiCoreV1SELinuxOptions)`

SetSeLinuxOptions sets SeLinuxOptions field to given value.

### HasSeLinuxOptions

`func (o *IoK8sApiCoreV1PodSecurityContext) HasSeLinuxOptions() bool`

HasSeLinuxOptions returns a boolean if a field has been set.

### GetSeccompProfile

`func (o *IoK8sApiCoreV1PodSecurityContext) GetSeccompProfile() IoK8sApiCoreV1SeccompProfile`

GetSeccompProfile returns the SeccompProfile field if non-nil, zero value otherwise.

### GetSeccompProfileOk

`func (o *IoK8sApiCoreV1PodSecurityContext) GetSeccompProfileOk() (*IoK8sApiCoreV1SeccompProfile, bool)`

GetSeccompProfileOk returns a tuple with the SeccompProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeccompProfile

`func (o *IoK8sApiCoreV1PodSecurityContext) SetSeccompProfile(v IoK8sApiCoreV1SeccompProfile)`

SetSeccompProfile sets SeccompProfile field to given value.

### HasSeccompProfile

`func (o *IoK8sApiCoreV1PodSecurityContext) HasSeccompProfile() bool`

HasSeccompProfile returns a boolean if a field has been set.

### GetSupplementalGroups

`func (o *IoK8sApiCoreV1PodSecurityContext) GetSupplementalGroups() []int64`

GetSupplementalGroups returns the SupplementalGroups field if non-nil, zero value otherwise.

### GetSupplementalGroupsOk

`func (o *IoK8sApiCoreV1PodSecurityContext) GetSupplementalGroupsOk() (*[]int64, bool)`

GetSupplementalGroupsOk returns a tuple with the SupplementalGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalGroups

`func (o *IoK8sApiCoreV1PodSecurityContext) SetSupplementalGroups(v []int64)`

SetSupplementalGroups sets SupplementalGroups field to given value.

### HasSupplementalGroups

`func (o *IoK8sApiCoreV1PodSecurityContext) HasSupplementalGroups() bool`

HasSupplementalGroups returns a boolean if a field has been set.

### GetSysctls

`func (o *IoK8sApiCoreV1PodSecurityContext) GetSysctls() []IoK8sApiCoreV1Sysctl`

GetSysctls returns the Sysctls field if non-nil, zero value otherwise.

### GetSysctlsOk

`func (o *IoK8sApiCoreV1PodSecurityContext) GetSysctlsOk() (*[]IoK8sApiCoreV1Sysctl, bool)`

GetSysctlsOk returns a tuple with the Sysctls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysctls

`func (o *IoK8sApiCoreV1PodSecurityContext) SetSysctls(v []IoK8sApiCoreV1Sysctl)`

SetSysctls sets Sysctls field to given value.

### HasSysctls

`func (o *IoK8sApiCoreV1PodSecurityContext) HasSysctls() bool`

HasSysctls returns a boolean if a field has been set.

### GetWindowsOptions

`func (o *IoK8sApiCoreV1PodSecurityContext) GetWindowsOptions() IoK8sApiCoreV1WindowsSecurityContextOptions`

GetWindowsOptions returns the WindowsOptions field if non-nil, zero value otherwise.

### GetWindowsOptionsOk

`func (o *IoK8sApiCoreV1PodSecurityContext) GetWindowsOptionsOk() (*IoK8sApiCoreV1WindowsSecurityContextOptions, bool)`

GetWindowsOptionsOk returns a tuple with the WindowsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsOptions

`func (o *IoK8sApiCoreV1PodSecurityContext) SetWindowsOptions(v IoK8sApiCoreV1WindowsSecurityContextOptions)`

SetWindowsOptions sets WindowsOptions field to given value.

### HasWindowsOptions

`func (o *IoK8sApiCoreV1PodSecurityContext) HasWindowsOptions() bool`

HasWindowsOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


