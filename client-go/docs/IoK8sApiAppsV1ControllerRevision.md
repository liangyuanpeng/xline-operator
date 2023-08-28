# IoK8sApiAppsV1ControllerRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Data** | Pointer to **map[string]interface{}** | RawExtension is used to hold extensions in external versions.  To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types.  // Internal package:   type MyAPIObject struct {   runtime.TypeMeta &#x60;json:\&quot;,inline\&quot;&#x60;   MyPlugin runtime.Object &#x60;json:\&quot;myPlugin\&quot;&#x60;  }   type PluginA struct {   AOption string &#x60;json:\&quot;aOption\&quot;&#x60;  }  // External package:   type MyAPIObject struct {   runtime.TypeMeta &#x60;json:\&quot;,inline\&quot;&#x60;   MyPlugin runtime.RawExtension &#x60;json:\&quot;myPlugin\&quot;&#x60;  }   type PluginA struct {   AOption string &#x60;json:\&quot;aOption\&quot;&#x60;  }  // On the wire, the JSON will look something like this:   {   \&quot;kind\&quot;:\&quot;MyAPIObject\&quot;,   \&quot;apiVersion\&quot;:\&quot;v1\&quot;,   \&quot;myPlugin\&quot;: {    \&quot;kind\&quot;:\&quot;PluginA\&quot;,    \&quot;aOption\&quot;:\&quot;foo\&quot;,   },  }  So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package&#39;s DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.) | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Revision** | **int64** | Revision indicates the revision of the state represented by Data. | 

## Methods

### NewIoK8sApiAppsV1ControllerRevision

`func NewIoK8sApiAppsV1ControllerRevision(revision int64, ) *IoK8sApiAppsV1ControllerRevision`

NewIoK8sApiAppsV1ControllerRevision instantiates a new IoK8sApiAppsV1ControllerRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1ControllerRevisionWithDefaults

`func NewIoK8sApiAppsV1ControllerRevisionWithDefaults() *IoK8sApiAppsV1ControllerRevision`

NewIoK8sApiAppsV1ControllerRevisionWithDefaults instantiates a new IoK8sApiAppsV1ControllerRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiAppsV1ControllerRevision) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiAppsV1ControllerRevision) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiAppsV1ControllerRevision) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiAppsV1ControllerRevision) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetData

`func (o *IoK8sApiAppsV1ControllerRevision) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IoK8sApiAppsV1ControllerRevision) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IoK8sApiAppsV1ControllerRevision) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *IoK8sApiAppsV1ControllerRevision) HasData() bool`

HasData returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiAppsV1ControllerRevision) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiAppsV1ControllerRevision) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiAppsV1ControllerRevision) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiAppsV1ControllerRevision) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiAppsV1ControllerRevision) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiAppsV1ControllerRevision) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiAppsV1ControllerRevision) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiAppsV1ControllerRevision) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRevision

`func (o *IoK8sApiAppsV1ControllerRevision) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *IoK8sApiAppsV1ControllerRevision) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *IoK8sApiAppsV1ControllerRevision) SetRevision(v int64)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


