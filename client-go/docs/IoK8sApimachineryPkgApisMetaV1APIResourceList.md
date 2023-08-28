# IoK8sApimachineryPkgApisMetaV1APIResourceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**GroupVersion** | **string** | groupVersion is the group and version this APIResourceList is for. | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Resources** | [**[]IoK8sApimachineryPkgApisMetaV1APIResource**](IoK8sApimachineryPkgApisMetaV1APIResource.md) | resources contains the name of the resources and if they are namespaced. | 

## Methods

### NewIoK8sApimachineryPkgApisMetaV1APIResourceList

`func NewIoK8sApimachineryPkgApisMetaV1APIResourceList(groupVersion string, resources []IoK8sApimachineryPkgApisMetaV1APIResource, ) *IoK8sApimachineryPkgApisMetaV1APIResourceList`

NewIoK8sApimachineryPkgApisMetaV1APIResourceList instantiates a new IoK8sApimachineryPkgApisMetaV1APIResourceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApimachineryPkgApisMetaV1APIResourceListWithDefaults

`func NewIoK8sApimachineryPkgApisMetaV1APIResourceListWithDefaults() *IoK8sApimachineryPkgApisMetaV1APIResourceList`

NewIoK8sApimachineryPkgApisMetaV1APIResourceListWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1APIResourceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetGroupVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) GetGroupVersion() string`

GetGroupVersion returns the GroupVersion field if non-nil, zero value otherwise.

### GetGroupVersionOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) GetGroupVersionOk() (*string, bool)`

GetGroupVersionOk returns a tuple with the GroupVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupVersion

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) SetGroupVersion(v string)`

SetGroupVersion sets GroupVersion field to given value.


### GetKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetResources

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) GetResources() []IoK8sApimachineryPkgApisMetaV1APIResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) GetResourcesOk() (*[]IoK8sApimachineryPkgApisMetaV1APIResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IoK8sApimachineryPkgApisMetaV1APIResourceList) SetResources(v []IoK8sApimachineryPkgApisMetaV1APIResource)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


