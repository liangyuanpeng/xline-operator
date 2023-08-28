# IoK8sApiCoreV1PreferredSchedulingTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preference** | [**IoK8sApiCoreV1NodeSelectorTerm**](IoK8sApiCoreV1NodeSelectorTerm.md) |  | 
**Weight** | **int32** | Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100. | 

## Methods

### NewIoK8sApiCoreV1PreferredSchedulingTerm

`func NewIoK8sApiCoreV1PreferredSchedulingTerm(preference IoK8sApiCoreV1NodeSelectorTerm, weight int32, ) *IoK8sApiCoreV1PreferredSchedulingTerm`

NewIoK8sApiCoreV1PreferredSchedulingTerm instantiates a new IoK8sApiCoreV1PreferredSchedulingTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PreferredSchedulingTermWithDefaults

`func NewIoK8sApiCoreV1PreferredSchedulingTermWithDefaults() *IoK8sApiCoreV1PreferredSchedulingTerm`

NewIoK8sApiCoreV1PreferredSchedulingTermWithDefaults instantiates a new IoK8sApiCoreV1PreferredSchedulingTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreference

`func (o *IoK8sApiCoreV1PreferredSchedulingTerm) GetPreference() IoK8sApiCoreV1NodeSelectorTerm`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *IoK8sApiCoreV1PreferredSchedulingTerm) GetPreferenceOk() (*IoK8sApiCoreV1NodeSelectorTerm, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *IoK8sApiCoreV1PreferredSchedulingTerm) SetPreference(v IoK8sApiCoreV1NodeSelectorTerm)`

SetPreference sets Preference field to given value.


### GetWeight

`func (o *IoK8sApiCoreV1PreferredSchedulingTerm) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *IoK8sApiCoreV1PreferredSchedulingTerm) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *IoK8sApiCoreV1PreferredSchedulingTerm) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


