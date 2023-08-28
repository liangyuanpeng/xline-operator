# IoK8sApiCoreV1VolumeDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicePath** | **string** | devicePath is the path inside of the container that the device will be mapped to. | 
**Name** | **string** | name must match the name of a persistentVolumeClaim in the pod | 

## Methods

### NewIoK8sApiCoreV1VolumeDevice

`func NewIoK8sApiCoreV1VolumeDevice(devicePath string, name string, ) *IoK8sApiCoreV1VolumeDevice`

NewIoK8sApiCoreV1VolumeDevice instantiates a new IoK8sApiCoreV1VolumeDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1VolumeDeviceWithDefaults

`func NewIoK8sApiCoreV1VolumeDeviceWithDefaults() *IoK8sApiCoreV1VolumeDevice`

NewIoK8sApiCoreV1VolumeDeviceWithDefaults instantiates a new IoK8sApiCoreV1VolumeDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicePath

`func (o *IoK8sApiCoreV1VolumeDevice) GetDevicePath() string`

GetDevicePath returns the DevicePath field if non-nil, zero value otherwise.

### GetDevicePathOk

`func (o *IoK8sApiCoreV1VolumeDevice) GetDevicePathOk() (*string, bool)`

GetDevicePathOk returns a tuple with the DevicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePath

`func (o *IoK8sApiCoreV1VolumeDevice) SetDevicePath(v string)`

SetDevicePath sets DevicePath field to given value.


### GetName

`func (o *IoK8sApiCoreV1VolumeDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1VolumeDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1VolumeDevice) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


