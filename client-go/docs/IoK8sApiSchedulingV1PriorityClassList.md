# IoK8sApiSchedulingV1PriorityClassList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]IoK8sApiSchedulingV1PriorityClass**](IoK8sApiSchedulingV1PriorityClass.md) | items is the list of PriorityClasses | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ListMeta**](IoK8sApimachineryPkgApisMetaV1ListMeta.md) |  | [optional] 

## Methods

### NewIoK8sApiSchedulingV1PriorityClassList

`func NewIoK8sApiSchedulingV1PriorityClassList(items []IoK8sApiSchedulingV1PriorityClass, ) *IoK8sApiSchedulingV1PriorityClassList`

NewIoK8sApiSchedulingV1PriorityClassList instantiates a new IoK8sApiSchedulingV1PriorityClassList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiSchedulingV1PriorityClassListWithDefaults

`func NewIoK8sApiSchedulingV1PriorityClassListWithDefaults() *IoK8sApiSchedulingV1PriorityClassList`

NewIoK8sApiSchedulingV1PriorityClassListWithDefaults instantiates a new IoK8sApiSchedulingV1PriorityClassList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiSchedulingV1PriorityClassList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiSchedulingV1PriorityClassList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiSchedulingV1PriorityClassList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiSchedulingV1PriorityClassList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *IoK8sApiSchedulingV1PriorityClassList) GetItems() []IoK8sApiSchedulingV1PriorityClass`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoK8sApiSchedulingV1PriorityClassList) GetItemsOk() (*[]IoK8sApiSchedulingV1PriorityClass, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoK8sApiSchedulingV1PriorityClassList) SetItems(v []IoK8sApiSchedulingV1PriorityClass)`

SetItems sets Items field to given value.


### GetKind

`func (o *IoK8sApiSchedulingV1PriorityClassList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiSchedulingV1PriorityClassList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiSchedulingV1PriorityClassList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiSchedulingV1PriorityClassList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiSchedulingV1PriorityClassList) GetMetadata() IoK8sApimachineryPkgApisMetaV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiSchedulingV1PriorityClassList) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiSchedulingV1PriorityClassList) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiSchedulingV1PriorityClassList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


