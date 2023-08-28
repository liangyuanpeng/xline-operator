# IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | categories is a list of grouped resources this custom resource belongs to (e.g. &#39;all&#39;). This is published in API discovery documents, and used by clients to support invocations like &#x60;kubectl get all&#x60;. | [optional] 
**Kind** | **string** | kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the &#x60;kind&#x60; attribute in API calls. | 
**ListKind** | Pointer to **string** | listKind is the serialized kind of the list for this resource. Defaults to \&quot;&#x60;kind&#x60;List\&quot;. | [optional] 
**Plural** | **string** | plural is the plural name of the resource to serve. The custom resources are served under &#x60;/apis/&lt;group&gt;/&lt;version&gt;/.../&lt;plural&gt;&#x60;. Must match the name of the CustomResourceDefinition (in the form &#x60;&lt;names.plural&gt;.&lt;group&gt;&#x60;). Must be all lowercase. | 
**ShortNames** | Pointer to **[]string** | shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like &#x60;kubectl get &lt;shortname&gt;&#x60;. It must be all lowercase. | [optional] 
**Singular** | Pointer to **string** | singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased &#x60;kind&#x60;. | [optional] 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames(kind string, plural string, ) *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNamesWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNamesWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNamesWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) SetKind(v string)`

SetKind sets Kind field to given value.


### GetListKind

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetListKind() string`

GetListKind returns the ListKind field if non-nil, zero value otherwise.

### GetListKindOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetListKindOk() (*string, bool)`

GetListKindOk returns a tuple with the ListKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListKind

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) SetListKind(v string)`

SetListKind sets ListKind field to given value.

### HasListKind

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) HasListKind() bool`

HasListKind returns a boolean if a field has been set.

### GetPlural

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetPlural() string`

GetPlural returns the Plural field if non-nil, zero value otherwise.

### GetPluralOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetPluralOk() (*string, bool)`

GetPluralOk returns a tuple with the Plural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlural

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) SetPlural(v string)`

SetPlural sets Plural field to given value.


### GetShortNames

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetShortNames() []string`

GetShortNames returns the ShortNames field if non-nil, zero value otherwise.

### GetShortNamesOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetShortNamesOk() (*[]string, bool)`

GetShortNamesOk returns a tuple with the ShortNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortNames

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) SetShortNames(v []string)`

SetShortNames sets ShortNames field to given value.

### HasShortNames

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) HasShortNames() bool`

HasShortNames returns a boolean if a field has been set.

### GetSingular

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetSingular() string`

GetSingular returns the Singular field if non-nil, zero value otherwise.

### GetSingularOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) GetSingularOk() (*string, bool)`

GetSingularOk returns a tuple with the Singular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingular

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) SetSingular(v string)`

SetSingular sets Singular field to given value.

### HasSingular

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames) HasSingular() bool`

HasSingular returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


