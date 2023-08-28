# IoK8sApiCoreV1EnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the environment variable. Must be a C_IDENTIFIER. | 
**Value** | Pointer to **string** | Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. \&quot;$$(VAR_NAME)\&quot; will produce the string literal \&quot;$(VAR_NAME)\&quot;. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \&quot;\&quot;. | [optional] 
**ValueFrom** | Pointer to [**IoK8sApiCoreV1EnvVarSource**](IoK8sApiCoreV1EnvVarSource.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1EnvVar

`func NewIoK8sApiCoreV1EnvVar(name string, ) *IoK8sApiCoreV1EnvVar`

NewIoK8sApiCoreV1EnvVar instantiates a new IoK8sApiCoreV1EnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1EnvVarWithDefaults

`func NewIoK8sApiCoreV1EnvVarWithDefaults() *IoK8sApiCoreV1EnvVar`

NewIoK8sApiCoreV1EnvVarWithDefaults instantiates a new IoK8sApiCoreV1EnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiCoreV1EnvVar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1EnvVar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1EnvVar) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *IoK8sApiCoreV1EnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IoK8sApiCoreV1EnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IoK8sApiCoreV1EnvVar) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IoK8sApiCoreV1EnvVar) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueFrom

`func (o *IoK8sApiCoreV1EnvVar) GetValueFrom() IoK8sApiCoreV1EnvVarSource`

GetValueFrom returns the ValueFrom field if non-nil, zero value otherwise.

### GetValueFromOk

`func (o *IoK8sApiCoreV1EnvVar) GetValueFromOk() (*IoK8sApiCoreV1EnvVarSource, bool)`

GetValueFromOk returns a tuple with the ValueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFrom

`func (o *IoK8sApiCoreV1EnvVar) SetValueFrom(v IoK8sApiCoreV1EnvVarSource)`

SetValueFrom sets ValueFrom field to given value.

### HasValueFrom

`func (o *IoK8sApiCoreV1EnvVar) HasValueFrom() bool`

HasValueFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


