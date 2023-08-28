# IoK8sApiCoreV1VolumeProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMap** | Pointer to [**IoK8sApiCoreV1ConfigMapProjection**](IoK8sApiCoreV1ConfigMapProjection.md) |  | [optional] 
**DownwardAPI** | Pointer to [**IoK8sApiCoreV1DownwardAPIProjection**](IoK8sApiCoreV1DownwardAPIProjection.md) |  | [optional] 
**Secret** | Pointer to [**IoK8sApiCoreV1SecretProjection**](IoK8sApiCoreV1SecretProjection.md) |  | [optional] 
**ServiceAccountToken** | Pointer to [**IoK8sApiCoreV1ServiceAccountTokenProjection**](IoK8sApiCoreV1ServiceAccountTokenProjection.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1VolumeProjection

`func NewIoK8sApiCoreV1VolumeProjection() *IoK8sApiCoreV1VolumeProjection`

NewIoK8sApiCoreV1VolumeProjection instantiates a new IoK8sApiCoreV1VolumeProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1VolumeProjectionWithDefaults

`func NewIoK8sApiCoreV1VolumeProjectionWithDefaults() *IoK8sApiCoreV1VolumeProjection`

NewIoK8sApiCoreV1VolumeProjectionWithDefaults instantiates a new IoK8sApiCoreV1VolumeProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMap

`func (o *IoK8sApiCoreV1VolumeProjection) GetConfigMap() IoK8sApiCoreV1ConfigMapProjection`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *IoK8sApiCoreV1VolumeProjection) GetConfigMapOk() (*IoK8sApiCoreV1ConfigMapProjection, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *IoK8sApiCoreV1VolumeProjection) SetConfigMap(v IoK8sApiCoreV1ConfigMapProjection)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *IoK8sApiCoreV1VolumeProjection) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.

### GetDownwardAPI

`func (o *IoK8sApiCoreV1VolumeProjection) GetDownwardAPI() IoK8sApiCoreV1DownwardAPIProjection`

GetDownwardAPI returns the DownwardAPI field if non-nil, zero value otherwise.

### GetDownwardAPIOk

`func (o *IoK8sApiCoreV1VolumeProjection) GetDownwardAPIOk() (*IoK8sApiCoreV1DownwardAPIProjection, bool)`

GetDownwardAPIOk returns a tuple with the DownwardAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownwardAPI

`func (o *IoK8sApiCoreV1VolumeProjection) SetDownwardAPI(v IoK8sApiCoreV1DownwardAPIProjection)`

SetDownwardAPI sets DownwardAPI field to given value.

### HasDownwardAPI

`func (o *IoK8sApiCoreV1VolumeProjection) HasDownwardAPI() bool`

HasDownwardAPI returns a boolean if a field has been set.

### GetSecret

`func (o *IoK8sApiCoreV1VolumeProjection) GetSecret() IoK8sApiCoreV1SecretProjection`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IoK8sApiCoreV1VolumeProjection) GetSecretOk() (*IoK8sApiCoreV1SecretProjection, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IoK8sApiCoreV1VolumeProjection) SetSecret(v IoK8sApiCoreV1SecretProjection)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IoK8sApiCoreV1VolumeProjection) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetServiceAccountToken

`func (o *IoK8sApiCoreV1VolumeProjection) GetServiceAccountToken() IoK8sApiCoreV1ServiceAccountTokenProjection`

GetServiceAccountToken returns the ServiceAccountToken field if non-nil, zero value otherwise.

### GetServiceAccountTokenOk

`func (o *IoK8sApiCoreV1VolumeProjection) GetServiceAccountTokenOk() (*IoK8sApiCoreV1ServiceAccountTokenProjection, bool)`

GetServiceAccountTokenOk returns a tuple with the ServiceAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountToken

`func (o *IoK8sApiCoreV1VolumeProjection) SetServiceAccountToken(v IoK8sApiCoreV1ServiceAccountTokenProjection)`

SetServiceAccountToken sets ServiceAccountToken field to given value.

### HasServiceAccountToken

`func (o *IoK8sApiCoreV1VolumeProjection) HasServiceAccountToken() bool`

HasServiceAccountToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


