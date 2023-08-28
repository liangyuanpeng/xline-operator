# IoK8sApiAdmissionregistrationV1RuleWithOperations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroups** | Pointer to **[]string** | APIGroups is the API groups the resources belong to. &#39;*&#39; is all groups. If &#39;*&#39; is present, the length of the slice must be one. Required. | [optional] 
**ApiVersions** | Pointer to **[]string** | APIVersions is the API versions the resources belong to. &#39;*&#39; is all versions. If &#39;*&#39; is present, the length of the slice must be one. Required. | [optional] 
**Operations** | Pointer to **[]string** | Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If &#39;*&#39; is present, the length of the slice must be one. Required. | [optional] 
**Resources** | Pointer to **[]string** | Resources is a list of resources this rule applies to.  For example: &#39;pods&#39; means pods. &#39;pods/log&#39; means the log subresource of pods. &#39;*&#39; means all resources, but not subresources. &#39;pods/_*&#39; means all subresources of pods. &#39;*_/scale&#39; means all scale subresources. &#39;*_/_*&#39; means all resources and their subresources.  If wildcard is present, the validation rule will ensure resources do not overlap with each other.  Depending on the enclosing object, subresources might not be allowed. Required. | [optional] 
**Scope** | Pointer to **string** | scope specifies the scope of this rule. Valid values are \&quot;Cluster\&quot;, \&quot;Namespaced\&quot;, and \&quot;*\&quot; \&quot;Cluster\&quot; means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. \&quot;Namespaced\&quot; means that only namespaced resources will match this rule. \&quot;*\&quot; means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is \&quot;*\&quot;. | [optional] 

## Methods

### NewIoK8sApiAdmissionregistrationV1RuleWithOperations

`func NewIoK8sApiAdmissionregistrationV1RuleWithOperations() *IoK8sApiAdmissionregistrationV1RuleWithOperations`

NewIoK8sApiAdmissionregistrationV1RuleWithOperations instantiates a new IoK8sApiAdmissionregistrationV1RuleWithOperations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAdmissionregistrationV1RuleWithOperationsWithDefaults

`func NewIoK8sApiAdmissionregistrationV1RuleWithOperationsWithDefaults() *IoK8sApiAdmissionregistrationV1RuleWithOperations`

NewIoK8sApiAdmissionregistrationV1RuleWithOperationsWithDefaults instantiates a new IoK8sApiAdmissionregistrationV1RuleWithOperations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroups

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) GetApiGroups() []string`

GetApiGroups returns the ApiGroups field if non-nil, zero value otherwise.

### GetApiGroupsOk

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) GetApiGroupsOk() (*[]string, bool)`

GetApiGroupsOk returns a tuple with the ApiGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroups

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) SetApiGroups(v []string)`

SetApiGroups sets ApiGroups field to given value.

### HasApiGroups

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) HasApiGroups() bool`

HasApiGroups returns a boolean if a field has been set.

### GetApiVersions

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) GetApiVersions() []string`

GetApiVersions returns the ApiVersions field if non-nil, zero value otherwise.

### GetApiVersionsOk

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) GetApiVersionsOk() (*[]string, bool)`

GetApiVersionsOk returns a tuple with the ApiVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersions

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) SetApiVersions(v []string)`

SetApiVersions sets ApiVersions field to given value.

### HasApiVersions

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) HasApiVersions() bool`

HasApiVersions returns a boolean if a field has been set.

### GetOperations

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) SetOperations(v []string)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetResources

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetScope

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IoK8sApiAdmissionregistrationV1RuleWithOperations) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


