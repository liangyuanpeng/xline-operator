# CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the environment variable. Must be a C_IDENTIFIER. | 
**Value** | Pointer to **string** | Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. \&quot;$$(VAR_NAME)\&quot; will produce the string literal \&quot;$(VAR_NAME)\&quot;. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \&quot;\&quot;. | [optional] 
**ValueFrom** | Pointer to [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFrom**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFrom.md) |  | [optional] 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner(name string, ) *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueFrom

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) GetValueFrom() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFrom`

GetValueFrom returns the ValueFrom field if non-nil, zero value otherwise.

### GetValueFromOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) GetValueFromOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFrom, bool)`

GetValueFromOk returns a tuple with the ValueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFrom

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) SetValueFrom(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInnerValueFrom)`

SetValueFrom sets ValueFrom field to given value.

### HasValueFrom

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerEnvInner) HasValueFrom() bool`

HasValueFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


