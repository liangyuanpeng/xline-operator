# IoK8sApiRbacV1ClusterRoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**RoleRef** | [**IoK8sApiRbacV1RoleRef**](IoK8sApiRbacV1RoleRef.md) |  | 
**Subjects** | Pointer to [**[]IoK8sApiRbacV1Subject**](IoK8sApiRbacV1Subject.md) | Subjects holds references to the objects the role applies to. | [optional] 

## Methods

### NewIoK8sApiRbacV1ClusterRoleBinding

`func NewIoK8sApiRbacV1ClusterRoleBinding(roleRef IoK8sApiRbacV1RoleRef, ) *IoK8sApiRbacV1ClusterRoleBinding`

NewIoK8sApiRbacV1ClusterRoleBinding instantiates a new IoK8sApiRbacV1ClusterRoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiRbacV1ClusterRoleBindingWithDefaults

`func NewIoK8sApiRbacV1ClusterRoleBindingWithDefaults() *IoK8sApiRbacV1ClusterRoleBinding`

NewIoK8sApiRbacV1ClusterRoleBindingWithDefaults instantiates a new IoK8sApiRbacV1ClusterRoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiRbacV1ClusterRoleBinding) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiRbacV1ClusterRoleBinding) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiRbacV1ClusterRoleBinding) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiRbacV1ClusterRoleBinding) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiRbacV1ClusterRoleBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiRbacV1ClusterRoleBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiRbacV1ClusterRoleBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiRbacV1ClusterRoleBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiRbacV1ClusterRoleBinding) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiRbacV1ClusterRoleBinding) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiRbacV1ClusterRoleBinding) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiRbacV1ClusterRoleBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRoleRef

`func (o *IoK8sApiRbacV1ClusterRoleBinding) GetRoleRef() IoK8sApiRbacV1RoleRef`

GetRoleRef returns the RoleRef field if non-nil, zero value otherwise.

### GetRoleRefOk

`func (o *IoK8sApiRbacV1ClusterRoleBinding) GetRoleRefOk() (*IoK8sApiRbacV1RoleRef, bool)`

GetRoleRefOk returns a tuple with the RoleRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleRef

`func (o *IoK8sApiRbacV1ClusterRoleBinding) SetRoleRef(v IoK8sApiRbacV1RoleRef)`

SetRoleRef sets RoleRef field to given value.


### GetSubjects

`func (o *IoK8sApiRbacV1ClusterRoleBinding) GetSubjects() []IoK8sApiRbacV1Subject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *IoK8sApiRbacV1ClusterRoleBinding) GetSubjectsOk() (*[]IoK8sApiRbacV1Subject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *IoK8sApiRbacV1ClusterRoleBinding) SetSubjects(v []IoK8sApiRbacV1Subject)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *IoK8sApiRbacV1ClusterRoleBinding) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


