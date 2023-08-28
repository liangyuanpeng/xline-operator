# IoK8sApimachineryPkgApisMetaV1StatusDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Causes** | Pointer to [**[]IoK8sApimachineryPkgApisMetaV1StatusCause**](IoK8sApimachineryPkgApisMetaV1StatusCause.md) | The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes. | [optional] 
**Group** | Pointer to **string** | The group attribute of the resource associated with the status StatusReason. | [optional] 
**Kind** | Pointer to **string** | The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Name** | Pointer to **string** | The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described). | [optional] 
**RetryAfterSeconds** | Pointer to **int32** | If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action. | [optional] 
**Uid** | Pointer to **string** | UID of the resource. (when there is a single resource which can be described). More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids | [optional] 

## Methods

### NewIoK8sApimachineryPkgApisMetaV1StatusDetails

`func NewIoK8sApimachineryPkgApisMetaV1StatusDetails() *IoK8sApimachineryPkgApisMetaV1StatusDetails`

NewIoK8sApimachineryPkgApisMetaV1StatusDetails instantiates a new IoK8sApimachineryPkgApisMetaV1StatusDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApimachineryPkgApisMetaV1StatusDetailsWithDefaults

`func NewIoK8sApimachineryPkgApisMetaV1StatusDetailsWithDefaults() *IoK8sApimachineryPkgApisMetaV1StatusDetails`

NewIoK8sApimachineryPkgApisMetaV1StatusDetailsWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1StatusDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCauses

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetCauses() []IoK8sApimachineryPkgApisMetaV1StatusCause`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetCausesOk() (*[]IoK8sApimachineryPkgApisMetaV1StatusCause, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetCauses(v []IoK8sApimachineryPkgApisMetaV1StatusCause)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetGroup

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetryAfterSeconds

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetRetryAfterSeconds() int32`

GetRetryAfterSeconds returns the RetryAfterSeconds field if non-nil, zero value otherwise.

### GetRetryAfterSecondsOk

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetRetryAfterSecondsOk() (*int32, bool)`

GetRetryAfterSecondsOk returns a tuple with the RetryAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfterSeconds

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetRetryAfterSeconds(v int32)`

SetRetryAfterSeconds sets RetryAfterSeconds field to given value.

### HasRetryAfterSeconds

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasRetryAfterSeconds() bool`

HasRetryAfterSeconds returns a boolean if a field has been set.

### GetUid

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


