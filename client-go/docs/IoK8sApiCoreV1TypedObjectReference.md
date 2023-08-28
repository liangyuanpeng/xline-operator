# IoK8sApiCoreV1TypedObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroup** | Pointer to **string** | APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required. | [optional] 
**Kind** | **string** | Kind is the type of resource being referenced | 
**Name** | **string** | Name is the name of resource being referenced | 
**Namespace** | Pointer to **string** | Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace&#39;s owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled. | [optional] 

## Methods

### NewIoK8sApiCoreV1TypedObjectReference

`func NewIoK8sApiCoreV1TypedObjectReference(kind string, name string, ) *IoK8sApiCoreV1TypedObjectReference`

NewIoK8sApiCoreV1TypedObjectReference instantiates a new IoK8sApiCoreV1TypedObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1TypedObjectReferenceWithDefaults

`func NewIoK8sApiCoreV1TypedObjectReferenceWithDefaults() *IoK8sApiCoreV1TypedObjectReference`

NewIoK8sApiCoreV1TypedObjectReferenceWithDefaults instantiates a new IoK8sApiCoreV1TypedObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroup

`func (o *IoK8sApiCoreV1TypedObjectReference) GetApiGroup() string`

GetApiGroup returns the ApiGroup field if non-nil, zero value otherwise.

### GetApiGroupOk

`func (o *IoK8sApiCoreV1TypedObjectReference) GetApiGroupOk() (*string, bool)`

GetApiGroupOk returns a tuple with the ApiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroup

`func (o *IoK8sApiCoreV1TypedObjectReference) SetApiGroup(v string)`

SetApiGroup sets ApiGroup field to given value.

### HasApiGroup

`func (o *IoK8sApiCoreV1TypedObjectReference) HasApiGroup() bool`

HasApiGroup returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiCoreV1TypedObjectReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiCoreV1TypedObjectReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiCoreV1TypedObjectReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *IoK8sApiCoreV1TypedObjectReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1TypedObjectReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1TypedObjectReference) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *IoK8sApiCoreV1TypedObjectReference) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoK8sApiCoreV1TypedObjectReference) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoK8sApiCoreV1TypedObjectReference) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoK8sApiCoreV1TypedObjectReference) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


