# IoK8sApiCertificatesV1CertificateSigningRequestSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationSeconds** | Pointer to **int32** | expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration.  The v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager.  Certificate signers may not honor this field for various reasons:    1. Old signer that is unaware of the field (such as the in-tree      implementations prior to v1.22)   2. Signer whose configured maximum is shorter than the requested duration   3. Signer whose configured minimum is longer than the requested duration  The minimum valid value for expirationSeconds is 600, i.e. 10 minutes. | [optional] 
**Extra** | Pointer to **map[string][]string** | extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable. | [optional] 
**Groups** | Pointer to **[]string** | groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable. | [optional] 
**Request** | **string** | request contains an x509 certificate signing request encoded in a \&quot;CERTIFICATE REQUEST\&quot; PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded. | 
**SignerName** | **string** | signerName indicates the requested signer, and is a qualified name.  List/watch requests for CertificateSigningRequests can filter on this field using a \&quot;spec.signerName&#x3D;NAME\&quot; fieldSelector.  Well-known Kubernetes signers are:  1. \&quot;kubernetes.io/kube-apiserver-client\&quot;: issues client certificates that can be used to authenticate to kube-apiserver.   Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the \&quot;csrsigning\&quot; controller in kube-controller-manager.  2. \&quot;kubernetes.io/kube-apiserver-client-kubelet\&quot;: issues client certificates that kubelets use to authenticate to kube-apiserver.   Requests for this signer can be auto-approved by the \&quot;csrapproving\&quot; controller in kube-controller-manager, and can be issued by the \&quot;csrsigning\&quot; controller in kube-controller-manager.  3. \&quot;kubernetes.io/kubelet-serving\&quot; issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.   Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the \&quot;csrsigning\&quot; controller in kube-controller-manager.  More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers  Custom signerNames can also be specified. The signer defines:  1. Trust distribution: how trust (CA bundles) are distributed.  2. Permitted subjects: and behavior when a disallowed subject is requested.  3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested.  4. Required, permitted, or forbidden key usages / extended key usages.  5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin.  6. Whether or not requests for CA certificates are allowed. | 
**Uid** | Pointer to **string** | uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable. | [optional] 
**Usages** | Pointer to **[]string** | usages specifies a set of key usages requested in the issued certificate.  Requests for TLS client certificates typically request: \&quot;digital signature\&quot;, \&quot;key encipherment\&quot;, \&quot;client auth\&quot;.  Requests for TLS serving certificates typically request: \&quot;key encipherment\&quot;, \&quot;digital signature\&quot;, \&quot;server auth\&quot;.  Valid values are:  \&quot;signing\&quot;, \&quot;digital signature\&quot;, \&quot;content commitment\&quot;,  \&quot;key encipherment\&quot;, \&quot;key agreement\&quot;, \&quot;data encipherment\&quot;,  \&quot;cert sign\&quot;, \&quot;crl sign\&quot;, \&quot;encipher only\&quot;, \&quot;decipher only\&quot;, \&quot;any\&quot;,  \&quot;server auth\&quot;, \&quot;client auth\&quot;,  \&quot;code signing\&quot;, \&quot;email protection\&quot;, \&quot;s/mime\&quot;,  \&quot;ipsec end system\&quot;, \&quot;ipsec tunnel\&quot;, \&quot;ipsec user\&quot;,  \&quot;timestamping\&quot;, \&quot;ocsp signing\&quot;, \&quot;microsoft sgc\&quot;, \&quot;netscape sgc\&quot; | [optional] 
**Username** | Pointer to **string** | username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable. | [optional] 

## Methods

### NewIoK8sApiCertificatesV1CertificateSigningRequestSpec

`func NewIoK8sApiCertificatesV1CertificateSigningRequestSpec(request string, signerName string, ) *IoK8sApiCertificatesV1CertificateSigningRequestSpec`

NewIoK8sApiCertificatesV1CertificateSigningRequestSpec instantiates a new IoK8sApiCertificatesV1CertificateSigningRequestSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCertificatesV1CertificateSigningRequestSpecWithDefaults

`func NewIoK8sApiCertificatesV1CertificateSigningRequestSpecWithDefaults() *IoK8sApiCertificatesV1CertificateSigningRequestSpec`

NewIoK8sApiCertificatesV1CertificateSigningRequestSpecWithDefaults instantiates a new IoK8sApiCertificatesV1CertificateSigningRequestSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationSeconds

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetExpirationSeconds() int32`

GetExpirationSeconds returns the ExpirationSeconds field if non-nil, zero value otherwise.

### GetExpirationSecondsOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetExpirationSecondsOk() (*int32, bool)`

GetExpirationSecondsOk returns a tuple with the ExpirationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationSeconds

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) SetExpirationSeconds(v int32)`

SetExpirationSeconds sets ExpirationSeconds field to given value.

### HasExpirationSeconds

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) HasExpirationSeconds() bool`

HasExpirationSeconds returns a boolean if a field has been set.

### GetExtra

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetExtra() map[string][]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetExtraOk() (*map[string][]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) SetExtra(v map[string][]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetGroups

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetRequest

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) SetRequest(v string)`

SetRequest sets Request field to given value.


### GetSignerName

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetSignerName() string`

GetSignerName returns the SignerName field if non-nil, zero value otherwise.

### GetSignerNameOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetSignerNameOk() (*string, bool)`

GetSignerNameOk returns a tuple with the SignerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerName

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) SetSignerName(v string)`

SetSignerName sets SignerName field to given value.


### GetUid

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsages

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetUsername

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IoK8sApiCertificatesV1CertificateSigningRequestSpec) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


