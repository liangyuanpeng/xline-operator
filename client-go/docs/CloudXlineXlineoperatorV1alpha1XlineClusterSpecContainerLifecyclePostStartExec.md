# CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **[]string** | Command is the command line to execute inside the container, the working directory for the command  is root (&#39;/&#39;) in the container&#39;s filesystem. The command is simply exec&#39;d, it is not run inside a shell, so traditional shell instructions (&#39;|&#39;, etc) won&#39;t work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy. | [optional] 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExecWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExecWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExecWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerLifecyclePostStartExec) HasCommand() bool`

HasCommand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


