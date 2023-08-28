# IoK8sApiFlowcontrolV1beta2Subject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**IoK8sApiFlowcontrolV1beta2GroupSubject**](IoK8sApiFlowcontrolV1beta2GroupSubject.md) |  | [optional] 
**Kind** | **string** | &#x60;kind&#x60; indicates which one of the other fields is non-empty. Required | 
**ServiceAccount** | Pointer to [**IoK8sApiFlowcontrolV1beta2ServiceAccountSubject**](IoK8sApiFlowcontrolV1beta2ServiceAccountSubject.md) |  | [optional] 
**User** | Pointer to [**IoK8sApiFlowcontrolV1beta2UserSubject**](IoK8sApiFlowcontrolV1beta2UserSubject.md) |  | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta2Subject

`func NewIoK8sApiFlowcontrolV1beta2Subject(kind string, ) *IoK8sApiFlowcontrolV1beta2Subject`

NewIoK8sApiFlowcontrolV1beta2Subject instantiates a new IoK8sApiFlowcontrolV1beta2Subject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta2SubjectWithDefaults

`func NewIoK8sApiFlowcontrolV1beta2SubjectWithDefaults() *IoK8sApiFlowcontrolV1beta2Subject`

NewIoK8sApiFlowcontrolV1beta2SubjectWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta2Subject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *IoK8sApiFlowcontrolV1beta2Subject) GetGroup() IoK8sApiFlowcontrolV1beta2GroupSubject`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IoK8sApiFlowcontrolV1beta2Subject) GetGroupOk() (*IoK8sApiFlowcontrolV1beta2GroupSubject, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IoK8sApiFlowcontrolV1beta2Subject) SetGroup(v IoK8sApiFlowcontrolV1beta2GroupSubject)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IoK8sApiFlowcontrolV1beta2Subject) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiFlowcontrolV1beta2Subject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiFlowcontrolV1beta2Subject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiFlowcontrolV1beta2Subject) SetKind(v string)`

SetKind sets Kind field to given value.


### GetServiceAccount

`func (o *IoK8sApiFlowcontrolV1beta2Subject) GetServiceAccount() IoK8sApiFlowcontrolV1beta2ServiceAccountSubject`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *IoK8sApiFlowcontrolV1beta2Subject) GetServiceAccountOk() (*IoK8sApiFlowcontrolV1beta2ServiceAccountSubject, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *IoK8sApiFlowcontrolV1beta2Subject) SetServiceAccount(v IoK8sApiFlowcontrolV1beta2ServiceAccountSubject)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *IoK8sApiFlowcontrolV1beta2Subject) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetUser

`func (o *IoK8sApiFlowcontrolV1beta2Subject) GetUser() IoK8sApiFlowcontrolV1beta2UserSubject`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IoK8sApiFlowcontrolV1beta2Subject) GetUserOk() (*IoK8sApiFlowcontrolV1beta2UserSubject, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IoK8sApiFlowcontrolV1beta2Subject) SetUser(v IoK8sApiFlowcontrolV1beta2UserSubject)`

SetUser sets User field to given value.

### HasUser

`func (o *IoK8sApiFlowcontrolV1beta2Subject) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


