# IoK8sApiAuthenticationV1TokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Spec** | [**IoK8sApiAuthenticationV1TokenRequestSpec**](IoK8sApiAuthenticationV1TokenRequestSpec.md) |  | 
**Status** | Pointer to [**IoK8sApiAuthenticationV1TokenRequestStatus**](IoK8sApiAuthenticationV1TokenRequestStatus.md) |  | [optional] 

## Methods

### NewIoK8sApiAuthenticationV1TokenRequest

`func NewIoK8sApiAuthenticationV1TokenRequest(spec IoK8sApiAuthenticationV1TokenRequestSpec, ) *IoK8sApiAuthenticationV1TokenRequest`

NewIoK8sApiAuthenticationV1TokenRequest instantiates a new IoK8sApiAuthenticationV1TokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthenticationV1TokenRequestWithDefaults

`func NewIoK8sApiAuthenticationV1TokenRequestWithDefaults() *IoK8sApiAuthenticationV1TokenRequest`

NewIoK8sApiAuthenticationV1TokenRequestWithDefaults instantiates a new IoK8sApiAuthenticationV1TokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiAuthenticationV1TokenRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiAuthenticationV1TokenRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiAuthenticationV1TokenRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiAuthenticationV1TokenRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiAuthenticationV1TokenRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiAuthenticationV1TokenRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiAuthenticationV1TokenRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiAuthenticationV1TokenRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiAuthenticationV1TokenRequest) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiAuthenticationV1TokenRequest) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiAuthenticationV1TokenRequest) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiAuthenticationV1TokenRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *IoK8sApiAuthenticationV1TokenRequest) GetSpec() IoK8sApiAuthenticationV1TokenRequestSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoK8sApiAuthenticationV1TokenRequest) GetSpecOk() (*IoK8sApiAuthenticationV1TokenRequestSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoK8sApiAuthenticationV1TokenRequest) SetSpec(v IoK8sApiAuthenticationV1TokenRequestSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *IoK8sApiAuthenticationV1TokenRequest) GetStatus() IoK8sApiAuthenticationV1TokenRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiAuthenticationV1TokenRequest) GetStatusOk() (*IoK8sApiAuthenticationV1TokenRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiAuthenticationV1TokenRequest) SetStatus(v IoK8sApiAuthenticationV1TokenRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoK8sApiAuthenticationV1TokenRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


