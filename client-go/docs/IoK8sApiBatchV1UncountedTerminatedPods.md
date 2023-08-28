# IoK8sApiBatchV1UncountedTerminatedPods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | Pointer to **[]string** | failed holds UIDs of failed Pods. | [optional] 
**Succeeded** | Pointer to **[]string** | succeeded holds UIDs of succeeded Pods. | [optional] 

## Methods

### NewIoK8sApiBatchV1UncountedTerminatedPods

`func NewIoK8sApiBatchV1UncountedTerminatedPods() *IoK8sApiBatchV1UncountedTerminatedPods`

NewIoK8sApiBatchV1UncountedTerminatedPods instantiates a new IoK8sApiBatchV1UncountedTerminatedPods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiBatchV1UncountedTerminatedPodsWithDefaults

`func NewIoK8sApiBatchV1UncountedTerminatedPodsWithDefaults() *IoK8sApiBatchV1UncountedTerminatedPods`

NewIoK8sApiBatchV1UncountedTerminatedPodsWithDefaults instantiates a new IoK8sApiBatchV1UncountedTerminatedPods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *IoK8sApiBatchV1UncountedTerminatedPods) GetFailed() []string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *IoK8sApiBatchV1UncountedTerminatedPods) GetFailedOk() (*[]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *IoK8sApiBatchV1UncountedTerminatedPods) SetFailed(v []string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *IoK8sApiBatchV1UncountedTerminatedPods) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetSucceeded

`func (o *IoK8sApiBatchV1UncountedTerminatedPods) GetSucceeded() []string`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *IoK8sApiBatchV1UncountedTerminatedPods) GetSucceededOk() (*[]string, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *IoK8sApiBatchV1UncountedTerminatedPods) SetSucceeded(v []string)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *IoK8sApiBatchV1UncountedTerminatedPods) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


