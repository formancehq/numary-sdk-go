# AccountCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Next** | **string** |  | 
**PageSize** | **int32** |  | 
**Previous** | **string** |  | 
**RemainingResults** | **int32** |  | 
**Total** | **int32** |  | 
**Data** | [**[]Account**](Account.md) |  | 

## Methods

### NewAccountCursor

`func NewAccountCursor(hasMore bool, next string, pageSize int32, previous string, remainingResults int32, total int32, data []Account, ) *AccountCursor`

NewAccountCursor instantiates a new AccountCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCursorWithDefaults

`func NewAccountCursorWithDefaults() *AccountCursor`

NewAccountCursorWithDefaults instantiates a new AccountCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *AccountCursor) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *AccountCursor) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *AccountCursor) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNext

`func (o *AccountCursor) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AccountCursor) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AccountCursor) SetNext(v string)`

SetNext sets Next field to given value.


### GetPageSize

`func (o *AccountCursor) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *AccountCursor) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *AccountCursor) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPrevious

`func (o *AccountCursor) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *AccountCursor) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *AccountCursor) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### GetRemainingResults

`func (o *AccountCursor) GetRemainingResults() int32`

GetRemainingResults returns the RemainingResults field if non-nil, zero value otherwise.

### GetRemainingResultsOk

`func (o *AccountCursor) GetRemainingResultsOk() (*int32, bool)`

GetRemainingResultsOk returns a tuple with the RemainingResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingResults

`func (o *AccountCursor) SetRemainingResults(v int32)`

SetRemainingResults sets RemainingResults field to given value.


### GetTotal

`func (o *AccountCursor) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountCursor) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountCursor) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetData

`func (o *AccountCursor) GetData() []Account`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountCursor) GetDataOk() (*[]Account, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountCursor) SetData(v []Account)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


