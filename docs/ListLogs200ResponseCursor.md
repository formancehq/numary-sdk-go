# ListLogs200ResponseCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | **int64** |  | 
**HasMore** | Pointer to **bool** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Data** | [**[]Log**](Log.md) |  | 

## Methods

### NewListLogs200ResponseCursor

`func NewListLogs200ResponseCursor(pageSize int64, data []Log, ) *ListLogs200ResponseCursor`

NewListLogs200ResponseCursor instantiates a new ListLogs200ResponseCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogs200ResponseCursorWithDefaults

`func NewListLogs200ResponseCursorWithDefaults() *ListLogs200ResponseCursor`

NewListLogs200ResponseCursorWithDefaults instantiates a new ListLogs200ResponseCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *ListLogs200ResponseCursor) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListLogs200ResponseCursor) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListLogs200ResponseCursor) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetHasMore

`func (o *ListLogs200ResponseCursor) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListLogs200ResponseCursor) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListLogs200ResponseCursor) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListLogs200ResponseCursor) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetPrevious

`func (o *ListLogs200ResponseCursor) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *ListLogs200ResponseCursor) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *ListLogs200ResponseCursor) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *ListLogs200ResponseCursor) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetNext

`func (o *ListLogs200ResponseCursor) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListLogs200ResponseCursor) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListLogs200ResponseCursor) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListLogs200ResponseCursor) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetData

`func (o *ListLogs200ResponseCursor) GetData() []Log`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListLogs200ResponseCursor) GetDataOk() (*[]Log, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListLogs200ResponseCursor) SetData(v []Log)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


