# IoK8sApiCoreV1NodeSystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** | The Architecture reported by the node | 
**BootID** | **string** | Boot ID reported by the node. | 
**ContainerRuntimeVersion** | **string** | ContainerRuntime Version reported by the node through runtime remote API (e.g. containerd://1.4.2). | 
**KernelVersion** | **string** | Kernel Version reported by the node from &#39;uname -r&#39; (e.g. 3.16.0-0.bpo.4-amd64). | 
**KubeProxyVersion** | **string** | KubeProxy Version reported by the node. | 
**KubeletVersion** | **string** | Kubelet Version reported by the node. | 
**MachineID** | **string** | MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html | 
**OperatingSystem** | **string** | The Operating System reported by the node | 
**OsImage** | **string** | OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)). | 
**SystemUUID** | **string** | SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-us/red_hat_subscription_management/1/html/rhsm/uuid | 

## Methods

### NewIoK8sApiCoreV1NodeSystemInfo

`func NewIoK8sApiCoreV1NodeSystemInfo(architecture string, bootID string, containerRuntimeVersion string, kernelVersion string, kubeProxyVersion string, kubeletVersion string, machineID string, operatingSystem string, osImage string, systemUUID string, ) *IoK8sApiCoreV1NodeSystemInfo`

NewIoK8sApiCoreV1NodeSystemInfo instantiates a new IoK8sApiCoreV1NodeSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1NodeSystemInfoWithDefaults

`func NewIoK8sApiCoreV1NodeSystemInfoWithDefaults() *IoK8sApiCoreV1NodeSystemInfo`

NewIoK8sApiCoreV1NodeSystemInfoWithDefaults instantiates a new IoK8sApiCoreV1NodeSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *IoK8sApiCoreV1NodeSystemInfo) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetBootID

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetBootID() string`

GetBootID returns the BootID field if non-nil, zero value otherwise.

### GetBootIDOk

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetBootIDOk() (*string, bool)`

GetBootIDOk returns a tuple with the BootID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootID

`func (o *IoK8sApiCoreV1NodeSystemInfo) SetBootID(v string)`

SetBootID sets BootID field to given value.


### GetContainerRuntimeVersion

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetContainerRuntimeVersion() string`

GetContainerRuntimeVersion returns the ContainerRuntimeVersion field if non-nil, zero value otherwise.

### GetContainerRuntimeVersionOk

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetContainerRuntimeVersionOk() (*string, bool)`

GetContainerRuntimeVersionOk returns a tuple with the ContainerRuntimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRuntimeVersion

`func (o *IoK8sApiCoreV1NodeSystemInfo) SetContainerRuntimeVersion(v string)`

SetContainerRuntimeVersion sets ContainerRuntimeVersion field to given value.


### GetKernelVersion

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *IoK8sApiCoreV1NodeSystemInfo) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.


### GetKubeProxyVersion

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetKubeProxyVersion() string`

GetKubeProxyVersion returns the KubeProxyVersion field if non-nil, zero value otherwise.

### GetKubeProxyVersionOk

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetKubeProxyVersionOk() (*string, bool)`

GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeProxyVersion

`func (o *IoK8sApiCoreV1NodeSystemInfo) SetKubeProxyVersion(v string)`

SetKubeProxyVersion sets KubeProxyVersion field to given value.


### GetKubeletVersion

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *IoK8sApiCoreV1NodeSystemInfo) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.


### GetMachineID

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetMachineID() string`

GetMachineID returns the MachineID field if non-nil, zero value otherwise.

### GetMachineIDOk

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetMachineIDOk() (*string, bool)`

GetMachineIDOk returns a tuple with the MachineID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineID

`func (o *IoK8sApiCoreV1NodeSystemInfo) SetMachineID(v string)`

SetMachineID sets MachineID field to given value.


### GetOperatingSystem

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *IoK8sApiCoreV1NodeSystemInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetOsImage

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetOsImage() string`

GetOsImage returns the OsImage field if non-nil, zero value otherwise.

### GetOsImageOk

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetOsImageOk() (*string, bool)`

GetOsImageOk returns a tuple with the OsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsImage

`func (o *IoK8sApiCoreV1NodeSystemInfo) SetOsImage(v string)`

SetOsImage sets OsImage field to given value.


### GetSystemUUID

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetSystemUUID() string`

GetSystemUUID returns the SystemUUID field if non-nil, zero value otherwise.

### GetSystemUUIDOk

`func (o *IoK8sApiCoreV1NodeSystemInfo) GetSystemUUIDOk() (*string, bool)`

GetSystemUUIDOk returns a tuple with the SystemUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUUID

`func (o *IoK8sApiCoreV1NodeSystemInfo) SetSystemUUID(v string)`

SetSystemUUID sets SystemUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


