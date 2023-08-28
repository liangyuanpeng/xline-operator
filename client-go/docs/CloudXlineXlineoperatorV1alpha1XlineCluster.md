# CloudXlineXlineoperatorV1alpha1XlineCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Spec** | [**CloudXlineXlineoperatorV1alpha1XlineClusterSpec**](CloudXlineXlineoperatorV1alpha1XlineClusterSpec.md) |  | 
**Status** | Pointer to **map[string]interface{}** | Xline cluster status | [optional] 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineCluster

`func NewCloudXlineXlineoperatorV1alpha1XlineCluster(spec CloudXlineXlineoperatorV1alpha1XlineClusterSpec, ) *CloudXlineXlineoperatorV1alpha1XlineCluster`

NewCloudXlineXlineoperatorV1alpha1XlineCluster instantiates a new CloudXlineXlineoperatorV1alpha1XlineCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineCluster`

NewCloudXlineXlineoperatorV1alpha1XlineClusterWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) GetSpec() CloudXlineXlineoperatorV1alpha1XlineClusterSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) GetSpecOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) SetSpec(v CloudXlineXlineoperatorV1alpha1XlineClusterSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudXlineXlineoperatorV1alpha1XlineCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


