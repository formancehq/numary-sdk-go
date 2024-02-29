/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// Contract struct for Contract
type Contract struct {
	Account *string `json:"account,omitempty"`
	Expr map[string]interface{} `json:"expr"`
}

// NewContract instantiates a new Contract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContract(expr map[string]interface{}) *Contract {
	this := Contract{}
	this.Expr = expr
	return &this
}

// NewContractWithDefaults instantiates a new Contract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractWithDefaults() *Contract {
	this := Contract{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Contract) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Contract) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *Contract) SetAccount(v string) {
	o.Account = &v
}

// GetExpr returns the Expr field value
func (o *Contract) GetExpr() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Expr
}

// GetExprOk returns a tuple with the Expr field value
// and a boolean to check if the value has been set.
func (o *Contract) GetExprOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Expr, true
}

// SetExpr sets field value
func (o *Contract) SetExpr(v map[string]interface{}) {
	o.Expr = v
}

func (o Contract) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["expr"] = o.Expr
	}
	return json.Marshal(toSerialize)
}

type NullableContract struct {
	value *Contract
	isSet bool
}

func (v NullableContract) Get() *Contract {
	return v.value
}

func (v *NullableContract) Set(val *Contract) {
	v.value = val
	v.isSet = true
}

func (v NullableContract) IsSet() bool {
	return v.isSet
}

func (v *NullableContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContract(val *Contract) *NullableContract {
	return &NullableContract{value: val, isSet: true}
}

func (v NullableContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


