/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// TransactionsCursorResponse struct for TransactionsCursorResponse
type TransactionsCursorResponse struct {
	Cursor TransactionsCursorResponseCursor `json:"cursor"`
}

// NewTransactionsCursorResponse instantiates a new TransactionsCursorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsCursorResponse(cursor TransactionsCursorResponseCursor) *TransactionsCursorResponse {
	this := TransactionsCursorResponse{}
	this.Cursor = cursor
	return &this
}

// NewTransactionsCursorResponseWithDefaults instantiates a new TransactionsCursorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsCursorResponseWithDefaults() *TransactionsCursorResponse {
	this := TransactionsCursorResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *TransactionsCursorResponse) GetCursor() TransactionsCursorResponseCursor {
	if o == nil {
		var ret TransactionsCursorResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *TransactionsCursorResponse) GetCursorOk() (*TransactionsCursorResponseCursor, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *TransactionsCursorResponse) SetCursor(v TransactionsCursorResponseCursor) {
	o.Cursor = v
}

func (o TransactionsCursorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsCursorResponse struct {
	value *TransactionsCursorResponse
	isSet bool
}

func (v NullableTransactionsCursorResponse) Get() *TransactionsCursorResponse {
	return v.value
}

func (v *NullableTransactionsCursorResponse) Set(val *TransactionsCursorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsCursorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsCursorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsCursorResponse(val *TransactionsCursorResponse) *NullableTransactionsCursorResponse {
	return &NullableTransactionsCursorResponse{value: val, isSet: true}
}

func (v NullableTransactionsCursorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsCursorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


