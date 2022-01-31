# TransactionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to [**Cursor**](Cursor.md) |  | [optional] 
**Data** | [**[]Transaction**](Transaction.md) |  | 

## Methods

### NewTransactionListResponse

`func NewTransactionListResponse(data []Transaction, ) *TransactionListResponse`

NewTransactionListResponse instantiates a new TransactionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionListResponseWithDefaults

`func NewTransactionListResponseWithDefaults() *TransactionListResponse`

NewTransactionListResponseWithDefaults instantiates a new TransactionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *TransactionListResponse) GetCursor() Cursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *TransactionListResponse) GetCursorOk() (*Cursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *TransactionListResponse) SetCursor(v Cursor)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *TransactionListResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetData

`func (o *TransactionListResponse) GetData() []Transaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionListResponse) GetDataOk() (*[]Transaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionListResponse) SetData(v []Transaction)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


