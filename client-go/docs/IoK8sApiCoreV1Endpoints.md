# IoK8sApiCoreV1Endpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Subsets** | Pointer to [**[]IoK8sApiCoreV1EndpointSubset**](IoK8sApiCoreV1EndpointSubset.md) | The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service. | [optional] 

## Methods

### NewIoK8sApiCoreV1Endpoints

`func NewIoK8sApiCoreV1Endpoints() *IoK8sApiCoreV1Endpoints`

NewIoK8sApiCoreV1Endpoints instantiates a new IoK8sApiCoreV1Endpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1EndpointsWithDefaults

`func NewIoK8sApiCoreV1EndpointsWithDefaults() *IoK8sApiCoreV1Endpoints`

NewIoK8sApiCoreV1EndpointsWithDefaults instantiates a new IoK8sApiCoreV1Endpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiCoreV1Endpoints) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiCoreV1Endpoints) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiCoreV1Endpoints) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiCoreV1Endpoints) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiCoreV1Endpoints) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiCoreV1Endpoints) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiCoreV1Endpoints) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiCoreV1Endpoints) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiCoreV1Endpoints) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiCoreV1Endpoints) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiCoreV1Endpoints) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiCoreV1Endpoints) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSubsets

`func (o *IoK8sApiCoreV1Endpoints) GetSubsets() []IoK8sApiCoreV1EndpointSubset`

GetSubsets returns the Subsets field if non-nil, zero value otherwise.

### GetSubsetsOk

`func (o *IoK8sApiCoreV1Endpoints) GetSubsetsOk() (*[]IoK8sApiCoreV1EndpointSubset, bool)`

GetSubsetsOk returns a tuple with the Subsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsets

`func (o *IoK8sApiCoreV1Endpoints) SetSubsets(v []IoK8sApiCoreV1EndpointSubset)`

SetSubsets sets Subsets field to given value.

### HasSubsets

`func (o *IoK8sApiCoreV1Endpoints) HasSubsets() bool`

HasSubsets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


