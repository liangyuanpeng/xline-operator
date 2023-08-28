# IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalPrinterColumns** | Pointer to [**[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition.md) | additionalPrinterColumns specifies additional columns returned in Table output. See https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables for details. If no columns are specified, a single column displaying the age of the custom resource is used. | [optional] 
**Deprecated** | Pointer to **bool** | deprecated indicates this version of the custom resource API is deprecated. When set to true, API requests to this version receive a warning header in the server response. Defaults to false. | [optional] 
**DeprecationWarning** | Pointer to **string** | deprecationWarning overrides the default warning returned to API clients. May only be set when &#x60;deprecated&#x60; is true. The default warning indicates this version is deprecated and recommends use of the newest served version of equal or greater stability, if one exists. | [optional] 
**Name** | **string** | name is the version name, e.g. “v1”, “v2beta1”, etc. The custom resources are served under this version at &#x60;/apis/&lt;group&gt;/&lt;version&gt;/...&#x60; if &#x60;served&#x60; is true. | 
**Schema** | Pointer to [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation.md) |  | [optional] 
**Served** | **bool** | served is a flag enabling/disabling this version from being served via REST APIs | 
**Storage** | **bool** | storage indicates this version should be used when persisting custom resources to storage. There must be exactly one version with storage&#x3D;true. | 
**Subresources** | Pointer to [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources.md) |  | [optional] 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion(name string, served bool, storage bool, ) *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersionWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersionWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersionWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalPrinterColumns

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetAdditionalPrinterColumns() []IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition`

GetAdditionalPrinterColumns returns the AdditionalPrinterColumns field if non-nil, zero value otherwise.

### GetAdditionalPrinterColumnsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetAdditionalPrinterColumnsOk() (*[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition, bool)`

GetAdditionalPrinterColumnsOk returns a tuple with the AdditionalPrinterColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPrinterColumns

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) SetAdditionalPrinterColumns(v []IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition)`

SetAdditionalPrinterColumns sets AdditionalPrinterColumns field to given value.

### HasAdditionalPrinterColumns

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) HasAdditionalPrinterColumns() bool`

HasAdditionalPrinterColumns returns a boolean if a field has been set.

### GetDeprecated

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDeprecationWarning

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetDeprecationWarning() string`

GetDeprecationWarning returns the DeprecationWarning field if non-nil, zero value otherwise.

### GetDeprecationWarningOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetDeprecationWarningOk() (*string, bool)`

GetDeprecationWarningOk returns a tuple with the DeprecationWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationWarning

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) SetDeprecationWarning(v string)`

SetDeprecationWarning sets DeprecationWarning field to given value.

### HasDeprecationWarning

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) HasDeprecationWarning() bool`

HasDeprecationWarning returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) SetName(v string)`

SetName sets Name field to given value.


### GetSchema

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetSchema() IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetSchemaOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) SetSchema(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetServed

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetServed() bool`

GetServed returns the Served field if non-nil, zero value otherwise.

### GetServedOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetServedOk() (*bool, bool)`

GetServedOk returns a tuple with the Served field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServed

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) SetServed(v bool)`

SetServed sets Served field to given value.


### GetStorage

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetStorage() bool`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetStorageOk() (*bool, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) SetStorage(v bool)`

SetStorage sets Storage field to given value.


### GetSubresources

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetSubresources() IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources`

GetSubresources returns the Subresources field if non-nil, zero value otherwise.

### GetSubresourcesOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) GetSubresourcesOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources, bool)`

GetSubresourcesOk returns a tuple with the Subresources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresources

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) SetSubresources(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources)`

SetSubresources sets Subresources field to given value.

### HasSubresources

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion) HasSubresources() bool`

HasSubresources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


