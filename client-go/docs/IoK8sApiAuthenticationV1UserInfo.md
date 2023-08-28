# IoK8sApiAuthenticationV1UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extra** | Pointer to **map[string][]string** | Any additional information provided by the authenticator. | [optional] 
**Groups** | Pointer to **[]string** | The names of groups this user is a part of. | [optional] 
**Uid** | Pointer to **string** | A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs. | [optional] 
**Username** | Pointer to **string** | The name that uniquely identifies this user among all active users. | [optional] 

## Methods

### NewIoK8sApiAuthenticationV1UserInfo

`func NewIoK8sApiAuthenticationV1UserInfo() *IoK8sApiAuthenticationV1UserInfo`

NewIoK8sApiAuthenticationV1UserInfo instantiates a new IoK8sApiAuthenticationV1UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthenticationV1UserInfoWithDefaults

`func NewIoK8sApiAuthenticationV1UserInfoWithDefaults() *IoK8sApiAuthenticationV1UserInfo`

NewIoK8sApiAuthenticationV1UserInfoWithDefaults instantiates a new IoK8sApiAuthenticationV1UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtra

`func (o *IoK8sApiAuthenticationV1UserInfo) GetExtra() map[string][]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *IoK8sApiAuthenticationV1UserInfo) GetExtraOk() (*map[string][]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *IoK8sApiAuthenticationV1UserInfo) SetExtra(v map[string][]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *IoK8sApiAuthenticationV1UserInfo) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetGroups

`func (o *IoK8sApiAuthenticationV1UserInfo) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IoK8sApiAuthenticationV1UserInfo) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IoK8sApiAuthenticationV1UserInfo) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IoK8sApiAuthenticationV1UserInfo) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUid

`func (o *IoK8sApiAuthenticationV1UserInfo) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IoK8sApiAuthenticationV1UserInfo) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IoK8sApiAuthenticationV1UserInfo) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IoK8sApiAuthenticationV1UserInfo) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *IoK8sApiAuthenticationV1UserInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IoK8sApiAuthenticationV1UserInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IoK8sApiAuthenticationV1UserInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IoK8sApiAuthenticationV1UserInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


