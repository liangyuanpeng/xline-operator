# IoK8sApiRbacV1AggregationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterRoleSelectors** | Pointer to [**[]IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) | ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole&#39;s permissions will be added | [optional] 

## Methods

### NewIoK8sApiRbacV1AggregationRule

`func NewIoK8sApiRbacV1AggregationRule() *IoK8sApiRbacV1AggregationRule`

NewIoK8sApiRbacV1AggregationRule instantiates a new IoK8sApiRbacV1AggregationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiRbacV1AggregationRuleWithDefaults

`func NewIoK8sApiRbacV1AggregationRuleWithDefaults() *IoK8sApiRbacV1AggregationRule`

NewIoK8sApiRbacV1AggregationRuleWithDefaults instantiates a new IoK8sApiRbacV1AggregationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterRoleSelectors

`func (o *IoK8sApiRbacV1AggregationRule) GetClusterRoleSelectors() []IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetClusterRoleSelectors returns the ClusterRoleSelectors field if non-nil, zero value otherwise.

### GetClusterRoleSelectorsOk

`func (o *IoK8sApiRbacV1AggregationRule) GetClusterRoleSelectorsOk() (*[]IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetClusterRoleSelectorsOk returns a tuple with the ClusterRoleSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRoleSelectors

`func (o *IoK8sApiRbacV1AggregationRule) SetClusterRoleSelectors(v []IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetClusterRoleSelectors sets ClusterRoleSelectors field to given value.

### HasClusterRoleSelectors

`func (o *IoK8sApiRbacV1AggregationRule) HasClusterRoleSelectors() bool`

HasClusterRoleSelectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


