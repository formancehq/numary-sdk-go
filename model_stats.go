/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// checks if the Stats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Stats{}

// Stats struct for Stats
type Stats struct {
	Accounts int64 `json:"accounts"`
	Transactions int64 `json:"transactions"`
}

// NewStats instantiates a new Stats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStats(accounts int64, transactions int64) *Stats {
	this := Stats{}
	this.Accounts = accounts
	this.Transactions = transactions
	return &this
}

// NewStatsWithDefaults instantiates a new Stats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsWithDefaults() *Stats {
	this := Stats{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *Stats) GetAccounts() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *Stats) GetAccountsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *Stats) SetAccounts(v int64) {
	o.Accounts = v
}

// GetTransactions returns the Transactions field value
func (o *Stats) GetTransactions() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *Stats) GetTransactionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *Stats) SetTransactions(v int64) {
	o.Transactions = v
}

func (o Stats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Stats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

type NullableStats struct {
	value *Stats
	isSet bool
}

func (v NullableStats) Get() *Stats {
	return v.value
}

func (v *NullableStats) Set(val *Stats) {
	v.value = val
	v.isSet = true
}

func (v NullableStats) IsSet() bool {
	return v.isSet
}

func (v *NullableStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStats(val *Stats) *NullableStats {
	return &NullableStats{value: val, isSet: true}
}

func (v NullableStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


