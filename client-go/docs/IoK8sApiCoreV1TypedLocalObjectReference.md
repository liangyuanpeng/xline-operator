# IoK8sApiCoreV1TypedLocalObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroup** | Pointer to **string** | APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required. | [optional] 
**Kind** | **string** | Kind is the type of resource being referenced | 
**Name** | **string** | Name is the name of resource being referenced | 

## Methods

### NewIoK8sApiCoreV1TypedLocalObjectReference

`func NewIoK8sApiCoreV1TypedLocalObjectReference(kind string, name string, ) *IoK8sApiCoreV1TypedLocalObjectReference`

NewIoK8sApiCoreV1TypedLocalObjectReference instantiates a new IoK8sApiCoreV1TypedLocalObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1TypedLocalObjectReferenceWithDefaults

`func NewIoK8sApiCoreV1TypedLocalObjectReferenceWithDefaults() *IoK8sApiCoreV1TypedLocalObjectReference`

NewIoK8sApiCoreV1TypedLocalObjectReferenceWithDefaults instantiates a new IoK8sApiCoreV1TypedLocalObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroup

`func (o *IoK8sApiCoreV1TypedLocalObjectReference) GetApiGroup() string`

GetApiGroup returns the ApiGroup field if non-nil, zero value otherwise.

### GetApiGroupOk

`func (o *IoK8sApiCoreV1TypedLocalObjectReference) GetApiGroupOk() (*string, bool)`

GetApiGroupOk returns a tuple with the ApiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroup

`func (o *IoK8sApiCoreV1TypedLocalObjectReference) SetApiGroup(v string)`

SetApiGroup sets ApiGroup field to given value.

### HasApiGroup

`func (o *IoK8sApiCoreV1TypedLocalObjectReference) HasApiGroup() bool`

HasApiGroup returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiCoreV1TypedLocalObjectReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiCoreV1TypedLocalObjectReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiCoreV1TypedLocalObjectReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *IoK8sApiCoreV1TypedLocalObjectReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1TypedLocalObjectReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1TypedLocalObjectReference) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


