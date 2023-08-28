# IoK8sApimachineryPkgApisMetaV1APIVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**ServerAddressByClientCIDRs** | [**[]IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR**](IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR.md) | a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP. | 
**Versions** | **[]string** | versions are the api versions that are available. | 

## Methods

### NewIoK8sApimachineryPkgApisMetaV1APIVersions

`func NewIoK8sApimachineryPkgApisMetaV1APIVersions(serverAddressByClientCIDRs []IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR, versions []string, ) *IoK8sApimachineryPkgApisMetaV1APIVersions`

NewIoK8sApimachineryPkgApisMetaV1APIVersions instantiates a new IoK8sApimachineryPkgApisMetaV1APIVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApimachineryPkgApisMetaV1APIVersionsWithDefaults

`func NewIoK8sApimachineryPkgApisMetaV1APIVersionsWithDefaults() *IoK8sApimachineryPkgApisMetaV1APIVersions`

NewIoK8sApimachineryPkgApisMetaV1APIVersionsWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1APIVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetServerAddressByClientCIDRs

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) GetServerAddressByClientCIDRs() []IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR`

GetServerAddressByClientCIDRs returns the ServerAddressByClientCIDRs field if non-nil, zero value otherwise.

### GetServerAddressByClientCIDRsOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) GetServerAddressByClientCIDRsOk() (*[]IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR, bool)`

GetServerAddressByClientCIDRsOk returns a tuple with the ServerAddressByClientCIDRs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddressByClientCIDRs

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) SetServerAddressByClientCIDRs(v []IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR)`

SetServerAddressByClientCIDRs sets ServerAddressByClientCIDRs field to given value.


### GetVersions

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *IoK8sApimachineryPkgApisMetaV1APIVersions) SetVersions(v []string)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


