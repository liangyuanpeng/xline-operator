# IoK8sApiDiscoveryV1EndpointHints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForZones** | Pointer to [**[]IoK8sApiDiscoveryV1ForZone**](IoK8sApiDiscoveryV1ForZone.md) | forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing. | [optional] 

## Methods

### NewIoK8sApiDiscoveryV1EndpointHints

`func NewIoK8sApiDiscoveryV1EndpointHints() *IoK8sApiDiscoveryV1EndpointHints`

NewIoK8sApiDiscoveryV1EndpointHints instantiates a new IoK8sApiDiscoveryV1EndpointHints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiDiscoveryV1EndpointHintsWithDefaults

`func NewIoK8sApiDiscoveryV1EndpointHintsWithDefaults() *IoK8sApiDiscoveryV1EndpointHints`

NewIoK8sApiDiscoveryV1EndpointHintsWithDefaults instantiates a new IoK8sApiDiscoveryV1EndpointHints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForZones

`func (o *IoK8sApiDiscoveryV1EndpointHints) GetForZones() []IoK8sApiDiscoveryV1ForZone`

GetForZones returns the ForZones field if non-nil, zero value otherwise.

### GetForZonesOk

`func (o *IoK8sApiDiscoveryV1EndpointHints) GetForZonesOk() (*[]IoK8sApiDiscoveryV1ForZone, bool)`

GetForZonesOk returns a tuple with the ForZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForZones

`func (o *IoK8sApiDiscoveryV1EndpointHints) SetForZones(v []IoK8sApiDiscoveryV1ForZone)`

SetForZones sets ForZones field to given value.

### HasForZones

`func (o *IoK8sApiDiscoveryV1EndpointHints) HasForZones() bool`

HasForZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


