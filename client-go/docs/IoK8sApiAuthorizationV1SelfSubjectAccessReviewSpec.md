# IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NonResourceAttributes** | Pointer to [**IoK8sApiAuthorizationV1NonResourceAttributes**](IoK8sApiAuthorizationV1NonResourceAttributes.md) |  | [optional] 
**ResourceAttributes** | Pointer to [**IoK8sApiAuthorizationV1ResourceAttributes**](IoK8sApiAuthorizationV1ResourceAttributes.md) |  | [optional] 

## Methods

### NewIoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec

`func NewIoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec() *IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec`

NewIoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec instantiates a new IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthorizationV1SelfSubjectAccessReviewSpecWithDefaults

`func NewIoK8sApiAuthorizationV1SelfSubjectAccessReviewSpecWithDefaults() *IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec`

NewIoK8sApiAuthorizationV1SelfSubjectAccessReviewSpecWithDefaults instantiates a new IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonResourceAttributes

`func (o *IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec) GetNonResourceAttributes() IoK8sApiAuthorizationV1NonResourceAttributes`

GetNonResourceAttributes returns the NonResourceAttributes field if non-nil, zero value otherwise.

### GetNonResourceAttributesOk

`func (o *IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec) GetNonResourceAttributesOk() (*IoK8sApiAuthorizationV1NonResourceAttributes, bool)`

GetNonResourceAttributesOk returns a tuple with the NonResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceAttributes

`func (o *IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec) SetNonResourceAttributes(v IoK8sApiAuthorizationV1NonResourceAttributes)`

SetNonResourceAttributes sets NonResourceAttributes field to given value.

### HasNonResourceAttributes

`func (o *IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec) HasNonResourceAttributes() bool`

HasNonResourceAttributes returns a boolean if a field has been set.

### GetResourceAttributes

`func (o *IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec) GetResourceAttributes() IoK8sApiAuthorizationV1ResourceAttributes`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec) GetResourceAttributesOk() (*IoK8sApiAuthorizationV1ResourceAttributes, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec) SetResourceAttributes(v IoK8sApiAuthorizationV1ResourceAttributes)`

SetResourceAttributes sets ResourceAttributes field to given value.

### HasResourceAttributes

`func (o *IoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec) HasResourceAttributes() bool`

HasResourceAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


