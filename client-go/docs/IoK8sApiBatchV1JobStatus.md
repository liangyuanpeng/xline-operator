# IoK8sApiBatchV1JobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | The number of pending and running pods. | [optional] 
**CompletedIndexes** | Pointer to **string** | completedIndexes holds the completed indexes when .spec.completionMode &#x3D; \&quot;Indexed\&quot; in a text format. The indexes are represented as decimal integers separated by commas. The numbers are listed in increasing order. Three or more consecutive numbers are compressed and represented by the first and last element of the series, separated by a hyphen. For example, if the completed indexes are 1, 3, 4, 5 and 7, they are represented as \&quot;1,3-5,7\&quot;. | [optional] 
**CompletionTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Conditions** | Pointer to [**[]IoK8sApiBatchV1JobCondition**](IoK8sApiBatchV1JobCondition.md) | The latest available observations of an object&#39;s current state. When a Job fails, one of the conditions will have type \&quot;Failed\&quot; and status true. When a Job is suspended, one of the conditions will have type \&quot;Suspended\&quot; and status true; when the Job is resumed, the status of this condition will become false. When a Job is completed, one of the conditions will have type \&quot;Complete\&quot; and status true. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/ | [optional] 
**Failed** | Pointer to **int32** | The number of pods which reached phase Failed. | [optional] 
**Ready** | Pointer to **int32** | The number of pods which have a Ready condition.  This field is beta-level. The job controller populates the field when the feature gate JobReadyPods is enabled (enabled by default). | [optional] 
**StartTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Succeeded** | Pointer to **int32** | The number of pods which reached phase Succeeded. | [optional] 
**UncountedTerminatedPods** | Pointer to [**IoK8sApiBatchV1UncountedTerminatedPods**](IoK8sApiBatchV1UncountedTerminatedPods.md) |  | [optional] 

## Methods

### NewIoK8sApiBatchV1JobStatus

`func NewIoK8sApiBatchV1JobStatus() *IoK8sApiBatchV1JobStatus`

NewIoK8sApiBatchV1JobStatus instantiates a new IoK8sApiBatchV1JobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiBatchV1JobStatusWithDefaults

`func NewIoK8sApiBatchV1JobStatusWithDefaults() *IoK8sApiBatchV1JobStatus`

NewIoK8sApiBatchV1JobStatusWithDefaults instantiates a new IoK8sApiBatchV1JobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *IoK8sApiBatchV1JobStatus) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *IoK8sApiBatchV1JobStatus) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *IoK8sApiBatchV1JobStatus) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *IoK8sApiBatchV1JobStatus) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCompletedIndexes

`func (o *IoK8sApiBatchV1JobStatus) GetCompletedIndexes() string`

GetCompletedIndexes returns the CompletedIndexes field if non-nil, zero value otherwise.

### GetCompletedIndexesOk

`func (o *IoK8sApiBatchV1JobStatus) GetCompletedIndexesOk() (*string, bool)`

GetCompletedIndexesOk returns a tuple with the CompletedIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedIndexes

`func (o *IoK8sApiBatchV1JobStatus) SetCompletedIndexes(v string)`

SetCompletedIndexes sets CompletedIndexes field to given value.

### HasCompletedIndexes

`func (o *IoK8sApiBatchV1JobStatus) HasCompletedIndexes() bool`

HasCompletedIndexes returns a boolean if a field has been set.

### GetCompletionTime

`func (o *IoK8sApiBatchV1JobStatus) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *IoK8sApiBatchV1JobStatus) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *IoK8sApiBatchV1JobStatus) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *IoK8sApiBatchV1JobStatus) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetConditions

`func (o *IoK8sApiBatchV1JobStatus) GetConditions() []IoK8sApiBatchV1JobCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiBatchV1JobStatus) GetConditionsOk() (*[]IoK8sApiBatchV1JobCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiBatchV1JobStatus) SetConditions(v []IoK8sApiBatchV1JobCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiBatchV1JobStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetFailed

`func (o *IoK8sApiBatchV1JobStatus) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *IoK8sApiBatchV1JobStatus) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *IoK8sApiBatchV1JobStatus) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *IoK8sApiBatchV1JobStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetReady

`func (o *IoK8sApiBatchV1JobStatus) GetReady() int32`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *IoK8sApiBatchV1JobStatus) GetReadyOk() (*int32, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *IoK8sApiBatchV1JobStatus) SetReady(v int32)`

SetReady sets Ready field to given value.

### HasReady

`func (o *IoK8sApiBatchV1JobStatus) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetStartTime

`func (o *IoK8sApiBatchV1JobStatus) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *IoK8sApiBatchV1JobStatus) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *IoK8sApiBatchV1JobStatus) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *IoK8sApiBatchV1JobStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSucceeded

`func (o *IoK8sApiBatchV1JobStatus) GetSucceeded() int32`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *IoK8sApiBatchV1JobStatus) GetSucceededOk() (*int32, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *IoK8sApiBatchV1JobStatus) SetSucceeded(v int32)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *IoK8sApiBatchV1JobStatus) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### GetUncountedTerminatedPods

`func (o *IoK8sApiBatchV1JobStatus) GetUncountedTerminatedPods() IoK8sApiBatchV1UncountedTerminatedPods`

GetUncountedTerminatedPods returns the UncountedTerminatedPods field if non-nil, zero value otherwise.

### GetUncountedTerminatedPodsOk

`func (o *IoK8sApiBatchV1JobStatus) GetUncountedTerminatedPodsOk() (*IoK8sApiBatchV1UncountedTerminatedPods, bool)`

GetUncountedTerminatedPodsOk returns a tuple with the UncountedTerminatedPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncountedTerminatedPods

`func (o *IoK8sApiBatchV1JobStatus) SetUncountedTerminatedPods(v IoK8sApiBatchV1UncountedTerminatedPods)`

SetUncountedTerminatedPods sets UncountedTerminatedPods field to given value.

### HasUncountedTerminatedPods

`func (o *IoK8sApiBatchV1JobStatus) HasUncountedTerminatedPods() bool`

HasUncountedTerminatedPods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


