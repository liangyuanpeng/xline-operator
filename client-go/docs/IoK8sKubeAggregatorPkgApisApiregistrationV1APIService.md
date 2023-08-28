# IoK8sKubeAggregatorPkgApisApiregistrationV1APIService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec**](IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec.md) |  | [optional] 
**Status** | Pointer to [**IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus**](IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus.md) |  | [optional] 

## Methods

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIService

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIService() *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIService instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1APIService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceWithDefaults

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceWithDefaults() *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceWithDefaults instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1APIService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) GetSpec() IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) GetSpecOk() (*IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) SetSpec(v IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) GetStatus() IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) GetStatusOk() (*IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) SetStatus(v IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


