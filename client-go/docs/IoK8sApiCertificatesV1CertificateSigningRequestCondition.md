# IoK8sApiCertificatesV1CertificateSigningRequestCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**LastUpdateTime** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | message contains a human readable message with details about the request state | [optional] 
**Reason** | Pointer to **string** | reason indicates a brief reason for the request state | [optional] 
**Status** | **string** | status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be \&quot;False\&quot; or \&quot;Unknown\&quot;. | 
**Type** | **string** | type of the condition. Known conditions are \&quot;Approved\&quot;, \&quot;Denied\&quot;, and \&quot;Failed\&quot;.  An \&quot;Approved\&quot; condition is added via the /approval subresource, indicating the request was approved and should be issued by the signer.  A \&quot;Denied\&quot; condition is added via the /approval subresource, indicating the request was denied and should not be issued by the signer.  A \&quot;Failed\&quot; condition is added via the /status subresource, indicating the signer failed to issue the certificate.  Approved and Denied conditions are mutually exclusive. Approved, Denied, and Failed conditions cannot be removed once added.  Only one condition of a given type is allowed. | 

## Methods

### NewIoK8sApiCertificatesV1CertificateSigningRequestCondition

`func NewIoK8sApiCertificatesV1CertificateSigningRequestCondition(status string, type_ string, ) *IoK8sApiCertificatesV1CertificateSigningRequestCondition`

NewIoK8sApiCertificatesV1CertificateSigningRequestCondition instantiates a new IoK8sApiCertificatesV1CertificateSigningRequestCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCertificatesV1CertificateSigningRequestConditionWithDefaults

`func NewIoK8sApiCertificatesV1CertificateSigningRequestConditionWithDefaults() *IoK8sApiCertificatesV1CertificateSigningRequestCondition`

NewIoK8sApiCertificatesV1CertificateSigningRequestConditionWithDefaults instantiates a new IoK8sApiCertificatesV1CertificateSigningRequestCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMessage

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


