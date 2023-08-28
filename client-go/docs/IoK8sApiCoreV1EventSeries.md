# IoK8sApiCoreV1EventSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of occurrences in this series up to the last heartbeat time | [optional] 
**LastObservedTime** | Pointer to **time.Time** | MicroTime is version of Time with microsecond level precision. | [optional] 

## Methods

### NewIoK8sApiCoreV1EventSeries

`func NewIoK8sApiCoreV1EventSeries() *IoK8sApiCoreV1EventSeries`

NewIoK8sApiCoreV1EventSeries instantiates a new IoK8sApiCoreV1EventSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1EventSeriesWithDefaults

`func NewIoK8sApiCoreV1EventSeriesWithDefaults() *IoK8sApiCoreV1EventSeries`

NewIoK8sApiCoreV1EventSeriesWithDefaults instantiates a new IoK8sApiCoreV1EventSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IoK8sApiCoreV1EventSeries) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IoK8sApiCoreV1EventSeries) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IoK8sApiCoreV1EventSeries) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IoK8sApiCoreV1EventSeries) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastObservedTime

`func (o *IoK8sApiCoreV1EventSeries) GetLastObservedTime() time.Time`

GetLastObservedTime returns the LastObservedTime field if non-nil, zero value otherwise.

### GetLastObservedTimeOk

`func (o *IoK8sApiCoreV1EventSeries) GetLastObservedTimeOk() (*time.Time, bool)`

GetLastObservedTimeOk returns a tuple with the LastObservedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastObservedTime

`func (o *IoK8sApiCoreV1EventSeries) SetLastObservedTime(v time.Time)`

SetLastObservedTime sets LastObservedTime field to given value.

### HasLastObservedTime

`func (o *IoK8sApiCoreV1EventSeries) HasLastObservedTime() bool`

HasLastObservedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


