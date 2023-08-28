# IoK8sApiDiscoveryV1EndpointSlice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressType** | **string** | addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.  Possible enum values:  - &#x60;\&quot;FQDN\&quot;&#x60; represents a FQDN.  - &#x60;\&quot;IPv4\&quot;&#x60; represents an IPv4 Address.  - &#x60;\&quot;IPv6\&quot;&#x60; represents an IPv6 Address. | 
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Endpoints** | [**[]IoK8sApiDiscoveryV1Endpoint**](IoK8sApiDiscoveryV1Endpoint.md) | endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints. | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Ports** | Pointer to [**[]IoK8sApiDiscoveryV1EndpointPort**](IoK8sApiDiscoveryV1EndpointPort.md) | ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates \&quot;all ports\&quot;. Each slice may include a maximum of 100 ports. | [optional] 

## Methods

### NewIoK8sApiDiscoveryV1EndpointSlice

`func NewIoK8sApiDiscoveryV1EndpointSlice(addressType string, endpoints []IoK8sApiDiscoveryV1Endpoint, ) *IoK8sApiDiscoveryV1EndpointSlice`

NewIoK8sApiDiscoveryV1EndpointSlice instantiates a new IoK8sApiDiscoveryV1EndpointSlice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiDiscoveryV1EndpointSliceWithDefaults

`func NewIoK8sApiDiscoveryV1EndpointSliceWithDefaults() *IoK8sApiDiscoveryV1EndpointSlice`

NewIoK8sApiDiscoveryV1EndpointSliceWithDefaults instantiates a new IoK8sApiDiscoveryV1EndpointSlice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressType

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *IoK8sApiDiscoveryV1EndpointSlice) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.


### GetApiVersion

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiDiscoveryV1EndpointSlice) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiDiscoveryV1EndpointSlice) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetEndpoints

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetEndpoints() []IoK8sApiDiscoveryV1Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetEndpointsOk() (*[]IoK8sApiDiscoveryV1Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *IoK8sApiDiscoveryV1EndpointSlice) SetEndpoints(v []IoK8sApiDiscoveryV1Endpoint)`

SetEndpoints sets Endpoints field to given value.


### GetKind

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiDiscoveryV1EndpointSlice) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiDiscoveryV1EndpointSlice) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiDiscoveryV1EndpointSlice) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiDiscoveryV1EndpointSlice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetPorts() []IoK8sApiDiscoveryV1EndpointPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *IoK8sApiDiscoveryV1EndpointSlice) GetPortsOk() (*[]IoK8sApiDiscoveryV1EndpointPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *IoK8sApiDiscoveryV1EndpointSlice) SetPorts(v []IoK8sApiDiscoveryV1EndpointPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *IoK8sApiDiscoveryV1EndpointSlice) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


