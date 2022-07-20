# CreateTransactions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to [**Cursor**](Cursor.md) |  | [optional] 
**Data** | [**[]Transaction**](Transaction.md) |  | 

## Methods

### NewCreateTransactions200Response

`func NewCreateTransactions200Response(data []Transaction, ) *CreateTransactions200Response`

NewCreateTransactions200Response instantiates a new CreateTransactions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransactions200ResponseWithDefaults

`func NewCreateTransactions200ResponseWithDefaults() *CreateTransactions200Response`

NewCreateTransactions200ResponseWithDefaults instantiates a new CreateTransactions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *CreateTransactions200Response) GetCursor() Cursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *CreateTransactions200Response) GetCursorOk() (*Cursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *CreateTransactions200Response) SetCursor(v Cursor)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *CreateTransactions200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetData

`func (o *CreateTransactions200Response) GetData() []Transaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTransactions200Response) GetDataOk() (*[]Transaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTransactions200Response) SetData(v []Transaction)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


