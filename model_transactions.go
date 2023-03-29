/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// checks if the Transactions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transactions{}

// Transactions struct for Transactions
type Transactions struct {
	Transactions []TransactionData `json:"transactions"`
}

// NewTransactions instantiates a new Transactions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactions(transactions []TransactionData) *Transactions {
	this := Transactions{}
	this.Transactions = transactions
	return &this
}

// NewTransactionsWithDefaults instantiates a new Transactions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsWithDefaults() *Transactions {
	this := Transactions{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *Transactions) GetTransactions() []TransactionData {
	if o == nil {
		var ret []TransactionData
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *Transactions) GetTransactionsOk() ([]TransactionData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *Transactions) SetTransactions(v []TransactionData) {
	o.Transactions = v
}

func (o Transactions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transactions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

type NullableTransactions struct {
	value *Transactions
	isSet bool
}

func (v NullableTransactions) Get() *Transactions {
	return v.value
}

func (v *NullableTransactions) Set(val *Transactions) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactions) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactions(val *Transactions) *NullableTransactions {
	return &NullableTransactions{value: val, isSet: true}
}

func (v NullableTransactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


