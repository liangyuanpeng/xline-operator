# IoK8sApiAuthenticationV1TokenReviewStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | Pointer to **[]string** | Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token&#39;s audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is \&quot;true\&quot;, the token is valid against the audience of the Kubernetes API server. | [optional] 
**Authenticated** | Pointer to **bool** | Authenticated indicates that the token was associated with a known user. | [optional] 
**Error** | Pointer to **string** | Error indicates that the token couldn&#39;t be checked | [optional] 
**User** | Pointer to [**IoK8sApiAuthenticationV1UserInfo**](IoK8sApiAuthenticationV1UserInfo.md) |  | [optional] 

## Methods

### NewIoK8sApiAuthenticationV1TokenReviewStatus

`func NewIoK8sApiAuthenticationV1TokenReviewStatus() *IoK8sApiAuthenticationV1TokenReviewStatus`

NewIoK8sApiAuthenticationV1TokenReviewStatus instantiates a new IoK8sApiAuthenticationV1TokenReviewStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthenticationV1TokenReviewStatusWithDefaults

`func NewIoK8sApiAuthenticationV1TokenReviewStatusWithDefaults() *IoK8sApiAuthenticationV1TokenReviewStatus`

NewIoK8sApiAuthenticationV1TokenReviewStatusWithDefaults instantiates a new IoK8sApiAuthenticationV1TokenReviewStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### GetAuthenticated

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### GetError

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetUser

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) GetUser() IoK8sApiAuthenticationV1UserInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) GetUserOk() (*IoK8sApiAuthenticationV1UserInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) SetUser(v IoK8sApiAuthenticationV1UserInfo)`

SetUser sets User field to given value.

### HasUser

`func (o *IoK8sApiAuthenticationV1TokenReviewStatus) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


