/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// TransactionResponse struct for TransactionResponse
type TransactionResponse struct {
	Data Transaction `json:"data"`
}

// NewTransactionResponse instantiates a new TransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionResponse(data Transaction) *TransactionResponse {
	this := TransactionResponse{}
	this.Data = data
	return &this
}

// NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionResponseWithDefaults() *TransactionResponse {
	this := TransactionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *TransactionResponse) GetData() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetDataOk() (*Transaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionResponse) SetData(v Transaction) {
	o.Data = v
}

func (o TransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionResponse struct {
	value *TransactionResponse
	isSet bool
}

func (v NullableTransactionResponse) Get() *TransactionResponse {
	return v.value
}

func (v *NullableTransactionResponse) Set(val *TransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResponse(val *TransactionResponse) *NullableTransactionResponse {
	return &NullableTransactionResponse{value: val, isSet: true}
}

func (v NullableTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


