# CloudXlineXlineoperatorV1alpha1XlineClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]CloudXlineXlineoperatorV1alpha1XlineCluster**](CloudXlineXlineoperatorV1alpha1XlineCluster.md) | List of xlineclusters. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ListMeta**](IoK8sApimachineryPkgApisMetaV1ListMeta.md) |  | [optional] 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterList

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterList(items []CloudXlineXlineoperatorV1alpha1XlineCluster, ) *CloudXlineXlineoperatorV1alpha1XlineClusterList`

NewCloudXlineXlineoperatorV1alpha1XlineClusterList instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterListWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterListWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterList`

NewCloudXlineXlineoperatorV1alpha1XlineClusterListWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetItems() []CloudXlineXlineoperatorV1alpha1XlineCluster`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetItemsOk() (*[]CloudXlineXlineoperatorV1alpha1XlineCluster, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) SetItems(v []CloudXlineXlineoperatorV1alpha1XlineCluster)`

SetItems sets Items field to given value.


### GetKind

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetMetadata() IoK8sApimachineryPkgApisMetaV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


