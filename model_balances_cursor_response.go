/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// BalancesCursorResponse struct for BalancesCursorResponse
type BalancesCursorResponse struct {
	Cursor BalancesCursorResponseCursor `json:"cursor"`
}

// NewBalancesCursorResponse instantiates a new BalancesCursorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalancesCursorResponse(cursor BalancesCursorResponseCursor) *BalancesCursorResponse {
	this := BalancesCursorResponse{}
	this.Cursor = cursor
	return &this
}

// NewBalancesCursorResponseWithDefaults instantiates a new BalancesCursorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalancesCursorResponseWithDefaults() *BalancesCursorResponse {
	this := BalancesCursorResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *BalancesCursorResponse) GetCursor() BalancesCursorResponseCursor {
	if o == nil {
		var ret BalancesCursorResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *BalancesCursorResponse) GetCursorOk() (*BalancesCursorResponseCursor, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *BalancesCursorResponse) SetCursor(v BalancesCursorResponseCursor) {
	o.Cursor = v
}

func (o BalancesCursorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableBalancesCursorResponse struct {
	value *BalancesCursorResponse
	isSet bool
}

func (v NullableBalancesCursorResponse) Get() *BalancesCursorResponse {
	return v.value
}

func (v *NullableBalancesCursorResponse) Set(val *BalancesCursorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancesCursorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancesCursorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancesCursorResponse(val *BalancesCursorResponse) *NullableBalancesCursorResponse {
	return &NullableBalancesCursorResponse{value: val, isSet: true}
}

func (v NullableBalancesCursorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancesCursorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


