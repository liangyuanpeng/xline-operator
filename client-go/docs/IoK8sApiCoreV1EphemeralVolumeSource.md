# IoK8sApiCoreV1EphemeralVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeClaimTemplate** | Pointer to [**IoK8sApiCoreV1PersistentVolumeClaimTemplate**](IoK8sApiCoreV1PersistentVolumeClaimTemplate.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1EphemeralVolumeSource

`func NewIoK8sApiCoreV1EphemeralVolumeSource() *IoK8sApiCoreV1EphemeralVolumeSource`

NewIoK8sApiCoreV1EphemeralVolumeSource instantiates a new IoK8sApiCoreV1EphemeralVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1EphemeralVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1EphemeralVolumeSourceWithDefaults() *IoK8sApiCoreV1EphemeralVolumeSource`

NewIoK8sApiCoreV1EphemeralVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1EphemeralVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeClaimTemplate

`func (o *IoK8sApiCoreV1EphemeralVolumeSource) GetVolumeClaimTemplate() IoK8sApiCoreV1PersistentVolumeClaimTemplate`

GetVolumeClaimTemplate returns the VolumeClaimTemplate field if non-nil, zero value otherwise.

### GetVolumeClaimTemplateOk

`func (o *IoK8sApiCoreV1EphemeralVolumeSource) GetVolumeClaimTemplateOk() (*IoK8sApiCoreV1PersistentVolumeClaimTemplate, bool)`

GetVolumeClaimTemplateOk returns a tuple with the VolumeClaimTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClaimTemplate

`func (o *IoK8sApiCoreV1EphemeralVolumeSource) SetVolumeClaimTemplate(v IoK8sApiCoreV1PersistentVolumeClaimTemplate)`

SetVolumeClaimTemplate sets VolumeClaimTemplate field to given value.

### HasVolumeClaimTemplate

`func (o *IoK8sApiCoreV1EphemeralVolumeSource) HasVolumeClaimTemplate() bool`

HasVolumeClaimTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


