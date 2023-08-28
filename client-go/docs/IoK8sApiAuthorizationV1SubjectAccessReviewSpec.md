# IoK8sApiAuthorizationV1SubjectAccessReviewSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extra** | Pointer to **map[string][]string** | Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here. | [optional] 
**Groups** | Pointer to **[]string** | Groups is the groups you&#39;re testing for. | [optional] 
**NonResourceAttributes** | Pointer to [**IoK8sApiAuthorizationV1NonResourceAttributes**](IoK8sApiAuthorizationV1NonResourceAttributes.md) |  | [optional] 
**ResourceAttributes** | Pointer to [**IoK8sApiAuthorizationV1ResourceAttributes**](IoK8sApiAuthorizationV1ResourceAttributes.md) |  | [optional] 
**Uid** | Pointer to **string** | UID information about the requesting user. | [optional] 
**User** | Pointer to **string** | User is the user you&#39;re testing for. If you specify \&quot;User\&quot; but not \&quot;Groups\&quot;, then is it interpreted as \&quot;What if User were not a member of any groups | [optional] 

## Methods

### NewIoK8sApiAuthorizationV1SubjectAccessReviewSpec

`func NewIoK8sApiAuthorizationV1SubjectAccessReviewSpec() *IoK8sApiAuthorizationV1SubjectAccessReviewSpec`

NewIoK8sApiAuthorizationV1SubjectAccessReviewSpec instantiates a new IoK8sApiAuthorizationV1SubjectAccessReviewSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthorizationV1SubjectAccessReviewSpecWithDefaults

`func NewIoK8sApiAuthorizationV1SubjectAccessReviewSpecWithDefaults() *IoK8sApiAuthorizationV1SubjectAccessReviewSpec`

NewIoK8sApiAuthorizationV1SubjectAccessReviewSpecWithDefaults instantiates a new IoK8sApiAuthorizationV1SubjectAccessReviewSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtra

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetExtra() map[string][]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetExtraOk() (*map[string][]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) SetExtra(v map[string][]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetGroups

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetNonResourceAttributes

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetNonResourceAttributes() IoK8sApiAuthorizationV1NonResourceAttributes`

GetNonResourceAttributes returns the NonResourceAttributes field if non-nil, zero value otherwise.

### GetNonResourceAttributesOk

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetNonResourceAttributesOk() (*IoK8sApiAuthorizationV1NonResourceAttributes, bool)`

GetNonResourceAttributesOk returns a tuple with the NonResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceAttributes

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) SetNonResourceAttributes(v IoK8sApiAuthorizationV1NonResourceAttributes)`

SetNonResourceAttributes sets NonResourceAttributes field to given value.

### HasNonResourceAttributes

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) HasNonResourceAttributes() bool`

HasNonResourceAttributes returns a boolean if a field has been set.

### GetResourceAttributes

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetResourceAttributes() IoK8sApiAuthorizationV1ResourceAttributes`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetResourceAttributesOk() (*IoK8sApiAuthorizationV1ResourceAttributes, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) SetResourceAttributes(v IoK8sApiAuthorizationV1ResourceAttributes)`

SetResourceAttributes sets ResourceAttributes field to given value.

### HasResourceAttributes

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) HasResourceAttributes() bool`

HasResourceAttributes returns a boolean if a field has been set.

### GetUid

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUser

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IoK8sApiAuthorizationV1SubjectAccessReviewSpec) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


