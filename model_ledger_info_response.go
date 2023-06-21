/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// checks if the LedgerInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LedgerInfoResponse{}

// LedgerInfoResponse struct for LedgerInfoResponse
type LedgerInfoResponse struct {
	Data *LedgerInfo `json:"data,omitempty"`
}

// NewLedgerInfoResponse instantiates a new LedgerInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerInfoResponse() *LedgerInfoResponse {
	this := LedgerInfoResponse{}
	return &this
}

// NewLedgerInfoResponseWithDefaults instantiates a new LedgerInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerInfoResponseWithDefaults() *LedgerInfoResponse {
	this := LedgerInfoResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LedgerInfoResponse) GetData() LedgerInfo {
	if o == nil || IsNil(o.Data) {
		var ret LedgerInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfoResponse) GetDataOk() (*LedgerInfo, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LedgerInfoResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LedgerInfo and assigns it to the Data field.
func (o *LedgerInfoResponse) SetData(v LedgerInfo) {
	o.Data = &v
}

func (o LedgerInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LedgerInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLedgerInfoResponse struct {
	value *LedgerInfoResponse
	isSet bool
}

func (v NullableLedgerInfoResponse) Get() *LedgerInfoResponse {
	return v.value
}

func (v *NullableLedgerInfoResponse) Set(val *LedgerInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerInfoResponse(val *LedgerInfoResponse) *NullableLedgerInfoResponse {
	return &NullableLedgerInfoResponse{value: val, isSet: true}
}

func (v NullableLedgerInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


