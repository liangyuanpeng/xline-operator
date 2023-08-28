# IoK8sApiNetworkingV1IngressClassList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]IoK8sApiNetworkingV1IngressClass**](IoK8sApiNetworkingV1IngressClass.md) | items is the list of IngressClasses. | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ListMeta**](IoK8sApimachineryPkgApisMetaV1ListMeta.md) |  | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IngressClassList

`func NewIoK8sApiNetworkingV1IngressClassList(items []IoK8sApiNetworkingV1IngressClass, ) *IoK8sApiNetworkingV1IngressClassList`

NewIoK8sApiNetworkingV1IngressClassList instantiates a new IoK8sApiNetworkingV1IngressClassList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IngressClassListWithDefaults

`func NewIoK8sApiNetworkingV1IngressClassListWithDefaults() *IoK8sApiNetworkingV1IngressClassList`

NewIoK8sApiNetworkingV1IngressClassListWithDefaults instantiates a new IoK8sApiNetworkingV1IngressClassList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiNetworkingV1IngressClassList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiNetworkingV1IngressClassList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiNetworkingV1IngressClassList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiNetworkingV1IngressClassList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *IoK8sApiNetworkingV1IngressClassList) GetItems() []IoK8sApiNetworkingV1IngressClass`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoK8sApiNetworkingV1IngressClassList) GetItemsOk() (*[]IoK8sApiNetworkingV1IngressClass, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoK8sApiNetworkingV1IngressClassList) SetItems(v []IoK8sApiNetworkingV1IngressClass)`

SetItems sets Items field to given value.


### GetKind

`func (o *IoK8sApiNetworkingV1IngressClassList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiNetworkingV1IngressClassList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiNetworkingV1IngressClassList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiNetworkingV1IngressClassList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiNetworkingV1IngressClassList) GetMetadata() IoK8sApimachineryPkgApisMetaV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiNetworkingV1IngressClassList) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiNetworkingV1IngressClassList) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiNetworkingV1IngressClassList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


