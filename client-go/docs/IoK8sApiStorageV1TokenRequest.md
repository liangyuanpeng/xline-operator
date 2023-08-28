# IoK8sApiStorageV1TokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** | audience is the intended audience of the token in \&quot;TokenRequestSpec\&quot;. It will default to the audiences of kube apiserver. | 
**ExpirationSeconds** | Pointer to **int64** | expirationSeconds is the duration of validity of the token in \&quot;TokenRequestSpec\&quot;. It has the same default value of \&quot;ExpirationSeconds\&quot; in \&quot;TokenRequestSpec\&quot;. | [optional] 

## Methods

### NewIoK8sApiStorageV1TokenRequest

`func NewIoK8sApiStorageV1TokenRequest(audience string, ) *IoK8sApiStorageV1TokenRequest`

NewIoK8sApiStorageV1TokenRequest instantiates a new IoK8sApiStorageV1TokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiStorageV1TokenRequestWithDefaults

`func NewIoK8sApiStorageV1TokenRequestWithDefaults() *IoK8sApiStorageV1TokenRequest`

NewIoK8sApiStorageV1TokenRequestWithDefaults instantiates a new IoK8sApiStorageV1TokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *IoK8sApiStorageV1TokenRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *IoK8sApiStorageV1TokenRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *IoK8sApiStorageV1TokenRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetExpirationSeconds

`func (o *IoK8sApiStorageV1TokenRequest) GetExpirationSeconds() int64`

GetExpirationSeconds returns the ExpirationSeconds field if non-nil, zero value otherwise.

### GetExpirationSecondsOk

`func (o *IoK8sApiStorageV1TokenRequest) GetExpirationSecondsOk() (*int64, bool)`

GetExpirationSecondsOk returns a tuple with the ExpirationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationSeconds

`func (o *IoK8sApiStorageV1TokenRequest) SetExpirationSeconds(v int64)`

SetExpirationSeconds sets ExpirationSeconds field to given value.

### HasExpirationSeconds

`func (o *IoK8sApiStorageV1TokenRequest) HasExpirationSeconds() bool`

HasExpirationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


