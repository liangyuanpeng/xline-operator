# IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec**](IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec.md) |  | [optional] 
**Status** | Pointer to [**IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus**](IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus.md) |  | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration

`func NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration() *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration`

NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration instantiates a new IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationWithDefaults

`func NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationWithDefaults() *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration`

NewIoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) GetSpec() IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) GetSpecOk() (*IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) SetSpec(v IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) GetStatus() IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) GetStatusOk() (*IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) SetStatus(v IoK8sApiFlowcontrolV1beta2PriorityLevelConfigurationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoK8sApiFlowcontrolV1beta2PriorityLevelConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


