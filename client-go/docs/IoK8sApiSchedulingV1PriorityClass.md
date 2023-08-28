# IoK8sApiSchedulingV1PriorityClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Description** | Pointer to **string** | description is an arbitrary string that usually provides guidelines on when this priority class should be used. | [optional] 
**GlobalDefault** | Pointer to **bool** | globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as &#x60;globalDefault&#x60;. However, if more than one PriorityClasses exists with their &#x60;globalDefault&#x60; field set to true, the smallest value of such global default PriorityClasses will be used as the default priority. | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**PreemptionPolicy** | Pointer to **string** | preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.  Possible enum values:  - &#x60;\&quot;Never\&quot;&#x60; means that pod never preempts other pods with lower priority.  - &#x60;\&quot;PreemptLowerPriority\&quot;&#x60; means that pod can preempt other pods with lower priority. | [optional] 
**Value** | **int32** | value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec. | 

## Methods

### NewIoK8sApiSchedulingV1PriorityClass

`func NewIoK8sApiSchedulingV1PriorityClass(value int32, ) *IoK8sApiSchedulingV1PriorityClass`

NewIoK8sApiSchedulingV1PriorityClass instantiates a new IoK8sApiSchedulingV1PriorityClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiSchedulingV1PriorityClassWithDefaults

`func NewIoK8sApiSchedulingV1PriorityClassWithDefaults() *IoK8sApiSchedulingV1PriorityClass`

NewIoK8sApiSchedulingV1PriorityClassWithDefaults instantiates a new IoK8sApiSchedulingV1PriorityClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiSchedulingV1PriorityClass) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiSchedulingV1PriorityClass) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiSchedulingV1PriorityClass) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiSchedulingV1PriorityClass) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDescription

`func (o *IoK8sApiSchedulingV1PriorityClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IoK8sApiSchedulingV1PriorityClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IoK8sApiSchedulingV1PriorityClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IoK8sApiSchedulingV1PriorityClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGlobalDefault

`func (o *IoK8sApiSchedulingV1PriorityClass) GetGlobalDefault() bool`

GetGlobalDefault returns the GlobalDefault field if non-nil, zero value otherwise.

### GetGlobalDefaultOk

`func (o *IoK8sApiSchedulingV1PriorityClass) GetGlobalDefaultOk() (*bool, bool)`

GetGlobalDefaultOk returns a tuple with the GlobalDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDefault

`func (o *IoK8sApiSchedulingV1PriorityClass) SetGlobalDefault(v bool)`

SetGlobalDefault sets GlobalDefault field to given value.

### HasGlobalDefault

`func (o *IoK8sApiSchedulingV1PriorityClass) HasGlobalDefault() bool`

HasGlobalDefault returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiSchedulingV1PriorityClass) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiSchedulingV1PriorityClass) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiSchedulingV1PriorityClass) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiSchedulingV1PriorityClass) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiSchedulingV1PriorityClass) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiSchedulingV1PriorityClass) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiSchedulingV1PriorityClass) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiSchedulingV1PriorityClass) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPreemptionPolicy

`func (o *IoK8sApiSchedulingV1PriorityClass) GetPreemptionPolicy() string`

GetPreemptionPolicy returns the PreemptionPolicy field if non-nil, zero value otherwise.

### GetPreemptionPolicyOk

`func (o *IoK8sApiSchedulingV1PriorityClass) GetPreemptionPolicyOk() (*string, bool)`

GetPreemptionPolicyOk returns a tuple with the PreemptionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptionPolicy

`func (o *IoK8sApiSchedulingV1PriorityClass) SetPreemptionPolicy(v string)`

SetPreemptionPolicy sets PreemptionPolicy field to given value.

### HasPreemptionPolicy

`func (o *IoK8sApiSchedulingV1PriorityClass) HasPreemptionPolicy() bool`

HasPreemptionPolicy returns a boolean if a field has been set.

### GetValue

`func (o *IoK8sApiSchedulingV1PriorityClass) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IoK8sApiSchedulingV1PriorityClass) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IoK8sApiSchedulingV1PriorityClass) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


