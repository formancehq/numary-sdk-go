/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// TransactionsCursorResponseCursor struct for TransactionsCursorResponseCursor
type TransactionsCursorResponseCursor struct {
	PageSize int64 `json:"pageSize"`
	HasMore bool `json:"hasMore"`
	Previous *string `json:"previous,omitempty"`
	Next *string `json:"next,omitempty"`
	Data []Transaction `json:"data"`
}

// NewTransactionsCursorResponseCursor instantiates a new TransactionsCursorResponseCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsCursorResponseCursor(pageSize int64, hasMore bool, data []Transaction) *TransactionsCursorResponseCursor {
	this := TransactionsCursorResponseCursor{}
	this.PageSize = pageSize
	this.HasMore = hasMore
	this.Data = data
	return &this
}

// NewTransactionsCursorResponseCursorWithDefaults instantiates a new TransactionsCursorResponseCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsCursorResponseCursorWithDefaults() *TransactionsCursorResponseCursor {
	this := TransactionsCursorResponseCursor{}
	return &this
}

// GetPageSize returns the PageSize field value
func (o *TransactionsCursorResponseCursor) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *TransactionsCursorResponseCursor) GetPageSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *TransactionsCursorResponseCursor) SetPageSize(v int64) {
	o.PageSize = v
}

// GetHasMore returns the HasMore field value
func (o *TransactionsCursorResponseCursor) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *TransactionsCursorResponseCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *TransactionsCursorResponseCursor) SetHasMore(v bool) {
	o.HasMore = v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *TransactionsCursorResponseCursor) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsCursorResponseCursor) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *TransactionsCursorResponseCursor) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *TransactionsCursorResponseCursor) SetPrevious(v string) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *TransactionsCursorResponseCursor) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsCursorResponseCursor) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *TransactionsCursorResponseCursor) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *TransactionsCursorResponseCursor) SetNext(v string) {
	o.Next = &v
}

// GetData returns the Data field value
func (o *TransactionsCursorResponseCursor) GetData() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionsCursorResponseCursor) GetDataOk() (*[]Transaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionsCursorResponseCursor) SetData(v []Transaction) {
	o.Data = v
}

func (o TransactionsCursorResponseCursor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pageSize"] = o.PageSize
	}
	if true {
		toSerialize["hasMore"] = o.HasMore
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsCursorResponseCursor struct {
	value *TransactionsCursorResponseCursor
	isSet bool
}

func (v NullableTransactionsCursorResponseCursor) Get() *TransactionsCursorResponseCursor {
	return v.value
}

func (v *NullableTransactionsCursorResponseCursor) Set(val *TransactionsCursorResponseCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsCursorResponseCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsCursorResponseCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsCursorResponseCursor(val *TransactionsCursorResponseCursor) *NullableTransactionsCursorResponseCursor {
	return &NullableTransactionsCursorResponseCursor{value: val, isSet: true}
}

func (v NullableTransactionsCursorResponseCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsCursorResponseCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


