# IoK8sApiCertificatesV1CertificateSigningRequestStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | certificate is populated with an issued certificate by the signer after an Approved condition is present. This field is set via the /status subresource. Once populated, this field is immutable.  If the certificate signing request is denied, a condition of type \&quot;Denied\&quot; is added and this field remains empty. If the signer cannot issue the certificate, a condition of type \&quot;Failed\&quot; is added and this field remains empty.  Validation requirements:  1. certificate must contain one or more PEM blocks.  2. All PEM blocks must have the \&quot;CERTIFICATE\&quot; label, contain no headers, and the encoded data   must be a BER-encoded ASN.1 Certificate structure as described in section 4 of RFC5280.  3. Non-PEM content may appear before or after the \&quot;CERTIFICATE\&quot; PEM blocks and is unvalidated,   to allow for explanatory text as described in section 5.2 of RFC7468.  If more than one PEM block is present, and the definition of the requested spec.signerName does not indicate otherwise, the first block is the issued certificate, and subsequent blocks should be treated as intermediate certificates and presented in TLS handshakes.  The certificate is encoded in PEM format.  When serialized as JSON or YAML, the data is additionally base64-encoded, so it consists of:      base64(     -----BEGIN CERTIFICATE-----     ...     -----END CERTIFICATE-----     ) | [optional] 
**Conditions** | Pointer to [**[]IoK8sApiCertificatesV1CertificateSigningRequestCondition**](IoK8sApiCertificatesV1CertificateSigningRequestCondition.md) | conditions applied to the request. Known conditions are \&quot;Approved\&quot;, \&quot;Denied\&quot;, and \&quot;Failed\&quot;. | [optional] 

## Methods

### NewIoK8sApiCertificatesV1CertificateSigningRequestStatus

`func NewIoK8sApiCertificatesV1CertificateSigningRequestStatus() *IoK8sApiCertificatesV1CertificateSigningRequestStatus`

NewIoK8sApiCertificatesV1CertificateSigningRequestStatus instantiates a new IoK8sApiCertificatesV1CertificateSigningRequestStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCertificatesV1CertificateSigningRequestStatusWithDefaults

`func NewIoK8sApiCertificatesV1CertificateSigningRequestStatusWithDefaults() *IoK8sApiCertificatesV1CertificateSigningRequestStatus`

NewIoK8sApiCertificatesV1CertificateSigningRequestStatusWithDefaults instantiates a new IoK8sApiCertificatesV1CertificateSigningRequestStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestStatus) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestStatus) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestStatus) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestStatus) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetConditions

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestStatus) GetConditions() []IoK8sApiCertificatesV1CertificateSigningRequestCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestStatus) GetConditionsOk() (*[]IoK8sApiCertificatesV1CertificateSigningRequestCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestStatus) SetConditions(v []IoK8sApiCertificatesV1CertificateSigningRequestCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


