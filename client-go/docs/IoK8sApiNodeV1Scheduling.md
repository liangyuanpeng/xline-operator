# IoK8sApiNodeV1Scheduling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeSelector** | Pointer to **map[string]string** | nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod&#39;s existing nodeSelector. Any conflicts will cause the pod to be rejected in admission. | [optional] 
**Tolerations** | Pointer to [**[]IoK8sApiCoreV1Toleration**](IoK8sApiCoreV1Toleration.md) | tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass. | [optional] 

## Methods

### NewIoK8sApiNodeV1Scheduling

`func NewIoK8sApiNodeV1Scheduling() *IoK8sApiNodeV1Scheduling`

NewIoK8sApiNodeV1Scheduling instantiates a new IoK8sApiNodeV1Scheduling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNodeV1SchedulingWithDefaults

`func NewIoK8sApiNodeV1SchedulingWithDefaults() *IoK8sApiNodeV1Scheduling`

NewIoK8sApiNodeV1SchedulingWithDefaults instantiates a new IoK8sApiNodeV1Scheduling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeSelector

`func (o *IoK8sApiNodeV1Scheduling) GetNodeSelector() map[string]string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *IoK8sApiNodeV1Scheduling) GetNodeSelectorOk() (*map[string]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *IoK8sApiNodeV1Scheduling) SetNodeSelector(v map[string]string)`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *IoK8sApiNodeV1Scheduling) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetTolerations

`func (o *IoK8sApiNodeV1Scheduling) GetTolerations() []IoK8sApiCoreV1Toleration`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *IoK8sApiNodeV1Scheduling) GetTolerationsOk() (*[]IoK8sApiCoreV1Toleration, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *IoK8sApiNodeV1Scheduling) SetTolerations(v []IoK8sApiCoreV1Toleration)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *IoK8sApiNodeV1Scheduling) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


