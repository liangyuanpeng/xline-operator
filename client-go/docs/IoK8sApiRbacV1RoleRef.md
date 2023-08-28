# IoK8sApiRbacV1RoleRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroup** | **string** | APIGroup is the group for the resource being referenced | 
**Kind** | **string** | Kind is the type of resource being referenced | 
**Name** | **string** | Name is the name of resource being referenced | 

## Methods

### NewIoK8sApiRbacV1RoleRef

`func NewIoK8sApiRbacV1RoleRef(apiGroup string, kind string, name string, ) *IoK8sApiRbacV1RoleRef`

NewIoK8sApiRbacV1RoleRef instantiates a new IoK8sApiRbacV1RoleRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiRbacV1RoleRefWithDefaults

`func NewIoK8sApiRbacV1RoleRefWithDefaults() *IoK8sApiRbacV1RoleRef`

NewIoK8sApiRbacV1RoleRefWithDefaults instantiates a new IoK8sApiRbacV1RoleRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroup

`func (o *IoK8sApiRbacV1RoleRef) GetApiGroup() string`

GetApiGroup returns the ApiGroup field if non-nil, zero value otherwise.

### GetApiGroupOk

`func (o *IoK8sApiRbacV1RoleRef) GetApiGroupOk() (*string, bool)`

GetApiGroupOk returns a tuple with the ApiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroup

`func (o *IoK8sApiRbacV1RoleRef) SetApiGroup(v string)`

SetApiGroup sets ApiGroup field to given value.


### GetKind

`func (o *IoK8sApiRbacV1RoleRef) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiRbacV1RoleRef) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiRbacV1RoleRef) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *IoK8sApiRbacV1RoleRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiRbacV1RoleRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiRbacV1RoleRef) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


