/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.0-beta.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// ListLogs200ResponseCursorAllOf struct for ListLogs200ResponseCursorAllOf
type ListLogs200ResponseCursorAllOf struct {
	Data []Log `json:"data"`
}

// NewListLogs200ResponseCursorAllOf instantiates a new ListLogs200ResponseCursorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLogs200ResponseCursorAllOf(data []Log) *ListLogs200ResponseCursorAllOf {
	this := ListLogs200ResponseCursorAllOf{}
	this.Data = data
	return &this
}

// NewListLogs200ResponseCursorAllOfWithDefaults instantiates a new ListLogs200ResponseCursorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLogs200ResponseCursorAllOfWithDefaults() *ListLogs200ResponseCursorAllOf {
	this := ListLogs200ResponseCursorAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *ListLogs200ResponseCursorAllOf) GetData() []Log {
	if o == nil {
		var ret []Log
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListLogs200ResponseCursorAllOf) GetDataOk() (*[]Log, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListLogs200ResponseCursorAllOf) SetData(v []Log) {
	o.Data = v
}

func (o ListLogs200ResponseCursorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListLogs200ResponseCursorAllOf struct {
	value *ListLogs200ResponseCursorAllOf
	isSet bool
}

func (v NullableListLogs200ResponseCursorAllOf) Get() *ListLogs200ResponseCursorAllOf {
	return v.value
}

func (v *NullableListLogs200ResponseCursorAllOf) Set(val *ListLogs200ResponseCursorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListLogs200ResponseCursorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListLogs200ResponseCursorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLogs200ResponseCursorAllOf(val *ListLogs200ResponseCursorAllOf) *NullableListLogs200ResponseCursorAllOf {
	return &NullableListLogs200ResponseCursorAllOf{value: val, isSet: true}
}

func (v NullableListLogs200ResponseCursorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLogs200ResponseCursorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


