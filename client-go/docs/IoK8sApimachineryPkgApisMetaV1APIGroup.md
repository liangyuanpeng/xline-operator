# IoK8sApimachineryPkgApisMetaV1APIGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Name** | **string** | name is the name of the group. | 
**PreferredVersion** | Pointer to [**IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery**](IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery.md) |  | [optional] 
**ServerAddressByClientCIDRs** | Pointer to [**[]IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR**](IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR.md) | a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP. | [optional] 
**Versions** | [**[]IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery**](IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery.md) | versions are the versions supported in this group. | 

## Methods

### NewIoK8sApimachineryPkgApisMetaV1APIGroup

`func NewIoK8sApimachineryPkgApisMetaV1APIGroup(name string, versions []IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery, ) *IoK8sApimachineryPkgApisMetaV1APIGroup`

NewIoK8sApimachineryPkgApisMetaV1APIGroup instantiates a new IoK8sApimachineryPkgApisMetaV1APIGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApimachineryPkgApisMetaV1APIGroupWithDefaults

`func NewIoK8sApimachineryPkgApisMetaV1APIGroupWithDefaults() *IoK8sApimachineryPkgApisMetaV1APIGroup`

NewIoK8sApimachineryPkgApisMetaV1APIGroupWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1APIGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) SetName(v string)`

SetName sets Name field to given value.


### GetPreferredVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetPreferredVersion() IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery`

GetPreferredVersion returns the PreferredVersion field if non-nil, zero value otherwise.

### GetPreferredVersionOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetPreferredVersionOk() (*IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery, bool)`

GetPreferredVersionOk returns a tuple with the PreferredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) SetPreferredVersion(v IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery)`

SetPreferredVersion sets PreferredVersion field to given value.

### HasPreferredVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) HasPreferredVersion() bool`

HasPreferredVersion returns a boolean if a field has been set.

### GetServerAddressByClientCIDRs

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetServerAddressByClientCIDRs() []IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR`

GetServerAddressByClientCIDRs returns the ServerAddressByClientCIDRs field if non-nil, zero value otherwise.

### GetServerAddressByClientCIDRsOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetServerAddressByClientCIDRsOk() (*[]IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR, bool)`

GetServerAddressByClientCIDRsOk returns a tuple with the ServerAddressByClientCIDRs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddressByClientCIDRs

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) SetServerAddressByClientCIDRs(v []IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR)`

SetServerAddressByClientCIDRs sets ServerAddressByClientCIDRs field to given value.

### HasServerAddressByClientCIDRs

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) HasServerAddressByClientCIDRs() bool`

HasServerAddressByClientCIDRs returns a boolean if a field has been set.

### GetVersions

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetVersions() []IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) GetVersionsOk() (*[]IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroup) SetVersions(v []IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


