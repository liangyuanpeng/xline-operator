# IoK8sApiCoreV1ContainerResizePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | Name of the resource to which this resource resize policy applies. Supported values: cpu, memory. | 
**RestartPolicy** | **string** | Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired. | 

## Methods

### NewIoK8sApiCoreV1ContainerResizePolicy

`func NewIoK8sApiCoreV1ContainerResizePolicy(resourceName string, restartPolicy string, ) *IoK8sApiCoreV1ContainerResizePolicy`

NewIoK8sApiCoreV1ContainerResizePolicy instantiates a new IoK8sApiCoreV1ContainerResizePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ContainerResizePolicyWithDefaults

`func NewIoK8sApiCoreV1ContainerResizePolicyWithDefaults() *IoK8sApiCoreV1ContainerResizePolicy`

NewIoK8sApiCoreV1ContainerResizePolicyWithDefaults instantiates a new IoK8sApiCoreV1ContainerResizePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *IoK8sApiCoreV1ContainerResizePolicy) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *IoK8sApiCoreV1ContainerResizePolicy) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *IoK8sApiCoreV1ContainerResizePolicy) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetRestartPolicy

`func (o *IoK8sApiCoreV1ContainerResizePolicy) GetRestartPolicy() string`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *IoK8sApiCoreV1ContainerResizePolicy) GetRestartPolicyOk() (*string, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *IoK8sApiCoreV1ContainerResizePolicy) SetRestartPolicy(v string)`

SetRestartPolicy sets RestartPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


