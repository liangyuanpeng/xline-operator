# IoK8sApiCoreV1SELinuxOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | Level is SELinux level label that applies to the container. | [optional] 
**Role** | Pointer to **string** | Role is a SELinux role label that applies to the container. | [optional] 
**Type** | Pointer to **string** | Type is a SELinux type label that applies to the container. | [optional] 
**User** | Pointer to **string** | User is a SELinux user label that applies to the container. | [optional] 

## Methods

### NewIoK8sApiCoreV1SELinuxOptions

`func NewIoK8sApiCoreV1SELinuxOptions() *IoK8sApiCoreV1SELinuxOptions`

NewIoK8sApiCoreV1SELinuxOptions instantiates a new IoK8sApiCoreV1SELinuxOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1SELinuxOptionsWithDefaults

`func NewIoK8sApiCoreV1SELinuxOptionsWithDefaults() *IoK8sApiCoreV1SELinuxOptions`

NewIoK8sApiCoreV1SELinuxOptionsWithDefaults instantiates a new IoK8sApiCoreV1SELinuxOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *IoK8sApiCoreV1SELinuxOptions) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *IoK8sApiCoreV1SELinuxOptions) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *IoK8sApiCoreV1SELinuxOptions) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *IoK8sApiCoreV1SELinuxOptions) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetRole

`func (o *IoK8sApiCoreV1SELinuxOptions) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IoK8sApiCoreV1SELinuxOptions) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IoK8sApiCoreV1SELinuxOptions) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IoK8sApiCoreV1SELinuxOptions) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiCoreV1SELinuxOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiCoreV1SELinuxOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiCoreV1SELinuxOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiCoreV1SELinuxOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *IoK8sApiCoreV1SELinuxOptions) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IoK8sApiCoreV1SELinuxOptions) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IoK8sApiCoreV1SELinuxOptions) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IoK8sApiCoreV1SELinuxOptions) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


