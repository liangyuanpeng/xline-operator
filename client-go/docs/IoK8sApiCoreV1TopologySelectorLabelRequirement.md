# IoK8sApiCoreV1TopologySelectorLabelRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The label key that the selector applies to. | 
**Values** | **[]string** | An array of string values. One value must match the label to be selected. Each entry in Values is ORed. | 

## Methods

### NewIoK8sApiCoreV1TopologySelectorLabelRequirement

`func NewIoK8sApiCoreV1TopologySelectorLabelRequirement(key string, values []string, ) *IoK8sApiCoreV1TopologySelectorLabelRequirement`

NewIoK8sApiCoreV1TopologySelectorLabelRequirement instantiates a new IoK8sApiCoreV1TopologySelectorLabelRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1TopologySelectorLabelRequirementWithDefaults

`func NewIoK8sApiCoreV1TopologySelectorLabelRequirementWithDefaults() *IoK8sApiCoreV1TopologySelectorLabelRequirement`

NewIoK8sApiCoreV1TopologySelectorLabelRequirementWithDefaults instantiates a new IoK8sApiCoreV1TopologySelectorLabelRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *IoK8sApiCoreV1TopologySelectorLabelRequirement) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IoK8sApiCoreV1TopologySelectorLabelRequirement) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IoK8sApiCoreV1TopologySelectorLabelRequirement) SetKey(v string)`

SetKey sets Key field to given value.


### GetValues

`func (o *IoK8sApiCoreV1TopologySelectorLabelRequirement) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *IoK8sApiCoreV1TopologySelectorLabelRequirement) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *IoK8sApiCoreV1TopologySelectorLabelRequirement) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


