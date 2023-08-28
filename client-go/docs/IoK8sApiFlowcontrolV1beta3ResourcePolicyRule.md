# IoK8sApiFlowcontrolV1beta3ResourcePolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroups** | **[]string** | &#x60;apiGroups&#x60; is a list of matching API groups and may not be empty. \&quot;*\&quot; matches all API groups and, if present, must be the only entry. Required. | 
**ClusterScope** | Pointer to **bool** | &#x60;clusterScope&#x60; indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the &#x60;namespaces&#x60; field must contain a non-empty list. | [optional] 
**Namespaces** | Pointer to **[]string** | &#x60;namespaces&#x60; is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains \&quot;*\&quot;.  Note that \&quot;*\&quot; matches any specified namespace but does not match a request that _does not specify_ a namespace (see the &#x60;clusterScope&#x60; field for that). This list may be empty, but only if &#x60;clusterScope&#x60; is true. | [optional] 
**Resources** | **[]string** | &#x60;resources&#x60; is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ \&quot;services\&quot;, \&quot;nodes/status\&quot; ].  This list may not be empty. \&quot;*\&quot; matches all resources and, if present, must be the only entry. Required. | 
**Verbs** | **[]string** | &#x60;verbs&#x60; is a list of matching verbs and may not be empty. \&quot;*\&quot; matches all verbs and, if present, must be the only entry. Required. | 

## Methods

### NewIoK8sApiFlowcontrolV1beta3ResourcePolicyRule

`func NewIoK8sApiFlowcontrolV1beta3ResourcePolicyRule(apiGroups []string, resources []string, verbs []string, ) *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule`

NewIoK8sApiFlowcontrolV1beta3ResourcePolicyRule instantiates a new IoK8sApiFlowcontrolV1beta3ResourcePolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta3ResourcePolicyRuleWithDefaults

`func NewIoK8sApiFlowcontrolV1beta3ResourcePolicyRuleWithDefaults() *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule`

NewIoK8sApiFlowcontrolV1beta3ResourcePolicyRuleWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta3ResourcePolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroups

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) GetApiGroups() []string`

GetApiGroups returns the ApiGroups field if non-nil, zero value otherwise.

### GetApiGroupsOk

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) GetApiGroupsOk() (*[]string, bool)`

GetApiGroupsOk returns a tuple with the ApiGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroups

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) SetApiGroups(v []string)`

SetApiGroups sets ApiGroups field to given value.


### GetClusterScope

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) GetClusterScope() bool`

GetClusterScope returns the ClusterScope field if non-nil, zero value otherwise.

### GetClusterScopeOk

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) GetClusterScopeOk() (*bool, bool)`

GetClusterScopeOk returns a tuple with the ClusterScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterScope

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) SetClusterScope(v bool)`

SetClusterScope sets ClusterScope field to given value.

### HasClusterScope

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) HasClusterScope() bool`

HasClusterScope returns a boolean if a field has been set.

### GetNamespaces

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetResources

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) SetResources(v []string)`

SetResources sets Resources field to given value.


### GetVerbs

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) GetVerbs() []string`

GetVerbs returns the Verbs field if non-nil, zero value otherwise.

### GetVerbsOk

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) GetVerbsOk() (*[]string, bool)`

GetVerbsOk returns a tuple with the Verbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbs

`func (o *IoK8sApiFlowcontrolV1beta3ResourcePolicyRule) SetVerbs(v []string)`

SetVerbs sets Verbs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


