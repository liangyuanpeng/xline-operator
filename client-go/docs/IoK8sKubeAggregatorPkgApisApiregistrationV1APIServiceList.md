# IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]IoK8sKubeAggregatorPkgApisApiregistrationV1APIService**](IoK8sKubeAggregatorPkgApisApiregistrationV1APIService.md) | Items is the list of APIService | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ListMeta**](IoK8sApimachineryPkgApisMetaV1ListMeta.md) |  | [optional] 

## Methods

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList(items []IoK8sKubeAggregatorPkgApisApiregistrationV1APIService, ) *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceListWithDefaults

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceListWithDefaults() *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceListWithDefaults instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) GetItems() []IoK8sKubeAggregatorPkgApisApiregistrationV1APIService`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) GetItemsOk() (*[]IoK8sKubeAggregatorPkgApisApiregistrationV1APIService, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) SetItems(v []IoK8sKubeAggregatorPkgApisApiregistrationV1APIService)`

SetItems sets Items field to given value.


### GetKind

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) GetMetadata() IoK8sApimachineryPkgApisMetaV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


