# IoK8sApiStorageV1VolumeAttachmentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attacher** | **string** | attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName(). | 
**NodeName** | **string** | nodeName represents the node that the volume should be attached to. | 
**Source** | [**IoK8sApiStorageV1VolumeAttachmentSource**](IoK8sApiStorageV1VolumeAttachmentSource.md) |  | 

## Methods

### NewIoK8sApiStorageV1VolumeAttachmentSpec

`func NewIoK8sApiStorageV1VolumeAttachmentSpec(attacher string, nodeName string, source IoK8sApiStorageV1VolumeAttachmentSource, ) *IoK8sApiStorageV1VolumeAttachmentSpec`

NewIoK8sApiStorageV1VolumeAttachmentSpec instantiates a new IoK8sApiStorageV1VolumeAttachmentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiStorageV1VolumeAttachmentSpecWithDefaults

`func NewIoK8sApiStorageV1VolumeAttachmentSpecWithDefaults() *IoK8sApiStorageV1VolumeAttachmentSpec`

NewIoK8sApiStorageV1VolumeAttachmentSpecWithDefaults instantiates a new IoK8sApiStorageV1VolumeAttachmentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttacher

`func (o *IoK8sApiStorageV1VolumeAttachmentSpec) GetAttacher() string`

GetAttacher returns the Attacher field if non-nil, zero value otherwise.

### GetAttacherOk

`func (o *IoK8sApiStorageV1VolumeAttachmentSpec) GetAttacherOk() (*string, bool)`

GetAttacherOk returns a tuple with the Attacher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttacher

`func (o *IoK8sApiStorageV1VolumeAttachmentSpec) SetAttacher(v string)`

SetAttacher sets Attacher field to given value.


### GetNodeName

`func (o *IoK8sApiStorageV1VolumeAttachmentSpec) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *IoK8sApiStorageV1VolumeAttachmentSpec) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *IoK8sApiStorageV1VolumeAttachmentSpec) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetSource

`func (o *IoK8sApiStorageV1VolumeAttachmentSpec) GetSource() IoK8sApiStorageV1VolumeAttachmentSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IoK8sApiStorageV1VolumeAttachmentSpec) GetSourceOk() (*IoK8sApiStorageV1VolumeAttachmentSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IoK8sApiStorageV1VolumeAttachmentSpec) SetSource(v IoK8sApiStorageV1VolumeAttachmentSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


