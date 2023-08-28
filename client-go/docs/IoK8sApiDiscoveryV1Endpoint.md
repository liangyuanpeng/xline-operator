# IoK8sApiDiscoveryV1Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** | addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100. These are all assumed to be fungible and clients may choose to only use the first element. Refer to: https://issue.k8s.io/106267 | 
**Conditions** | Pointer to [**IoK8sApiDiscoveryV1EndpointConditions**](IoK8sApiDiscoveryV1EndpointConditions.md) |  | [optional] 
**DeprecatedTopology** | Pointer to **map[string]string** | deprecatedTopology contains topology information part of the v1beta1 API. This field is deprecated, and will be removed when the v1beta1 API is removed (no sooner than kubernetes v1.24).  While this field can hold values, it is not writable through the v1 API, and any attempts to write to it will be silently ignored. Topology information can be found in the zone and nodeName fields instead. | [optional] 
**Hints** | Pointer to [**IoK8sApiDiscoveryV1EndpointHints**](IoK8sApiDiscoveryV1EndpointHints.md) |  | [optional] 
**Hostname** | Pointer to **string** | hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must be lowercase and pass DNS Label (RFC 1123) validation. | [optional] 
**NodeName** | Pointer to **string** | nodeName represents the name of the Node hosting this endpoint. This can be used to determine endpoints local to a Node. | [optional] 
**TargetRef** | Pointer to [**IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) |  | [optional] 
**Zone** | Pointer to **string** | zone is the name of the Zone this endpoint exists in. | [optional] 

## Methods

### NewIoK8sApiDiscoveryV1Endpoint

`func NewIoK8sApiDiscoveryV1Endpoint(addresses []string, ) *IoK8sApiDiscoveryV1Endpoint`

NewIoK8sApiDiscoveryV1Endpoint instantiates a new IoK8sApiDiscoveryV1Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiDiscoveryV1EndpointWithDefaults

`func NewIoK8sApiDiscoveryV1EndpointWithDefaults() *IoK8sApiDiscoveryV1Endpoint`

NewIoK8sApiDiscoveryV1EndpointWithDefaults instantiates a new IoK8sApiDiscoveryV1Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *IoK8sApiDiscoveryV1Endpoint) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *IoK8sApiDiscoveryV1Endpoint) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *IoK8sApiDiscoveryV1Endpoint) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetConditions

`func (o *IoK8sApiDiscoveryV1Endpoint) GetConditions() IoK8sApiDiscoveryV1EndpointConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiDiscoveryV1Endpoint) GetConditionsOk() (*IoK8sApiDiscoveryV1EndpointConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiDiscoveryV1Endpoint) SetConditions(v IoK8sApiDiscoveryV1EndpointConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiDiscoveryV1Endpoint) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDeprecatedTopology

`func (o *IoK8sApiDiscoveryV1Endpoint) GetDeprecatedTopology() map[string]string`

GetDeprecatedTopology returns the DeprecatedTopology field if non-nil, zero value otherwise.

### GetDeprecatedTopologyOk

`func (o *IoK8sApiDiscoveryV1Endpoint) GetDeprecatedTopologyOk() (*map[string]string, bool)`

GetDeprecatedTopologyOk returns a tuple with the DeprecatedTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedTopology

`func (o *IoK8sApiDiscoveryV1Endpoint) SetDeprecatedTopology(v map[string]string)`

SetDeprecatedTopology sets DeprecatedTopology field to given value.

### HasDeprecatedTopology

`func (o *IoK8sApiDiscoveryV1Endpoint) HasDeprecatedTopology() bool`

HasDeprecatedTopology returns a boolean if a field has been set.

### GetHints

`func (o *IoK8sApiDiscoveryV1Endpoint) GetHints() IoK8sApiDiscoveryV1EndpointHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *IoK8sApiDiscoveryV1Endpoint) GetHintsOk() (*IoK8sApiDiscoveryV1EndpointHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *IoK8sApiDiscoveryV1Endpoint) SetHints(v IoK8sApiDiscoveryV1EndpointHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *IoK8sApiDiscoveryV1Endpoint) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetHostname

`func (o *IoK8sApiDiscoveryV1Endpoint) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *IoK8sApiDiscoveryV1Endpoint) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *IoK8sApiDiscoveryV1Endpoint) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *IoK8sApiDiscoveryV1Endpoint) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNodeName

`func (o *IoK8sApiDiscoveryV1Endpoint) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *IoK8sApiDiscoveryV1Endpoint) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *IoK8sApiDiscoveryV1Endpoint) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *IoK8sApiDiscoveryV1Endpoint) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetTargetRef

`func (o *IoK8sApiDiscoveryV1Endpoint) GetTargetRef() IoK8sApiCoreV1ObjectReference`

GetTargetRef returns the TargetRef field if non-nil, zero value otherwise.

### GetTargetRefOk

`func (o *IoK8sApiDiscoveryV1Endpoint) GetTargetRefOk() (*IoK8sApiCoreV1ObjectReference, bool)`

GetTargetRefOk returns a tuple with the TargetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRef

`func (o *IoK8sApiDiscoveryV1Endpoint) SetTargetRef(v IoK8sApiCoreV1ObjectReference)`

SetTargetRef sets TargetRef field to given value.

### HasTargetRef

`func (o *IoK8sApiDiscoveryV1Endpoint) HasTargetRef() bool`

HasTargetRef returns a boolean if a field has been set.

### GetZone

`func (o *IoK8sApiDiscoveryV1Endpoint) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *IoK8sApiDiscoveryV1Endpoint) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *IoK8sApiDiscoveryV1Endpoint) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *IoK8sApiDiscoveryV1Endpoint) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


