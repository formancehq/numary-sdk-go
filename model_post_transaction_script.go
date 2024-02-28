/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// PostTransactionScript struct for PostTransactionScript
type PostTransactionScript struct {
	Plain string `json:"plain"`
	Vars *map[string]interface{} `json:"vars,omitempty"`
}

// NewPostTransactionScript instantiates a new PostTransactionScript object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTransactionScript(plain string) *PostTransactionScript {
	this := PostTransactionScript{}
	this.Plain = plain
	return &this
}

// NewPostTransactionScriptWithDefaults instantiates a new PostTransactionScript object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTransactionScriptWithDefaults() *PostTransactionScript {
	this := PostTransactionScript{}
	return &this
}

// GetPlain returns the Plain field value
func (o *PostTransactionScript) GetPlain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plain
}

// GetPlainOk returns a tuple with the Plain field value
// and a boolean to check if the value has been set.
func (o *PostTransactionScript) GetPlainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Plain, true
}

// SetPlain sets field value
func (o *PostTransactionScript) SetPlain(v string) {
	o.Plain = v
}

// GetVars returns the Vars field value if set, zero value otherwise.
func (o *PostTransactionScript) GetVars() map[string]interface{} {
	if o == nil || o.Vars == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Vars
}

// GetVarsOk returns a tuple with the Vars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionScript) GetVarsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Vars == nil {
		return nil, false
	}
	return o.Vars, true
}

// HasVars returns a boolean if a field has been set.
func (o *PostTransactionScript) HasVars() bool {
	if o != nil && o.Vars != nil {
		return true
	}

	return false
}

// SetVars gets a reference to the given map[string]interface{} and assigns it to the Vars field.
func (o *PostTransactionScript) SetVars(v map[string]interface{}) {
	o.Vars = &v
}

func (o PostTransactionScript) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plain"] = o.Plain
	}
	if o.Vars != nil {
		toSerialize["vars"] = o.Vars
	}
	return json.Marshal(toSerialize)
}

type NullablePostTransactionScript struct {
	value *PostTransactionScript
	isSet bool
}

func (v NullablePostTransactionScript) Get() *PostTransactionScript {
	return v.value
}

func (v *NullablePostTransactionScript) Set(val *PostTransactionScript) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTransactionScript) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTransactionScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTransactionScript(val *PostTransactionScript) *NullablePostTransactionScript {
	return &NullablePostTransactionScript{value: val, isSet: true}
}

func (v NullablePostTransactionScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTransactionScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


