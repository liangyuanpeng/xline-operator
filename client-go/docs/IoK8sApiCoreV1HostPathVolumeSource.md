# IoK8sApiCoreV1HostPathVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | 
**Type** | Pointer to **string** | type for HostPath Volume Defaults to \&quot;\&quot; More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath  Possible enum values:  - &#x60;\&quot;\&quot;&#x60; For backwards compatible, leave it empty if unset  - &#x60;\&quot;BlockDevice\&quot;&#x60; A block device must exist at the given path  - &#x60;\&quot;CharDevice\&quot;&#x60; A character device must exist at the given path  - &#x60;\&quot;Directory\&quot;&#x60; A directory must exist at the given path  - &#x60;\&quot;DirectoryOrCreate\&quot;&#x60; If nothing exists at the given path, an empty directory will be created there as needed with file mode 0755, having the same group and ownership with Kubelet.  - &#x60;\&quot;File\&quot;&#x60; A file must exist at the given path  - &#x60;\&quot;FileOrCreate\&quot;&#x60; If nothing exists at the given path, an empty file will be created there as needed with file mode 0644, having the same group and ownership with Kubelet.  - &#x60;\&quot;Socket\&quot;&#x60; A UNIX socket must exist at the given path | [optional] 

## Methods

### NewIoK8sApiCoreV1HostPathVolumeSource

`func NewIoK8sApiCoreV1HostPathVolumeSource(path string, ) *IoK8sApiCoreV1HostPathVolumeSource`

NewIoK8sApiCoreV1HostPathVolumeSource instantiates a new IoK8sApiCoreV1HostPathVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1HostPathVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1HostPathVolumeSourceWithDefaults() *IoK8sApiCoreV1HostPathVolumeSource`

NewIoK8sApiCoreV1HostPathVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1HostPathVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *IoK8sApiCoreV1HostPathVolumeSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoK8sApiCoreV1HostPathVolumeSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoK8sApiCoreV1HostPathVolumeSource) SetPath(v string)`

SetPath sets Path field to given value.


### GetType

`func (o *IoK8sApiCoreV1HostPathVolumeSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiCoreV1HostPathVolumeSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiCoreV1HostPathVolumeSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiCoreV1HostPathVolumeSource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


