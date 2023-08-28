# IoK8sApiAuthenticationV1TokenRequestStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTimestamp** | **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | 
**Token** | **string** | Token is the opaque bearer token. | 

## Methods

### NewIoK8sApiAuthenticationV1TokenRequestStatus

`func NewIoK8sApiAuthenticationV1TokenRequestStatus(expirationTimestamp time.Time, token string, ) *IoK8sApiAuthenticationV1TokenRequestStatus`

NewIoK8sApiAuthenticationV1TokenRequestStatus instantiates a new IoK8sApiAuthenticationV1TokenRequestStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthenticationV1TokenRequestStatusWithDefaults

`func NewIoK8sApiAuthenticationV1TokenRequestStatusWithDefaults() *IoK8sApiAuthenticationV1TokenRequestStatus`

NewIoK8sApiAuthenticationV1TokenRequestStatusWithDefaults instantiates a new IoK8sApiAuthenticationV1TokenRequestStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTimestamp

`func (o *IoK8sApiAuthenticationV1TokenRequestStatus) GetExpirationTimestamp() time.Time`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *IoK8sApiAuthenticationV1TokenRequestStatus) GetExpirationTimestampOk() (*time.Time, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *IoK8sApiAuthenticationV1TokenRequestStatus) SetExpirationTimestamp(v time.Time)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### GetToken

`func (o *IoK8sApiAuthenticationV1TokenRequestStatus) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *IoK8sApiAuthenticationV1TokenRequestStatus) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *IoK8sApiAuthenticationV1TokenRequestStatus) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


