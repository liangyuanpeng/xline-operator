# IoK8sApiAuthorizationV1SelfSubjectRulesReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Spec** | [**IoK8sApiAuthorizationV1SelfSubjectRulesReviewSpec**](IoK8sApiAuthorizationV1SelfSubjectRulesReviewSpec.md) |  | 
**Status** | Pointer to [**IoK8sApiAuthorizationV1SubjectRulesReviewStatus**](IoK8sApiAuthorizationV1SubjectRulesReviewStatus.md) |  | [optional] 

## Methods

### NewIoK8sApiAuthorizationV1SelfSubjectRulesReview

`func NewIoK8sApiAuthorizationV1SelfSubjectRulesReview(spec IoK8sApiAuthorizationV1SelfSubjectRulesReviewSpec, ) *IoK8sApiAuthorizationV1SelfSubjectRulesReview`

NewIoK8sApiAuthorizationV1SelfSubjectRulesReview instantiates a new IoK8sApiAuthorizationV1SelfSubjectRulesReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthorizationV1SelfSubjectRulesReviewWithDefaults

`func NewIoK8sApiAuthorizationV1SelfSubjectRulesReviewWithDefaults() *IoK8sApiAuthorizationV1SelfSubjectRulesReview`

NewIoK8sApiAuthorizationV1SelfSubjectRulesReviewWithDefaults instantiates a new IoK8sApiAuthorizationV1SelfSubjectRulesReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) GetSpec() IoK8sApiAuthorizationV1SelfSubjectRulesReviewSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) GetSpecOk() (*IoK8sApiAuthorizationV1SelfSubjectRulesReviewSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) SetSpec(v IoK8sApiAuthorizationV1SelfSubjectRulesReviewSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) GetStatus() IoK8sApiAuthorizationV1SubjectRulesReviewStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) GetStatusOk() (*IoK8sApiAuthorizationV1SubjectRulesReviewStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) SetStatus(v IoK8sApiAuthorizationV1SubjectRulesReviewStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoK8sApiAuthorizationV1SelfSubjectRulesReview) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


