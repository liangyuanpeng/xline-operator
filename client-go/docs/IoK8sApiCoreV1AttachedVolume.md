# IoK8sApiCoreV1AttachedVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicePath** | **string** | DevicePath represents the device path where the volume should be available | 
**Name** | **string** | Name of the attached volume | 

## Methods

### NewIoK8sApiCoreV1AttachedVolume

`func NewIoK8sApiCoreV1AttachedVolume(devicePath string, name string, ) *IoK8sApiCoreV1AttachedVolume`

NewIoK8sApiCoreV1AttachedVolume instantiates a new IoK8sApiCoreV1AttachedVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1AttachedVolumeWithDefaults

`func NewIoK8sApiCoreV1AttachedVolumeWithDefaults() *IoK8sApiCoreV1AttachedVolume`

NewIoK8sApiCoreV1AttachedVolumeWithDefaults instantiates a new IoK8sApiCoreV1AttachedVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicePath

`func (o *IoK8sApiCoreV1AttachedVolume) GetDevicePath() string`

GetDevicePath returns the DevicePath field if non-nil, zero value otherwise.

### GetDevicePathOk

`func (o *IoK8sApiCoreV1AttachedVolume) GetDevicePathOk() (*string, bool)`

GetDevicePathOk returns a tuple with the DevicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePath

`func (o *IoK8sApiCoreV1AttachedVolume) SetDevicePath(v string)`

SetDevicePath sets DevicePath field to given value.


### GetName

`func (o *IoK8sApiCoreV1AttachedVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1AttachedVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1AttachedVolume) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


