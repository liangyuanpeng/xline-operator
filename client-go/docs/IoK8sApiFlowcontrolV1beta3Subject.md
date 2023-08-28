# IoK8sApiFlowcontrolV1beta3Subject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**IoK8sApiFlowcontrolV1beta3GroupSubject**](IoK8sApiFlowcontrolV1beta3GroupSubject.md) |  | [optional] 
**Kind** | **string** | &#x60;kind&#x60; indicates which one of the other fields is non-empty. Required | 
**ServiceAccount** | Pointer to [**IoK8sApiFlowcontrolV1beta3ServiceAccountSubject**](IoK8sApiFlowcontrolV1beta3ServiceAccountSubject.md) |  | [optional] 
**User** | Pointer to [**IoK8sApiFlowcontrolV1beta3UserSubject**](IoK8sApiFlowcontrolV1beta3UserSubject.md) |  | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta3Subject

`func NewIoK8sApiFlowcontrolV1beta3Subject(kind string, ) *IoK8sApiFlowcontrolV1beta3Subject`

NewIoK8sApiFlowcontrolV1beta3Subject instantiates a new IoK8sApiFlowcontrolV1beta3Subject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta3SubjectWithDefaults

`func NewIoK8sApiFlowcontrolV1beta3SubjectWithDefaults() *IoK8sApiFlowcontrolV1beta3Subject`

NewIoK8sApiFlowcontrolV1beta3SubjectWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta3Subject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *IoK8sApiFlowcontrolV1beta3Subject) GetGroup() IoK8sApiFlowcontrolV1beta3GroupSubject`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IoK8sApiFlowcontrolV1beta3Subject) GetGroupOk() (*IoK8sApiFlowcontrolV1beta3GroupSubject, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IoK8sApiFlowcontrolV1beta3Subject) SetGroup(v IoK8sApiFlowcontrolV1beta3GroupSubject)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IoK8sApiFlowcontrolV1beta3Subject) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiFlowcontrolV1beta3Subject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiFlowcontrolV1beta3Subject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiFlowcontrolV1beta3Subject) SetKind(v string)`

SetKind sets Kind field to given value.


### GetServiceAccount

`func (o *IoK8sApiFlowcontrolV1beta3Subject) GetServiceAccount() IoK8sApiFlowcontrolV1beta3ServiceAccountSubject`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *IoK8sApiFlowcontrolV1beta3Subject) GetServiceAccountOk() (*IoK8sApiFlowcontrolV1beta3ServiceAccountSubject, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *IoK8sApiFlowcontrolV1beta3Subject) SetServiceAccount(v IoK8sApiFlowcontrolV1beta3ServiceAccountSubject)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *IoK8sApiFlowcontrolV1beta3Subject) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetUser

`func (o *IoK8sApiFlowcontrolV1beta3Subject) GetUser() IoK8sApiFlowcontrolV1beta3UserSubject`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IoK8sApiFlowcontrolV1beta3Subject) GetUserOk() (*IoK8sApiFlowcontrolV1beta3UserSubject, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IoK8sApiFlowcontrolV1beta3Subject) SetUser(v IoK8sApiFlowcontrolV1beta3UserSubject)`

SetUser sets User field to given value.

### HasUser

`func (o *IoK8sApiFlowcontrolV1beta3Subject) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


