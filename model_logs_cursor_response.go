/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.0-rc.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// LogsCursorResponse struct for LogsCursorResponse
type LogsCursorResponse struct {
	Cursor LogsCursorResponseCursor `json:"cursor"`
}

// NewLogsCursorResponse instantiates a new LogsCursorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsCursorResponse(cursor LogsCursorResponseCursor) *LogsCursorResponse {
	this := LogsCursorResponse{}
	this.Cursor = cursor
	return &this
}

// NewLogsCursorResponseWithDefaults instantiates a new LogsCursorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsCursorResponseWithDefaults() *LogsCursorResponse {
	this := LogsCursorResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *LogsCursorResponse) GetCursor() LogsCursorResponseCursor {
	if o == nil {
		var ret LogsCursorResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *LogsCursorResponse) GetCursorOk() (*LogsCursorResponseCursor, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *LogsCursorResponse) SetCursor(v LogsCursorResponseCursor) {
	o.Cursor = v
}

func (o LogsCursorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableLogsCursorResponse struct {
	value *LogsCursorResponse
	isSet bool
}

func (v NullableLogsCursorResponse) Get() *LogsCursorResponse {
	return v.value
}

func (v *NullableLogsCursorResponse) Set(val *LogsCursorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsCursorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsCursorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsCursorResponse(val *LogsCursorResponse) *NullableLogsCursorResponse {
	return &NullableLogsCursorResponse{value: val, isSet: true}
}

func (v NullableLogsCursorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsCursorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


