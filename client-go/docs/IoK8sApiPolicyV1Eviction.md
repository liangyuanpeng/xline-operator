# IoK8sApiPolicyV1Eviction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**DeleteOptions** | Pointer to [**IoK8sApimachineryPkgApisMetaV1DeleteOptions**](IoK8sApimachineryPkgApisMetaV1DeleteOptions.md) |  | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 

## Methods

### NewIoK8sApiPolicyV1Eviction

`func NewIoK8sApiPolicyV1Eviction() *IoK8sApiPolicyV1Eviction`

NewIoK8sApiPolicyV1Eviction instantiates a new IoK8sApiPolicyV1Eviction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiPolicyV1EvictionWithDefaults

`func NewIoK8sApiPolicyV1EvictionWithDefaults() *IoK8sApiPolicyV1Eviction`

NewIoK8sApiPolicyV1EvictionWithDefaults instantiates a new IoK8sApiPolicyV1Eviction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiPolicyV1Eviction) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiPolicyV1Eviction) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiPolicyV1Eviction) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiPolicyV1Eviction) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDeleteOptions

`func (o *IoK8sApiPolicyV1Eviction) GetDeleteOptions() IoK8sApimachineryPkgApisMetaV1DeleteOptions`

GetDeleteOptions returns the DeleteOptions field if non-nil, zero value otherwise.

### GetDeleteOptionsOk

`func (o *IoK8sApiPolicyV1Eviction) GetDeleteOptionsOk() (*IoK8sApimachineryPkgApisMetaV1DeleteOptions, bool)`

GetDeleteOptionsOk returns a tuple with the DeleteOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOptions

`func (o *IoK8sApiPolicyV1Eviction) SetDeleteOptions(v IoK8sApimachineryPkgApisMetaV1DeleteOptions)`

SetDeleteOptions sets DeleteOptions field to given value.

### HasDeleteOptions

`func (o *IoK8sApiPolicyV1Eviction) HasDeleteOptions() bool`

HasDeleteOptions returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiPolicyV1Eviction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiPolicyV1Eviction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiPolicyV1Eviction) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiPolicyV1Eviction) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiPolicyV1Eviction) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiPolicyV1Eviction) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiPolicyV1Eviction) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiPolicyV1Eviction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


