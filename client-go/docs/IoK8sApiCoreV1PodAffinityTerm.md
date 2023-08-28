# IoK8sApiCoreV1PodAffinityTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelSelector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**NamespaceSelector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**Namespaces** | Pointer to **[]string** | namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means \&quot;this pod&#39;s namespace\&quot;. | [optional] 
**TopologyKey** | **string** | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed. | 

## Methods

### NewIoK8sApiCoreV1PodAffinityTerm

`func NewIoK8sApiCoreV1PodAffinityTerm(topologyKey string, ) *IoK8sApiCoreV1PodAffinityTerm`

NewIoK8sApiCoreV1PodAffinityTerm instantiates a new IoK8sApiCoreV1PodAffinityTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PodAffinityTermWithDefaults

`func NewIoK8sApiCoreV1PodAffinityTermWithDefaults() *IoK8sApiCoreV1PodAffinityTerm`

NewIoK8sApiCoreV1PodAffinityTermWithDefaults instantiates a new IoK8sApiCoreV1PodAffinityTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelSelector

`func (o *IoK8sApiCoreV1PodAffinityTerm) GetLabelSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetLabelSelector returns the LabelSelector field if non-nil, zero value otherwise.

### GetLabelSelectorOk

`func (o *IoK8sApiCoreV1PodAffinityTerm) GetLabelSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetLabelSelectorOk returns a tuple with the LabelSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSelector

`func (o *IoK8sApiCoreV1PodAffinityTerm) SetLabelSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetLabelSelector sets LabelSelector field to given value.

### HasLabelSelector

`func (o *IoK8sApiCoreV1PodAffinityTerm) HasLabelSelector() bool`

HasLabelSelector returns a boolean if a field has been set.

### GetNamespaceSelector

`func (o *IoK8sApiCoreV1PodAffinityTerm) GetNamespaceSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetNamespaceSelector returns the NamespaceSelector field if non-nil, zero value otherwise.

### GetNamespaceSelectorOk

`func (o *IoK8sApiCoreV1PodAffinityTerm) GetNamespaceSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetNamespaceSelectorOk returns a tuple with the NamespaceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceSelector

`func (o *IoK8sApiCoreV1PodAffinityTerm) SetNamespaceSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetNamespaceSelector sets NamespaceSelector field to given value.

### HasNamespaceSelector

`func (o *IoK8sApiCoreV1PodAffinityTerm) HasNamespaceSelector() bool`

HasNamespaceSelector returns a boolean if a field has been set.

### GetNamespaces

`func (o *IoK8sApiCoreV1PodAffinityTerm) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *IoK8sApiCoreV1PodAffinityTerm) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *IoK8sApiCoreV1PodAffinityTerm) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *IoK8sApiCoreV1PodAffinityTerm) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetTopologyKey

`func (o *IoK8sApiCoreV1PodAffinityTerm) GetTopologyKey() string`

GetTopologyKey returns the TopologyKey field if non-nil, zero value otherwise.

### GetTopologyKeyOk

`func (o *IoK8sApiCoreV1PodAffinityTerm) GetTopologyKeyOk() (*string, bool)`

GetTopologyKeyOk returns a tuple with the TopologyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyKey

`func (o *IoK8sApiCoreV1PodAffinityTerm) SetTopologyKey(v string)`

SetTopologyKey sets TopologyKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


