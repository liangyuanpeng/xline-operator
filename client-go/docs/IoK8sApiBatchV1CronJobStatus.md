# IoK8sApiBatchV1CronJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to [**[]IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) | A list of pointers to currently running jobs. | [optional] 
**LastScheduleTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**LastSuccessfulTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 

## Methods

### NewIoK8sApiBatchV1CronJobStatus

`func NewIoK8sApiBatchV1CronJobStatus() *IoK8sApiBatchV1CronJobStatus`

NewIoK8sApiBatchV1CronJobStatus instantiates a new IoK8sApiBatchV1CronJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiBatchV1CronJobStatusWithDefaults

`func NewIoK8sApiBatchV1CronJobStatusWithDefaults() *IoK8sApiBatchV1CronJobStatus`

NewIoK8sApiBatchV1CronJobStatusWithDefaults instantiates a new IoK8sApiBatchV1CronJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *IoK8sApiBatchV1CronJobStatus) GetActive() []IoK8sApiCoreV1ObjectReference`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *IoK8sApiBatchV1CronJobStatus) GetActiveOk() (*[]IoK8sApiCoreV1ObjectReference, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *IoK8sApiBatchV1CronJobStatus) SetActive(v []IoK8sApiCoreV1ObjectReference)`

SetActive sets Active field to given value.

### HasActive

`func (o *IoK8sApiBatchV1CronJobStatus) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastScheduleTime

`func (o *IoK8sApiBatchV1CronJobStatus) GetLastScheduleTime() time.Time`

GetLastScheduleTime returns the LastScheduleTime field if non-nil, zero value otherwise.

### GetLastScheduleTimeOk

`func (o *IoK8sApiBatchV1CronJobStatus) GetLastScheduleTimeOk() (*time.Time, bool)`

GetLastScheduleTimeOk returns a tuple with the LastScheduleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScheduleTime

`func (o *IoK8sApiBatchV1CronJobStatus) SetLastScheduleTime(v time.Time)`

SetLastScheduleTime sets LastScheduleTime field to given value.

### HasLastScheduleTime

`func (o *IoK8sApiBatchV1CronJobStatus) HasLastScheduleTime() bool`

HasLastScheduleTime returns a boolean if a field has been set.

### GetLastSuccessfulTime

`func (o *IoK8sApiBatchV1CronJobStatus) GetLastSuccessfulTime() time.Time`

GetLastSuccessfulTime returns the LastSuccessfulTime field if non-nil, zero value otherwise.

### GetLastSuccessfulTimeOk

`func (o *IoK8sApiBatchV1CronJobStatus) GetLastSuccessfulTimeOk() (*time.Time, bool)`

GetLastSuccessfulTimeOk returns a tuple with the LastSuccessfulTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulTime

`func (o *IoK8sApiBatchV1CronJobStatus) SetLastSuccessfulTime(v time.Time)`

SetLastSuccessfulTime sets LastSuccessfulTime field to given value.

### HasLastSuccessfulTime

`func (o *IoK8sApiBatchV1CronJobStatus) HasLastSuccessfulTime() bool`

HasLastSuccessfulTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


