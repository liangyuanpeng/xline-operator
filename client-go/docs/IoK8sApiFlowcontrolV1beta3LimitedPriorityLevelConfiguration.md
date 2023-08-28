# IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BorrowingLimitPercent** | Pointer to **int32** | &#x60;borrowingLimitPercent&#x60;, if present, configures a limit on how many seats this priority level can borrow from other priority levels. The limit is known as this level&#39;s BorrowingConcurrencyLimit (BorrowingCL) and is a limit on the total number of seats that this level may borrow at any one time. This field holds the ratio of that limit to the level&#39;s nominal concurrency limit. When this field is non-nil, it must hold a non-negative integer and the limit is calculated as follows.  BorrowingCL(i) &#x3D; round( NominalCL(i) * borrowingLimitPercent(i)/100.0 )  The value of this field can be more than 100, implying that this priority level can borrow a number of seats that is greater than its own nominal concurrency limit (NominalCL). When this field is left &#x60;nil&#x60;, the limit is effectively infinite. | [optional] 
**LendablePercent** | Pointer to **int32** | &#x60;lendablePercent&#x60; prescribes the fraction of the level&#39;s NominalCL that can be borrowed by other priority levels. The value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level&#39;s LendableConcurrencyLimit (LendableCL), is defined as follows.  LendableCL(i) &#x3D; round( NominalCL(i) * lendablePercent(i)/100.0 ) | [optional] 
**LimitResponse** | Pointer to [**IoK8sApiFlowcontrolV1beta3LimitResponse**](IoK8sApiFlowcontrolV1beta3LimitResponse.md) |  | [optional] 
**NominalConcurrencyShares** | Pointer to **int32** | &#x60;nominalConcurrencyShares&#x60; (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats available at this priority level. This is used both for requests dispatched from this priority level as well as requests dispatched from other priority levels borrowing seats from this level. The server&#39;s concurrency limit (ServerCL) is divided among the Limited priority levels in proportion to their NCS values:  NominalCL(i)  &#x3D; ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs &#x3D; sum[limited priority level k] NCS(k)  Bigger numbers mean a larger nominal concurrency limit, at the expense of every other Limited priority level. This field has a default value of 30. | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration

`func NewIoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration() *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration`

NewIoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration instantiates a new IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfigurationWithDefaults

`func NewIoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfigurationWithDefaults() *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration`

NewIoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfigurationWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBorrowingLimitPercent

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) GetBorrowingLimitPercent() int32`

GetBorrowingLimitPercent returns the BorrowingLimitPercent field if non-nil, zero value otherwise.

### GetBorrowingLimitPercentOk

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) GetBorrowingLimitPercentOk() (*int32, bool)`

GetBorrowingLimitPercentOk returns a tuple with the BorrowingLimitPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowingLimitPercent

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) SetBorrowingLimitPercent(v int32)`

SetBorrowingLimitPercent sets BorrowingLimitPercent field to given value.

### HasBorrowingLimitPercent

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) HasBorrowingLimitPercent() bool`

HasBorrowingLimitPercent returns a boolean if a field has been set.

### GetLendablePercent

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) GetLendablePercent() int32`

GetLendablePercent returns the LendablePercent field if non-nil, zero value otherwise.

### GetLendablePercentOk

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) GetLendablePercentOk() (*int32, bool)`

GetLendablePercentOk returns a tuple with the LendablePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLendablePercent

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) SetLendablePercent(v int32)`

SetLendablePercent sets LendablePercent field to given value.

### HasLendablePercent

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) HasLendablePercent() bool`

HasLendablePercent returns a boolean if a field has been set.

### GetLimitResponse

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) GetLimitResponse() IoK8sApiFlowcontrolV1beta3LimitResponse`

GetLimitResponse returns the LimitResponse field if non-nil, zero value otherwise.

### GetLimitResponseOk

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) GetLimitResponseOk() (*IoK8sApiFlowcontrolV1beta3LimitResponse, bool)`

GetLimitResponseOk returns a tuple with the LimitResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitResponse

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) SetLimitResponse(v IoK8sApiFlowcontrolV1beta3LimitResponse)`

SetLimitResponse sets LimitResponse field to given value.

### HasLimitResponse

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) HasLimitResponse() bool`

HasLimitResponse returns a boolean if a field has been set.

### GetNominalConcurrencyShares

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) GetNominalConcurrencyShares() int32`

GetNominalConcurrencyShares returns the NominalConcurrencyShares field if non-nil, zero value otherwise.

### GetNominalConcurrencySharesOk

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) GetNominalConcurrencySharesOk() (*int32, bool)`

GetNominalConcurrencySharesOk returns a tuple with the NominalConcurrencyShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominalConcurrencyShares

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) SetNominalConcurrencyShares(v int32)`

SetNominalConcurrencyShares sets NominalConcurrencyShares field to given value.

### HasNominalConcurrencyShares

`func (o *IoK8sApiFlowcontrolV1beta3LimitedPriorityLevelConfiguration) HasNominalConcurrencyShares() bool`

HasNominalConcurrencyShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


