# IoK8sApiCoreV1WeightedPodAffinityTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodAffinityTerm** | [**IoK8sApiCoreV1PodAffinityTerm**](IoK8sApiCoreV1PodAffinityTerm.md) |  | 
**Weight** | **int32** | weight associated with matching the corresponding podAffinityTerm, in the range 1-100. | 

## Methods

### NewIoK8sApiCoreV1WeightedPodAffinityTerm

`func NewIoK8sApiCoreV1WeightedPodAffinityTerm(podAffinityTerm IoK8sApiCoreV1PodAffinityTerm, weight int32, ) *IoK8sApiCoreV1WeightedPodAffinityTerm`

NewIoK8sApiCoreV1WeightedPodAffinityTerm instantiates a new IoK8sApiCoreV1WeightedPodAffinityTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1WeightedPodAffinityTermWithDefaults

`func NewIoK8sApiCoreV1WeightedPodAffinityTermWithDefaults() *IoK8sApiCoreV1WeightedPodAffinityTerm`

NewIoK8sApiCoreV1WeightedPodAffinityTermWithDefaults instantiates a new IoK8sApiCoreV1WeightedPodAffinityTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodAffinityTerm

`func (o *IoK8sApiCoreV1WeightedPodAffinityTerm) GetPodAffinityTerm() IoK8sApiCoreV1PodAffinityTerm`

GetPodAffinityTerm returns the PodAffinityTerm field if non-nil, zero value otherwise.

### GetPodAffinityTermOk

`func (o *IoK8sApiCoreV1WeightedPodAffinityTerm) GetPodAffinityTermOk() (*IoK8sApiCoreV1PodAffinityTerm, bool)`

GetPodAffinityTermOk returns a tuple with the PodAffinityTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAffinityTerm

`func (o *IoK8sApiCoreV1WeightedPodAffinityTerm) SetPodAffinityTerm(v IoK8sApiCoreV1PodAffinityTerm)`

SetPodAffinityTerm sets PodAffinityTerm field to given value.


### GetWeight

`func (o *IoK8sApiCoreV1WeightedPodAffinityTerm) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *IoK8sApiCoreV1WeightedPodAffinityTerm) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *IoK8sApiCoreV1WeightedPodAffinityTerm) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


