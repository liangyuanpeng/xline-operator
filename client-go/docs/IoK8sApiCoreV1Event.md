# IoK8sApiCoreV1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | What action was taken/failed regarding to the Regarding object. | [optional] 
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Count** | Pointer to **int32** | The number of times this event has occurred. | [optional] 
**EventTime** | Pointer to **time.Time** | MicroTime is version of Time with microsecond level precision. | [optional] 
**FirstTimestamp** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**InvolvedObject** | [**IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) |  | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**LastTimestamp** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | A human-readable description of the status of this operation. | [optional] 
**Metadata** | [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | 
**Reason** | Pointer to **string** | This should be a short, machine understandable string that gives the reason for the transition into the object&#39;s current status. | [optional] 
**Related** | Pointer to [**IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) |  | [optional] 
**ReportingComponent** | Pointer to **string** | Name of the controller that emitted this Event, e.g. &#x60;kubernetes.io/kubelet&#x60;. | [optional] 
**ReportingInstance** | Pointer to **string** | ID of the controller instance, e.g. &#x60;kubelet-xyzf&#x60;. | [optional] 
**Series** | Pointer to [**IoK8sApiCoreV1EventSeries**](IoK8sApiCoreV1EventSeries.md) |  | [optional] 
**Source** | Pointer to [**IoK8sApiCoreV1EventSource**](IoK8sApiCoreV1EventSource.md) |  | [optional] 
**Type** | Pointer to **string** | Type of this event (Normal, Warning), new types could be added in the future | [optional] 

## Methods

### NewIoK8sApiCoreV1Event

`func NewIoK8sApiCoreV1Event(involvedObject IoK8sApiCoreV1ObjectReference, metadata IoK8sApimachineryPkgApisMetaV1ObjectMeta, ) *IoK8sApiCoreV1Event`

NewIoK8sApiCoreV1Event instantiates a new IoK8sApiCoreV1Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1EventWithDefaults

`func NewIoK8sApiCoreV1EventWithDefaults() *IoK8sApiCoreV1Event`

NewIoK8sApiCoreV1EventWithDefaults instantiates a new IoK8sApiCoreV1Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IoK8sApiCoreV1Event) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IoK8sApiCoreV1Event) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IoK8sApiCoreV1Event) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IoK8sApiCoreV1Event) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetApiVersion

`func (o *IoK8sApiCoreV1Event) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiCoreV1Event) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiCoreV1Event) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiCoreV1Event) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCount

`func (o *IoK8sApiCoreV1Event) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IoK8sApiCoreV1Event) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IoK8sApiCoreV1Event) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IoK8sApiCoreV1Event) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetEventTime

`func (o *IoK8sApiCoreV1Event) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *IoK8sApiCoreV1Event) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *IoK8sApiCoreV1Event) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *IoK8sApiCoreV1Event) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetFirstTimestamp

`func (o *IoK8sApiCoreV1Event) GetFirstTimestamp() time.Time`

GetFirstTimestamp returns the FirstTimestamp field if non-nil, zero value otherwise.

### GetFirstTimestampOk

`func (o *IoK8sApiCoreV1Event) GetFirstTimestampOk() (*time.Time, bool)`

GetFirstTimestampOk returns a tuple with the FirstTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimestamp

`func (o *IoK8sApiCoreV1Event) SetFirstTimestamp(v time.Time)`

SetFirstTimestamp sets FirstTimestamp field to given value.

### HasFirstTimestamp

`func (o *IoK8sApiCoreV1Event) HasFirstTimestamp() bool`

HasFirstTimestamp returns a boolean if a field has been set.

### GetInvolvedObject

`func (o *IoK8sApiCoreV1Event) GetInvolvedObject() IoK8sApiCoreV1ObjectReference`

GetInvolvedObject returns the InvolvedObject field if non-nil, zero value otherwise.

### GetInvolvedObjectOk

`func (o *IoK8sApiCoreV1Event) GetInvolvedObjectOk() (*IoK8sApiCoreV1ObjectReference, bool)`

GetInvolvedObjectOk returns a tuple with the InvolvedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedObject

`func (o *IoK8sApiCoreV1Event) SetInvolvedObject(v IoK8sApiCoreV1ObjectReference)`

SetInvolvedObject sets InvolvedObject field to given value.


### GetKind

`func (o *IoK8sApiCoreV1Event) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiCoreV1Event) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiCoreV1Event) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiCoreV1Event) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastTimestamp

`func (o *IoK8sApiCoreV1Event) GetLastTimestamp() time.Time`

GetLastTimestamp returns the LastTimestamp field if non-nil, zero value otherwise.

### GetLastTimestampOk

`func (o *IoK8sApiCoreV1Event) GetLastTimestampOk() (*time.Time, bool)`

GetLastTimestampOk returns a tuple with the LastTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimestamp

`func (o *IoK8sApiCoreV1Event) SetLastTimestamp(v time.Time)`

SetLastTimestamp sets LastTimestamp field to given value.

### HasLastTimestamp

`func (o *IoK8sApiCoreV1Event) HasLastTimestamp() bool`

HasLastTimestamp returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiCoreV1Event) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiCoreV1Event) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiCoreV1Event) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiCoreV1Event) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiCoreV1Event) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiCoreV1Event) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiCoreV1Event) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetReason

`func (o *IoK8sApiCoreV1Event) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiCoreV1Event) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiCoreV1Event) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiCoreV1Event) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRelated

`func (o *IoK8sApiCoreV1Event) GetRelated() IoK8sApiCoreV1ObjectReference`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *IoK8sApiCoreV1Event) GetRelatedOk() (*IoK8sApiCoreV1ObjectReference, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *IoK8sApiCoreV1Event) SetRelated(v IoK8sApiCoreV1ObjectReference)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *IoK8sApiCoreV1Event) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetReportingComponent

`func (o *IoK8sApiCoreV1Event) GetReportingComponent() string`

GetReportingComponent returns the ReportingComponent field if non-nil, zero value otherwise.

### GetReportingComponentOk

`func (o *IoK8sApiCoreV1Event) GetReportingComponentOk() (*string, bool)`

GetReportingComponentOk returns a tuple with the ReportingComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingComponent

`func (o *IoK8sApiCoreV1Event) SetReportingComponent(v string)`

SetReportingComponent sets ReportingComponent field to given value.

### HasReportingComponent

`func (o *IoK8sApiCoreV1Event) HasReportingComponent() bool`

HasReportingComponent returns a boolean if a field has been set.

### GetReportingInstance

`func (o *IoK8sApiCoreV1Event) GetReportingInstance() string`

GetReportingInstance returns the ReportingInstance field if non-nil, zero value otherwise.

### GetReportingInstanceOk

`func (o *IoK8sApiCoreV1Event) GetReportingInstanceOk() (*string, bool)`

GetReportingInstanceOk returns a tuple with the ReportingInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingInstance

`func (o *IoK8sApiCoreV1Event) SetReportingInstance(v string)`

SetReportingInstance sets ReportingInstance field to given value.

### HasReportingInstance

`func (o *IoK8sApiCoreV1Event) HasReportingInstance() bool`

HasReportingInstance returns a boolean if a field has been set.

### GetSeries

`func (o *IoK8sApiCoreV1Event) GetSeries() IoK8sApiCoreV1EventSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *IoK8sApiCoreV1Event) GetSeriesOk() (*IoK8sApiCoreV1EventSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *IoK8sApiCoreV1Event) SetSeries(v IoK8sApiCoreV1EventSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *IoK8sApiCoreV1Event) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetSource

`func (o *IoK8sApiCoreV1Event) GetSource() IoK8sApiCoreV1EventSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IoK8sApiCoreV1Event) GetSourceOk() (*IoK8sApiCoreV1EventSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IoK8sApiCoreV1Event) SetSource(v IoK8sApiCoreV1EventSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *IoK8sApiCoreV1Event) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiCoreV1Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiCoreV1Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiCoreV1Event) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiCoreV1Event) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


