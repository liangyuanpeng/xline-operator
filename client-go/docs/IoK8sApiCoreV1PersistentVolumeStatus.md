# IoK8sApiCoreV1PersistentVolumeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | message is a human-readable message indicating details about why the volume is in this state. | [optional] 
**Phase** | Pointer to **string** | phase indicates if a volume is available, bound to a claim, or released by a claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase  Possible enum values:  - &#x60;\&quot;Available\&quot;&#x60; used for PersistentVolumes that are not yet bound Available volumes are held by the binder and matched to PersistentVolumeClaims  - &#x60;\&quot;Bound\&quot;&#x60; used for PersistentVolumes that are bound  - &#x60;\&quot;Failed\&quot;&#x60; used for PersistentVolumes that failed to be correctly recycled or deleted after being released from a claim  - &#x60;\&quot;Pending\&quot;&#x60; used for PersistentVolumes that are not available  - &#x60;\&quot;Released\&quot;&#x60; used for PersistentVolumes where the bound PersistentVolumeClaim was deleted released volumes must be recycled before becoming available again this phase is used by the persistent volume claim binder to signal to another process to reclaim the resource | [optional] 
**Reason** | Pointer to **string** | reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI. | [optional] 

## Methods

### NewIoK8sApiCoreV1PersistentVolumeStatus

`func NewIoK8sApiCoreV1PersistentVolumeStatus() *IoK8sApiCoreV1PersistentVolumeStatus`

NewIoK8sApiCoreV1PersistentVolumeStatus instantiates a new IoK8sApiCoreV1PersistentVolumeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PersistentVolumeStatusWithDefaults

`func NewIoK8sApiCoreV1PersistentVolumeStatusWithDefaults() *IoK8sApiCoreV1PersistentVolumeStatus`

NewIoK8sApiCoreV1PersistentVolumeStatusWithDefaults instantiates a new IoK8sApiCoreV1PersistentVolumeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPhase

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiCoreV1PersistentVolumeStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


