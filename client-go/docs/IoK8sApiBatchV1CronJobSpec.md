# IoK8sApiBatchV1CronJobSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConcurrencyPolicy** | Pointer to **string** | Specifies how to treat concurrent executions of a Job. Valid values are:  - \&quot;Allow\&quot; (default): allows CronJobs to run concurrently; - \&quot;Forbid\&quot;: forbids concurrent runs, skipping next run if previous run hasn&#39;t finished yet; - \&quot;Replace\&quot;: cancels currently running job and replaces it with a new one  Possible enum values:  - &#x60;\&quot;Allow\&quot;&#x60; allows CronJobs to run concurrently.  - &#x60;\&quot;Forbid\&quot;&#x60; forbids concurrent runs, skipping next run if previous hasn&#39;t finished yet.  - &#x60;\&quot;Replace\&quot;&#x60; cancels currently running job and replaces it with a new one. | [optional] 
**FailedJobsHistoryLimit** | Pointer to **int32** | The number of failed finished jobs to retain. Value must be non-negative integer. Defaults to 1. | [optional] 
**JobTemplate** | [**IoK8sApiBatchV1JobTemplateSpec**](IoK8sApiBatchV1JobTemplateSpec.md) |  | 
**Schedule** | **string** | The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron. | 
**StartingDeadlineSeconds** | Pointer to **int64** | Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones. | [optional] 
**SuccessfulJobsHistoryLimit** | Pointer to **int32** | The number of successful finished jobs to retain. Value must be non-negative integer. Defaults to 3. | [optional] 
**Suspend** | Pointer to **bool** | This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false. | [optional] 
**TimeZone** | Pointer to **string** | The time zone name for the given schedule, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. If not specified, this will default to the time zone of the kube-controller-manager process. The set of valid time zone names and the time zone offset is loaded from the system-wide time zone database by the API server during CronJob validation and the controller manager during execution. If no system-wide time zone database can be found a bundled version of the database is used instead. If the time zone name becomes invalid during the lifetime of a CronJob or due to a change in host configuration, the controller will stop creating new new Jobs and will create a system event with the reason UnknownTimeZone. More information can be found in https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones | [optional] 

## Methods

### NewIoK8sApiBatchV1CronJobSpec

`func NewIoK8sApiBatchV1CronJobSpec(jobTemplate IoK8sApiBatchV1JobTemplateSpec, schedule string, ) *IoK8sApiBatchV1CronJobSpec`

NewIoK8sApiBatchV1CronJobSpec instantiates a new IoK8sApiBatchV1CronJobSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiBatchV1CronJobSpecWithDefaults

`func NewIoK8sApiBatchV1CronJobSpecWithDefaults() *IoK8sApiBatchV1CronJobSpec`

NewIoK8sApiBatchV1CronJobSpecWithDefaults instantiates a new IoK8sApiBatchV1CronJobSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrencyPolicy

`func (o *IoK8sApiBatchV1CronJobSpec) GetConcurrencyPolicy() string`

GetConcurrencyPolicy returns the ConcurrencyPolicy field if non-nil, zero value otherwise.

### GetConcurrencyPolicyOk

`func (o *IoK8sApiBatchV1CronJobSpec) GetConcurrencyPolicyOk() (*string, bool)`

GetConcurrencyPolicyOk returns a tuple with the ConcurrencyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyPolicy

`func (o *IoK8sApiBatchV1CronJobSpec) SetConcurrencyPolicy(v string)`

SetConcurrencyPolicy sets ConcurrencyPolicy field to given value.

### HasConcurrencyPolicy

`func (o *IoK8sApiBatchV1CronJobSpec) HasConcurrencyPolicy() bool`

HasConcurrencyPolicy returns a boolean if a field has been set.

### GetFailedJobsHistoryLimit

`func (o *IoK8sApiBatchV1CronJobSpec) GetFailedJobsHistoryLimit() int32`

GetFailedJobsHistoryLimit returns the FailedJobsHistoryLimit field if non-nil, zero value otherwise.

### GetFailedJobsHistoryLimitOk

`func (o *IoK8sApiBatchV1CronJobSpec) GetFailedJobsHistoryLimitOk() (*int32, bool)`

GetFailedJobsHistoryLimitOk returns a tuple with the FailedJobsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJobsHistoryLimit

`func (o *IoK8sApiBatchV1CronJobSpec) SetFailedJobsHistoryLimit(v int32)`

SetFailedJobsHistoryLimit sets FailedJobsHistoryLimit field to given value.

### HasFailedJobsHistoryLimit

`func (o *IoK8sApiBatchV1CronJobSpec) HasFailedJobsHistoryLimit() bool`

HasFailedJobsHistoryLimit returns a boolean if a field has been set.

### GetJobTemplate

`func (o *IoK8sApiBatchV1CronJobSpec) GetJobTemplate() IoK8sApiBatchV1JobTemplateSpec`

GetJobTemplate returns the JobTemplate field if non-nil, zero value otherwise.

### GetJobTemplateOk

`func (o *IoK8sApiBatchV1CronJobSpec) GetJobTemplateOk() (*IoK8sApiBatchV1JobTemplateSpec, bool)`

GetJobTemplateOk returns a tuple with the JobTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTemplate

`func (o *IoK8sApiBatchV1CronJobSpec) SetJobTemplate(v IoK8sApiBatchV1JobTemplateSpec)`

SetJobTemplate sets JobTemplate field to given value.


### GetSchedule

`func (o *IoK8sApiBatchV1CronJobSpec) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *IoK8sApiBatchV1CronJobSpec) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *IoK8sApiBatchV1CronJobSpec) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetStartingDeadlineSeconds

`func (o *IoK8sApiBatchV1CronJobSpec) GetStartingDeadlineSeconds() int64`

GetStartingDeadlineSeconds returns the StartingDeadlineSeconds field if non-nil, zero value otherwise.

### GetStartingDeadlineSecondsOk

`func (o *IoK8sApiBatchV1CronJobSpec) GetStartingDeadlineSecondsOk() (*int64, bool)`

GetStartingDeadlineSecondsOk returns a tuple with the StartingDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingDeadlineSeconds

`func (o *IoK8sApiBatchV1CronJobSpec) SetStartingDeadlineSeconds(v int64)`

SetStartingDeadlineSeconds sets StartingDeadlineSeconds field to given value.

### HasStartingDeadlineSeconds

`func (o *IoK8sApiBatchV1CronJobSpec) HasStartingDeadlineSeconds() bool`

HasStartingDeadlineSeconds returns a boolean if a field has been set.

### GetSuccessfulJobsHistoryLimit

`func (o *IoK8sApiBatchV1CronJobSpec) GetSuccessfulJobsHistoryLimit() int32`

GetSuccessfulJobsHistoryLimit returns the SuccessfulJobsHistoryLimit field if non-nil, zero value otherwise.

### GetSuccessfulJobsHistoryLimitOk

`func (o *IoK8sApiBatchV1CronJobSpec) GetSuccessfulJobsHistoryLimitOk() (*int32, bool)`

GetSuccessfulJobsHistoryLimitOk returns a tuple with the SuccessfulJobsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulJobsHistoryLimit

`func (o *IoK8sApiBatchV1CronJobSpec) SetSuccessfulJobsHistoryLimit(v int32)`

SetSuccessfulJobsHistoryLimit sets SuccessfulJobsHistoryLimit field to given value.

### HasSuccessfulJobsHistoryLimit

`func (o *IoK8sApiBatchV1CronJobSpec) HasSuccessfulJobsHistoryLimit() bool`

HasSuccessfulJobsHistoryLimit returns a boolean if a field has been set.

### GetSuspend

`func (o *IoK8sApiBatchV1CronJobSpec) GetSuspend() bool`

GetSuspend returns the Suspend field if non-nil, zero value otherwise.

### GetSuspendOk

`func (o *IoK8sApiBatchV1CronJobSpec) GetSuspendOk() (*bool, bool)`

GetSuspendOk returns a tuple with the Suspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspend

`func (o *IoK8sApiBatchV1CronJobSpec) SetSuspend(v bool)`

SetSuspend sets Suspend field to given value.

### HasSuspend

`func (o *IoK8sApiBatchV1CronJobSpec) HasSuspend() bool`

HasSuspend returns a boolean if a field has been set.

### GetTimeZone

`func (o *IoK8sApiBatchV1CronJobSpec) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *IoK8sApiBatchV1CronJobSpec) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *IoK8sApiBatchV1CronJobSpec) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *IoK8sApiBatchV1CronJobSpec) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


