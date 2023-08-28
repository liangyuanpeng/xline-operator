# IoK8sApiAutoscalingV1Scale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**IoK8sApiAutoscalingV1ScaleSpec**](IoK8sApiAutoscalingV1ScaleSpec.md) |  | [optional] 
**Status** | Pointer to [**IoK8sApiAutoscalingV1ScaleStatus**](IoK8sApiAutoscalingV1ScaleStatus.md) |  | [optional] 

## Methods

### NewIoK8sApiAutoscalingV1Scale

`func NewIoK8sApiAutoscalingV1Scale() *IoK8sApiAutoscalingV1Scale`

NewIoK8sApiAutoscalingV1Scale instantiates a new IoK8sApiAutoscalingV1Scale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV1ScaleWithDefaults

`func NewIoK8sApiAutoscalingV1ScaleWithDefaults() *IoK8sApiAutoscalingV1Scale`

NewIoK8sApiAutoscalingV1ScaleWithDefaults instantiates a new IoK8sApiAutoscalingV1Scale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiAutoscalingV1Scale) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiAutoscalingV1Scale) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiAutoscalingV1Scale) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiAutoscalingV1Scale) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiAutoscalingV1Scale) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiAutoscalingV1Scale) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiAutoscalingV1Scale) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiAutoscalingV1Scale) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiAutoscalingV1Scale) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiAutoscalingV1Scale) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiAutoscalingV1Scale) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiAutoscalingV1Scale) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *IoK8sApiAutoscalingV1Scale) GetSpec() IoK8sApiAutoscalingV1ScaleSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoK8sApiAutoscalingV1Scale) GetSpecOk() (*IoK8sApiAutoscalingV1ScaleSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoK8sApiAutoscalingV1Scale) SetSpec(v IoK8sApiAutoscalingV1ScaleSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IoK8sApiAutoscalingV1Scale) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiAutoscalingV1Scale) GetStatus() IoK8sApiAutoscalingV1ScaleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiAutoscalingV1Scale) GetStatusOk() (*IoK8sApiAutoscalingV1ScaleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiAutoscalingV1Scale) SetStatus(v IoK8sApiAutoscalingV1ScaleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoK8sApiAutoscalingV1Scale) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


