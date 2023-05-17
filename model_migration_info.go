/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
	"time"
)

// checks if the MigrationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrationInfo{}

// MigrationInfo struct for MigrationInfo
type MigrationInfo struct {
	Version *int64 `json:"version,omitempty"`
	Name *string `json:"name,omitempty"`
	Date *time.Time `json:"date,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewMigrationInfo instantiates a new MigrationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationInfo() *MigrationInfo {
	this := MigrationInfo{}
	return &this
}

// NewMigrationInfoWithDefaults instantiates a new MigrationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationInfoWithDefaults() *MigrationInfo {
	this := MigrationInfo{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MigrationInfo) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationInfo) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MigrationInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *MigrationInfo) SetVersion(v int64) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MigrationInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MigrationInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MigrationInfo) SetName(v string) {
	o.Name = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *MigrationInfo) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationInfo) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *MigrationInfo) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *MigrationInfo) SetDate(v time.Time) {
	o.Date = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MigrationInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MigrationInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MigrationInfo) SetState(v string) {
	o.State = &v
}

func (o MigrationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableMigrationInfo struct {
	value *MigrationInfo
	isSet bool
}

func (v NullableMigrationInfo) Get() *MigrationInfo {
	return v.value
}

func (v *NullableMigrationInfo) Set(val *MigrationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationInfo(val *MigrationInfo) *NullableMigrationInfo {
	return &NullableMigrationInfo{value: val, isSet: true}
}

func (v NullableMigrationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


