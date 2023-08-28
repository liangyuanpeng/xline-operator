# IoK8sApiFlowcontrolV1beta2QueuingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandSize** | Pointer to **int32** | &#x60;handSize&#x60; is a small positive number that configures the shuffle sharding of requests into queues.  When enqueuing a request at this priority level the request&#39;s flow identifier (a string pair) is hashed and the hash value is used to shuffle the list of queues and deal a hand of the size specified here.  The request is put into one of the shortest queues in that hand. &#x60;handSize&#x60; must be no larger than &#x60;queues&#x60;, and should be significantly smaller (so that a few heavy flows do not saturate most of the queues).  See the user-facing documentation for more extensive guidance on setting this field.  This field has a default value of 8. | [optional] 
**QueueLengthLimit** | Pointer to **int32** | &#x60;queueLengthLimit&#x60; is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected.  This value must be positive.  If not specified, it will be defaulted to 50. | [optional] 
**Queues** | Pointer to **int32** | &#x60;queues&#x60; is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive.  Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant.  This field has a default value of 64. | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta2QueuingConfiguration

`func NewIoK8sApiFlowcontrolV1beta2QueuingConfiguration() *IoK8sApiFlowcontrolV1beta2QueuingConfiguration`

NewIoK8sApiFlowcontrolV1beta2QueuingConfiguration instantiates a new IoK8sApiFlowcontrolV1beta2QueuingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta2QueuingConfigurationWithDefaults

`func NewIoK8sApiFlowcontrolV1beta2QueuingConfigurationWithDefaults() *IoK8sApiFlowcontrolV1beta2QueuingConfiguration`

NewIoK8sApiFlowcontrolV1beta2QueuingConfigurationWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta2QueuingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandSize

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) GetHandSize() int32`

GetHandSize returns the HandSize field if non-nil, zero value otherwise.

### GetHandSizeOk

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) GetHandSizeOk() (*int32, bool)`

GetHandSizeOk returns a tuple with the HandSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandSize

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) SetHandSize(v int32)`

SetHandSize sets HandSize field to given value.

### HasHandSize

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) HasHandSize() bool`

HasHandSize returns a boolean if a field has been set.

### GetQueueLengthLimit

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) GetQueueLengthLimit() int32`

GetQueueLengthLimit returns the QueueLengthLimit field if non-nil, zero value otherwise.

### GetQueueLengthLimitOk

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) GetQueueLengthLimitOk() (*int32, bool)`

GetQueueLengthLimitOk returns a tuple with the QueueLengthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueLengthLimit

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) SetQueueLengthLimit(v int32)`

SetQueueLengthLimit sets QueueLengthLimit field to given value.

### HasQueueLengthLimit

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) HasQueueLengthLimit() bool`

HasQueueLengthLimit returns a boolean if a field has been set.

### GetQueues

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) GetQueues() int32`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) GetQueuesOk() (*int32, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) SetQueues(v int32)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *IoK8sApiFlowcontrolV1beta2QueuingConfiguration) HasQueues() bool`

HasQueues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


