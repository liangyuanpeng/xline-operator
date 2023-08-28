# IoK8sApiNetworkingV1IngressClassParametersReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroup** | Pointer to **string** | apiGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required. | [optional] 
**Kind** | **string** | kind is the type of resource being referenced. | 
**Name** | **string** | name is the name of resource being referenced. | 
**Namespace** | Pointer to **string** | namespace is the namespace of the resource being referenced. This field is required when scope is set to \&quot;Namespace\&quot; and must be unset when scope is set to \&quot;Cluster\&quot;. | [optional] 
**Scope** | Pointer to **string** | scope represents if this refers to a cluster or namespace scoped resource. This may be set to \&quot;Cluster\&quot; (default) or \&quot;Namespace\&quot;. | [optional] 

## Methods

### NewIoK8sApiNetworkingV1IngressClassParametersReference

`func NewIoK8sApiNetworkingV1IngressClassParametersReference(kind string, name string, ) *IoK8sApiNetworkingV1IngressClassParametersReference`

NewIoK8sApiNetworkingV1IngressClassParametersReference instantiates a new IoK8sApiNetworkingV1IngressClassParametersReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1IngressClassParametersReferenceWithDefaults

`func NewIoK8sApiNetworkingV1IngressClassParametersReferenceWithDefaults() *IoK8sApiNetworkingV1IngressClassParametersReference`

NewIoK8sApiNetworkingV1IngressClassParametersReferenceWithDefaults instantiates a new IoK8sApiNetworkingV1IngressClassParametersReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroup

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) GetApiGroup() string`

GetApiGroup returns the ApiGroup field if non-nil, zero value otherwise.

### GetApiGroupOk

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) GetApiGroupOk() (*string, bool)`

GetApiGroupOk returns a tuple with the ApiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroup

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) SetApiGroup(v string)`

SetApiGroup sets ApiGroup field to given value.

### HasApiGroup

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) HasApiGroup() bool`

HasApiGroup returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetScope

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IoK8sApiNetworkingV1IngressClassParametersReference) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


