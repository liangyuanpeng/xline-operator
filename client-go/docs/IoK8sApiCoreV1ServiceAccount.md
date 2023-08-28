# IoK8sApiCoreV1ServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**AutomountServiceAccountToken** | Pointer to **bool** | AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level. | [optional] 
**ImagePullSecrets** | Pointer to [**[]IoK8sApiCoreV1LocalObjectReference**](IoK8sApiCoreV1LocalObjectReference.md) | ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Secrets** | Pointer to [**[]IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) | Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a \&quot;kubernetes.io/enforce-mountable-secrets\&quot; annotation set to \&quot;true\&quot;. This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret | [optional] 

## Methods

### NewIoK8sApiCoreV1ServiceAccount

`func NewIoK8sApiCoreV1ServiceAccount() *IoK8sApiCoreV1ServiceAccount`

NewIoK8sApiCoreV1ServiceAccount instantiates a new IoK8sApiCoreV1ServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ServiceAccountWithDefaults

`func NewIoK8sApiCoreV1ServiceAccountWithDefaults() *IoK8sApiCoreV1ServiceAccount`

NewIoK8sApiCoreV1ServiceAccountWithDefaults instantiates a new IoK8sApiCoreV1ServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiCoreV1ServiceAccount) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiCoreV1ServiceAccount) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiCoreV1ServiceAccount) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiCoreV1ServiceAccount) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAutomountServiceAccountToken

`func (o *IoK8sApiCoreV1ServiceAccount) GetAutomountServiceAccountToken() bool`

GetAutomountServiceAccountToken returns the AutomountServiceAccountToken field if non-nil, zero value otherwise.

### GetAutomountServiceAccountTokenOk

`func (o *IoK8sApiCoreV1ServiceAccount) GetAutomountServiceAccountTokenOk() (*bool, bool)`

GetAutomountServiceAccountTokenOk returns a tuple with the AutomountServiceAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomountServiceAccountToken

`func (o *IoK8sApiCoreV1ServiceAccount) SetAutomountServiceAccountToken(v bool)`

SetAutomountServiceAccountToken sets AutomountServiceAccountToken field to given value.

### HasAutomountServiceAccountToken

`func (o *IoK8sApiCoreV1ServiceAccount) HasAutomountServiceAccountToken() bool`

HasAutomountServiceAccountToken returns a boolean if a field has been set.

### GetImagePullSecrets

`func (o *IoK8sApiCoreV1ServiceAccount) GetImagePullSecrets() []IoK8sApiCoreV1LocalObjectReference`

GetImagePullSecrets returns the ImagePullSecrets field if non-nil, zero value otherwise.

### GetImagePullSecretsOk

`func (o *IoK8sApiCoreV1ServiceAccount) GetImagePullSecretsOk() (*[]IoK8sApiCoreV1LocalObjectReference, bool)`

GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullSecrets

`func (o *IoK8sApiCoreV1ServiceAccount) SetImagePullSecrets(v []IoK8sApiCoreV1LocalObjectReference)`

SetImagePullSecrets sets ImagePullSecrets field to given value.

### HasImagePullSecrets

`func (o *IoK8sApiCoreV1ServiceAccount) HasImagePullSecrets() bool`

HasImagePullSecrets returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiCoreV1ServiceAccount) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiCoreV1ServiceAccount) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiCoreV1ServiceAccount) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiCoreV1ServiceAccount) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiCoreV1ServiceAccount) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiCoreV1ServiceAccount) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiCoreV1ServiceAccount) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiCoreV1ServiceAccount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSecrets

`func (o *IoK8sApiCoreV1ServiceAccount) GetSecrets() []IoK8sApiCoreV1ObjectReference`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *IoK8sApiCoreV1ServiceAccount) GetSecretsOk() (*[]IoK8sApiCoreV1ObjectReference, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *IoK8sApiCoreV1ServiceAccount) SetSecrets(v []IoK8sApiCoreV1ObjectReference)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *IoK8sApiCoreV1ServiceAccount) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


