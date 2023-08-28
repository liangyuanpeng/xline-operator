# IoK8sApiFlowcontrolV1beta3GroupSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is the user group that matches, or \&quot;*\&quot; to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required. | 

## Methods

### NewIoK8sApiFlowcontrolV1beta3GroupSubject

`func NewIoK8sApiFlowcontrolV1beta3GroupSubject(name string, ) *IoK8sApiFlowcontrolV1beta3GroupSubject`

NewIoK8sApiFlowcontrolV1beta3GroupSubject instantiates a new IoK8sApiFlowcontrolV1beta3GroupSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta3GroupSubjectWithDefaults

`func NewIoK8sApiFlowcontrolV1beta3GroupSubjectWithDefaults() *IoK8sApiFlowcontrolV1beta3GroupSubject`

NewIoK8sApiFlowcontrolV1beta3GroupSubjectWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta3GroupSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiFlowcontrolV1beta3GroupSubject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiFlowcontrolV1beta3GroupSubject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiFlowcontrolV1beta3GroupSubject) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


