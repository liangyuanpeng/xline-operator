# IoK8sApiRbacV1ClusterRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationRule** | Pointer to [**IoK8sApiRbacV1AggregationRule**](IoK8sApiRbacV1AggregationRule.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Rules** | Pointer to [**[]IoK8sApiRbacV1PolicyRule**](IoK8sApiRbacV1PolicyRule.md) | Rules holds all the PolicyRules for this ClusterRole | [optional] 

## Methods

### NewIoK8sApiRbacV1ClusterRole

`func NewIoK8sApiRbacV1ClusterRole() *IoK8sApiRbacV1ClusterRole`

NewIoK8sApiRbacV1ClusterRole instantiates a new IoK8sApiRbacV1ClusterRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiRbacV1ClusterRoleWithDefaults

`func NewIoK8sApiRbacV1ClusterRoleWithDefaults() *IoK8sApiRbacV1ClusterRole`

NewIoK8sApiRbacV1ClusterRoleWithDefaults instantiates a new IoK8sApiRbacV1ClusterRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationRule

`func (o *IoK8sApiRbacV1ClusterRole) GetAggregationRule() IoK8sApiRbacV1AggregationRule`

GetAggregationRule returns the AggregationRule field if non-nil, zero value otherwise.

### GetAggregationRuleOk

`func (o *IoK8sApiRbacV1ClusterRole) GetAggregationRuleOk() (*IoK8sApiRbacV1AggregationRule, bool)`

GetAggregationRuleOk returns a tuple with the AggregationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationRule

`func (o *IoK8sApiRbacV1ClusterRole) SetAggregationRule(v IoK8sApiRbacV1AggregationRule)`

SetAggregationRule sets AggregationRule field to given value.

### HasAggregationRule

`func (o *IoK8sApiRbacV1ClusterRole) HasAggregationRule() bool`

HasAggregationRule returns a boolean if a field has been set.

### GetApiVersion

`func (o *IoK8sApiRbacV1ClusterRole) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiRbacV1ClusterRole) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiRbacV1ClusterRole) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiRbacV1ClusterRole) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiRbacV1ClusterRole) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiRbacV1ClusterRole) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiRbacV1ClusterRole) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiRbacV1ClusterRole) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiRbacV1ClusterRole) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiRbacV1ClusterRole) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiRbacV1ClusterRole) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiRbacV1ClusterRole) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRules

`func (o *IoK8sApiRbacV1ClusterRole) GetRules() []IoK8sApiRbacV1PolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *IoK8sApiRbacV1ClusterRole) GetRulesOk() (*[]IoK8sApiRbacV1PolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *IoK8sApiRbacV1ClusterRole) SetRules(v []IoK8sApiRbacV1PolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *IoK8sApiRbacV1ClusterRole) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


