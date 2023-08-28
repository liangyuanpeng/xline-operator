# IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversion** | Pointer to [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion.md) |  | [optional] 
**Group** | **string** | group is the API group of the defined custom resource. The custom resources are served under &#x60;/apis/&lt;group&gt;/...&#x60;. Must match the name of the CustomResourceDefinition (in the form &#x60;&lt;names.plural&gt;.&lt;group&gt;&#x60;). | 
**Names** | [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames.md) |  | 
**PreserveUnknownFields** | Pointer to **bool** | preserveUnknownFields indicates that object fields which are not specified in the OpenAPI schema should be preserved when persisting to storage. apiVersion, kind, metadata and known fields inside metadata are always preserved. This field is deprecated in favor of setting &#x60;x-preserve-unknown-fields&#x60; to true in &#x60;spec.versions[*].schema.openAPIV3Schema&#x60;. See https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#field-pruning for details. | [optional] 
**Scope** | **string** | scope indicates whether the defined custom resource is cluster- or namespace-scoped. Allowed values are &#x60;Cluster&#x60; and &#x60;Namespaced&#x60;. | 
**Versions** | [**[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion.md) | versions is the list of all API versions of the defined custom resource. Version names are used to compute the order in which served versions are listed in API discovery. If the version string is \&quot;kube-like\&quot;, it will sort above non \&quot;kube-like\&quot; version strings, which are ordered lexicographically. \&quot;Kube-like\&quot; versions start with a \&quot;v\&quot;, then are followed by a number (the major version), then optionally the string \&quot;alpha\&quot; or \&quot;beta\&quot; and another number (the minor version). These are sorted first by GA &gt; beta &gt; alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10. | 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec(group string, names IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames, scope string, versions []IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion, ) *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpecWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpecWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpecWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversion

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetConversion() IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion`

GetConversion returns the Conversion field if non-nil, zero value otherwise.

### GetConversionOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetConversionOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion, bool)`

GetConversionOk returns a tuple with the Conversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversion

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) SetConversion(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion)`

SetConversion sets Conversion field to given value.

### HasConversion

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) HasConversion() bool`

HasConversion returns a boolean if a field has been set.

### GetGroup

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetNames

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetNames() IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetNamesOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) SetNames(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames)`

SetNames sets Names field to given value.


### GetPreserveUnknownFields

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetPreserveUnknownFields() bool`

GetPreserveUnknownFields returns the PreserveUnknownFields field if non-nil, zero value otherwise.

### GetPreserveUnknownFieldsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetPreserveUnknownFieldsOk() (*bool, bool)`

GetPreserveUnknownFieldsOk returns a tuple with the PreserveUnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveUnknownFields

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) SetPreserveUnknownFields(v bool)`

SetPreserveUnknownFields sets PreserveUnknownFields field to given value.

### HasPreserveUnknownFields

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) HasPreserveUnknownFields() bool`

HasPreserveUnknownFields returns a boolean if a field has been set.

### GetScope

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) SetScope(v string)`

SetScope sets Scope field to given value.


### GetVersions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetVersions() []IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) GetVersionsOk() (*[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) SetVersions(v []IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


