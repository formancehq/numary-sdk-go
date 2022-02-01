# TransactionCommitError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Postings** | [**[]Posting**](Posting.md) |  | 
**Reference** | Pointer to **string** |  | [optional] 
**Timestamp** | **time.Time** |  | 
**Txid** | **int32** |  | 
**ErrorCode** | Pointer to [**ErrorCode**](ErrorCode.md) |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionCommitError

`func NewTransactionCommitError(postings []Posting, timestamp time.Time, txid int32, ) *TransactionCommitError`

NewTransactionCommitError instantiates a new TransactionCommitError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCommitErrorWithDefaults

`func NewTransactionCommitErrorWithDefaults() *TransactionCommitError`

NewTransactionCommitErrorWithDefaults instantiates a new TransactionCommitError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *TransactionCommitError) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransactionCommitError) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransactionCommitError) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransactionCommitError) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TransactionCommitError) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TransactionCommitError) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPostings

`func (o *TransactionCommitError) GetPostings() []Posting`

GetPostings returns the Postings field if non-nil, zero value otherwise.

### GetPostingsOk

`func (o *TransactionCommitError) GetPostingsOk() (*[]Posting, bool)`

GetPostingsOk returns a tuple with the Postings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostings

`func (o *TransactionCommitError) SetPostings(v []Posting)`

SetPostings sets Postings field to given value.


### GetReference

`func (o *TransactionCommitError) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TransactionCommitError) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TransactionCommitError) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *TransactionCommitError) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTimestamp

`func (o *TransactionCommitError) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TransactionCommitError) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TransactionCommitError) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTxid

`func (o *TransactionCommitError) GetTxid() int32`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *TransactionCommitError) GetTxidOk() (*int32, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *TransactionCommitError) SetTxid(v int32)`

SetTxid sets Txid field to given value.


### GetErrorCode

`func (o *TransactionCommitError) GetErrorCode() ErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TransactionCommitError) GetErrorCodeOk() (*ErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TransactionCommitError) SetErrorCode(v ErrorCode)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TransactionCommitError) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TransactionCommitError) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TransactionCommitError) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TransactionCommitError) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TransactionCommitError) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


