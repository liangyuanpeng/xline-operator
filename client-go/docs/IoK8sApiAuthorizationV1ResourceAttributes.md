# IoK8sApiAuthorizationV1ResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Group is the API Group of the Resource.  \&quot;*\&quot; means all. | [optional] 
**Name** | Pointer to **string** | Name is the name of the resource being requested for a \&quot;get\&quot; or deleted for a \&quot;delete\&quot;. \&quot;\&quot; (empty) means all. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces \&quot;\&quot; (empty) is defaulted for LocalSubjectAccessReviews \&quot;\&quot; (empty) is empty for cluster-scoped resources \&quot;\&quot; (empty) means \&quot;all\&quot; for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview | [optional] 
**Resource** | Pointer to **string** | Resource is one of the existing resource types.  \&quot;*\&quot; means all. | [optional] 
**Subresource** | Pointer to **string** | Subresource is one of the existing resource types.  \&quot;\&quot; means none. | [optional] 
**Verb** | Pointer to **string** | Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  \&quot;*\&quot; means all. | [optional] 
**Version** | Pointer to **string** | Version is the API Version of the Resource.  \&quot;*\&quot; means all. | [optional] 

## Methods

### NewIoK8sApiAuthorizationV1ResourceAttributes

`func NewIoK8sApiAuthorizationV1ResourceAttributes() *IoK8sApiAuthorizationV1ResourceAttributes`

NewIoK8sApiAuthorizationV1ResourceAttributes instantiates a new IoK8sApiAuthorizationV1ResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthorizationV1ResourceAttributesWithDefaults

`func NewIoK8sApiAuthorizationV1ResourceAttributesWithDefaults() *IoK8sApiAuthorizationV1ResourceAttributes`

NewIoK8sApiAuthorizationV1ResourceAttributesWithDefaults instantiates a new IoK8sApiAuthorizationV1ResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetResource

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSubresource

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetSubresource() string`

GetSubresource returns the Subresource field if non-nil, zero value otherwise.

### GetSubresourceOk

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetSubresourceOk() (*string, bool)`

GetSubresourceOk returns a tuple with the Subresource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresource

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) SetSubresource(v string)`

SetSubresource sets Subresource field to given value.

### HasSubresource

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) HasSubresource() bool`

HasSubresource returns a boolean if a field has been set.

### GetVerb

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### GetVersion

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IoK8sApiAuthorizationV1ResourceAttributes) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


