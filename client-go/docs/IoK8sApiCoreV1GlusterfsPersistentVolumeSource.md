# IoK8sApiCoreV1GlusterfsPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | **string** | endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | 
**EndpointsNamespace** | Pointer to **string** | endpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty, the EndpointNamespace defaults to the same namespace as the bound PVC. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | [optional] 
**Path** | **string** | path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | 
**ReadOnly** | Pointer to **bool** | readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | [optional] 

## Methods

### NewIoK8sApiCoreV1GlusterfsPersistentVolumeSource

`func NewIoK8sApiCoreV1GlusterfsPersistentVolumeSource(endpoints string, path string, ) *IoK8sApiCoreV1GlusterfsPersistentVolumeSource`

NewIoK8sApiCoreV1GlusterfsPersistentVolumeSource instantiates a new IoK8sApiCoreV1GlusterfsPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1GlusterfsPersistentVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1GlusterfsPersistentVolumeSourceWithDefaults() *IoK8sApiCoreV1GlusterfsPersistentVolumeSource`

NewIoK8sApiCoreV1GlusterfsPersistentVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1GlusterfsPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) GetEndpoints() string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) GetEndpointsOk() (*string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) SetEndpoints(v string)`

SetEndpoints sets Endpoints field to given value.


### GetEndpointsNamespace

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) GetEndpointsNamespace() string`

GetEndpointsNamespace returns the EndpointsNamespace field if non-nil, zero value otherwise.

### GetEndpointsNamespaceOk

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) GetEndpointsNamespaceOk() (*string, bool)`

GetEndpointsNamespaceOk returns a tuple with the EndpointsNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointsNamespace

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) SetEndpointsNamespace(v string)`

SetEndpointsNamespace sets EndpointsNamespace field to given value.

### HasEndpointsNamespace

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) HasEndpointsNamespace() bool`

HasEndpointsNamespace returns a boolean if a field has been set.

### GetPath

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) SetPath(v string)`

SetPath sets Path field to given value.


### GetReadOnly

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1GlusterfsPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


