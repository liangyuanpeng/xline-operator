# IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description is a human readable description of this column. | [optional] 
**Format** | Pointer to **string** | format is an optional OpenAPI type definition for this column. The &#39;name&#39; format is applied to the primary identifier column to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details. | [optional] 
**JsonPath** | **string** | jsonPath is a simple JSON path (i.e. with array notation) which is evaluated against each custom resource to produce the value for this column. | 
**Name** | **string** | name is a human readable name for the column. | 
**Priority** | Pointer to **int32** | priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a priority greater than 0. | [optional] 
**Type** | **string** | type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details. | 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition(jsonPath string, name string, type_ string, ) *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinitionWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinitionWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinitionWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormat

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetJsonPath

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetJsonPath() string`

GetJsonPath returns the JsonPath field if non-nil, zero value otherwise.

### GetJsonPathOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetJsonPathOk() (*string, bool)`

GetJsonPathOk returns a tuple with the JsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonPath

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) SetJsonPath(v string)`

SetJsonPath sets JsonPath field to given value.


### GetName

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceColumnDefinition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


