# IoK8sApiCoreV1Affinity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeAffinity** | Pointer to [**IoK8sApiCoreV1NodeAffinity**](IoK8sApiCoreV1NodeAffinity.md) |  | [optional] 
**PodAffinity** | Pointer to [**IoK8sApiCoreV1PodAffinity**](IoK8sApiCoreV1PodAffinity.md) |  | [optional] 
**PodAntiAffinity** | Pointer to [**IoK8sApiCoreV1PodAntiAffinity**](IoK8sApiCoreV1PodAntiAffinity.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1Affinity

`func NewIoK8sApiCoreV1Affinity() *IoK8sApiCoreV1Affinity`

NewIoK8sApiCoreV1Affinity instantiates a new IoK8sApiCoreV1Affinity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1AffinityWithDefaults

`func NewIoK8sApiCoreV1AffinityWithDefaults() *IoK8sApiCoreV1Affinity`

NewIoK8sApiCoreV1AffinityWithDefaults instantiates a new IoK8sApiCoreV1Affinity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeAffinity

`func (o *IoK8sApiCoreV1Affinity) GetNodeAffinity() IoK8sApiCoreV1NodeAffinity`

GetNodeAffinity returns the NodeAffinity field if non-nil, zero value otherwise.

### GetNodeAffinityOk

`func (o *IoK8sApiCoreV1Affinity) GetNodeAffinityOk() (*IoK8sApiCoreV1NodeAffinity, bool)`

GetNodeAffinityOk returns a tuple with the NodeAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAffinity

`func (o *IoK8sApiCoreV1Affinity) SetNodeAffinity(v IoK8sApiCoreV1NodeAffinity)`

SetNodeAffinity sets NodeAffinity field to given value.

### HasNodeAffinity

`func (o *IoK8sApiCoreV1Affinity) HasNodeAffinity() bool`

HasNodeAffinity returns a boolean if a field has been set.

### GetPodAffinity

`func (o *IoK8sApiCoreV1Affinity) GetPodAffinity() IoK8sApiCoreV1PodAffinity`

GetPodAffinity returns the PodAffinity field if non-nil, zero value otherwise.

### GetPodAffinityOk

`func (o *IoK8sApiCoreV1Affinity) GetPodAffinityOk() (*IoK8sApiCoreV1PodAffinity, bool)`

GetPodAffinityOk returns a tuple with the PodAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAffinity

`func (o *IoK8sApiCoreV1Affinity) SetPodAffinity(v IoK8sApiCoreV1PodAffinity)`

SetPodAffinity sets PodAffinity field to given value.

### HasPodAffinity

`func (o *IoK8sApiCoreV1Affinity) HasPodAffinity() bool`

HasPodAffinity returns a boolean if a field has been set.

### GetPodAntiAffinity

`func (o *IoK8sApiCoreV1Affinity) GetPodAntiAffinity() IoK8sApiCoreV1PodAntiAffinity`

GetPodAntiAffinity returns the PodAntiAffinity field if non-nil, zero value otherwise.

### GetPodAntiAffinityOk

`func (o *IoK8sApiCoreV1Affinity) GetPodAntiAffinityOk() (*IoK8sApiCoreV1PodAntiAffinity, bool)`

GetPodAntiAffinityOk returns a tuple with the PodAntiAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAntiAffinity

`func (o *IoK8sApiCoreV1Affinity) SetPodAntiAffinity(v IoK8sApiCoreV1PodAntiAffinity)`

SetPodAntiAffinity sets PodAntiAffinity field to given value.

### HasPodAntiAffinity

`func (o *IoK8sApiCoreV1Affinity) HasPodAntiAffinity() bool`

HasPodAntiAffinity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


