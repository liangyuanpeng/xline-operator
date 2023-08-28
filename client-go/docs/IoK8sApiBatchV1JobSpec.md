# IoK8sApiBatchV1JobSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDeadlineSeconds** | Pointer to **int64** | Specifies the duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it; value must be positive integer. If a Job is suspended (at creation or through an update), this timer will effectively be stopped and reset when the Job is resumed again. | [optional] 
**BackoffLimit** | Pointer to **int32** | Specifies the number of retries before marking this job failed. Defaults to 6 | [optional] 
**CompletionMode** | Pointer to **string** | completionMode specifies how Pod completions are tracked. It can be &#x60;NonIndexed&#x60; (default) or &#x60;Indexed&#x60;.  &#x60;NonIndexed&#x60; means that the Job is considered complete when there have been .spec.completions successfully completed Pods. Each Pod completion is homologous to each other.  &#x60;Indexed&#x60; means that the Pods of a Job get an associated completion index from 0 to (.spec.completions - 1), available in the annotation batch.kubernetes.io/job-completion-index. The Job is considered complete when there is one successfully completed Pod for each index. When value is &#x60;Indexed&#x60;, .spec.completions must be specified and &#x60;.spec.parallelism&#x60; must be less than or equal to 10^5. In addition, The Pod name takes the form &#x60;$(job-name)-$(index)-$(random-string)&#x60;, the Pod hostname takes the form &#x60;$(job-name)-$(index)&#x60;.  More completion modes can be added in the future. If the Job controller observes a mode that it doesn&#39;t recognize, which is possible during upgrades due to version skew, the controller skips updates for the Job.  Possible enum values:  - &#x60;\&quot;Indexed\&quot;&#x60; is a Job completion mode. In this mode, the Pods of a Job get an associated completion index from 0 to (.spec.completions - 1). The Job is considered complete when a Pod completes for each completion index.  - &#x60;\&quot;NonIndexed\&quot;&#x60; is a Job completion mode. In this mode, the Job is considered complete when there have been .spec.completions successfully completed Pods. Pod completions are homologous to each other. | [optional] 
**Completions** | Pointer to **int32** | Specifies the desired number of successfully finished pods the job should be run with.  Setting to null means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/ | [optional] 
**ManualSelector** | Pointer to **bool** | manualSelector controls generation of pod labels and pod selectors. Leave &#x60;manualSelector&#x60; unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see &#x60;manualSelector&#x3D;true&#x60; in jobs that were created with the old &#x60;extensions/v1beta1&#x60; API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector | [optional] 
**Parallelism** | Pointer to **int32** | Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) &lt; .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/ | [optional] 
**PodFailurePolicy** | Pointer to [**IoK8sApiBatchV1PodFailurePolicy**](IoK8sApiBatchV1PodFailurePolicy.md) |  | [optional] 
**Selector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**Suspend** | Pointer to **bool** | suspend specifies whether the Job controller should create Pods or not. If a Job is created with suspend set to true, no Pods are created by the Job controller. If a Job is suspended after creation (i.e. the flag goes from false to true), the Job controller will delete all active Pods associated with this Job. Users must design their workload to gracefully handle this. Suspending a Job will reset the StartTime field of the Job, effectively resetting the ActiveDeadlineSeconds timer too. Defaults to false. | [optional] 
**Template** | [**IoK8sApiCoreV1PodTemplateSpec**](IoK8sApiCoreV1PodTemplateSpec.md) |  | 
**TtlSecondsAfterFinished** | Pointer to **int32** | ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won&#39;t be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. | [optional] 

## Methods

### NewIoK8sApiBatchV1JobSpec

`func NewIoK8sApiBatchV1JobSpec(template IoK8sApiCoreV1PodTemplateSpec, ) *IoK8sApiBatchV1JobSpec`

NewIoK8sApiBatchV1JobSpec instantiates a new IoK8sApiBatchV1JobSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiBatchV1JobSpecWithDefaults

`func NewIoK8sApiBatchV1JobSpecWithDefaults() *IoK8sApiBatchV1JobSpec`

NewIoK8sApiBatchV1JobSpecWithDefaults instantiates a new IoK8sApiBatchV1JobSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDeadlineSeconds

`func (o *IoK8sApiBatchV1JobSpec) GetActiveDeadlineSeconds() int64`

GetActiveDeadlineSeconds returns the ActiveDeadlineSeconds field if non-nil, zero value otherwise.

### GetActiveDeadlineSecondsOk

`func (o *IoK8sApiBatchV1JobSpec) GetActiveDeadlineSecondsOk() (*int64, bool)`

GetActiveDeadlineSecondsOk returns a tuple with the ActiveDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeadlineSeconds

`func (o *IoK8sApiBatchV1JobSpec) SetActiveDeadlineSeconds(v int64)`

SetActiveDeadlineSeconds sets ActiveDeadlineSeconds field to given value.

### HasActiveDeadlineSeconds

`func (o *IoK8sApiBatchV1JobSpec) HasActiveDeadlineSeconds() bool`

HasActiveDeadlineSeconds returns a boolean if a field has been set.

### GetBackoffLimit

`func (o *IoK8sApiBatchV1JobSpec) GetBackoffLimit() int32`

GetBackoffLimit returns the BackoffLimit field if non-nil, zero value otherwise.

### GetBackoffLimitOk

`func (o *IoK8sApiBatchV1JobSpec) GetBackoffLimitOk() (*int32, bool)`

GetBackoffLimitOk returns a tuple with the BackoffLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoffLimit

`func (o *IoK8sApiBatchV1JobSpec) SetBackoffLimit(v int32)`

SetBackoffLimit sets BackoffLimit field to given value.

### HasBackoffLimit

`func (o *IoK8sApiBatchV1JobSpec) HasBackoffLimit() bool`

HasBackoffLimit returns a boolean if a field has been set.

### GetCompletionMode

`func (o *IoK8sApiBatchV1JobSpec) GetCompletionMode() string`

GetCompletionMode returns the CompletionMode field if non-nil, zero value otherwise.

### GetCompletionModeOk

`func (o *IoK8sApiBatchV1JobSpec) GetCompletionModeOk() (*string, bool)`

GetCompletionModeOk returns a tuple with the CompletionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionMode

`func (o *IoK8sApiBatchV1JobSpec) SetCompletionMode(v string)`

SetCompletionMode sets CompletionMode field to given value.

### HasCompletionMode

`func (o *IoK8sApiBatchV1JobSpec) HasCompletionMode() bool`

HasCompletionMode returns a boolean if a field has been set.

### GetCompletions

`func (o *IoK8sApiBatchV1JobSpec) GetCompletions() int32`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *IoK8sApiBatchV1JobSpec) GetCompletionsOk() (*int32, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *IoK8sApiBatchV1JobSpec) SetCompletions(v int32)`

SetCompletions sets Completions field to given value.

### HasCompletions

`func (o *IoK8sApiBatchV1JobSpec) HasCompletions() bool`

HasCompletions returns a boolean if a field has been set.

### GetManualSelector

`func (o *IoK8sApiBatchV1JobSpec) GetManualSelector() bool`

GetManualSelector returns the ManualSelector field if non-nil, zero value otherwise.

### GetManualSelectorOk

`func (o *IoK8sApiBatchV1JobSpec) GetManualSelectorOk() (*bool, bool)`

GetManualSelectorOk returns a tuple with the ManualSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualSelector

`func (o *IoK8sApiBatchV1JobSpec) SetManualSelector(v bool)`

SetManualSelector sets ManualSelector field to given value.

### HasManualSelector

`func (o *IoK8sApiBatchV1JobSpec) HasManualSelector() bool`

HasManualSelector returns a boolean if a field has been set.

### GetParallelism

`func (o *IoK8sApiBatchV1JobSpec) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *IoK8sApiBatchV1JobSpec) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *IoK8sApiBatchV1JobSpec) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *IoK8sApiBatchV1JobSpec) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetPodFailurePolicy

`func (o *IoK8sApiBatchV1JobSpec) GetPodFailurePolicy() IoK8sApiBatchV1PodFailurePolicy`

GetPodFailurePolicy returns the PodFailurePolicy field if non-nil, zero value otherwise.

### GetPodFailurePolicyOk

`func (o *IoK8sApiBatchV1JobSpec) GetPodFailurePolicyOk() (*IoK8sApiBatchV1PodFailurePolicy, bool)`

GetPodFailurePolicyOk returns a tuple with the PodFailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodFailurePolicy

`func (o *IoK8sApiBatchV1JobSpec) SetPodFailurePolicy(v IoK8sApiBatchV1PodFailurePolicy)`

SetPodFailurePolicy sets PodFailurePolicy field to given value.

### HasPodFailurePolicy

`func (o *IoK8sApiBatchV1JobSpec) HasPodFailurePolicy() bool`

HasPodFailurePolicy returns a boolean if a field has been set.

### GetSelector

`func (o *IoK8sApiBatchV1JobSpec) GetSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiBatchV1JobSpec) GetSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiBatchV1JobSpec) SetSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *IoK8sApiBatchV1JobSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSuspend

`func (o *IoK8sApiBatchV1JobSpec) GetSuspend() bool`

GetSuspend returns the Suspend field if non-nil, zero value otherwise.

### GetSuspendOk

`func (o *IoK8sApiBatchV1JobSpec) GetSuspendOk() (*bool, bool)`

GetSuspendOk returns a tuple with the Suspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspend

`func (o *IoK8sApiBatchV1JobSpec) SetSuspend(v bool)`

SetSuspend sets Suspend field to given value.

### HasSuspend

`func (o *IoK8sApiBatchV1JobSpec) HasSuspend() bool`

HasSuspend returns a boolean if a field has been set.

### GetTemplate

`func (o *IoK8sApiBatchV1JobSpec) GetTemplate() IoK8sApiCoreV1PodTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoK8sApiBatchV1JobSpec) GetTemplateOk() (*IoK8sApiCoreV1PodTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoK8sApiBatchV1JobSpec) SetTemplate(v IoK8sApiCoreV1PodTemplateSpec)`

SetTemplate sets Template field to given value.


### GetTtlSecondsAfterFinished

`func (o *IoK8sApiBatchV1JobSpec) GetTtlSecondsAfterFinished() int32`

GetTtlSecondsAfterFinished returns the TtlSecondsAfterFinished field if non-nil, zero value otherwise.

### GetTtlSecondsAfterFinishedOk

`func (o *IoK8sApiBatchV1JobSpec) GetTtlSecondsAfterFinishedOk() (*int32, bool)`

GetTtlSecondsAfterFinishedOk returns a tuple with the TtlSecondsAfterFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSecondsAfterFinished

`func (o *IoK8sApiBatchV1JobSpec) SetTtlSecondsAfterFinished(v int32)`

SetTtlSecondsAfterFinished sets TtlSecondsAfterFinished field to given value.

### HasTtlSecondsAfterFinished

`func (o *IoK8sApiBatchV1JobSpec) HasTtlSecondsAfterFinished() bool`

HasTtlSecondsAfterFinished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


