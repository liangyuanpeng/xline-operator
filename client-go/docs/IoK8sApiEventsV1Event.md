# IoK8sApiEventsV1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters. | [optional] 
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**DeprecatedCount** | Pointer to **int32** | deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type. | [optional] 
**DeprecatedFirstTimestamp** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**DeprecatedLastTimestamp** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**DeprecatedSource** | Pointer to [**IoK8sApiCoreV1EventSource**](IoK8sApiCoreV1EventSource.md) |  | [optional] 
**EventTime** | **time.Time** | MicroTime is version of Time with microsecond level precision. | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Note** | Pointer to **string** | note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB. | [optional] 
**Reason** | Pointer to **string** | reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters. | [optional] 
**Regarding** | Pointer to [**IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) |  | [optional] 
**Related** | Pointer to [**IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) |  | [optional] 
**ReportingController** | Pointer to **string** | reportingController is the name of the controller that emitted this Event, e.g. &#x60;kubernetes.io/kubelet&#x60;. This field cannot be empty for new Events. | [optional] 
**ReportingInstance** | Pointer to **string** | reportingInstance is the ID of the controller instance, e.g. &#x60;kubelet-xyzf&#x60;. This field cannot be empty for new Events and it can have at most 128 characters. | [optional] 
**Series** | Pointer to [**IoK8sApiEventsV1EventSeries**](IoK8sApiEventsV1EventSeries.md) |  | [optional] 
**Type** | Pointer to **string** | type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events. | [optional] 

## Methods

### NewIoK8sApiEventsV1Event

`func NewIoK8sApiEventsV1Event(eventTime time.Time, ) *IoK8sApiEventsV1Event`

NewIoK8sApiEventsV1Event instantiates a new IoK8sApiEventsV1Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiEventsV1EventWithDefaults

`func NewIoK8sApiEventsV1EventWithDefaults() *IoK8sApiEventsV1Event`

NewIoK8sApiEventsV1EventWithDefaults instantiates a new IoK8sApiEventsV1Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IoK8sApiEventsV1Event) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IoK8sApiEventsV1Event) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IoK8sApiEventsV1Event) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IoK8sApiEventsV1Event) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetApiVersion

`func (o *IoK8sApiEventsV1Event) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiEventsV1Event) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiEventsV1Event) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiEventsV1Event) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDeprecatedCount

`func (o *IoK8sApiEventsV1Event) GetDeprecatedCount() int32`

GetDeprecatedCount returns the DeprecatedCount field if non-nil, zero value otherwise.

### GetDeprecatedCountOk

`func (o *IoK8sApiEventsV1Event) GetDeprecatedCountOk() (*int32, bool)`

GetDeprecatedCountOk returns a tuple with the DeprecatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedCount

`func (o *IoK8sApiEventsV1Event) SetDeprecatedCount(v int32)`

SetDeprecatedCount sets DeprecatedCount field to given value.

### HasDeprecatedCount

`func (o *IoK8sApiEventsV1Event) HasDeprecatedCount() bool`

HasDeprecatedCount returns a boolean if a field has been set.

### GetDeprecatedFirstTimestamp

`func (o *IoK8sApiEventsV1Event) GetDeprecatedFirstTimestamp() time.Time`

GetDeprecatedFirstTimestamp returns the DeprecatedFirstTimestamp field if non-nil, zero value otherwise.

### GetDeprecatedFirstTimestampOk

`func (o *IoK8sApiEventsV1Event) GetDeprecatedFirstTimestampOk() (*time.Time, bool)`

GetDeprecatedFirstTimestampOk returns a tuple with the DeprecatedFirstTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedFirstTimestamp

`func (o *IoK8sApiEventsV1Event) SetDeprecatedFirstTimestamp(v time.Time)`

SetDeprecatedFirstTimestamp sets DeprecatedFirstTimestamp field to given value.

### HasDeprecatedFirstTimestamp

`func (o *IoK8sApiEventsV1Event) HasDeprecatedFirstTimestamp() bool`

HasDeprecatedFirstTimestamp returns a boolean if a field has been set.

### GetDeprecatedLastTimestamp

`func (o *IoK8sApiEventsV1Event) GetDeprecatedLastTimestamp() time.Time`

GetDeprecatedLastTimestamp returns the DeprecatedLastTimestamp field if non-nil, zero value otherwise.

### GetDeprecatedLastTimestampOk

`func (o *IoK8sApiEventsV1Event) GetDeprecatedLastTimestampOk() (*time.Time, bool)`

GetDeprecatedLastTimestampOk returns a tuple with the DeprecatedLastTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedLastTimestamp

`func (o *IoK8sApiEventsV1Event) SetDeprecatedLastTimestamp(v time.Time)`

SetDeprecatedLastTimestamp sets DeprecatedLastTimestamp field to given value.

### HasDeprecatedLastTimestamp

`func (o *IoK8sApiEventsV1Event) HasDeprecatedLastTimestamp() bool`

HasDeprecatedLastTimestamp returns a boolean if a field has been set.

### GetDeprecatedSource

`func (o *IoK8sApiEventsV1Event) GetDeprecatedSource() IoK8sApiCoreV1EventSource`

GetDeprecatedSource returns the DeprecatedSource field if non-nil, zero value otherwise.

### GetDeprecatedSourceOk

`func (o *IoK8sApiEventsV1Event) GetDeprecatedSourceOk() (*IoK8sApiCoreV1EventSource, bool)`

GetDeprecatedSourceOk returns a tuple with the DeprecatedSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedSource

`func (o *IoK8sApiEventsV1Event) SetDeprecatedSource(v IoK8sApiCoreV1EventSource)`

SetDeprecatedSource sets DeprecatedSource field to given value.

### HasDeprecatedSource

`func (o *IoK8sApiEventsV1Event) HasDeprecatedSource() bool`

HasDeprecatedSource returns a boolean if a field has been set.

### GetEventTime

`func (o *IoK8sApiEventsV1Event) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *IoK8sApiEventsV1Event) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *IoK8sApiEventsV1Event) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetKind

`func (o *IoK8sApiEventsV1Event) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiEventsV1Event) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiEventsV1Event) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiEventsV1Event) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiEventsV1Event) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiEventsV1Event) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiEventsV1Event) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiEventsV1Event) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNote

`func (o *IoK8sApiEventsV1Event) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *IoK8sApiEventsV1Event) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *IoK8sApiEventsV1Event) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *IoK8sApiEventsV1Event) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiEventsV1Event) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiEventsV1Event) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiEventsV1Event) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiEventsV1Event) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRegarding

`func (o *IoK8sApiEventsV1Event) GetRegarding() IoK8sApiCoreV1ObjectReference`

GetRegarding returns the Regarding field if non-nil, zero value otherwise.

### GetRegardingOk

`func (o *IoK8sApiEventsV1Event) GetRegardingOk() (*IoK8sApiCoreV1ObjectReference, bool)`

GetRegardingOk returns a tuple with the Regarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegarding

`func (o *IoK8sApiEventsV1Event) SetRegarding(v IoK8sApiCoreV1ObjectReference)`

SetRegarding sets Regarding field to given value.

### HasRegarding

`func (o *IoK8sApiEventsV1Event) HasRegarding() bool`

HasRegarding returns a boolean if a field has been set.

### GetRelated

`func (o *IoK8sApiEventsV1Event) GetRelated() IoK8sApiCoreV1ObjectReference`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *IoK8sApiEventsV1Event) GetRelatedOk() (*IoK8sApiCoreV1ObjectReference, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *IoK8sApiEventsV1Event) SetRelated(v IoK8sApiCoreV1ObjectReference)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *IoK8sApiEventsV1Event) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetReportingController

`func (o *IoK8sApiEventsV1Event) GetReportingController() string`

GetReportingController returns the ReportingController field if non-nil, zero value otherwise.

### GetReportingControllerOk

`func (o *IoK8sApiEventsV1Event) GetReportingControllerOk() (*string, bool)`

GetReportingControllerOk returns a tuple with the ReportingController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingController

`func (o *IoK8sApiEventsV1Event) SetReportingController(v string)`

SetReportingController sets ReportingController field to given value.

### HasReportingController

`func (o *IoK8sApiEventsV1Event) HasReportingController() bool`

HasReportingController returns a boolean if a field has been set.

### GetReportingInstance

`func (o *IoK8sApiEventsV1Event) GetReportingInstance() string`

GetReportingInstance returns the ReportingInstance field if non-nil, zero value otherwise.

### GetReportingInstanceOk

`func (o *IoK8sApiEventsV1Event) GetReportingInstanceOk() (*string, bool)`

GetReportingInstanceOk returns a tuple with the ReportingInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingInstance

`func (o *IoK8sApiEventsV1Event) SetReportingInstance(v string)`

SetReportingInstance sets ReportingInstance field to given value.

### HasReportingInstance

`func (o *IoK8sApiEventsV1Event) HasReportingInstance() bool`

HasReportingInstance returns a boolean if a field has been set.

### GetSeries

`func (o *IoK8sApiEventsV1Event) GetSeries() IoK8sApiEventsV1EventSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *IoK8sApiEventsV1Event) GetSeriesOk() (*IoK8sApiEventsV1EventSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *IoK8sApiEventsV1Event) SetSeries(v IoK8sApiEventsV1EventSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *IoK8sApiEventsV1Event) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiEventsV1Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiEventsV1Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiEventsV1Event) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiEventsV1Event) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


