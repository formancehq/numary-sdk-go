/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.8.0-rc.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// GetBalances200ResponseCursorAllOf struct for GetBalances200ResponseCursorAllOf
type GetBalances200ResponseCursorAllOf struct {
	Data []map[string]map[string]int32 `json:"data"`
}

// NewGetBalances200ResponseCursorAllOf instantiates a new GetBalances200ResponseCursorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalances200ResponseCursorAllOf(data []map[string]map[string]int32) *GetBalances200ResponseCursorAllOf {
	this := GetBalances200ResponseCursorAllOf{}
	this.Data = data
	return &this
}

// NewGetBalances200ResponseCursorAllOfWithDefaults instantiates a new GetBalances200ResponseCursorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalances200ResponseCursorAllOfWithDefaults() *GetBalances200ResponseCursorAllOf {
	this := GetBalances200ResponseCursorAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *GetBalances200ResponseCursorAllOf) GetData() []map[string]map[string]int32 {
	if o == nil {
		var ret []map[string]map[string]int32
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetBalances200ResponseCursorAllOf) GetDataOk() (*[]map[string]map[string]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetBalances200ResponseCursorAllOf) SetData(v []map[string]map[string]int32) {
	o.Data = v
}

func (o GetBalances200ResponseCursorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetBalances200ResponseCursorAllOf struct {
	value *GetBalances200ResponseCursorAllOf
	isSet bool
}

func (v NullableGetBalances200ResponseCursorAllOf) Get() *GetBalances200ResponseCursorAllOf {
	return v.value
}

func (v *NullableGetBalances200ResponseCursorAllOf) Set(val *GetBalances200ResponseCursorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalances200ResponseCursorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalances200ResponseCursorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalances200ResponseCursorAllOf(val *GetBalances200ResponseCursorAllOf) *NullableGetBalances200ResponseCursorAllOf {
	return &NullableGetBalances200ResponseCursorAllOf{value: val, isSet: true}
}

func (v NullableGetBalances200ResponseCursorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalances200ResponseCursorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


