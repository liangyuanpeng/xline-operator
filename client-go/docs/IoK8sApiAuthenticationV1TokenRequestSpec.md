# IoK8sApiAuthenticationV1TokenRequestSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | **[]string** | Audiences are the intendend audiences of the token. A recipient of a token must identify themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences. | 
**BoundObjectRef** | Pointer to [**IoK8sApiAuthenticationV1BoundObjectReference**](IoK8sApiAuthenticationV1BoundObjectReference.md) |  | [optional] 
**ExpirationSeconds** | Pointer to **int64** | ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the &#39;expiration&#39; field in a response. | [optional] 

## Methods

### NewIoK8sApiAuthenticationV1TokenRequestSpec

`func NewIoK8sApiAuthenticationV1TokenRequestSpec(audiences []string, ) *IoK8sApiAuthenticationV1TokenRequestSpec`

NewIoK8sApiAuthenticationV1TokenRequestSpec instantiates a new IoK8sApiAuthenticationV1TokenRequestSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthenticationV1TokenRequestSpecWithDefaults

`func NewIoK8sApiAuthenticationV1TokenRequestSpecWithDefaults() *IoK8sApiAuthenticationV1TokenRequestSpec`

NewIoK8sApiAuthenticationV1TokenRequestSpecWithDefaults instantiates a new IoK8sApiAuthenticationV1TokenRequestSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.


### GetBoundObjectRef

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) GetBoundObjectRef() IoK8sApiAuthenticationV1BoundObjectReference`

GetBoundObjectRef returns the BoundObjectRef field if non-nil, zero value otherwise.

### GetBoundObjectRefOk

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) GetBoundObjectRefOk() (*IoK8sApiAuthenticationV1BoundObjectReference, bool)`

GetBoundObjectRefOk returns a tuple with the BoundObjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundObjectRef

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) SetBoundObjectRef(v IoK8sApiAuthenticationV1BoundObjectReference)`

SetBoundObjectRef sets BoundObjectRef field to given value.

### HasBoundObjectRef

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) HasBoundObjectRef() bool`

HasBoundObjectRef returns a boolean if a field has been set.

### GetExpirationSeconds

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) GetExpirationSeconds() int64`

GetExpirationSeconds returns the ExpirationSeconds field if non-nil, zero value otherwise.

### GetExpirationSecondsOk

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) GetExpirationSecondsOk() (*int64, bool)`

GetExpirationSecondsOk returns a tuple with the ExpirationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationSeconds

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) SetExpirationSeconds(v int64)`

SetExpirationSeconds sets ExpirationSeconds field to given value.

### HasExpirationSeconds

`func (o *IoK8sApiAuthenticationV1TokenRequestSpec) HasExpirationSeconds() bool`

HasExpirationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


