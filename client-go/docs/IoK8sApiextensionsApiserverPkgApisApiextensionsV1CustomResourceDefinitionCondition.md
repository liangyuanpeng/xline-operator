# IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | message is a human-readable message indicating details about last transition. | [optional] 
**Reason** | Pointer to **string** | reason is a unique, one-word, CamelCase reason for the condition&#39;s last transition. | [optional] 
**Status** | **string** | status is the status of the condition. Can be True, False, Unknown. | 
**Type** | **string** | type is the type of the condition. Types include Established, NamesAccepted and Terminating. | 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition(status string, type_ string, ) *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionConditionWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionConditionWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionConditionWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


