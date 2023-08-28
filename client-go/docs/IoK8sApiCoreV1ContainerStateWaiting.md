# IoK8sApiCoreV1ContainerStateWaiting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message regarding why the container is not yet running. | [optional] 
**Reason** | Pointer to **string** | (brief) reason the container is not yet running. | [optional] 

## Methods

### NewIoK8sApiCoreV1ContainerStateWaiting

`func NewIoK8sApiCoreV1ContainerStateWaiting() *IoK8sApiCoreV1ContainerStateWaiting`

NewIoK8sApiCoreV1ContainerStateWaiting instantiates a new IoK8sApiCoreV1ContainerStateWaiting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ContainerStateWaitingWithDefaults

`func NewIoK8sApiCoreV1ContainerStateWaitingWithDefaults() *IoK8sApiCoreV1ContainerStateWaiting`

NewIoK8sApiCoreV1ContainerStateWaitingWithDefaults instantiates a new IoK8sApiCoreV1ContainerStateWaiting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *IoK8sApiCoreV1ContainerStateWaiting) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiCoreV1ContainerStateWaiting) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiCoreV1ContainerStateWaiting) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiCoreV1ContainerStateWaiting) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiCoreV1ContainerStateWaiting) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiCoreV1ContainerStateWaiting) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiCoreV1ContainerStateWaiting) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiCoreV1ContainerStateWaiting) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


