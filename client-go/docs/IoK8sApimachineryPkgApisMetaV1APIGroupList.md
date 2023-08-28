# IoK8sApimachineryPkgApisMetaV1APIGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Groups** | [**[]IoK8sApimachineryPkgApisMetaV1APIGroup**](IoK8sApimachineryPkgApisMetaV1APIGroup.md) | groups is a list of APIGroup. | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 

## Methods

### NewIoK8sApimachineryPkgApisMetaV1APIGroupList

`func NewIoK8sApimachineryPkgApisMetaV1APIGroupList(groups []IoK8sApimachineryPkgApisMetaV1APIGroup, ) *IoK8sApimachineryPkgApisMetaV1APIGroupList`

NewIoK8sApimachineryPkgApisMetaV1APIGroupList instantiates a new IoK8sApimachineryPkgApisMetaV1APIGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApimachineryPkgApisMetaV1APIGroupListWithDefaults

`func NewIoK8sApimachineryPkgApisMetaV1APIGroupListWithDefaults() *IoK8sApimachineryPkgApisMetaV1APIGroupList`

NewIoK8sApimachineryPkgApisMetaV1APIGroupListWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1APIGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetGroups

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) GetGroups() []IoK8sApimachineryPkgApisMetaV1APIGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) GetGroupsOk() (*[]IoK8sApimachineryPkgApisMetaV1APIGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) SetGroups(v []IoK8sApimachineryPkgApisMetaV1APIGroup)`

SetGroups sets Groups field to given value.


### GetKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIGroupList) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


