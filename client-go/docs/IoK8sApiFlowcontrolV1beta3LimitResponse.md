# IoK8sApiFlowcontrolV1beta3LimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queuing** | Pointer to [**IoK8sApiFlowcontrolV1beta3QueuingConfiguration**](IoK8sApiFlowcontrolV1beta3QueuingConfiguration.md) |  | [optional] 
**Type** | **string** | &#x60;type&#x60; is \&quot;Queue\&quot; or \&quot;Reject\&quot;. \&quot;Queue\&quot; means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. \&quot;Reject\&quot; means that requests that can not be executed upon arrival are rejected. Required. | 

## Methods

### NewIoK8sApiFlowcontrolV1beta3LimitResponse

`func NewIoK8sApiFlowcontrolV1beta3LimitResponse(type_ string, ) *IoK8sApiFlowcontrolV1beta3LimitResponse`

NewIoK8sApiFlowcontrolV1beta3LimitResponse instantiates a new IoK8sApiFlowcontrolV1beta3LimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta3LimitResponseWithDefaults

`func NewIoK8sApiFlowcontrolV1beta3LimitResponseWithDefaults() *IoK8sApiFlowcontrolV1beta3LimitResponse`

NewIoK8sApiFlowcontrolV1beta3LimitResponseWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta3LimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueuing

`func (o *IoK8sApiFlowcontrolV1beta3LimitResponse) GetQueuing() IoK8sApiFlowcontrolV1beta3QueuingConfiguration`

GetQueuing returns the Queuing field if non-nil, zero value otherwise.

### GetQueuingOk

`func (o *IoK8sApiFlowcontrolV1beta3LimitResponse) GetQueuingOk() (*IoK8sApiFlowcontrolV1beta3QueuingConfiguration, bool)`

GetQueuingOk returns a tuple with the Queuing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuing

`func (o *IoK8sApiFlowcontrolV1beta3LimitResponse) SetQueuing(v IoK8sApiFlowcontrolV1beta3QueuingConfiguration)`

SetQueuing sets Queuing field to given value.

### HasQueuing

`func (o *IoK8sApiFlowcontrolV1beta3LimitResponse) HasQueuing() bool`

HasQueuing returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiFlowcontrolV1beta3LimitResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiFlowcontrolV1beta3LimitResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiFlowcontrolV1beta3LimitResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


