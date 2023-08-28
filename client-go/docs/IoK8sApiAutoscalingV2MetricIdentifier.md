# IoK8sApiAutoscalingV2MetricIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is the name of the given metric | 
**Selector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 

## Methods

### NewIoK8sApiAutoscalingV2MetricIdentifier

`func NewIoK8sApiAutoscalingV2MetricIdentifier(name string, ) *IoK8sApiAutoscalingV2MetricIdentifier`

NewIoK8sApiAutoscalingV2MetricIdentifier instantiates a new IoK8sApiAutoscalingV2MetricIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV2MetricIdentifierWithDefaults

`func NewIoK8sApiAutoscalingV2MetricIdentifierWithDefaults() *IoK8sApiAutoscalingV2MetricIdentifier`

NewIoK8sApiAutoscalingV2MetricIdentifierWithDefaults instantiates a new IoK8sApiAutoscalingV2MetricIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiAutoscalingV2MetricIdentifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiAutoscalingV2MetricIdentifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiAutoscalingV2MetricIdentifier) SetName(v string)`

SetName sets Name field to given value.


### GetSelector

`func (o *IoK8sApiAutoscalingV2MetricIdentifier) GetSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiAutoscalingV2MetricIdentifier) GetSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiAutoscalingV2MetricIdentifier) SetSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *IoK8sApiAutoscalingV2MetricIdentifier) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


