# IoK8sApiCoreV1ComponentCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Condition error code for a component. For example, a health check error code. | [optional] 
**Message** | Pointer to **string** | Message about the condition for a component. For example, information about a health check. | [optional] 
**Status** | **string** | Status of the condition for a component. Valid values for \&quot;Healthy\&quot;: \&quot;True\&quot;, \&quot;False\&quot;, or \&quot;Unknown\&quot;. | 
**Type** | **string** | Type of condition for a component. Valid value: \&quot;Healthy\&quot; | 

## Methods

### NewIoK8sApiCoreV1ComponentCondition

`func NewIoK8sApiCoreV1ComponentCondition(status string, type_ string, ) *IoK8sApiCoreV1ComponentCondition`

NewIoK8sApiCoreV1ComponentCondition instantiates a new IoK8sApiCoreV1ComponentCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ComponentConditionWithDefaults

`func NewIoK8sApiCoreV1ComponentConditionWithDefaults() *IoK8sApiCoreV1ComponentCondition`

NewIoK8sApiCoreV1ComponentConditionWithDefaults instantiates a new IoK8sApiCoreV1ComponentCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *IoK8sApiCoreV1ComponentCondition) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IoK8sApiCoreV1ComponentCondition) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IoK8sApiCoreV1ComponentCondition) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *IoK8sApiCoreV1ComponentCondition) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiCoreV1ComponentCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiCoreV1ComponentCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiCoreV1ComponentCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiCoreV1ComponentCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiCoreV1ComponentCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiCoreV1ComponentCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiCoreV1ComponentCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *IoK8sApiCoreV1ComponentCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiCoreV1ComponentCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiCoreV1ComponentCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


