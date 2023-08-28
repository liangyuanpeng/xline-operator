# IoK8sApiCoreV1SecretReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name is unique within a namespace to reference a secret resource. | [optional] 
**Namespace** | Pointer to **string** | namespace defines the space within which the secret name must be unique. | [optional] 

## Methods

### NewIoK8sApiCoreV1SecretReference

`func NewIoK8sApiCoreV1SecretReference() *IoK8sApiCoreV1SecretReference`

NewIoK8sApiCoreV1SecretReference instantiates a new IoK8sApiCoreV1SecretReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1SecretReferenceWithDefaults

`func NewIoK8sApiCoreV1SecretReferenceWithDefaults() *IoK8sApiCoreV1SecretReference`

NewIoK8sApiCoreV1SecretReferenceWithDefaults instantiates a new IoK8sApiCoreV1SecretReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiCoreV1SecretReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1SecretReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1SecretReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiCoreV1SecretReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *IoK8sApiCoreV1SecretReference) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoK8sApiCoreV1SecretReference) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoK8sApiCoreV1SecretReference) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoK8sApiCoreV1SecretReference) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


