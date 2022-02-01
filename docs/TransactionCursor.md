# TransactionCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Next** | Pointer to **string** |  | [optional] 
**PageSize** | **int32** |  | 
**Previous** | Pointer to **string** |  | [optional] 
**RemainingResults** | **int32** |  | 
**Total** | **int32** |  | 
**Data** | [**[]Transaction**](Transaction.md) |  | 

## Methods

### NewTransactionCursor

`func NewTransactionCursor(hasMore bool, pageSize int32, remainingResults int32, total int32, data []Transaction, ) *TransactionCursor`

NewTransactionCursor instantiates a new TransactionCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCursorWithDefaults

`func NewTransactionCursorWithDefaults() *TransactionCursor`

NewTransactionCursorWithDefaults instantiates a new TransactionCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *TransactionCursor) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *TransactionCursor) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *TransactionCursor) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNext

`func (o *TransactionCursor) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *TransactionCursor) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *TransactionCursor) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *TransactionCursor) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPageSize

`func (o *TransactionCursor) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *TransactionCursor) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *TransactionCursor) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPrevious

`func (o *TransactionCursor) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *TransactionCursor) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *TransactionCursor) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *TransactionCursor) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetRemainingResults

`func (o *TransactionCursor) GetRemainingResults() int32`

GetRemainingResults returns the RemainingResults field if non-nil, zero value otherwise.

### GetRemainingResultsOk

`func (o *TransactionCursor) GetRemainingResultsOk() (*int32, bool)`

GetRemainingResultsOk returns a tuple with the RemainingResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingResults

`func (o *TransactionCursor) SetRemainingResults(v int32)`

SetRemainingResults sets RemainingResults field to given value.


### GetTotal

`func (o *TransactionCursor) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TransactionCursor) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TransactionCursor) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetData

`func (o *TransactionCursor) GetData() []Transaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionCursor) GetDataOk() (*[]Transaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionCursor) SetData(v []Transaction)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


