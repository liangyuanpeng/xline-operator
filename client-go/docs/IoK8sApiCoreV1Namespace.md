# IoK8sApiCoreV1Namespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**IoK8sApiCoreV1NamespaceSpec**](IoK8sApiCoreV1NamespaceSpec.md) |  | [optional] 
**Status** | Pointer to [**IoK8sApiCoreV1NamespaceStatus**](IoK8sApiCoreV1NamespaceStatus.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1Namespace

`func NewIoK8sApiCoreV1Namespace() *IoK8sApiCoreV1Namespace`

NewIoK8sApiCoreV1Namespace instantiates a new IoK8sApiCoreV1Namespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1NamespaceWithDefaults

`func NewIoK8sApiCoreV1NamespaceWithDefaults() *IoK8sApiCoreV1Namespace`

NewIoK8sApiCoreV1NamespaceWithDefaults instantiates a new IoK8sApiCoreV1Namespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiCoreV1Namespace) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiCoreV1Namespace) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiCoreV1Namespace) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiCoreV1Namespace) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiCoreV1Namespace) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiCoreV1Namespace) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiCoreV1Namespace) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiCoreV1Namespace) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiCoreV1Namespace) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiCoreV1Namespace) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiCoreV1Namespace) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiCoreV1Namespace) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *IoK8sApiCoreV1Namespace) GetSpec() IoK8sApiCoreV1NamespaceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoK8sApiCoreV1Namespace) GetSpecOk() (*IoK8sApiCoreV1NamespaceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoK8sApiCoreV1Namespace) SetSpec(v IoK8sApiCoreV1NamespaceSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IoK8sApiCoreV1Namespace) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiCoreV1Namespace) GetStatus() IoK8sApiCoreV1NamespaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiCoreV1Namespace) GetStatusOk() (*IoK8sApiCoreV1NamespaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiCoreV1Namespace) SetStatus(v IoK8sApiCoreV1NamespaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoK8sApiCoreV1Namespace) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


