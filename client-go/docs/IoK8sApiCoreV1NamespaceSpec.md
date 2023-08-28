# IoK8sApiCoreV1NamespaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finalizers** | Pointer to **[]string** | Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/ | [optional] 

## Methods

### NewIoK8sApiCoreV1NamespaceSpec

`func NewIoK8sApiCoreV1NamespaceSpec() *IoK8sApiCoreV1NamespaceSpec`

NewIoK8sApiCoreV1NamespaceSpec instantiates a new IoK8sApiCoreV1NamespaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1NamespaceSpecWithDefaults

`func NewIoK8sApiCoreV1NamespaceSpecWithDefaults() *IoK8sApiCoreV1NamespaceSpec`

NewIoK8sApiCoreV1NamespaceSpecWithDefaults instantiates a new IoK8sApiCoreV1NamespaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalizers

`func (o *IoK8sApiCoreV1NamespaceSpec) GetFinalizers() []string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *IoK8sApiCoreV1NamespaceSpec) GetFinalizersOk() (*[]string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *IoK8sApiCoreV1NamespaceSpec) SetFinalizers(v []string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *IoK8sApiCoreV1NamespaceSpec) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


