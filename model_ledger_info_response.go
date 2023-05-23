/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

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
	if o == nil || o.Data == nil {
		var ret LedgerInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfoResponse) GetDataOk() (*LedgerInfo, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LedgerInfoResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given LedgerInfo and assigns it to the Data field.
func (o *LedgerInfoResponse) SetData(v LedgerInfo) {
	o.Data = &v
}

func (o LedgerInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
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


