/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.7.3-rc.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// ListAccounts200Response struct for ListAccounts200Response
type ListAccounts200Response struct {
	Cursor ListAccounts200ResponseCursor `json:"cursor"`
}

// NewListAccounts200Response instantiates a new ListAccounts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccounts200Response(cursor ListAccounts200ResponseCursor) *ListAccounts200Response {
	this := ListAccounts200Response{}
	this.Cursor = cursor
	return &this
}

// NewListAccounts200ResponseWithDefaults instantiates a new ListAccounts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccounts200ResponseWithDefaults() *ListAccounts200Response {
	this := ListAccounts200Response{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *ListAccounts200Response) GetCursor() ListAccounts200ResponseCursor {
	if o == nil {
		var ret ListAccounts200ResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *ListAccounts200Response) GetCursorOk() (*ListAccounts200ResponseCursor, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *ListAccounts200Response) SetCursor(v ListAccounts200ResponseCursor) {
	o.Cursor = v
}

func (o ListAccounts200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableListAccounts200Response struct {
	value *ListAccounts200Response
	isSet bool
}

func (v NullableListAccounts200Response) Get() *ListAccounts200Response {
	return v.value
}

func (v *NullableListAccounts200Response) Set(val *ListAccounts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccounts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccounts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccounts200Response(val *ListAccounts200Response) *NullableListAccounts200Response {
	return &NullableListAccounts200Response{value: val, isSet: true}
}

func (v NullableListAccounts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccounts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

