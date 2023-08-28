# IoK8sApiAutoscalingV1CrossVersionObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | apiVersion is the API version of the referent | [optional] 
**Kind** | **string** | kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | 
**Name** | **string** | name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | 

## Methods

### NewIoK8sApiAutoscalingV1CrossVersionObjectReference

`func NewIoK8sApiAutoscalingV1CrossVersionObjectReference(kind string, name string, ) *IoK8sApiAutoscalingV1CrossVersionObjectReference`

NewIoK8sApiAutoscalingV1CrossVersionObjectReference instantiates a new IoK8sApiAutoscalingV1CrossVersionObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAutoscalingV1CrossVersionObjectReferenceWithDefaults

`func NewIoK8sApiAutoscalingV1CrossVersionObjectReferenceWithDefaults() *IoK8sApiAutoscalingV1CrossVersionObjectReference`

NewIoK8sApiAutoscalingV1CrossVersionObjectReferenceWithDefaults instantiates a new IoK8sApiAutoscalingV1CrossVersionObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiAutoscalingV1CrossVersionObjectReference) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiAutoscalingV1CrossVersionObjectReference) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiAutoscalingV1CrossVersionObjectReference) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiAutoscalingV1CrossVersionObjectReference) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiAutoscalingV1CrossVersionObjectReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiAutoscalingV1CrossVersionObjectReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiAutoscalingV1CrossVersionObjectReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *IoK8sApiAutoscalingV1CrossVersionObjectReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiAutoscalingV1CrossVersionObjectReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiAutoscalingV1CrossVersionObjectReference) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


