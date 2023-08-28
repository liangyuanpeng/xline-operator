# IoK8sApiAuthorizationV1LocalSubjectAccessReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Spec** | [**IoK8sApiAuthorizationV1SubjectAccessReviewSpec**](IoK8sApiAuthorizationV1SubjectAccessReviewSpec.md) |  | 
**Status** | Pointer to [**IoK8sApiAuthorizationV1SubjectAccessReviewStatus**](IoK8sApiAuthorizationV1SubjectAccessReviewStatus.md) |  | [optional] 

## Methods

### NewIoK8sApiAuthorizationV1LocalSubjectAccessReview

`func NewIoK8sApiAuthorizationV1LocalSubjectAccessReview(spec IoK8sApiAuthorizationV1SubjectAccessReviewSpec, ) *IoK8sApiAuthorizationV1LocalSubjectAccessReview`

NewIoK8sApiAuthorizationV1LocalSubjectAccessReview instantiates a new IoK8sApiAuthorizationV1LocalSubjectAccessReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthorizationV1LocalSubjectAccessReviewWithDefaults

`func NewIoK8sApiAuthorizationV1LocalSubjectAccessReviewWithDefaults() *IoK8sApiAuthorizationV1LocalSubjectAccessReview`

NewIoK8sApiAuthorizationV1LocalSubjectAccessReviewWithDefaults instantiates a new IoK8sApiAuthorizationV1LocalSubjectAccessReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) GetSpec() IoK8sApiAuthorizationV1SubjectAccessReviewSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) GetSpecOk() (*IoK8sApiAuthorizationV1SubjectAccessReviewSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) SetSpec(v IoK8sApiAuthorizationV1SubjectAccessReviewSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) GetStatus() IoK8sApiAuthorizationV1SubjectAccessReviewStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) GetStatusOk() (*IoK8sApiAuthorizationV1SubjectAccessReviewStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) SetStatus(v IoK8sApiAuthorizationV1SubjectAccessReviewStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoK8sApiAuthorizationV1LocalSubjectAccessReview) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


