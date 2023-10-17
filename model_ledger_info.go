/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// LedgerInfo struct for LedgerInfo
type LedgerInfo struct {
	Name *string `json:"name,omitempty"`
	Storage *LedgerInfoStorage `json:"storage,omitempty"`
}

// NewLedgerInfo instantiates a new LedgerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerInfo() *LedgerInfo {
	this := LedgerInfo{}
	return &this
}

// NewLedgerInfoWithDefaults instantiates a new LedgerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerInfoWithDefaults() *LedgerInfo {
	this := LedgerInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LedgerInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LedgerInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LedgerInfo) SetName(v string) {
	o.Name = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *LedgerInfo) GetStorage() LedgerInfoStorage {
	if o == nil || o.Storage == nil {
		var ret LedgerInfoStorage
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetStorageOk() (*LedgerInfoStorage, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *LedgerInfo) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given LedgerInfoStorage and assigns it to the Storage field.
func (o *LedgerInfo) SetStorage(v LedgerInfoStorage) {
	o.Storage = &v
}

func (o LedgerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return json.Marshal(toSerialize)
}

type NullableLedgerInfo struct {
	value *LedgerInfo
	isSet bool
}

func (v NullableLedgerInfo) Get() *LedgerInfo {
	return v.value
}

func (v *NullableLedgerInfo) Set(val *LedgerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerInfo(val *LedgerInfo) *NullableLedgerInfo {
	return &NullableLedgerInfo{value: val, isSet: true}
}

func (v NullableLedgerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


