# IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**AdditionalItems** | Pointer to **map[string]interface{}** | JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value. Defaults to true for the boolean property. | [optional] 
**AdditionalPropertiesField** | Pointer to **map[string]interface{}** | JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value. Defaults to true for the boolean property. | [optional] 
**AllOf** | Pointer to [**[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps.md) |  | [optional] 
**AnyOf** | Pointer to [**[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps.md) |  | [optional] 
**Default** | Pointer to **map[string]interface{}** | JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil. | [optional] 
**Definitions** | Pointer to [**map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps.md) |  | [optional] 
**Dependencies** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enum** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Example** | Pointer to **map[string]interface{}** | JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil. | [optional] 
**ExclusiveMaximum** | Pointer to **bool** |  | [optional] 
**ExclusiveMinimum** | Pointer to **bool** |  | [optional] 
**ExternalDocs** | Pointer to [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1ExternalDocumentation**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1ExternalDocumentation.md) |  | [optional] 
**Format** | Pointer to **string** | format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:  - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like \&quot;0321751043\&quot; or \&quot;978-0321751041\&quot; - isbn10: an ISBN10 number string like \&quot;0321751043\&quot; - isbn13: an ISBN13 number string like \&quot;978-0321751041\&quot; - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^\\d{3}[- ]?\\d{2}[- ]?\\d{4}$ - hexcolor: an hexadecimal color code like \&quot;#FFFFFF: following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like \&quot;rgb(255,255,2559\&quot; - byte: base64 encoded binary data - password: any kind of string - date: a date string like \&quot;2006-01-02\&quot; as defined by full-date in RFC3339 - duration: a duration string like \&quot;22 ns\&quot; as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like \&quot;2014-12-15T19:30:20.000Z\&quot; as defined by date-time in RFC3339. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Items** | Pointer to **map[string]interface{}** | JSONSchemaPropsOrArray represents a value that can either be a JSONSchemaProps or an array of JSONSchemaProps. Mainly here for serialization purposes. | [optional] 
**MaxItems** | Pointer to **int64** |  | [optional] 
**MaxLength** | Pointer to **int64** |  | [optional] 
**MaxProperties** | Pointer to **int64** |  | [optional] 
**Maximum** | Pointer to **float64** |  | [optional] 
**MinItems** | Pointer to **int64** |  | [optional] 
**MinLength** | Pointer to **int64** |  | [optional] 
**MinProperties** | Pointer to **int64** |  | [optional] 
**Minimum** | Pointer to **float64** |  | [optional] 
**MultipleOf** | Pointer to **float64** |  | [optional] 
**Not** | Pointer to [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps.md) |  | [optional] 
**Nullable** | Pointer to **bool** |  | [optional] 
**OneOf** | Pointer to [**[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps.md) |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**PatternProperties** | Pointer to [**map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps.md) |  | [optional] 
**Properties** | Pointer to [**map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps.md) |  | [optional] 
**Required** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UniqueItems** | Pointer to **bool** |  | [optional] 
**XKubernetesEmbeddedResource** | Pointer to **bool** | x-kubernetes-embedded-resource defines that the value is an embedded Kubernetes runtime.Object, with TypeMeta and ObjectMeta. The type must be object. It is allowed to further restrict the embedded object. kind, apiVersion and metadata are validated automatically. x-kubernetes-preserve-unknown-fields is allowed to be true, but does not have to be if the object is fully specified (up to kind, apiVersion, metadata). | [optional] 
**XKubernetesIntOrString** | Pointer to **bool** | x-kubernetes-int-or-string specifies that this value is either an integer or a string. If this is true, an empty type is allowed and type as child of anyOf is permitted if following one of the following patterns:  1) anyOf:    - type: integer    - type: string 2) allOf:    - anyOf:      - type: integer      - type: string    - ... zero or more | [optional] 
**XKubernetesListMapKeys** | Pointer to **[]string** | x-kubernetes-list-map-keys annotates an array with the x-kubernetes-list-type &#x60;map&#x60; by specifying the keys used as the index of the map.  This tag MUST only be used on lists that have the \&quot;x-kubernetes-list-type\&quot; extension set to \&quot;map\&quot;. Also, the values specified for this attribute must be a scalar typed field of the child structure (no nesting is supported).  The properties specified must either be required or have a default value, to ensure those properties are present for all list items. | [optional] 
**XKubernetesListType** | Pointer to **string** | x-kubernetes-list-type annotates an array to further describe its topology. This extension must only be used on lists and may have 3 possible values:  1) &#x60;atomic&#x60;: the list is treated as a single entity, like a scalar.      Atomic lists will be entirely replaced when updated. This extension      may be used on any type of list (struct, scalar, ...). 2) &#x60;set&#x60;:      Sets are lists that must not have multiple items with the same value. Each      value must be a scalar, an object with x-kubernetes-map-type &#x60;atomic&#x60; or an      array with x-kubernetes-list-type &#x60;atomic&#x60;. 3) &#x60;map&#x60;:      These lists are like maps in that their elements have a non-index key      used to identify them. Order is preserved upon merge. The map tag      must only be used on a list with elements of type object. Defaults to atomic for arrays. | [optional] 
**XKubernetesMapType** | Pointer to **string** | x-kubernetes-map-type annotates an object to further describe its topology. This extension must only be used when type is object and may have 2 possible values:  1) &#x60;granular&#x60;:      These maps are actual maps (key-value pairs) and each fields are independent      from each other (they can each be manipulated by separate actors). This is      the default behaviour for all maps. 2) &#x60;atomic&#x60;: the list is treated as a single entity, like a scalar.      Atomic maps will be entirely replaced when updated. | [optional] 
**XKubernetesPreserveUnknownFields** | Pointer to **bool** | x-kubernetes-preserve-unknown-fields stops the API server decoding step from pruning fields which are not specified in the validation schema. This affects fields recursively, but switches back to normal pruning behaviour if nested properties or additionalProperties are specified in the schema. This can either be true or undefined. False is forbidden. | [optional] 
**XKubernetesValidations** | Pointer to [**[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1ValidationRule**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1ValidationRule.md) | x-kubernetes-validations describes a list of validation rules written in the CEL expression language. This field is an alpha-level. Using this field requires the feature gate &#x60;CustomResourceValidationExpressions&#x60; to be enabled. | [optional] 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaPropsWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaPropsWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaPropsWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSchema

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAdditionalItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetAdditionalItems() map[string]interface{}`

GetAdditionalItems returns the AdditionalItems field if non-nil, zero value otherwise.

### GetAdditionalItemsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetAdditionalItemsOk() (*map[string]interface{}, bool)`

GetAdditionalItemsOk returns a tuple with the AdditionalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetAdditionalItems(v map[string]interface{})`

SetAdditionalItems sets AdditionalItems field to given value.

### HasAdditionalItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasAdditionalItems() bool`

HasAdditionalItems returns a boolean if a field has been set.

### GetAdditionalPropertiesField

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetAdditionalPropertiesField() map[string]interface{}`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetAdditionalPropertiesFieldOk() (*map[string]interface{}, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetAdditionalPropertiesField(v map[string]interface{})`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### GetAllOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetAllOf() []IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps`

GetAllOf returns the AllOf field if non-nil, zero value otherwise.

### GetAllOfOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetAllOfOk() (*[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps, bool)`

GetAllOfOk returns a tuple with the AllOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetAllOf(v []IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps)`

SetAllOf sets AllOf field to given value.

### HasAllOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasAllOf() bool`

HasAllOf returns a boolean if a field has been set.

### GetAnyOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetAnyOf() []IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps`

GetAnyOf returns the AnyOf field if non-nil, zero value otherwise.

### GetAnyOfOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetAnyOfOk() (*[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps, bool)`

GetAnyOfOk returns a tuple with the AnyOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetAnyOf(v []IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps)`

SetAnyOf sets AnyOf field to given value.

### HasAnyOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasAnyOf() bool`

HasAnyOf returns a boolean if a field has been set.

### GetDefault

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetDefault(v map[string]interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDefinitions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetDefinitions() map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetDefinitionsOk() (*map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetDefinitions(v map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetDependencies

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetDependencies() map[string]map[string]interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetDependenciesOk() (*map[string]map[string]interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetDependencies(v map[string]map[string]interface{})`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetDescription

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetEnum() []map[string]interface{}`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetEnumOk() (*[]map[string]interface{}, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetEnum(v []map[string]interface{})`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### GetExample

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetExample() map[string]interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetExampleOk() (*map[string]interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetExample(v map[string]interface{})`

SetExample sets Example field to given value.

### HasExample

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetExclusiveMaximum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetExclusiveMaximum() bool`

GetExclusiveMaximum returns the ExclusiveMaximum field if non-nil, zero value otherwise.

### GetExclusiveMaximumOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetExclusiveMaximumOk() (*bool, bool)`

GetExclusiveMaximumOk returns a tuple with the ExclusiveMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMaximum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetExclusiveMaximum(v bool)`

SetExclusiveMaximum sets ExclusiveMaximum field to given value.

### HasExclusiveMaximum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasExclusiveMaximum() bool`

HasExclusiveMaximum returns a boolean if a field has been set.

### GetExclusiveMinimum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetExclusiveMinimum() bool`

GetExclusiveMinimum returns the ExclusiveMinimum field if non-nil, zero value otherwise.

### GetExclusiveMinimumOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetExclusiveMinimumOk() (*bool, bool)`

GetExclusiveMinimumOk returns a tuple with the ExclusiveMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMinimum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetExclusiveMinimum(v bool)`

SetExclusiveMinimum sets ExclusiveMinimum field to given value.

### HasExclusiveMinimum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasExclusiveMinimum() bool`

HasExclusiveMinimum returns a boolean if a field has been set.

### GetExternalDocs

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetExternalDocs() IoK8sApiextensionsApiserverPkgApisApiextensionsV1ExternalDocumentation`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetExternalDocsOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1ExternalDocumentation, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetExternalDocs(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1ExternalDocumentation)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetFormat

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetId

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMaxItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMaxItems() int64`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMaxItemsOk() (*int64, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetMaxItems(v int64)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### GetMaxLength

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMaxLength() int64`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMaxLengthOk() (*int64, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetMaxLength(v int64)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMaxProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMaxProperties() int64`

GetMaxProperties returns the MaxProperties field if non-nil, zero value otherwise.

### GetMaxPropertiesOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMaxPropertiesOk() (*int64, bool)`

GetMaxPropertiesOk returns a tuple with the MaxProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetMaxProperties(v int64)`

SetMaxProperties sets MaxProperties field to given value.

### HasMaxProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasMaxProperties() bool`

HasMaxProperties returns a boolean if a field has been set.

### GetMaximum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMaximum() float64`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMaximumOk() (*float64, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetMaximum(v float64)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMinItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMinItems() int64`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMinItemsOk() (*int64, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetMinItems(v int64)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### GetMinLength

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMinLength() int64`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMinLengthOk() (*int64, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetMinLength(v int64)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMinProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMinProperties() int64`

GetMinProperties returns the MinProperties field if non-nil, zero value otherwise.

### GetMinPropertiesOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMinPropertiesOk() (*int64, bool)`

GetMinPropertiesOk returns a tuple with the MinProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetMinProperties(v int64)`

SetMinProperties sets MinProperties field to given value.

### HasMinProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasMinProperties() bool`

HasMinProperties returns a boolean if a field has been set.

### GetMinimum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMinimum() float64`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMinimumOk() (*float64, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetMinimum(v float64)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetMultipleOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMultipleOf() float64`

GetMultipleOf returns the MultipleOf field if non-nil, zero value otherwise.

### GetMultipleOfOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetMultipleOfOk() (*float64, bool)`

GetMultipleOfOk returns a tuple with the MultipleOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetMultipleOf(v float64)`

SetMultipleOf sets MultipleOf field to given value.

### HasMultipleOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasMultipleOf() bool`

HasMultipleOf returns a boolean if a field has been set.

### GetNot

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetNot() IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetNotOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetNot(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps)`

SetNot sets Not field to given value.

### HasNot

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasNot() bool`

HasNot returns a boolean if a field has been set.

### GetNullable

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetNullable(v bool)`

SetNullable sets Nullable field to given value.

### HasNullable

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasNullable() bool`

HasNullable returns a boolean if a field has been set.

### GetOneOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetOneOf() []IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps`

GetOneOf returns the OneOf field if non-nil, zero value otherwise.

### GetOneOfOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetOneOfOk() (*[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps, bool)`

GetOneOfOk returns a tuple with the OneOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetOneOf(v []IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps)`

SetOneOf sets OneOf field to given value.

### HasOneOf

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasOneOf() bool`

HasOneOf returns a boolean if a field has been set.

### GetPattern

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPatternProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetPatternProperties() map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps`

GetPatternProperties returns the PatternProperties field if non-nil, zero value otherwise.

### GetPatternPropertiesOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetPatternPropertiesOk() (*map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps, bool)`

GetPatternPropertiesOk returns a tuple with the PatternProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetPatternProperties(v map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps)`

SetPatternProperties sets PatternProperties field to given value.

### HasPatternProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasPatternProperties() bool`

HasPatternProperties returns a boolean if a field has been set.

### GetProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetProperties() map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetPropertiesOk() (*map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetProperties(v map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRequired

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetTitle

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUniqueItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetUniqueItems() bool`

GetUniqueItems returns the UniqueItems field if non-nil, zero value otherwise.

### GetUniqueItemsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetUniqueItemsOk() (*bool, bool)`

GetUniqueItemsOk returns a tuple with the UniqueItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetUniqueItems(v bool)`

SetUniqueItems sets UniqueItems field to given value.

### HasUniqueItems

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasUniqueItems() bool`

HasUniqueItems returns a boolean if a field has been set.

### GetXKubernetesEmbeddedResource

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesEmbeddedResource() bool`

GetXKubernetesEmbeddedResource returns the XKubernetesEmbeddedResource field if non-nil, zero value otherwise.

### GetXKubernetesEmbeddedResourceOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesEmbeddedResourceOk() (*bool, bool)`

GetXKubernetesEmbeddedResourceOk returns a tuple with the XKubernetesEmbeddedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXKubernetesEmbeddedResource

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetXKubernetesEmbeddedResource(v bool)`

SetXKubernetesEmbeddedResource sets XKubernetesEmbeddedResource field to given value.

### HasXKubernetesEmbeddedResource

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasXKubernetesEmbeddedResource() bool`

HasXKubernetesEmbeddedResource returns a boolean if a field has been set.

### GetXKubernetesIntOrString

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesIntOrString() bool`

GetXKubernetesIntOrString returns the XKubernetesIntOrString field if non-nil, zero value otherwise.

### GetXKubernetesIntOrStringOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesIntOrStringOk() (*bool, bool)`

GetXKubernetesIntOrStringOk returns a tuple with the XKubernetesIntOrString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXKubernetesIntOrString

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetXKubernetesIntOrString(v bool)`

SetXKubernetesIntOrString sets XKubernetesIntOrString field to given value.

### HasXKubernetesIntOrString

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasXKubernetesIntOrString() bool`

HasXKubernetesIntOrString returns a boolean if a field has been set.

### GetXKubernetesListMapKeys

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesListMapKeys() []string`

GetXKubernetesListMapKeys returns the XKubernetesListMapKeys field if non-nil, zero value otherwise.

### GetXKubernetesListMapKeysOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesListMapKeysOk() (*[]string, bool)`

GetXKubernetesListMapKeysOk returns a tuple with the XKubernetesListMapKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXKubernetesListMapKeys

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetXKubernetesListMapKeys(v []string)`

SetXKubernetesListMapKeys sets XKubernetesListMapKeys field to given value.

### HasXKubernetesListMapKeys

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasXKubernetesListMapKeys() bool`

HasXKubernetesListMapKeys returns a boolean if a field has been set.

### GetXKubernetesListType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesListType() string`

GetXKubernetesListType returns the XKubernetesListType field if non-nil, zero value otherwise.

### GetXKubernetesListTypeOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesListTypeOk() (*string, bool)`

GetXKubernetesListTypeOk returns a tuple with the XKubernetesListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXKubernetesListType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetXKubernetesListType(v string)`

SetXKubernetesListType sets XKubernetesListType field to given value.

### HasXKubernetesListType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasXKubernetesListType() bool`

HasXKubernetesListType returns a boolean if a field has been set.

### GetXKubernetesMapType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesMapType() string`

GetXKubernetesMapType returns the XKubernetesMapType field if non-nil, zero value otherwise.

### GetXKubernetesMapTypeOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesMapTypeOk() (*string, bool)`

GetXKubernetesMapTypeOk returns a tuple with the XKubernetesMapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXKubernetesMapType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetXKubernetesMapType(v string)`

SetXKubernetesMapType sets XKubernetesMapType field to given value.

### HasXKubernetesMapType

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasXKubernetesMapType() bool`

HasXKubernetesMapType returns a boolean if a field has been set.

### GetXKubernetesPreserveUnknownFields

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesPreserveUnknownFields() bool`

GetXKubernetesPreserveUnknownFields returns the XKubernetesPreserveUnknownFields field if non-nil, zero value otherwise.

### GetXKubernetesPreserveUnknownFieldsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesPreserveUnknownFieldsOk() (*bool, bool)`

GetXKubernetesPreserveUnknownFieldsOk returns a tuple with the XKubernetesPreserveUnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXKubernetesPreserveUnknownFields

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetXKubernetesPreserveUnknownFields(v bool)`

SetXKubernetesPreserveUnknownFields sets XKubernetesPreserveUnknownFields field to given value.

### HasXKubernetesPreserveUnknownFields

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasXKubernetesPreserveUnknownFields() bool`

HasXKubernetesPreserveUnknownFields returns a boolean if a field has been set.

### GetXKubernetesValidations

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesValidations() []IoK8sApiextensionsApiserverPkgApisApiextensionsV1ValidationRule`

GetXKubernetesValidations returns the XKubernetesValidations field if non-nil, zero value otherwise.

### GetXKubernetesValidationsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) GetXKubernetesValidationsOk() (*[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1ValidationRule, bool)`

GetXKubernetesValidationsOk returns a tuple with the XKubernetesValidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXKubernetesValidations

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) SetXKubernetesValidations(v []IoK8sApiextensionsApiserverPkgApisApiextensionsV1ValidationRule)`

SetXKubernetesValidations sets XKubernetesValidations field to given value.

### HasXKubernetesValidations

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) HasXKubernetesValidations() bool`

HasXKubernetesValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


