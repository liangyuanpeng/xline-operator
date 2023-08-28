# IoK8sApiCoreV1PersistentVolumeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** | accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes | [optional] 
**AwsElasticBlockStore** | Pointer to [**IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource**](IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource.md) |  | [optional] 
**AzureDisk** | Pointer to [**IoK8sApiCoreV1AzureDiskVolumeSource**](IoK8sApiCoreV1AzureDiskVolumeSource.md) |  | [optional] 
**AzureFile** | Pointer to [**IoK8sApiCoreV1AzureFilePersistentVolumeSource**](IoK8sApiCoreV1AzureFilePersistentVolumeSource.md) |  | [optional] 
**Capacity** | Pointer to **map[string]string** | capacity is the description of the persistent volume&#39;s resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity | [optional] 
**Cephfs** | Pointer to [**IoK8sApiCoreV1CephFSPersistentVolumeSource**](IoK8sApiCoreV1CephFSPersistentVolumeSource.md) |  | [optional] 
**Cinder** | Pointer to [**IoK8sApiCoreV1CinderPersistentVolumeSource**](IoK8sApiCoreV1CinderPersistentVolumeSource.md) |  | [optional] 
**ClaimRef** | Pointer to [**IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) |  | [optional] 
**Csi** | Pointer to [**IoK8sApiCoreV1CSIPersistentVolumeSource**](IoK8sApiCoreV1CSIPersistentVolumeSource.md) |  | [optional] 
**Fc** | Pointer to [**IoK8sApiCoreV1FCVolumeSource**](IoK8sApiCoreV1FCVolumeSource.md) |  | [optional] 
**FlexVolume** | Pointer to [**IoK8sApiCoreV1FlexPersistentVolumeSource**](IoK8sApiCoreV1FlexPersistentVolumeSource.md) |  | [optional] 
**Flocker** | Pointer to [**IoK8sApiCoreV1FlockerVolumeSource**](IoK8sApiCoreV1FlockerVolumeSource.md) |  | [optional] 
**GcePersistentDisk** | Pointer to [**IoK8sApiCoreV1GCEPersistentDiskVolumeSource**](IoK8sApiCoreV1GCEPersistentDiskVolumeSource.md) |  | [optional] 
**Glusterfs** | Pointer to [**IoK8sApiCoreV1GlusterfsPersistentVolumeSource**](IoK8sApiCoreV1GlusterfsPersistentVolumeSource.md) |  | [optional] 
**HostPath** | Pointer to [**IoK8sApiCoreV1HostPathVolumeSource**](IoK8sApiCoreV1HostPathVolumeSource.md) |  | [optional] 
**Iscsi** | Pointer to [**IoK8sApiCoreV1ISCSIPersistentVolumeSource**](IoK8sApiCoreV1ISCSIPersistentVolumeSource.md) |  | [optional] 
**Local** | Pointer to [**IoK8sApiCoreV1LocalVolumeSource**](IoK8sApiCoreV1LocalVolumeSource.md) |  | [optional] 
**MountOptions** | Pointer to **[]string** | mountOptions is the list of mount options, e.g. [\&quot;ro\&quot;, \&quot;soft\&quot;]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options | [optional] 
**Nfs** | Pointer to [**IoK8sApiCoreV1NFSVolumeSource**](IoK8sApiCoreV1NFSVolumeSource.md) |  | [optional] 
**NodeAffinity** | Pointer to [**IoK8sApiCoreV1VolumeNodeAffinity**](IoK8sApiCoreV1VolumeNodeAffinity.md) |  | [optional] 
**PersistentVolumeReclaimPolicy** | Pointer to **string** | persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming  Possible enum values:  - &#x60;\&quot;Delete\&quot;&#x60; means the volume will be deleted from Kubernetes on release from its claim. The volume plugin must support Deletion.  - &#x60;\&quot;Recycle\&quot;&#x60; means the volume will be recycled back into the pool of unbound persistent volumes on release from its claim. The volume plugin must support Recycling.  - &#x60;\&quot;Retain\&quot;&#x60; means the volume will be left in its current phase (Released) for manual reclamation by the administrator. The default policy is Retain. | [optional] 
**PhotonPersistentDisk** | Pointer to [**IoK8sApiCoreV1PhotonPersistentDiskVolumeSource**](IoK8sApiCoreV1PhotonPersistentDiskVolumeSource.md) |  | [optional] 
**PortworxVolume** | Pointer to [**IoK8sApiCoreV1PortworxVolumeSource**](IoK8sApiCoreV1PortworxVolumeSource.md) |  | [optional] 
**Quobyte** | Pointer to [**IoK8sApiCoreV1QuobyteVolumeSource**](IoK8sApiCoreV1QuobyteVolumeSource.md) |  | [optional] 
**Rbd** | Pointer to [**IoK8sApiCoreV1RBDPersistentVolumeSource**](IoK8sApiCoreV1RBDPersistentVolumeSource.md) |  | [optional] 
**ScaleIO** | Pointer to [**IoK8sApiCoreV1ScaleIOPersistentVolumeSource**](IoK8sApiCoreV1ScaleIOPersistentVolumeSource.md) |  | [optional] 
**StorageClassName** | Pointer to **string** | storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass. | [optional] 
**Storageos** | Pointer to [**IoK8sApiCoreV1StorageOSPersistentVolumeSource**](IoK8sApiCoreV1StorageOSPersistentVolumeSource.md) |  | [optional] 
**VolumeMode** | Pointer to **string** | volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.  Possible enum values:  - &#x60;\&quot;Block\&quot;&#x60; means the volume will not be formatted with a filesystem and will remain a raw block device.  - &#x60;\&quot;Filesystem\&quot;&#x60; means the volume will be or is formatted with a filesystem. | [optional] 
**VsphereVolume** | Pointer to [**IoK8sApiCoreV1VsphereVirtualDiskVolumeSource**](IoK8sApiCoreV1VsphereVirtualDiskVolumeSource.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1PersistentVolumeSpec

`func NewIoK8sApiCoreV1PersistentVolumeSpec() *IoK8sApiCoreV1PersistentVolumeSpec`

NewIoK8sApiCoreV1PersistentVolumeSpec instantiates a new IoK8sApiCoreV1PersistentVolumeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PersistentVolumeSpecWithDefaults

`func NewIoK8sApiCoreV1PersistentVolumeSpecWithDefaults() *IoK8sApiCoreV1PersistentVolumeSpec`

NewIoK8sApiCoreV1PersistentVolumeSpecWithDefaults instantiates a new IoK8sApiCoreV1PersistentVolumeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetAwsElasticBlockStore

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetAwsElasticBlockStore() IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource`

GetAwsElasticBlockStore returns the AwsElasticBlockStore field if non-nil, zero value otherwise.

### GetAwsElasticBlockStoreOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetAwsElasticBlockStoreOk() (*IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource, bool)`

GetAwsElasticBlockStoreOk returns a tuple with the AwsElasticBlockStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsElasticBlockStore

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetAwsElasticBlockStore(v IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource)`

SetAwsElasticBlockStore sets AwsElasticBlockStore field to given value.

### HasAwsElasticBlockStore

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasAwsElasticBlockStore() bool`

HasAwsElasticBlockStore returns a boolean if a field has been set.

### GetAzureDisk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetAzureDisk() IoK8sApiCoreV1AzureDiskVolumeSource`

GetAzureDisk returns the AzureDisk field if non-nil, zero value otherwise.

### GetAzureDiskOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetAzureDiskOk() (*IoK8sApiCoreV1AzureDiskVolumeSource, bool)`

GetAzureDiskOk returns a tuple with the AzureDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDisk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetAzureDisk(v IoK8sApiCoreV1AzureDiskVolumeSource)`

SetAzureDisk sets AzureDisk field to given value.

### HasAzureDisk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasAzureDisk() bool`

HasAzureDisk returns a boolean if a field has been set.

### GetAzureFile

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetAzureFile() IoK8sApiCoreV1AzureFilePersistentVolumeSource`

GetAzureFile returns the AzureFile field if non-nil, zero value otherwise.

### GetAzureFileOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetAzureFileOk() (*IoK8sApiCoreV1AzureFilePersistentVolumeSource, bool)`

GetAzureFileOk returns a tuple with the AzureFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureFile

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetAzureFile(v IoK8sApiCoreV1AzureFilePersistentVolumeSource)`

SetAzureFile sets AzureFile field to given value.

### HasAzureFile

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasAzureFile() bool`

HasAzureFile returns a boolean if a field has been set.

### GetCapacity

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetCapacity() map[string]string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetCapacityOk() (*map[string]string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetCapacity(v map[string]string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCephfs

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetCephfs() IoK8sApiCoreV1CephFSPersistentVolumeSource`

GetCephfs returns the Cephfs field if non-nil, zero value otherwise.

### GetCephfsOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetCephfsOk() (*IoK8sApiCoreV1CephFSPersistentVolumeSource, bool)`

GetCephfsOk returns a tuple with the Cephfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephfs

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetCephfs(v IoK8sApiCoreV1CephFSPersistentVolumeSource)`

SetCephfs sets Cephfs field to given value.

### HasCephfs

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasCephfs() bool`

HasCephfs returns a boolean if a field has been set.

### GetCinder

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetCinder() IoK8sApiCoreV1CinderPersistentVolumeSource`

GetCinder returns the Cinder field if non-nil, zero value otherwise.

### GetCinderOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetCinderOk() (*IoK8sApiCoreV1CinderPersistentVolumeSource, bool)`

GetCinderOk returns a tuple with the Cinder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCinder

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetCinder(v IoK8sApiCoreV1CinderPersistentVolumeSource)`

SetCinder sets Cinder field to given value.

### HasCinder

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasCinder() bool`

HasCinder returns a boolean if a field has been set.

### GetClaimRef

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetClaimRef() IoK8sApiCoreV1ObjectReference`

GetClaimRef returns the ClaimRef field if non-nil, zero value otherwise.

### GetClaimRefOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetClaimRefOk() (*IoK8sApiCoreV1ObjectReference, bool)`

GetClaimRefOk returns a tuple with the ClaimRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimRef

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetClaimRef(v IoK8sApiCoreV1ObjectReference)`

SetClaimRef sets ClaimRef field to given value.

### HasClaimRef

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasClaimRef() bool`

HasClaimRef returns a boolean if a field has been set.

### GetCsi

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetCsi() IoK8sApiCoreV1CSIPersistentVolumeSource`

GetCsi returns the Csi field if non-nil, zero value otherwise.

### GetCsiOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetCsiOk() (*IoK8sApiCoreV1CSIPersistentVolumeSource, bool)`

GetCsiOk returns a tuple with the Csi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsi

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetCsi(v IoK8sApiCoreV1CSIPersistentVolumeSource)`

SetCsi sets Csi field to given value.

### HasCsi

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasCsi() bool`

HasCsi returns a boolean if a field has been set.

### GetFc

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetFc() IoK8sApiCoreV1FCVolumeSource`

GetFc returns the Fc field if non-nil, zero value otherwise.

### GetFcOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetFcOk() (*IoK8sApiCoreV1FCVolumeSource, bool)`

GetFcOk returns a tuple with the Fc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFc

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetFc(v IoK8sApiCoreV1FCVolumeSource)`

SetFc sets Fc field to given value.

### HasFc

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasFc() bool`

HasFc returns a boolean if a field has been set.

### GetFlexVolume

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetFlexVolume() IoK8sApiCoreV1FlexPersistentVolumeSource`

GetFlexVolume returns the FlexVolume field if non-nil, zero value otherwise.

### GetFlexVolumeOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetFlexVolumeOk() (*IoK8sApiCoreV1FlexPersistentVolumeSource, bool)`

GetFlexVolumeOk returns a tuple with the FlexVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexVolume

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetFlexVolume(v IoK8sApiCoreV1FlexPersistentVolumeSource)`

SetFlexVolume sets FlexVolume field to given value.

### HasFlexVolume

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasFlexVolume() bool`

HasFlexVolume returns a boolean if a field has been set.

### GetFlocker

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetFlocker() IoK8sApiCoreV1FlockerVolumeSource`

GetFlocker returns the Flocker field if non-nil, zero value otherwise.

### GetFlockerOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetFlockerOk() (*IoK8sApiCoreV1FlockerVolumeSource, bool)`

GetFlockerOk returns a tuple with the Flocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlocker

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetFlocker(v IoK8sApiCoreV1FlockerVolumeSource)`

SetFlocker sets Flocker field to given value.

### HasFlocker

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasFlocker() bool`

HasFlocker returns a boolean if a field has been set.

### GetGcePersistentDisk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetGcePersistentDisk() IoK8sApiCoreV1GCEPersistentDiskVolumeSource`

GetGcePersistentDisk returns the GcePersistentDisk field if non-nil, zero value otherwise.

### GetGcePersistentDiskOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetGcePersistentDiskOk() (*IoK8sApiCoreV1GCEPersistentDiskVolumeSource, bool)`

GetGcePersistentDiskOk returns a tuple with the GcePersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcePersistentDisk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetGcePersistentDisk(v IoK8sApiCoreV1GCEPersistentDiskVolumeSource)`

SetGcePersistentDisk sets GcePersistentDisk field to given value.

### HasGcePersistentDisk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasGcePersistentDisk() bool`

HasGcePersistentDisk returns a boolean if a field has been set.

### GetGlusterfs

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetGlusterfs() IoK8sApiCoreV1GlusterfsPersistentVolumeSource`

GetGlusterfs returns the Glusterfs field if non-nil, zero value otherwise.

### GetGlusterfsOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetGlusterfsOk() (*IoK8sApiCoreV1GlusterfsPersistentVolumeSource, bool)`

GetGlusterfsOk returns a tuple with the Glusterfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlusterfs

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetGlusterfs(v IoK8sApiCoreV1GlusterfsPersistentVolumeSource)`

SetGlusterfs sets Glusterfs field to given value.

### HasGlusterfs

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasGlusterfs() bool`

HasGlusterfs returns a boolean if a field has been set.

### GetHostPath

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetHostPath() IoK8sApiCoreV1HostPathVolumeSource`

GetHostPath returns the HostPath field if non-nil, zero value otherwise.

### GetHostPathOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetHostPathOk() (*IoK8sApiCoreV1HostPathVolumeSource, bool)`

GetHostPathOk returns a tuple with the HostPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPath

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetHostPath(v IoK8sApiCoreV1HostPathVolumeSource)`

SetHostPath sets HostPath field to given value.

### HasHostPath

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasHostPath() bool`

HasHostPath returns a boolean if a field has been set.

### GetIscsi

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetIscsi() IoK8sApiCoreV1ISCSIPersistentVolumeSource`

GetIscsi returns the Iscsi field if non-nil, zero value otherwise.

### GetIscsiOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetIscsiOk() (*IoK8sApiCoreV1ISCSIPersistentVolumeSource, bool)`

GetIscsiOk returns a tuple with the Iscsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsi

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetIscsi(v IoK8sApiCoreV1ISCSIPersistentVolumeSource)`

SetIscsi sets Iscsi field to given value.

### HasIscsi

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasIscsi() bool`

HasIscsi returns a boolean if a field has been set.

### GetLocal

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetLocal() IoK8sApiCoreV1LocalVolumeSource`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetLocalOk() (*IoK8sApiCoreV1LocalVolumeSource, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetLocal(v IoK8sApiCoreV1LocalVolumeSource)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetMountOptions

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetMountOptions() []string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetMountOptionsOk() (*[]string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetMountOptions(v []string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.

### GetNfs

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetNfs() IoK8sApiCoreV1NFSVolumeSource`

GetNfs returns the Nfs field if non-nil, zero value otherwise.

### GetNfsOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetNfsOk() (*IoK8sApiCoreV1NFSVolumeSource, bool)`

GetNfsOk returns a tuple with the Nfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfs

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetNfs(v IoK8sApiCoreV1NFSVolumeSource)`

SetNfs sets Nfs field to given value.

### HasNfs

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasNfs() bool`

HasNfs returns a boolean if a field has been set.

### GetNodeAffinity

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetNodeAffinity() IoK8sApiCoreV1VolumeNodeAffinity`

GetNodeAffinity returns the NodeAffinity field if non-nil, zero value otherwise.

### GetNodeAffinityOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetNodeAffinityOk() (*IoK8sApiCoreV1VolumeNodeAffinity, bool)`

GetNodeAffinityOk returns a tuple with the NodeAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAffinity

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetNodeAffinity(v IoK8sApiCoreV1VolumeNodeAffinity)`

SetNodeAffinity sets NodeAffinity field to given value.

### HasNodeAffinity

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasNodeAffinity() bool`

HasNodeAffinity returns a boolean if a field has been set.

### GetPersistentVolumeReclaimPolicy

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetPersistentVolumeReclaimPolicy() string`

GetPersistentVolumeReclaimPolicy returns the PersistentVolumeReclaimPolicy field if non-nil, zero value otherwise.

### GetPersistentVolumeReclaimPolicyOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetPersistentVolumeReclaimPolicyOk() (*string, bool)`

GetPersistentVolumeReclaimPolicyOk returns a tuple with the PersistentVolumeReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeReclaimPolicy

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetPersistentVolumeReclaimPolicy(v string)`

SetPersistentVolumeReclaimPolicy sets PersistentVolumeReclaimPolicy field to given value.

### HasPersistentVolumeReclaimPolicy

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasPersistentVolumeReclaimPolicy() bool`

HasPersistentVolumeReclaimPolicy returns a boolean if a field has been set.

### GetPhotonPersistentDisk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetPhotonPersistentDisk() IoK8sApiCoreV1PhotonPersistentDiskVolumeSource`

GetPhotonPersistentDisk returns the PhotonPersistentDisk field if non-nil, zero value otherwise.

### GetPhotonPersistentDiskOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetPhotonPersistentDiskOk() (*IoK8sApiCoreV1PhotonPersistentDiskVolumeSource, bool)`

GetPhotonPersistentDiskOk returns a tuple with the PhotonPersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotonPersistentDisk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetPhotonPersistentDisk(v IoK8sApiCoreV1PhotonPersistentDiskVolumeSource)`

SetPhotonPersistentDisk sets PhotonPersistentDisk field to given value.

### HasPhotonPersistentDisk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasPhotonPersistentDisk() bool`

HasPhotonPersistentDisk returns a boolean if a field has been set.

### GetPortworxVolume

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetPortworxVolume() IoK8sApiCoreV1PortworxVolumeSource`

GetPortworxVolume returns the PortworxVolume field if non-nil, zero value otherwise.

### GetPortworxVolumeOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetPortworxVolumeOk() (*IoK8sApiCoreV1PortworxVolumeSource, bool)`

GetPortworxVolumeOk returns a tuple with the PortworxVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortworxVolume

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetPortworxVolume(v IoK8sApiCoreV1PortworxVolumeSource)`

SetPortworxVolume sets PortworxVolume field to given value.

### HasPortworxVolume

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasPortworxVolume() bool`

HasPortworxVolume returns a boolean if a field has been set.

### GetQuobyte

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetQuobyte() IoK8sApiCoreV1QuobyteVolumeSource`

GetQuobyte returns the Quobyte field if non-nil, zero value otherwise.

### GetQuobyteOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetQuobyteOk() (*IoK8sApiCoreV1QuobyteVolumeSource, bool)`

GetQuobyteOk returns a tuple with the Quobyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuobyte

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetQuobyte(v IoK8sApiCoreV1QuobyteVolumeSource)`

SetQuobyte sets Quobyte field to given value.

### HasQuobyte

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasQuobyte() bool`

HasQuobyte returns a boolean if a field has been set.

### GetRbd

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetRbd() IoK8sApiCoreV1RBDPersistentVolumeSource`

GetRbd returns the Rbd field if non-nil, zero value otherwise.

### GetRbdOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetRbdOk() (*IoK8sApiCoreV1RBDPersistentVolumeSource, bool)`

GetRbdOk returns a tuple with the Rbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbd

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetRbd(v IoK8sApiCoreV1RBDPersistentVolumeSource)`

SetRbd sets Rbd field to given value.

### HasRbd

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasRbd() bool`

HasRbd returns a boolean if a field has been set.

### GetScaleIO

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetScaleIO() IoK8sApiCoreV1ScaleIOPersistentVolumeSource`

GetScaleIO returns the ScaleIO field if non-nil, zero value otherwise.

### GetScaleIOOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetScaleIOOk() (*IoK8sApiCoreV1ScaleIOPersistentVolumeSource, bool)`

GetScaleIOOk returns a tuple with the ScaleIO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleIO

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetScaleIO(v IoK8sApiCoreV1ScaleIOPersistentVolumeSource)`

SetScaleIO sets ScaleIO field to given value.

### HasScaleIO

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasScaleIO() bool`

HasScaleIO returns a boolean if a field has been set.

### GetStorageClassName

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### GetStorageos

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetStorageos() IoK8sApiCoreV1StorageOSPersistentVolumeSource`

GetStorageos returns the Storageos field if non-nil, zero value otherwise.

### GetStorageosOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetStorageosOk() (*IoK8sApiCoreV1StorageOSPersistentVolumeSource, bool)`

GetStorageosOk returns a tuple with the Storageos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageos

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetStorageos(v IoK8sApiCoreV1StorageOSPersistentVolumeSource)`

SetStorageos sets Storageos field to given value.

### HasStorageos

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasStorageos() bool`

HasStorageos returns a boolean if a field has been set.

### GetVolumeMode

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetVolumeMode() string`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetVolumeModeOk() (*string, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetVolumeMode(v string)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.

### GetVsphereVolume

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetVsphereVolume() IoK8sApiCoreV1VsphereVirtualDiskVolumeSource`

GetVsphereVolume returns the VsphereVolume field if non-nil, zero value otherwise.

### GetVsphereVolumeOk

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) GetVsphereVolumeOk() (*IoK8sApiCoreV1VsphereVirtualDiskVolumeSource, bool)`

GetVsphereVolumeOk returns a tuple with the VsphereVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereVolume

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) SetVsphereVolume(v IoK8sApiCoreV1VsphereVirtualDiskVolumeSource)`

SetVsphereVolume sets VsphereVolume field to given value.

### HasVsphereVolume

`func (o *IoK8sApiCoreV1PersistentVolumeSpec) HasVsphereVolume() bool`

HasVsphereVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


