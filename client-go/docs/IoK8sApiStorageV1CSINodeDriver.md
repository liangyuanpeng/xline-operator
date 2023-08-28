# IoK8sApiStorageV1CSINodeDriver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocatable** | Pointer to [**IoK8sApiStorageV1VolumeNodeResources**](IoK8sApiStorageV1VolumeNodeResources.md) |  | [optional] 
**Name** | **string** | name represents the name of the CSI driver that this object refers to. This MUST be the same name returned by the CSI GetPluginName() call for that driver. | 
**NodeID** | **string** | nodeID of the node from the driver point of view. This field enables Kubernetes to communicate with storage systems that do not share the same nomenclature for nodes. For example, Kubernetes may refer to a given node as \&quot;node1\&quot;, but the storage system may refer to the same node as \&quot;nodeA\&quot;. When Kubernetes issues a command to the storage system to attach a volume to a specific node, it can use this field to refer to the node name using the ID that the storage system will understand, e.g. \&quot;nodeA\&quot; instead of \&quot;node1\&quot;. This field is required. | 
**TopologyKeys** | Pointer to **[]string** | topologyKeys is the list of keys supported by the driver. When a driver is initialized on a cluster, it provides a set of topology keys that it understands (e.g. \&quot;company.com/zone\&quot;, \&quot;company.com/region\&quot;). When a driver is initialized on a node, it provides the same topology keys along with values. Kubelet will expose these topology keys as labels on its own node object. When Kubernetes does topology aware provisioning, it can use this list to determine which labels it should retrieve from the node object and pass back to the driver. It is possible for different nodes to use different topology keys. This can be empty if driver does not support topology. | [optional] 

## Methods

### NewIoK8sApiStorageV1CSINodeDriver

`func NewIoK8sApiStorageV1CSINodeDriver(name string, nodeID string, ) *IoK8sApiStorageV1CSINodeDriver`

NewIoK8sApiStorageV1CSINodeDriver instantiates a new IoK8sApiStorageV1CSINodeDriver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiStorageV1CSINodeDriverWithDefaults

`func NewIoK8sApiStorageV1CSINodeDriverWithDefaults() *IoK8sApiStorageV1CSINodeDriver`

NewIoK8sApiStorageV1CSINodeDriverWithDefaults instantiates a new IoK8sApiStorageV1CSINodeDriver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatable

`func (o *IoK8sApiStorageV1CSINodeDriver) GetAllocatable() IoK8sApiStorageV1VolumeNodeResources`

GetAllocatable returns the Allocatable field if non-nil, zero value otherwise.

### GetAllocatableOk

`func (o *IoK8sApiStorageV1CSINodeDriver) GetAllocatableOk() (*IoK8sApiStorageV1VolumeNodeResources, bool)`

GetAllocatableOk returns a tuple with the Allocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatable

`func (o *IoK8sApiStorageV1CSINodeDriver) SetAllocatable(v IoK8sApiStorageV1VolumeNodeResources)`

SetAllocatable sets Allocatable field to given value.

### HasAllocatable

`func (o *IoK8sApiStorageV1CSINodeDriver) HasAllocatable() bool`

HasAllocatable returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiStorageV1CSINodeDriver) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiStorageV1CSINodeDriver) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiStorageV1CSINodeDriver) SetName(v string)`

SetName sets Name field to given value.


### GetNodeID

`func (o *IoK8sApiStorageV1CSINodeDriver) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *IoK8sApiStorageV1CSINodeDriver) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *IoK8sApiStorageV1CSINodeDriver) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.


### GetTopologyKeys

`func (o *IoK8sApiStorageV1CSINodeDriver) GetTopologyKeys() []string`

GetTopologyKeys returns the TopologyKeys field if non-nil, zero value otherwise.

### GetTopologyKeysOk

`func (o *IoK8sApiStorageV1CSINodeDriver) GetTopologyKeysOk() (*[]string, bool)`

GetTopologyKeysOk returns a tuple with the TopologyKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyKeys

`func (o *IoK8sApiStorageV1CSINodeDriver) SetTopologyKeys(v []string)`

SetTopologyKeys sets TopologyKeys field to given value.

### HasTopologyKeys

`func (o *IoK8sApiStorageV1CSINodeDriver) HasTopologyKeys() bool`

HasTopologyKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


