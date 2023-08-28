# IoK8sApiRbacV1Subject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroup** | Pointer to **string** | APIGroup holds the API group of the referenced subject. Defaults to \&quot;\&quot; for ServiceAccount subjects. Defaults to \&quot;rbac.authorization.k8s.io\&quot; for User and Group subjects. | [optional] 
**Kind** | **string** | Kind of object being referenced. Values defined by this API group are \&quot;User\&quot;, \&quot;Group\&quot;, and \&quot;ServiceAccount\&quot;. If the Authorizer does not recognized the kind value, the Authorizer should report an error. | 
**Name** | **string** | Name of the object being referenced. | 
**Namespace** | Pointer to **string** | Namespace of the referenced object.  If the object kind is non-namespace, such as \&quot;User\&quot; or \&quot;Group\&quot;, and this value is not empty the Authorizer should report an error. | [optional] 

## Methods

### NewIoK8sApiRbacV1Subject

`func NewIoK8sApiRbacV1Subject(kind string, name string, ) *IoK8sApiRbacV1Subject`

NewIoK8sApiRbacV1Subject instantiates a new IoK8sApiRbacV1Subject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiRbacV1SubjectWithDefaults

`func NewIoK8sApiRbacV1SubjectWithDefaults() *IoK8sApiRbacV1Subject`

NewIoK8sApiRbacV1SubjectWithDefaults instantiates a new IoK8sApiRbacV1Subject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroup

`func (o *IoK8sApiRbacV1Subject) GetApiGroup() string`

GetApiGroup returns the ApiGroup field if non-nil, zero value otherwise.

### GetApiGroupOk

`func (o *IoK8sApiRbacV1Subject) GetApiGroupOk() (*string, bool)`

GetApiGroupOk returns a tuple with the ApiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroup

`func (o *IoK8sApiRbacV1Subject) SetApiGroup(v string)`

SetApiGroup sets ApiGroup field to given value.

### HasApiGroup

`func (o *IoK8sApiRbacV1Subject) HasApiGroup() bool`

HasApiGroup returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiRbacV1Subject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiRbacV1Subject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiRbacV1Subject) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *IoK8sApiRbacV1Subject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiRbacV1Subject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiRbacV1Subject) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *IoK8sApiRbacV1Subject) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoK8sApiRbacV1Subject) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoK8sApiRbacV1Subject) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoK8sApiRbacV1Subject) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


