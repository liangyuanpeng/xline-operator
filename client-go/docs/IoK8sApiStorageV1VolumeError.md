# IoK8sApiStorageV1VolumeError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | message represents the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information. | [optional] 
**Time** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 

## Methods

### NewIoK8sApiStorageV1VolumeError

`func NewIoK8sApiStorageV1VolumeError() *IoK8sApiStorageV1VolumeError`

NewIoK8sApiStorageV1VolumeError instantiates a new IoK8sApiStorageV1VolumeError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiStorageV1VolumeErrorWithDefaults

`func NewIoK8sApiStorageV1VolumeErrorWithDefaults() *IoK8sApiStorageV1VolumeError`

NewIoK8sApiStorageV1VolumeErrorWithDefaults instantiates a new IoK8sApiStorageV1VolumeError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *IoK8sApiStorageV1VolumeError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiStorageV1VolumeError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiStorageV1VolumeError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiStorageV1VolumeError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTime

`func (o *IoK8sApiStorageV1VolumeError) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IoK8sApiStorageV1VolumeError) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IoK8sApiStorageV1VolumeError) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *IoK8sApiStorageV1VolumeError) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


