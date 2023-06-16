/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// checks if the LedgerInfoStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LedgerInfoStorage{}

// LedgerInfoStorage struct for LedgerInfoStorage
type LedgerInfoStorage struct {
	Migrations []MigrationInfo `json:"migrations,omitempty"`
}

// NewLedgerInfoStorage instantiates a new LedgerInfoStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerInfoStorage() *LedgerInfoStorage {
	this := LedgerInfoStorage{}
	return &this
}

// NewLedgerInfoStorageWithDefaults instantiates a new LedgerInfoStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerInfoStorageWithDefaults() *LedgerInfoStorage {
	this := LedgerInfoStorage{}
	return &this
}

// GetMigrations returns the Migrations field value if set, zero value otherwise.
func (o *LedgerInfoStorage) GetMigrations() []MigrationInfo {
	if o == nil || IsNil(o.Migrations) {
		var ret []MigrationInfo
		return ret
	}
	return o.Migrations
}

// GetMigrationsOk returns a tuple with the Migrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfoStorage) GetMigrationsOk() ([]MigrationInfo, bool) {
	if o == nil || IsNil(o.Migrations) {
		return nil, false
	}
	return o.Migrations, true
}

// HasMigrations returns a boolean if a field has been set.
func (o *LedgerInfoStorage) HasMigrations() bool {
	if o != nil && !IsNil(o.Migrations) {
		return true
	}

	return false
}

// SetMigrations gets a reference to the given []MigrationInfo and assigns it to the Migrations field.
func (o *LedgerInfoStorage) SetMigrations(v []MigrationInfo) {
	o.Migrations = v
}

func (o LedgerInfoStorage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LedgerInfoStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Migrations) {
		toSerialize["migrations"] = o.Migrations
	}
	return toSerialize, nil
}

type NullableLedgerInfoStorage struct {
	value *LedgerInfoStorage
	isSet bool
}

func (v NullableLedgerInfoStorage) Get() *LedgerInfoStorage {
	return v.value
}

func (v *NullableLedgerInfoStorage) Set(val *LedgerInfoStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerInfoStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerInfoStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerInfoStorage(val *LedgerInfoStorage) *NullableLedgerInfoStorage {
	return &NullableLedgerInfoStorage{value: val, isSet: true}
}

func (v NullableLedgerInfoStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerInfoStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


