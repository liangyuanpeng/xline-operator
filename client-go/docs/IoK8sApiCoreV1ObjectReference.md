# IoK8sApiCoreV1ObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API version of the referent. | [optional] 
**FieldPath** | Pointer to **string** | If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: \&quot;spec.containers{name}\&quot; (where \&quot;name\&quot; refers to the name of the container that triggered the event) or if no container name is specified \&quot;spec.containers[2]\&quot; (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. | [optional] 
**Kind** | Pointer to **string** | Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Name** | Pointer to **string** | Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [optional] 
**Namespace** | Pointer to **string** | Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ | [optional] 
**ResourceVersion** | Pointer to **string** | Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency | [optional] 
**Uid** | Pointer to **string** | UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids | [optional] 

## Methods

### NewIoK8sApiCoreV1ObjectReference

`func NewIoK8sApiCoreV1ObjectReference() *IoK8sApiCoreV1ObjectReference`

NewIoK8sApiCoreV1ObjectReference instantiates a new IoK8sApiCoreV1ObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ObjectReferenceWithDefaults

`func NewIoK8sApiCoreV1ObjectReferenceWithDefaults() *IoK8sApiCoreV1ObjectReference`

NewIoK8sApiCoreV1ObjectReferenceWithDefaults instantiates a new IoK8sApiCoreV1ObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiCoreV1ObjectReference) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiCoreV1ObjectReference) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiCoreV1ObjectReference) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiCoreV1ObjectReference) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetFieldPath

`func (o *IoK8sApiCoreV1ObjectReference) GetFieldPath() string`

GetFieldPath returns the FieldPath field if non-nil, zero value otherwise.

### GetFieldPathOk

`func (o *IoK8sApiCoreV1ObjectReference) GetFieldPathOk() (*string, bool)`

GetFieldPathOk returns a tuple with the FieldPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldPath

`func (o *IoK8sApiCoreV1ObjectReference) SetFieldPath(v string)`

SetFieldPath sets FieldPath field to given value.

### HasFieldPath

`func (o *IoK8sApiCoreV1ObjectReference) HasFieldPath() bool`

HasFieldPath returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiCoreV1ObjectReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiCoreV1ObjectReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiCoreV1ObjectReference) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiCoreV1ObjectReference) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiCoreV1ObjectReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1ObjectReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1ObjectReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiCoreV1ObjectReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *IoK8sApiCoreV1ObjectReference) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoK8sApiCoreV1ObjectReference) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoK8sApiCoreV1ObjectReference) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoK8sApiCoreV1ObjectReference) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetResourceVersion

`func (o *IoK8sApiCoreV1ObjectReference) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *IoK8sApiCoreV1ObjectReference) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *IoK8sApiCoreV1ObjectReference) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *IoK8sApiCoreV1ObjectReference) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetUid

`func (o *IoK8sApiCoreV1ObjectReference) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IoK8sApiCoreV1ObjectReference) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IoK8sApiCoreV1ObjectReference) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IoK8sApiCoreV1ObjectReference) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


