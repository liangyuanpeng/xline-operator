# IoK8sApiCoreV1PersistentVolumeClaimTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Spec** | [**IoK8sApiCoreV1PersistentVolumeClaimSpec**](IoK8sApiCoreV1PersistentVolumeClaimSpec.md) |  | 

## Methods

### NewIoK8sApiCoreV1PersistentVolumeClaimTemplate

`func NewIoK8sApiCoreV1PersistentVolumeClaimTemplate(spec IoK8sApiCoreV1PersistentVolumeClaimSpec, ) *IoK8sApiCoreV1PersistentVolumeClaimTemplate`

NewIoK8sApiCoreV1PersistentVolumeClaimTemplate instantiates a new IoK8sApiCoreV1PersistentVolumeClaimTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PersistentVolumeClaimTemplateWithDefaults

`func NewIoK8sApiCoreV1PersistentVolumeClaimTemplateWithDefaults() *IoK8sApiCoreV1PersistentVolumeClaimTemplate`

NewIoK8sApiCoreV1PersistentVolumeClaimTemplateWithDefaults instantiates a new IoK8sApiCoreV1PersistentVolumeClaimTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *IoK8sApiCoreV1PersistentVolumeClaimTemplate) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimTemplate) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiCoreV1PersistentVolumeClaimTemplate) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiCoreV1PersistentVolumeClaimTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *IoK8sApiCoreV1PersistentVolumeClaimTemplate) GetSpec() IoK8sApiCoreV1PersistentVolumeClaimSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimTemplate) GetSpecOk() (*IoK8sApiCoreV1PersistentVolumeClaimSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoK8sApiCoreV1PersistentVolumeClaimTemplate) SetSpec(v IoK8sApiCoreV1PersistentVolumeClaimSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


