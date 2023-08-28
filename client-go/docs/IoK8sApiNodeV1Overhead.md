# IoK8sApiNodeV1Overhead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodFixed** | Pointer to **map[string]string** | podFixed represents the fixed resource overhead associated with running a pod. | [optional] 

## Methods

### NewIoK8sApiNodeV1Overhead

`func NewIoK8sApiNodeV1Overhead() *IoK8sApiNodeV1Overhead`

NewIoK8sApiNodeV1Overhead instantiates a new IoK8sApiNodeV1Overhead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNodeV1OverheadWithDefaults

`func NewIoK8sApiNodeV1OverheadWithDefaults() *IoK8sApiNodeV1Overhead`

NewIoK8sApiNodeV1OverheadWithDefaults instantiates a new IoK8sApiNodeV1Overhead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodFixed

`func (o *IoK8sApiNodeV1Overhead) GetPodFixed() map[string]string`

GetPodFixed returns the PodFixed field if non-nil, zero value otherwise.

### GetPodFixedOk

`func (o *IoK8sApiNodeV1Overhead) GetPodFixedOk() (*map[string]string, bool)`

GetPodFixedOk returns a tuple with the PodFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodFixed

`func (o *IoK8sApiNodeV1Overhead) SetPodFixed(v map[string]string)`

SetPodFixed sets PodFixed field to given value.

### HasPodFixed

`func (o *IoK8sApiNodeV1Overhead) HasPodFixed() bool`

HasPodFixed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


