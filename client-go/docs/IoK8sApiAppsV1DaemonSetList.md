# IoK8sApiAppsV1DaemonSetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]IoK8sApiAppsV1DaemonSet**](IoK8sApiAppsV1DaemonSet.md) | A list of daemon sets. | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ListMeta**](IoK8sApimachineryPkgApisMetaV1ListMeta.md) |  | [optional] 

## Methods

### NewIoK8sApiAppsV1DaemonSetList

`func NewIoK8sApiAppsV1DaemonSetList(items []IoK8sApiAppsV1DaemonSet, ) *IoK8sApiAppsV1DaemonSetList`

NewIoK8sApiAppsV1DaemonSetList instantiates a new IoK8sApiAppsV1DaemonSetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1DaemonSetListWithDefaults

`func NewIoK8sApiAppsV1DaemonSetListWithDefaults() *IoK8sApiAppsV1DaemonSetList`

NewIoK8sApiAppsV1DaemonSetListWithDefaults instantiates a new IoK8sApiAppsV1DaemonSetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiAppsV1DaemonSetList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiAppsV1DaemonSetList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiAppsV1DaemonSetList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiAppsV1DaemonSetList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *IoK8sApiAppsV1DaemonSetList) GetItems() []IoK8sApiAppsV1DaemonSet`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoK8sApiAppsV1DaemonSetList) GetItemsOk() (*[]IoK8sApiAppsV1DaemonSet, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoK8sApiAppsV1DaemonSetList) SetItems(v []IoK8sApiAppsV1DaemonSet)`

SetItems sets Items field to given value.


### GetKind

`func (o *IoK8sApiAppsV1DaemonSetList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiAppsV1DaemonSetList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiAppsV1DaemonSetList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiAppsV1DaemonSetList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiAppsV1DaemonSetList) GetMetadata() IoK8sApimachineryPkgApisMetaV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiAppsV1DaemonSetList) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiAppsV1DaemonSetList) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiAppsV1DaemonSetList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


