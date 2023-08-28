# IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssuredConcurrencyShares** | Pointer to **int32** | &#x60;assuredConcurrencyShares&#x60; (ACS) configures the execution limit, which is a limit on the number of requests of this priority level that may be exeucting at a given time.  ACS must be a positive number. The server&#39;s concurrency limit (SCL) is divided among the concurrency-controlled priority levels in proportion to their assured concurrency shares. This produces the assured concurrency value (ACV) --- the number of requests that may be executing at a time --- for each such priority level:              ACV(l) &#x3D; ceil( SCL * ACS(l) / ( sum[priority levels k] ACS(k) ) )  bigger numbers of ACS mean more reserved concurrent requests (at the expense of every other PL). This field has a default value of 30. | [optional] 
**BorrowingLimitPercent** | Pointer to **int32** | &#x60;borrowingLimitPercent&#x60;, if present, configures a limit on how many seats this priority level can borrow from other priority levels. The limit is known as this level&#39;s BorrowingConcurrencyLimit (BorrowingCL) and is a limit on the total number of seats that this level may borrow at any one time. This field holds the ratio of that limit to the level&#39;s nominal concurrency limit. When this field is non-nil, it must hold a non-negative integer and the limit is calculated as follows.  BorrowingCL(i) &#x3D; round( NominalCL(i) * borrowingLimitPercent(i)/100.0 )  The value of this field can be more than 100, implying that this priority level can borrow a number of seats that is greater than its own nominal concurrency limit (NominalCL). When this field is left &#x60;nil&#x60;, the limit is effectively infinite. | [optional] 
**LendablePercent** | Pointer to **int32** | &#x60;lendablePercent&#x60; prescribes the fraction of the level&#39;s NominalCL that can be borrowed by other priority levels. The value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level&#39;s LendableConcurrencyLimit (LendableCL), is defined as follows.  LendableCL(i) &#x3D; round( NominalCL(i) * lendablePercent(i)/100.0 ) | [optional] 
**LimitResponse** | Pointer to [**IoK8sApiFlowcontrolV1beta2LimitResponse**](IoK8sApiFlowcontrolV1beta2LimitResponse.md) |  | [optional] 

## Methods

### NewIoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration

`func NewIoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration() *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration`

NewIoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration instantiates a new IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfigurationWithDefaults

`func NewIoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfigurationWithDefaults() *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration`

NewIoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfigurationWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssuredConcurrencyShares

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) GetAssuredConcurrencyShares() int32`

GetAssuredConcurrencyShares returns the AssuredConcurrencyShares field if non-nil, zero value otherwise.

### GetAssuredConcurrencySharesOk

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) GetAssuredConcurrencySharesOk() (*int32, bool)`

GetAssuredConcurrencySharesOk returns a tuple with the AssuredConcurrencyShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuredConcurrencyShares

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) SetAssuredConcurrencyShares(v int32)`

SetAssuredConcurrencyShares sets AssuredConcurrencyShares field to given value.

### HasAssuredConcurrencyShares

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) HasAssuredConcurrencyShares() bool`

HasAssuredConcurrencyShares returns a boolean if a field has been set.

### GetBorrowingLimitPercent

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) GetBorrowingLimitPercent() int32`

GetBorrowingLimitPercent returns the BorrowingLimitPercent field if non-nil, zero value otherwise.

### GetBorrowingLimitPercentOk

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) GetBorrowingLimitPercentOk() (*int32, bool)`

GetBorrowingLimitPercentOk returns a tuple with the BorrowingLimitPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowingLimitPercent

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) SetBorrowingLimitPercent(v int32)`

SetBorrowingLimitPercent sets BorrowingLimitPercent field to given value.

### HasBorrowingLimitPercent

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) HasBorrowingLimitPercent() bool`

HasBorrowingLimitPercent returns a boolean if a field has been set.

### GetLendablePercent

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) GetLendablePercent() int32`

GetLendablePercent returns the LendablePercent field if non-nil, zero value otherwise.

### GetLendablePercentOk

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) GetLendablePercentOk() (*int32, bool)`

GetLendablePercentOk returns a tuple with the LendablePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLendablePercent

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) SetLendablePercent(v int32)`

SetLendablePercent sets LendablePercent field to given value.

### HasLendablePercent

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) HasLendablePercent() bool`

HasLendablePercent returns a boolean if a field has been set.

### GetLimitResponse

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) GetLimitResponse() IoK8sApiFlowcontrolV1beta2LimitResponse`

GetLimitResponse returns the LimitResponse field if non-nil, zero value otherwise.

### GetLimitResponseOk

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) GetLimitResponseOk() (*IoK8sApiFlowcontrolV1beta2LimitResponse, bool)`

GetLimitResponseOk returns a tuple with the LimitResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitResponse

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) SetLimitResponse(v IoK8sApiFlowcontrolV1beta2LimitResponse)`

SetLimitResponse sets LimitResponse field to given value.

### HasLimitResponse

`func (o *IoK8sApiFlowcontrolV1beta2LimitedPriorityLevelConfiguration) HasLimitResponse() bool`

HasLimitResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


