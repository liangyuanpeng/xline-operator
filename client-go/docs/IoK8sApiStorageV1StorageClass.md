# IoK8sApiStorageV1StorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowVolumeExpansion** | Pointer to **bool** | allowVolumeExpansion shows whether the storage class allow volume expand. | [optional] 
**AllowedTopologies** | Pointer to [**[]IoK8sApiCoreV1TopologySelectorTerm**](IoK8sApiCoreV1TopologySelectorTerm.md) | allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature. | [optional] 
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**MountOptions** | Pointer to **[]string** | mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. [\&quot;ro\&quot;, \&quot;soft\&quot;]. Not validated - mount of the PVs will simply fail if one is invalid. | [optional] 
**Parameters** | Pointer to **map[string]string** | parameters holds the parameters for the provisioner that should create volumes of this storage class. | [optional] 
**Provisioner** | **string** | provisioner indicates the type of the provisioner. | 
**ReclaimPolicy** | Pointer to **string** | reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.  Possible enum values:  - &#x60;\&quot;Delete\&quot;&#x60; means the volume will be deleted from Kubernetes on release from its claim. The volume plugin must support Deletion.  - &#x60;\&quot;Recycle\&quot;&#x60; means the volume will be recycled back into the pool of unbound persistent volumes on release from its claim. The volume plugin must support Recycling.  - &#x60;\&quot;Retain\&quot;&#x60; means the volume will be left in its current phase (Released) for manual reclamation by the administrator. The default policy is Retain. | [optional] 
**VolumeBindingMode** | Pointer to **string** | volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.  Possible enum values:  - &#x60;\&quot;Immediate\&quot;&#x60; indicates that PersistentVolumeClaims should be immediately provisioned and bound. This is the default mode.  - &#x60;\&quot;WaitForFirstConsumer\&quot;&#x60; indicates that PersistentVolumeClaims should not be provisioned and bound until the first Pod is created that references the PeristentVolumeClaim. The volume provisioning and binding will occur during Pod scheduing. | [optional] 

## Methods

### NewIoK8sApiStorageV1StorageClass

`func NewIoK8sApiStorageV1StorageClass(provisioner string, ) *IoK8sApiStorageV1StorageClass`

NewIoK8sApiStorageV1StorageClass instantiates a new IoK8sApiStorageV1StorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiStorageV1StorageClassWithDefaults

`func NewIoK8sApiStorageV1StorageClassWithDefaults() *IoK8sApiStorageV1StorageClass`

NewIoK8sApiStorageV1StorageClassWithDefaults instantiates a new IoK8sApiStorageV1StorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowVolumeExpansion

`func (o *IoK8sApiStorageV1StorageClass) GetAllowVolumeExpansion() bool`

GetAllowVolumeExpansion returns the AllowVolumeExpansion field if non-nil, zero value otherwise.

### GetAllowVolumeExpansionOk

`func (o *IoK8sApiStorageV1StorageClass) GetAllowVolumeExpansionOk() (*bool, bool)`

GetAllowVolumeExpansionOk returns a tuple with the AllowVolumeExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVolumeExpansion

`func (o *IoK8sApiStorageV1StorageClass) SetAllowVolumeExpansion(v bool)`

SetAllowVolumeExpansion sets AllowVolumeExpansion field to given value.

### HasAllowVolumeExpansion

`func (o *IoK8sApiStorageV1StorageClass) HasAllowVolumeExpansion() bool`

HasAllowVolumeExpansion returns a boolean if a field has been set.

### GetAllowedTopologies

`func (o *IoK8sApiStorageV1StorageClass) GetAllowedTopologies() []IoK8sApiCoreV1TopologySelectorTerm`

GetAllowedTopologies returns the AllowedTopologies field if non-nil, zero value otherwise.

### GetAllowedTopologiesOk

`func (o *IoK8sApiStorageV1StorageClass) GetAllowedTopologiesOk() (*[]IoK8sApiCoreV1TopologySelectorTerm, bool)`

GetAllowedTopologiesOk returns a tuple with the AllowedTopologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTopologies

`func (o *IoK8sApiStorageV1StorageClass) SetAllowedTopologies(v []IoK8sApiCoreV1TopologySelectorTerm)`

SetAllowedTopologies sets AllowedTopologies field to given value.

### HasAllowedTopologies

`func (o *IoK8sApiStorageV1StorageClass) HasAllowedTopologies() bool`

HasAllowedTopologies returns a boolean if a field has been set.

### GetApiVersion

`func (o *IoK8sApiStorageV1StorageClass) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiStorageV1StorageClass) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiStorageV1StorageClass) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiStorageV1StorageClass) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiStorageV1StorageClass) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiStorageV1StorageClass) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiStorageV1StorageClass) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiStorageV1StorageClass) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiStorageV1StorageClass) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiStorageV1StorageClass) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiStorageV1StorageClass) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiStorageV1StorageClass) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMountOptions

`func (o *IoK8sApiStorageV1StorageClass) GetMountOptions() []string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *IoK8sApiStorageV1StorageClass) GetMountOptionsOk() (*[]string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *IoK8sApiStorageV1StorageClass) SetMountOptions(v []string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *IoK8sApiStorageV1StorageClass) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.

### GetParameters

`func (o *IoK8sApiStorageV1StorageClass) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoK8sApiStorageV1StorageClass) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoK8sApiStorageV1StorageClass) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoK8sApiStorageV1StorageClass) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetProvisioner

`func (o *IoK8sApiStorageV1StorageClass) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *IoK8sApiStorageV1StorageClass) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *IoK8sApiStorageV1StorageClass) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.


### GetReclaimPolicy

`func (o *IoK8sApiStorageV1StorageClass) GetReclaimPolicy() string`

GetReclaimPolicy returns the ReclaimPolicy field if non-nil, zero value otherwise.

### GetReclaimPolicyOk

`func (o *IoK8sApiStorageV1StorageClass) GetReclaimPolicyOk() (*string, bool)`

GetReclaimPolicyOk returns a tuple with the ReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimPolicy

`func (o *IoK8sApiStorageV1StorageClass) SetReclaimPolicy(v string)`

SetReclaimPolicy sets ReclaimPolicy field to given value.

### HasReclaimPolicy

`func (o *IoK8sApiStorageV1StorageClass) HasReclaimPolicy() bool`

HasReclaimPolicy returns a boolean if a field has been set.

### GetVolumeBindingMode

`func (o *IoK8sApiStorageV1StorageClass) GetVolumeBindingMode() string`

GetVolumeBindingMode returns the VolumeBindingMode field if non-nil, zero value otherwise.

### GetVolumeBindingModeOk

`func (o *IoK8sApiStorageV1StorageClass) GetVolumeBindingModeOk() (*string, bool)`

GetVolumeBindingModeOk returns a tuple with the VolumeBindingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBindingMode

`func (o *IoK8sApiStorageV1StorageClass) SetVolumeBindingMode(v string)`

SetVolumeBindingMode sets VolumeBindingMode field to given value.

### HasVolumeBindingMode

`func (o *IoK8sApiStorageV1StorageClass) HasVolumeBindingMode() bool`

HasVolumeBindingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


