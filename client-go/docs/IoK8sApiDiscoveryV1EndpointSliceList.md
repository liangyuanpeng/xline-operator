# IoK8sApiDiscoveryV1EndpointSliceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]IoK8sApiDiscoveryV1EndpointSlice**](IoK8sApiDiscoveryV1EndpointSlice.md) | items is the list of endpoint slices | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ListMeta**](IoK8sApimachineryPkgApisMetaV1ListMeta.md) |  | [optional] 

## Methods

### NewIoK8sApiDiscoveryV1EndpointSliceList

`func NewIoK8sApiDiscoveryV1EndpointSliceList(items []IoK8sApiDiscoveryV1EndpointSlice, ) *IoK8sApiDiscoveryV1EndpointSliceList`

NewIoK8sApiDiscoveryV1EndpointSliceList instantiates a new IoK8sApiDiscoveryV1EndpointSliceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiDiscoveryV1EndpointSliceListWithDefaults

`func NewIoK8sApiDiscoveryV1EndpointSliceListWithDefaults() *IoK8sApiDiscoveryV1EndpointSliceList`

NewIoK8sApiDiscoveryV1EndpointSliceListWithDefaults instantiates a new IoK8sApiDiscoveryV1EndpointSliceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) GetItems() []IoK8sApiDiscoveryV1EndpointSlice`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) GetItemsOk() (*[]IoK8sApiDiscoveryV1EndpointSlice, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) SetItems(v []IoK8sApiDiscoveryV1EndpointSlice)`

SetItems sets Items field to given value.


### GetKind

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) GetMetadata() IoK8sApimachineryPkgApisMetaV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiDiscoveryV1EndpointSliceList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


