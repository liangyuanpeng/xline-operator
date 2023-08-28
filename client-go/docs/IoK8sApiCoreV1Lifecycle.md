# IoK8sApiCoreV1Lifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostStart** | Pointer to [**IoK8sApiCoreV1LifecycleHandler**](IoK8sApiCoreV1LifecycleHandler.md) |  | [optional] 
**PreStop** | Pointer to [**IoK8sApiCoreV1LifecycleHandler**](IoK8sApiCoreV1LifecycleHandler.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1Lifecycle

`func NewIoK8sApiCoreV1Lifecycle() *IoK8sApiCoreV1Lifecycle`

NewIoK8sApiCoreV1Lifecycle instantiates a new IoK8sApiCoreV1Lifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1LifecycleWithDefaults

`func NewIoK8sApiCoreV1LifecycleWithDefaults() *IoK8sApiCoreV1Lifecycle`

NewIoK8sApiCoreV1LifecycleWithDefaults instantiates a new IoK8sApiCoreV1Lifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostStart

`func (o *IoK8sApiCoreV1Lifecycle) GetPostStart() IoK8sApiCoreV1LifecycleHandler`

GetPostStart returns the PostStart field if non-nil, zero value otherwise.

### GetPostStartOk

`func (o *IoK8sApiCoreV1Lifecycle) GetPostStartOk() (*IoK8sApiCoreV1LifecycleHandler, bool)`

GetPostStartOk returns a tuple with the PostStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStart

`func (o *IoK8sApiCoreV1Lifecycle) SetPostStart(v IoK8sApiCoreV1LifecycleHandler)`

SetPostStart sets PostStart field to given value.

### HasPostStart

`func (o *IoK8sApiCoreV1Lifecycle) HasPostStart() bool`

HasPostStart returns a boolean if a field has been set.

### GetPreStop

`func (o *IoK8sApiCoreV1Lifecycle) GetPreStop() IoK8sApiCoreV1LifecycleHandler`

GetPreStop returns the PreStop field if non-nil, zero value otherwise.

### GetPreStopOk

`func (o *IoK8sApiCoreV1Lifecycle) GetPreStopOk() (*IoK8sApiCoreV1LifecycleHandler, bool)`

GetPreStopOk returns a tuple with the PreStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStop

`func (o *IoK8sApiCoreV1Lifecycle) SetPreStop(v IoK8sApiCoreV1LifecycleHandler)`

SetPreStop sets PreStop field to given value.

### HasPreStop

`func (o *IoK8sApiCoreV1Lifecycle) HasPreStop() bool`

HasPreStop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


