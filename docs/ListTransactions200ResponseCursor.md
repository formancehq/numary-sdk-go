# ListTransactions200ResponseCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | **int32** |  | 
**HasMore** | **bool** |  | 
**Data** | [**[]Transaction**](Transaction.md) |  | 

## Methods

### NewListTransactions200ResponseCursor

`func NewListTransactions200ResponseCursor(pageSize int32, hasMore bool, data []Transaction, ) *ListTransactions200ResponseCursor`

NewListTransactions200ResponseCursor instantiates a new ListTransactions200ResponseCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactions200ResponseCursorWithDefaults

`func NewListTransactions200ResponseCursorWithDefaults() *ListTransactions200ResponseCursor`

NewListTransactions200ResponseCursorWithDefaults instantiates a new ListTransactions200ResponseCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *ListTransactions200ResponseCursor) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListTransactions200ResponseCursor) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListTransactions200ResponseCursor) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetHasMore

`func (o *ListTransactions200ResponseCursor) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListTransactions200ResponseCursor) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListTransactions200ResponseCursor) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *ListTransactions200ResponseCursor) GetData() []Transaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTransactions200ResponseCursor) GetDataOk() (*[]Transaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTransactions200ResponseCursor) SetData(v []Transaction)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


