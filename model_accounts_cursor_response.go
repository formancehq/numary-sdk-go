/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// AccountsCursorResponse struct for AccountsCursorResponse
type AccountsCursorResponse struct {
	Cursor AccountsCursorResponseCursor `json:"cursor"`
}

// NewAccountsCursorResponse instantiates a new AccountsCursorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsCursorResponse(cursor AccountsCursorResponseCursor) *AccountsCursorResponse {
	this := AccountsCursorResponse{}
	this.Cursor = cursor
	return &this
}

// NewAccountsCursorResponseWithDefaults instantiates a new AccountsCursorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsCursorResponseWithDefaults() *AccountsCursorResponse {
	this := AccountsCursorResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *AccountsCursorResponse) GetCursor() AccountsCursorResponseCursor {
	if o == nil {
		var ret AccountsCursorResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *AccountsCursorResponse) GetCursorOk() (*AccountsCursorResponseCursor, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *AccountsCursorResponse) SetCursor(v AccountsCursorResponseCursor) {
	o.Cursor = v
}

func (o AccountsCursorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableAccountsCursorResponse struct {
	value *AccountsCursorResponse
	isSet bool
}

func (v NullableAccountsCursorResponse) Get() *AccountsCursorResponse {
	return v.value
}

func (v *NullableAccountsCursorResponse) Set(val *AccountsCursorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsCursorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsCursorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsCursorResponse(val *AccountsCursorResponse) *NullableAccountsCursorResponse {
	return &NullableAccountsCursorResponse{value: val, isSet: true}
}

func (v NullableAccountsCursorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsCursorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


