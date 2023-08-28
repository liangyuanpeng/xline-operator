# IoK8sApiFlowcontrolV1beta3FlowSchemaStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IoK8sApiFlowcontrolV1beta3FlowSchemaCondition**](IoK8sApiFlowcontrolV1beta3FlowSchemaCondition.md) | &#x60;conditions&#x60; is a list of the current states of FlowSchema. | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta3FlowSchemaStatus

`func NewIoK8sApiFlowcontrolV1beta3FlowSchemaStatus() *IoK8sApiFlowcontrolV1beta3FlowSchemaStatus`

NewIoK8sApiFlowcontrolV1beta3FlowSchemaStatus instantiates a new IoK8sApiFlowcontrolV1beta3FlowSchemaStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta3FlowSchemaStatusWithDefaults

`func NewIoK8sApiFlowcontrolV1beta3FlowSchemaStatusWithDefaults() *IoK8sApiFlowcontrolV1beta3FlowSchemaStatus`

NewIoK8sApiFlowcontrolV1beta3FlowSchemaStatusWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta3FlowSchemaStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaStatus) GetConditions() []IoK8sApiFlowcontrolV1beta3FlowSchemaCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaStatus) GetConditionsOk() (*[]IoK8sApiFlowcontrolV1beta3FlowSchemaCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaStatus) SetConditions(v []IoK8sApiFlowcontrolV1beta3FlowSchemaCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiFlowcontrolV1beta3FlowSchemaStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


