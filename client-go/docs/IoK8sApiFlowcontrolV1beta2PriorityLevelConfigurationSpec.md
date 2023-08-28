# IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limited** | Pointer to [**IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration**](IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration.md) |  | [optional] 
**Type** | **string** | &#x60;type&#x60; indicates whether this priority level is subject to limitation on request execution.  A value of &#x60;\&quot;Exempt\&quot;&#x60; means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of &#x60;\&quot;Limited\&quot;&#x60; means that (a) requests of this priority level _are_ subject to limits and (b) some of the server&#39;s limited capacity is made available exclusively to this priority level. Required. | 

## Methods

### NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec

`func NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec(type_ string, ) *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec`

NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec instantiates a new IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpecWithDefaults

`func NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpecWithDefaults() *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec`

NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpecWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimited

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec) GetLimited() IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration`

GetLimited returns the Limited field if non-nil, zero value otherwise.

### GetLimitedOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec) GetLimitedOk() (*IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration, bool)`

GetLimitedOk returns a tuple with the Limited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimited

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec) SetLimited(v IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration)`

SetLimited sets Limited field to given value.

### HasLimited

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec) HasLimited() bool`

HasLimited returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


