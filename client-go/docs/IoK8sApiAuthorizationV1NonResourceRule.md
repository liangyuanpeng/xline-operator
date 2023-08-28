# IoK8sApiAuthorizationV1NonResourceRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NonResourceURLs** | Pointer to **[]string** | NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path.  \&quot;*\&quot; means all. | [optional] 
**Verbs** | **[]string** | Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options.  \&quot;*\&quot; means all. | 

## Methods

### NewIoK8sApiAuthorizationV1NonResourceRule

`func NewIoK8sApiAuthorizationV1NonResourceRule(verbs []string, ) *IoK8sApiAuthorizationV1NonResourceRule`

NewIoK8sApiAuthorizationV1NonResourceRule instantiates a new IoK8sApiAuthorizationV1NonResourceRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthorizationV1NonResourceRuleWithDefaults

`func NewIoK8sApiAuthorizationV1NonResourceRuleWithDefaults() *IoK8sApiAuthorizationV1NonResourceRule`

NewIoK8sApiAuthorizationV1NonResourceRuleWithDefaults instantiates a new IoK8sApiAuthorizationV1NonResourceRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonResourceURLs

`func (o *IoK8sApiAuthorizationV1NonResourceRule) GetNonResourceURLs() []string`

GetNonResourceURLs returns the NonResourceURLs field if non-nil, zero value otherwise.

### GetNonResourceURLsOk

`func (o *IoK8sApiAuthorizationV1NonResourceRule) GetNonResourceURLsOk() (*[]string, bool)`

GetNonResourceURLsOk returns a tuple with the NonResourceURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceURLs

`func (o *IoK8sApiAuthorizationV1NonResourceRule) SetNonResourceURLs(v []string)`

SetNonResourceURLs sets NonResourceURLs field to given value.

### HasNonResourceURLs

`func (o *IoK8sApiAuthorizationV1NonResourceRule) HasNonResourceURLs() bool`

HasNonResourceURLs returns a boolean if a field has been set.

### GetVerbs

`func (o *IoK8sApiAuthorizationV1NonResourceRule) GetVerbs() []string`

GetVerbs returns the Verbs field if non-nil, zero value otherwise.

### GetVerbsOk

`func (o *IoK8sApiAuthorizationV1NonResourceRule) GetVerbsOk() (*[]string, bool)`

GetVerbsOk returns a tuple with the Verbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbs

`func (o *IoK8sApiAuthorizationV1NonResourceRule) SetVerbs(v []string)`

SetVerbs sets Verbs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


