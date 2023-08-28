# IoK8sApiCoreV1Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsElasticBlockStore** | Pointer to [**IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource**](IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource.md) |  | [optional] 
**AzureDisk** | Pointer to [**IoK8sApiCoreV1AzureDiskVolumeSource**](IoK8sApiCoreV1AzureDiskVolumeSource.md) |  | [optional] 
**AzureFile** | Pointer to [**IoK8sApiCoreV1AzureFileVolumeSource**](IoK8sApiCoreV1AzureFileVolumeSource.md) |  | [optional] 
**Cephfs** | Pointer to [**IoK8sApiCoreV1CephFSVolumeSource**](IoK8sApiCoreV1CephFSVolumeSource.md) |  | [optional] 
**Cinder** | Pointer to [**IoK8sApiCoreV1CinderVolumeSource**](IoK8sApiCoreV1CinderVolumeSource.md) |  | [optional] 
**ConfigMap** | Pointer to [**IoK8sApiCoreV1ConfigMapVolumeSource**](IoK8sApiCoreV1ConfigMapVolumeSource.md) |  | [optional] 
**Csi** | Pointer to [**IoK8sApiCoreV1CSIVolumeSource**](IoK8sApiCoreV1CSIVolumeSource.md) |  | [optional] 
**DownwardAPI** | Pointer to [**IoK8sApiCoreV1DownwardAPIVolumeSource**](IoK8sApiCoreV1DownwardAPIVolumeSource.md) |  | [optional] 
**EmptyDir** | Pointer to [**IoK8sApiCoreV1EmptyDirVolumeSource**](IoK8sApiCoreV1EmptyDirVolumeSource.md) |  | [optional] 
**Ephemeral** | Pointer to [**IoK8sApiCoreV1EphemeralVolumeSource**](IoK8sApiCoreV1EphemeralVolumeSource.md) |  | [optional] 
**Fc** | Pointer to [**IoK8sApiCoreV1FCVolumeSource**](IoK8sApiCoreV1FCVolumeSource.md) |  | [optional] 
**FlexVolume** | Pointer to [**IoK8sApiCoreV1FlexVolumeSource**](IoK8sApiCoreV1FlexVolumeSource.md) |  | [optional] 
**Flocker** | Pointer to [**IoK8sApiCoreV1FlockerVolumeSource**](IoK8sApiCoreV1FlockerVolumeSource.md) |  | [optional] 
**GcePersistentDisk** | Pointer to [**IoK8sApiCoreV1GCEPersistentDiskVolumeSource**](IoK8sApiCoreV1GCEPersistentDiskVolumeSource.md) |  | [optional] 
**GitRepo** | Pointer to [**IoK8sApiCoreV1GitRepoVolumeSource**](IoK8sApiCoreV1GitRepoVolumeSource.md) |  | [optional] 
**Glusterfs** | Pointer to [**IoK8sApiCoreV1GlusterfsVolumeSource**](IoK8sApiCoreV1GlusterfsVolumeSource.md) |  | [optional] 
**HostPath** | Pointer to [**IoK8sApiCoreV1HostPathVolumeSource**](IoK8sApiCoreV1HostPathVolumeSource.md) |  | [optional] 
**Iscsi** | Pointer to [**IoK8sApiCoreV1ISCSIVolumeSource**](IoK8sApiCoreV1ISCSIVolumeSource.md) |  | [optional] 
**Name** | **string** | name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | 
**Nfs** | Pointer to [**IoK8sApiCoreV1NFSVolumeSource**](IoK8sApiCoreV1NFSVolumeSource.md) |  | [optional] 
**PersistentVolumeClaim** | Pointer to [**IoK8sApiCoreV1PersistentVolumeClaimVolumeSource**](IoK8sApiCoreV1PersistentVolumeClaimVolumeSource.md) |  | [optional] 
**PhotonPersistentDisk** | Pointer to [**IoK8sApiCoreV1PhotonPersistentDiskVolumeSource**](IoK8sApiCoreV1PhotonPersistentDiskVolumeSource.md) |  | [optional] 
**PortworxVolume** | Pointer to [**IoK8sApiCoreV1PortworxVolumeSource**](IoK8sApiCoreV1PortworxVolumeSource.md) |  | [optional] 
**Projected** | Pointer to [**IoK8sApiCoreV1ProjectedVolumeSource**](IoK8sApiCoreV1ProjectedVolumeSource.md) |  | [optional] 
**Quobyte** | Pointer to [**IoK8sApiCoreV1QuobyteVolumeSource**](IoK8sApiCoreV1QuobyteVolumeSource.md) |  | [optional] 
**Rbd** | Pointer to [**IoK8sApiCoreV1RBDVolumeSource**](IoK8sApiCoreV1RBDVolumeSource.md) |  | [optional] 
**ScaleIO** | Pointer to [**IoK8sApiCoreV1ScaleIOVolumeSource**](IoK8sApiCoreV1ScaleIOVolumeSource.md) |  | [optional] 
**Secret** | Pointer to [**IoK8sApiCoreV1SecretVolumeSource**](IoK8sApiCoreV1SecretVolumeSource.md) |  | [optional] 
**Storageos** | Pointer to [**IoK8sApiCoreV1StorageOSVolumeSource**](IoK8sApiCoreV1StorageOSVolumeSource.md) |  | [optional] 
**VsphereVolume** | Pointer to [**IoK8sApiCoreV1VsphereVirtualDiskVolumeSource**](IoK8sApiCoreV1VsphereVirtualDiskVolumeSource.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1Volume

`func NewIoK8sApiCoreV1Volume(name string, ) *IoK8sApiCoreV1Volume`

NewIoK8sApiCoreV1Volume instantiates a new IoK8sApiCoreV1Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1VolumeWithDefaults

`func NewIoK8sApiCoreV1VolumeWithDefaults() *IoK8sApiCoreV1Volume`

NewIoK8sApiCoreV1VolumeWithDefaults instantiates a new IoK8sApiCoreV1Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsElasticBlockStore

`func (o *IoK8sApiCoreV1Volume) GetAwsElasticBlockStore() IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource`

GetAwsElasticBlockStore returns the AwsElasticBlockStore field if non-nil, zero value otherwise.

### GetAwsElasticBlockStoreOk

`func (o *IoK8sApiCoreV1Volume) GetAwsElasticBlockStoreOk() (*IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource, bool)`

GetAwsElasticBlockStoreOk returns a tuple with the AwsElasticBlockStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsElasticBlockStore

`func (o *IoK8sApiCoreV1Volume) SetAwsElasticBlockStore(v IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource)`

SetAwsElasticBlockStore sets AwsElasticBlockStore field to given value.

### HasAwsElasticBlockStore

`func (o *IoK8sApiCoreV1Volume) HasAwsElasticBlockStore() bool`

HasAwsElasticBlockStore returns a boolean if a field has been set.

### GetAzureDisk

`func (o *IoK8sApiCoreV1Volume) GetAzureDisk() IoK8sApiCoreV1AzureDiskVolumeSource`

GetAzureDisk returns the AzureDisk field if non-nil, zero value otherwise.

### GetAzureDiskOk

`func (o *IoK8sApiCoreV1Volume) GetAzureDiskOk() (*IoK8sApiCoreV1AzureDiskVolumeSource, bool)`

GetAzureDiskOk returns a tuple with the AzureDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDisk

`func (o *IoK8sApiCoreV1Volume) SetAzureDisk(v IoK8sApiCoreV1AzureDiskVolumeSource)`

SetAzureDisk sets AzureDisk field to given value.

### HasAzureDisk

`func (o *IoK8sApiCoreV1Volume) HasAzureDisk() bool`

HasAzureDisk returns a boolean if a field has been set.

### GetAzureFile

`func (o *IoK8sApiCoreV1Volume) GetAzureFile() IoK8sApiCoreV1AzureFileVolumeSource`

GetAzureFile returns the AzureFile field if non-nil, zero value otherwise.

### GetAzureFileOk

`func (o *IoK8sApiCoreV1Volume) GetAzureFileOk() (*IoK8sApiCoreV1AzureFileVolumeSource, bool)`

GetAzureFileOk returns a tuple with the AzureFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureFile

`func (o *IoK8sApiCoreV1Volume) SetAzureFile(v IoK8sApiCoreV1AzureFileVolumeSource)`

SetAzureFile sets AzureFile field to given value.

### HasAzureFile

`func (o *IoK8sApiCoreV1Volume) HasAzureFile() bool`

HasAzureFile returns a boolean if a field has been set.

### GetCephfs

`func (o *IoK8sApiCoreV1Volume) GetCephfs() IoK8sApiCoreV1CephFSVolumeSource`

GetCephfs returns the Cephfs field if non-nil, zero value otherwise.

### GetCephfsOk

`func (o *IoK8sApiCoreV1Volume) GetCephfsOk() (*IoK8sApiCoreV1CephFSVolumeSource, bool)`

GetCephfsOk returns a tuple with the Cephfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephfs

`func (o *IoK8sApiCoreV1Volume) SetCephfs(v IoK8sApiCoreV1CephFSVolumeSource)`

SetCephfs sets Cephfs field to given value.

### HasCephfs

`func (o *IoK8sApiCoreV1Volume) HasCephfs() bool`

HasCephfs returns a boolean if a field has been set.

### GetCinder

`func (o *IoK8sApiCoreV1Volume) GetCinder() IoK8sApiCoreV1CinderVolumeSource`

GetCinder returns the Cinder field if non-nil, zero value otherwise.

### GetCinderOk

`func (o *IoK8sApiCoreV1Volume) GetCinderOk() (*IoK8sApiCoreV1CinderVolumeSource, bool)`

GetCinderOk returns a tuple with the Cinder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCinder

`func (o *IoK8sApiCoreV1Volume) SetCinder(v IoK8sApiCoreV1CinderVolumeSource)`

SetCinder sets Cinder field to given value.

### HasCinder

`func (o *IoK8sApiCoreV1Volume) HasCinder() bool`

HasCinder returns a boolean if a field has been set.

### GetConfigMap

`func (o *IoK8sApiCoreV1Volume) GetConfigMap() IoK8sApiCoreV1ConfigMapVolumeSource`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *IoK8sApiCoreV1Volume) GetConfigMapOk() (*IoK8sApiCoreV1ConfigMapVolumeSource, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *IoK8sApiCoreV1Volume) SetConfigMap(v IoK8sApiCoreV1ConfigMapVolumeSource)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *IoK8sApiCoreV1Volume) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.

### GetCsi

`func (o *IoK8sApiCoreV1Volume) GetCsi() IoK8sApiCoreV1CSIVolumeSource`

GetCsi returns the Csi field if non-nil, zero value otherwise.

### GetCsiOk

`func (o *IoK8sApiCoreV1Volume) GetCsiOk() (*IoK8sApiCoreV1CSIVolumeSource, bool)`

GetCsiOk returns a tuple with the Csi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsi

`func (o *IoK8sApiCoreV1Volume) SetCsi(v IoK8sApiCoreV1CSIVolumeSource)`

SetCsi sets Csi field to given value.

### HasCsi

`func (o *IoK8sApiCoreV1Volume) HasCsi() bool`

HasCsi returns a boolean if a field has been set.

### GetDownwardAPI

`func (o *IoK8sApiCoreV1Volume) GetDownwardAPI() IoK8sApiCoreV1DownwardAPIVolumeSource`

GetDownwardAPI returns the DownwardAPI field if non-nil, zero value otherwise.

### GetDownwardAPIOk

`func (o *IoK8sApiCoreV1Volume) GetDownwardAPIOk() (*IoK8sApiCoreV1DownwardAPIVolumeSource, bool)`

GetDownwardAPIOk returns a tuple with the DownwardAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownwardAPI

`func (o *IoK8sApiCoreV1Volume) SetDownwardAPI(v IoK8sApiCoreV1DownwardAPIVolumeSource)`

SetDownwardAPI sets DownwardAPI field to given value.

### HasDownwardAPI

`func (o *IoK8sApiCoreV1Volume) HasDownwardAPI() bool`

HasDownwardAPI returns a boolean if a field has been set.

### GetEmptyDir

`func (o *IoK8sApiCoreV1Volume) GetEmptyDir() IoK8sApiCoreV1EmptyDirVolumeSource`

GetEmptyDir returns the EmptyDir field if non-nil, zero value otherwise.

### GetEmptyDirOk

`func (o *IoK8sApiCoreV1Volume) GetEmptyDirOk() (*IoK8sApiCoreV1EmptyDirVolumeSource, bool)`

GetEmptyDirOk returns a tuple with the EmptyDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyDir

`func (o *IoK8sApiCoreV1Volume) SetEmptyDir(v IoK8sApiCoreV1EmptyDirVolumeSource)`

SetEmptyDir sets EmptyDir field to given value.

### HasEmptyDir

`func (o *IoK8sApiCoreV1Volume) HasEmptyDir() bool`

HasEmptyDir returns a boolean if a field has been set.

### GetEphemeral

`func (o *IoK8sApiCoreV1Volume) GetEphemeral() IoK8sApiCoreV1EphemeralVolumeSource`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *IoK8sApiCoreV1Volume) GetEphemeralOk() (*IoK8sApiCoreV1EphemeralVolumeSource, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *IoK8sApiCoreV1Volume) SetEphemeral(v IoK8sApiCoreV1EphemeralVolumeSource)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *IoK8sApiCoreV1Volume) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetFc

`func (o *IoK8sApiCoreV1Volume) GetFc() IoK8sApiCoreV1FCVolumeSource`

GetFc returns the Fc field if non-nil, zero value otherwise.

### GetFcOk

`func (o *IoK8sApiCoreV1Volume) GetFcOk() (*IoK8sApiCoreV1FCVolumeSource, bool)`

GetFcOk returns a tuple with the Fc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFc

`func (o *IoK8sApiCoreV1Volume) SetFc(v IoK8sApiCoreV1FCVolumeSource)`

SetFc sets Fc field to given value.

### HasFc

`func (o *IoK8sApiCoreV1Volume) HasFc() bool`

HasFc returns a boolean if a field has been set.

### GetFlexVolume

`func (o *IoK8sApiCoreV1Volume) GetFlexVolume() IoK8sApiCoreV1FlexVolumeSource`

GetFlexVolume returns the FlexVolume field if non-nil, zero value otherwise.

### GetFlexVolumeOk

`func (o *IoK8sApiCoreV1Volume) GetFlexVolumeOk() (*IoK8sApiCoreV1FlexVolumeSource, bool)`

GetFlexVolumeOk returns a tuple with the FlexVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexVolume

`func (o *IoK8sApiCoreV1Volume) SetFlexVolume(v IoK8sApiCoreV1FlexVolumeSource)`

SetFlexVolume sets FlexVolume field to given value.

### HasFlexVolume

`func (o *IoK8sApiCoreV1Volume) HasFlexVolume() bool`

HasFlexVolume returns a boolean if a field has been set.

### GetFlocker

`func (o *IoK8sApiCoreV1Volume) GetFlocker() IoK8sApiCoreV1FlockerVolumeSource`

GetFlocker returns the Flocker field if non-nil, zero value otherwise.

### GetFlockerOk

`func (o *IoK8sApiCoreV1Volume) GetFlockerOk() (*IoK8sApiCoreV1FlockerVolumeSource, bool)`

GetFlockerOk returns a tuple with the Flocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlocker

`func (o *IoK8sApiCoreV1Volume) SetFlocker(v IoK8sApiCoreV1FlockerVolumeSource)`

SetFlocker sets Flocker field to given value.

### HasFlocker

`func (o *IoK8sApiCoreV1Volume) HasFlocker() bool`

HasFlocker returns a boolean if a field has been set.

### GetGcePersistentDisk

`func (o *IoK8sApiCoreV1Volume) GetGcePersistentDisk() IoK8sApiCoreV1GCEPersistentDiskVolumeSource`

GetGcePersistentDisk returns the GcePersistentDisk field if non-nil, zero value otherwise.

### GetGcePersistentDiskOk

`func (o *IoK8sApiCoreV1Volume) GetGcePersistentDiskOk() (*IoK8sApiCoreV1GCEPersistentDiskVolumeSource, bool)`

GetGcePersistentDiskOk returns a tuple with the GcePersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcePersistentDisk

`func (o *IoK8sApiCoreV1Volume) SetGcePersistentDisk(v IoK8sApiCoreV1GCEPersistentDiskVolumeSource)`

SetGcePersistentDisk sets GcePersistentDisk field to given value.

### HasGcePersistentDisk

`func (o *IoK8sApiCoreV1Volume) HasGcePersistentDisk() bool`

HasGcePersistentDisk returns a boolean if a field has been set.

### GetGitRepo

`func (o *IoK8sApiCoreV1Volume) GetGitRepo() IoK8sApiCoreV1GitRepoVolumeSource`

GetGitRepo returns the GitRepo field if non-nil, zero value otherwise.

### GetGitRepoOk

`func (o *IoK8sApiCoreV1Volume) GetGitRepoOk() (*IoK8sApiCoreV1GitRepoVolumeSource, bool)`

GetGitRepoOk returns a tuple with the GitRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepo

`func (o *IoK8sApiCoreV1Volume) SetGitRepo(v IoK8sApiCoreV1GitRepoVolumeSource)`

SetGitRepo sets GitRepo field to given value.

### HasGitRepo

`func (o *IoK8sApiCoreV1Volume) HasGitRepo() bool`

HasGitRepo returns a boolean if a field has been set.

### GetGlusterfs

`func (o *IoK8sApiCoreV1Volume) GetGlusterfs() IoK8sApiCoreV1GlusterfsVolumeSource`

GetGlusterfs returns the Glusterfs field if non-nil, zero value otherwise.

### GetGlusterfsOk

`func (o *IoK8sApiCoreV1Volume) GetGlusterfsOk() (*IoK8sApiCoreV1GlusterfsVolumeSource, bool)`

GetGlusterfsOk returns a tuple with the Glusterfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlusterfs

`func (o *IoK8sApiCoreV1Volume) SetGlusterfs(v IoK8sApiCoreV1GlusterfsVolumeSource)`

SetGlusterfs sets Glusterfs field to given value.

### HasGlusterfs

`func (o *IoK8sApiCoreV1Volume) HasGlusterfs() bool`

HasGlusterfs returns a boolean if a field has been set.

### GetHostPath

`func (o *IoK8sApiCoreV1Volume) GetHostPath() IoK8sApiCoreV1HostPathVolumeSource`

GetHostPath returns the HostPath field if non-nil, zero value otherwise.

### GetHostPathOk

`func (o *IoK8sApiCoreV1Volume) GetHostPathOk() (*IoK8sApiCoreV1HostPathVolumeSource, bool)`

GetHostPathOk returns a tuple with the HostPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPath

`func (o *IoK8sApiCoreV1Volume) SetHostPath(v IoK8sApiCoreV1HostPathVolumeSource)`

SetHostPath sets HostPath field to given value.

### HasHostPath

`func (o *IoK8sApiCoreV1Volume) HasHostPath() bool`

HasHostPath returns a boolean if a field has been set.

### GetIscsi

`func (o *IoK8sApiCoreV1Volume) GetIscsi() IoK8sApiCoreV1ISCSIVolumeSource`

GetIscsi returns the Iscsi field if non-nil, zero value otherwise.

### GetIscsiOk

`func (o *IoK8sApiCoreV1Volume) GetIscsiOk() (*IoK8sApiCoreV1ISCSIVolumeSource, bool)`

GetIscsiOk returns a tuple with the Iscsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsi

`func (o *IoK8sApiCoreV1Volume) SetIscsi(v IoK8sApiCoreV1ISCSIVolumeSource)`

SetIscsi sets Iscsi field to given value.

### HasIscsi

`func (o *IoK8sApiCoreV1Volume) HasIscsi() bool`

HasIscsi returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiCoreV1Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1Volume) SetName(v string)`

SetName sets Name field to given value.


### GetNfs

`func (o *IoK8sApiCoreV1Volume) GetNfs() IoK8sApiCoreV1NFSVolumeSource`

GetNfs returns the Nfs field if non-nil, zero value otherwise.

### GetNfsOk

`func (o *IoK8sApiCoreV1Volume) GetNfsOk() (*IoK8sApiCoreV1NFSVolumeSource, bool)`

GetNfsOk returns a tuple with the Nfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfs

`func (o *IoK8sApiCoreV1Volume) SetNfs(v IoK8sApiCoreV1NFSVolumeSource)`

SetNfs sets Nfs field to given value.

### HasNfs

`func (o *IoK8sApiCoreV1Volume) HasNfs() bool`

HasNfs returns a boolean if a field has been set.

### GetPersistentVolumeClaim

`func (o *IoK8sApiCoreV1Volume) GetPersistentVolumeClaim() IoK8sApiCoreV1PersistentVolumeClaimVolumeSource`

GetPersistentVolumeClaim returns the PersistentVolumeClaim field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimOk

`func (o *IoK8sApiCoreV1Volume) GetPersistentVolumeClaimOk() (*IoK8sApiCoreV1PersistentVolumeClaimVolumeSource, bool)`

GetPersistentVolumeClaimOk returns a tuple with the PersistentVolumeClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaim

`func (o *IoK8sApiCoreV1Volume) SetPersistentVolumeClaim(v IoK8sApiCoreV1PersistentVolumeClaimVolumeSource)`

SetPersistentVolumeClaim sets PersistentVolumeClaim field to given value.

### HasPersistentVolumeClaim

`func (o *IoK8sApiCoreV1Volume) HasPersistentVolumeClaim() bool`

HasPersistentVolumeClaim returns a boolean if a field has been set.

### GetPhotonPersistentDisk

`func (o *IoK8sApiCoreV1Volume) GetPhotonPersistentDisk() IoK8sApiCoreV1PhotonPersistentDiskVolumeSource`

GetPhotonPersistentDisk returns the PhotonPersistentDisk field if non-nil, zero value otherwise.

### GetPhotonPersistentDiskOk

`func (o *IoK8sApiCoreV1Volume) GetPhotonPersistentDiskOk() (*IoK8sApiCoreV1PhotonPersistentDiskVolumeSource, bool)`

GetPhotonPersistentDiskOk returns a tuple with the PhotonPersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotonPersistentDisk

`func (o *IoK8sApiCoreV1Volume) SetPhotonPersistentDisk(v IoK8sApiCoreV1PhotonPersistentDiskVolumeSource)`

SetPhotonPersistentDisk sets PhotonPersistentDisk field to given value.

### HasPhotonPersistentDisk

`func (o *IoK8sApiCoreV1Volume) HasPhotonPersistentDisk() bool`

HasPhotonPersistentDisk returns a boolean if a field has been set.

### GetPortworxVolume

`func (o *IoK8sApiCoreV1Volume) GetPortworxVolume() IoK8sApiCoreV1PortworxVolumeSource`

GetPortworxVolume returns the PortworxVolume field if non-nil, zero value otherwise.

### GetPortworxVolumeOk

`func (o *IoK8sApiCoreV1Volume) GetPortworxVolumeOk() (*IoK8sApiCoreV1PortworxVolumeSource, bool)`

GetPortworxVolumeOk returns a tuple with the PortworxVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortworxVolume

`func (o *IoK8sApiCoreV1Volume) SetPortworxVolume(v IoK8sApiCoreV1PortworxVolumeSource)`

SetPortworxVolume sets PortworxVolume field to given value.

### HasPortworxVolume

`func (o *IoK8sApiCoreV1Volume) HasPortworxVolume() bool`

HasPortworxVolume returns a boolean if a field has been set.

### GetProjected

`func (o *IoK8sApiCoreV1Volume) GetProjected() IoK8sApiCoreV1ProjectedVolumeSource`

GetProjected returns the Projected field if non-nil, zero value otherwise.

### GetProjectedOk

`func (o *IoK8sApiCoreV1Volume) GetProjectedOk() (*IoK8sApiCoreV1ProjectedVolumeSource, bool)`

GetProjectedOk returns a tuple with the Projected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjected

`func (o *IoK8sApiCoreV1Volume) SetProjected(v IoK8sApiCoreV1ProjectedVolumeSource)`

SetProjected sets Projected field to given value.

### HasProjected

`func (o *IoK8sApiCoreV1Volume) HasProjected() bool`

HasProjected returns a boolean if a field has been set.

### GetQuobyte

`func (o *IoK8sApiCoreV1Volume) GetQuobyte() IoK8sApiCoreV1QuobyteVolumeSource`

GetQuobyte returns the Quobyte field if non-nil, zero value otherwise.

### GetQuobyteOk

`func (o *IoK8sApiCoreV1Volume) GetQuobyteOk() (*IoK8sApiCoreV1QuobyteVolumeSource, bool)`

GetQuobyteOk returns a tuple with the Quobyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuobyte

`func (o *IoK8sApiCoreV1Volume) SetQuobyte(v IoK8sApiCoreV1QuobyteVolumeSource)`

SetQuobyte sets Quobyte field to given value.

### HasQuobyte

`func (o *IoK8sApiCoreV1Volume) HasQuobyte() bool`

HasQuobyte returns a boolean if a field has been set.

### GetRbd

`func (o *IoK8sApiCoreV1Volume) GetRbd() IoK8sApiCoreV1RBDVolumeSource`

GetRbd returns the Rbd field if non-nil, zero value otherwise.

### GetRbdOk

`func (o *IoK8sApiCoreV1Volume) GetRbdOk() (*IoK8sApiCoreV1RBDVolumeSource, bool)`

GetRbdOk returns a tuple with the Rbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbd

`func (o *IoK8sApiCoreV1Volume) SetRbd(v IoK8sApiCoreV1RBDVolumeSource)`

SetRbd sets Rbd field to given value.

### HasRbd

`func (o *IoK8sApiCoreV1Volume) HasRbd() bool`

HasRbd returns a boolean if a field has been set.

### GetScaleIO

`func (o *IoK8sApiCoreV1Volume) GetScaleIO() IoK8sApiCoreV1ScaleIOVolumeSource`

GetScaleIO returns the ScaleIO field if non-nil, zero value otherwise.

### GetScaleIOOk

`func (o *IoK8sApiCoreV1Volume) GetScaleIOOk() (*IoK8sApiCoreV1ScaleIOVolumeSource, bool)`

GetScaleIOOk returns a tuple with the ScaleIO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleIO

`func (o *IoK8sApiCoreV1Volume) SetScaleIO(v IoK8sApiCoreV1ScaleIOVolumeSource)`

SetScaleIO sets ScaleIO field to given value.

### HasScaleIO

`func (o *IoK8sApiCoreV1Volume) HasScaleIO() bool`

HasScaleIO returns a boolean if a field has been set.

### GetSecret

`func (o *IoK8sApiCoreV1Volume) GetSecret() IoK8sApiCoreV1SecretVolumeSource`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IoK8sApiCoreV1Volume) GetSecretOk() (*IoK8sApiCoreV1SecretVolumeSource, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IoK8sApiCoreV1Volume) SetSecret(v IoK8sApiCoreV1SecretVolumeSource)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IoK8sApiCoreV1Volume) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStorageos

`func (o *IoK8sApiCoreV1Volume) GetStorageos() IoK8sApiCoreV1StorageOSVolumeSource`

GetStorageos returns the Storageos field if non-nil, zero value otherwise.

### GetStorageosOk

`func (o *IoK8sApiCoreV1Volume) GetStorageosOk() (*IoK8sApiCoreV1StorageOSVolumeSource, bool)`

GetStorageosOk returns a tuple with the Storageos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageos

`func (o *IoK8sApiCoreV1Volume) SetStorageos(v IoK8sApiCoreV1StorageOSVolumeSource)`

SetStorageos sets Storageos field to given value.

### HasStorageos

`func (o *IoK8sApiCoreV1Volume) HasStorageos() bool`

HasStorageos returns a boolean if a field has been set.

### GetVsphereVolume

`func (o *IoK8sApiCoreV1Volume) GetVsphereVolume() IoK8sApiCoreV1VsphereVirtualDiskVolumeSource`

GetVsphereVolume returns the VsphereVolume field if non-nil, zero value otherwise.

### GetVsphereVolumeOk

`func (o *IoK8sApiCoreV1Volume) GetVsphereVolumeOk() (*IoK8sApiCoreV1VsphereVirtualDiskVolumeSource, bool)`

GetVsphereVolumeOk returns a tuple with the VsphereVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereVolume

`func (o *IoK8sApiCoreV1Volume) SetVsphereVolume(v IoK8sApiCoreV1VsphereVirtualDiskVolumeSource)`

SetVsphereVolume sets VsphereVolume field to given value.

### HasVsphereVolume

`func (o *IoK8sApiCoreV1Volume) HasVsphereVolume() bool`

HasVsphereVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


