# Cursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Next** | **string** |  | 
**PageSize** | **int32** |  | 
**Previous** | **string** |  | 
**RemainingResults** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewCursor

`func NewCursor(hasMore bool, next string, pageSize int32, previous string, remainingResults int32, total int32, ) *Cursor`

NewCursor instantiates a new Cursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorWithDefaults

`func NewCursorWithDefaults() *Cursor`

NewCursorWithDefaults instantiates a new Cursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *Cursor) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *Cursor) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *Cursor) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNext

`func (o *Cursor) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Cursor) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Cursor) SetNext(v string)`

SetNext sets Next field to given value.


### GetPageSize

`func (o *Cursor) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Cursor) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Cursor) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPrevious

`func (o *Cursor) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *Cursor) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *Cursor) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### GetRemainingResults

`func (o *Cursor) GetRemainingResults() int32`

GetRemainingResults returns the RemainingResults field if non-nil, zero value otherwise.

### GetRemainingResultsOk

`func (o *Cursor) GetRemainingResultsOk() (*int32, bool)`

GetRemainingResultsOk returns a tuple with the RemainingResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingResults

`func (o *Cursor) SetRemainingResults(v int32)`

SetRemainingResults sets RemainingResults field to given value.


### GetTotal

`func (o *Cursor) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Cursor) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Cursor) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


