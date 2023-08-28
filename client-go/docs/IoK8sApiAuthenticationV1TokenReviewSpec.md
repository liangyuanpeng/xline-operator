# IoK8sApiAuthenticationV1TokenReviewSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | Pointer to **[]string** | Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver. | [optional] 
**Token** | Pointer to **string** | Token is the opaque bearer token. | [optional] 

## Methods

### NewIoK8sApiAuthenticationV1TokenReviewSpec

`func NewIoK8sApiAuthenticationV1TokenReviewSpec() *IoK8sApiAuthenticationV1TokenReviewSpec`

NewIoK8sApiAuthenticationV1TokenReviewSpec instantiates a new IoK8sApiAuthenticationV1TokenReviewSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthenticationV1TokenReviewSpecWithDefaults

`func NewIoK8sApiAuthenticationV1TokenReviewSpecWithDefaults() *IoK8sApiAuthenticationV1TokenReviewSpec`

NewIoK8sApiAuthenticationV1TokenReviewSpecWithDefaults instantiates a new IoK8sApiAuthenticationV1TokenReviewSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *IoK8sApiAuthenticationV1TokenReviewSpec) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *IoK8sApiAuthenticationV1TokenReviewSpec) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *IoK8sApiAuthenticationV1TokenReviewSpec) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *IoK8sApiAuthenticationV1TokenReviewSpec) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### GetToken

`func (o *IoK8sApiAuthenticationV1TokenReviewSpec) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *IoK8sApiAuthenticationV1TokenReviewSpec) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *IoK8sApiAuthenticationV1TokenReviewSpec) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *IoK8sApiAuthenticationV1TokenReviewSpec) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


