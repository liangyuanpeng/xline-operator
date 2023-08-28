# IoK8sApiEventsV1EventSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count is the number of occurrences in this series up to the last heartbeat time. | 
**LastObservedTime** | **time.Time** | MicroTime is version of Time with microsecond level precision. | 

## Methods

### NewIoK8sApiEventsV1EventSeries

`func NewIoK8sApiEventsV1EventSeries(count int32, lastObservedTime time.Time, ) *IoK8sApiEventsV1EventSeries`

NewIoK8sApiEventsV1EventSeries instantiates a new IoK8sApiEventsV1EventSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiEventsV1EventSeriesWithDefaults

`func NewIoK8sApiEventsV1EventSeriesWithDefaults() *IoK8sApiEventsV1EventSeries`

NewIoK8sApiEventsV1EventSeriesWithDefaults instantiates a new IoK8sApiEventsV1EventSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IoK8sApiEventsV1EventSeries) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IoK8sApiEventsV1EventSeries) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IoK8sApiEventsV1EventSeries) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLastObservedTime

`func (o *IoK8sApiEventsV1EventSeries) GetLastObservedTime() time.Time`

GetLastObservedTime returns the LastObservedTime field if non-nil, zero value otherwise.

### GetLastObservedTimeOk

`func (o *IoK8sApiEventsV1EventSeries) GetLastObservedTimeOk() (*time.Time, bool)`

GetLastObservedTimeOk returns a tuple with the LastObservedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastObservedTime

`func (o *IoK8sApiEventsV1EventSeries) SetLastObservedTime(v time.Time)`

SetLastObservedTime sets LastObservedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


