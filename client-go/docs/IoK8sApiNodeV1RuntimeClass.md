# IoK8sApiNodeV1RuntimeClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Handler** | **string** | handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node &amp; CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called \&quot;runc\&quot; might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable. | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Overhead** | Pointer to [**IoK8sApiNodeV1Overhead**](IoK8sApiNodeV1Overhead.md) |  | [optional] 
**Scheduling** | Pointer to [**IoK8sApiNodeV1Scheduling**](IoK8sApiNodeV1Scheduling.md) |  | [optional] 

## Methods

### NewIoK8sApiNodeV1RuntimeClass

`func NewIoK8sApiNodeV1RuntimeClass(handler string, ) *IoK8sApiNodeV1RuntimeClass`

NewIoK8sApiNodeV1RuntimeClass instantiates a new IoK8sApiNodeV1RuntimeClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNodeV1RuntimeClassWithDefaults

`func NewIoK8sApiNodeV1RuntimeClassWithDefaults() *IoK8sApiNodeV1RuntimeClass`

NewIoK8sApiNodeV1RuntimeClassWithDefaults instantiates a new IoK8sApiNodeV1RuntimeClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiNodeV1RuntimeClass) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiNodeV1RuntimeClass) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiNodeV1RuntimeClass) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiNodeV1RuntimeClass) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetHandler

`func (o *IoK8sApiNodeV1RuntimeClass) GetHandler() string`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *IoK8sApiNodeV1RuntimeClass) GetHandlerOk() (*string, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandler

`func (o *IoK8sApiNodeV1RuntimeClass) SetHandler(v string)`

SetHandler sets Handler field to given value.


### GetKind

`func (o *IoK8sApiNodeV1RuntimeClass) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiNodeV1RuntimeClass) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiNodeV1RuntimeClass) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiNodeV1RuntimeClass) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiNodeV1RuntimeClass) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiNodeV1RuntimeClass) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiNodeV1RuntimeClass) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiNodeV1RuntimeClass) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOverhead

`func (o *IoK8sApiNodeV1RuntimeClass) GetOverhead() IoK8sApiNodeV1Overhead`

GetOverhead returns the Overhead field if non-nil, zero value otherwise.

### GetOverheadOk

`func (o *IoK8sApiNodeV1RuntimeClass) GetOverheadOk() (*IoK8sApiNodeV1Overhead, bool)`

GetOverheadOk returns a tuple with the Overhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverhead

`func (o *IoK8sApiNodeV1RuntimeClass) SetOverhead(v IoK8sApiNodeV1Overhead)`

SetOverhead sets Overhead field to given value.

### HasOverhead

`func (o *IoK8sApiNodeV1RuntimeClass) HasOverhead() bool`

HasOverhead returns a boolean if a field has been set.

### GetScheduling

`func (o *IoK8sApiNodeV1RuntimeClass) GetScheduling() IoK8sApiNodeV1Scheduling`

GetScheduling returns the Scheduling field if non-nil, zero value otherwise.

### GetSchedulingOk

`func (o *IoK8sApiNodeV1RuntimeClass) GetSchedulingOk() (*IoK8sApiNodeV1Scheduling, bool)`

GetSchedulingOk returns a tuple with the Scheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduling

`func (o *IoK8sApiNodeV1RuntimeClass) SetScheduling(v IoK8sApiNodeV1Scheduling)`

SetScheduling sets Scheduling field to given value.

### HasScheduling

`func (o *IoK8sApiNodeV1RuntimeClass) HasScheduling() bool`

HasScheduling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


