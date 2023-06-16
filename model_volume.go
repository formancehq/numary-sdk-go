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

// checks if the Volume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Volume{}

// Volume struct for Volume
type Volume struct {
	Input int64 `json:"input"`
	Output int64 `json:"output"`
	Balance *int64 `json:"balance,omitempty"`
}

// NewVolume instantiates a new Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolume(input int64, output int64) *Volume {
	this := Volume{}
	this.Input = input
	this.Output = output
	return &this
}

// NewVolumeWithDefaults instantiates a new Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeWithDefaults() *Volume {
	this := Volume{}
	return &this
}

// GetInput returns the Input field value
func (o *Volume) GetInput() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *Volume) GetInputOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *Volume) SetInput(v int64) {
	o.Input = v
}

// GetOutput returns the Output field value
func (o *Volume) GetOutput() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Output
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
func (o *Volume) GetOutputOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Output, true
}

// SetOutput sets field value
func (o *Volume) SetOutput(v int64) {
	o.Output = v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *Volume) GetBalance() int64 {
	if o == nil || IsNil(o.Balance) {
		var ret int64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetBalanceOk() (*int64, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *Volume) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int64 and assigns it to the Balance field.
func (o *Volume) SetBalance(v int64) {
	o.Balance = &v
}

func (o Volume) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Volume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input"] = o.Input
	toSerialize["output"] = o.Output
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	return toSerialize, nil
}

type NullableVolume struct {
	value *Volume
	isSet bool
}

func (v NullableVolume) Get() *Volume {
	return v.value
}

func (v *NullableVolume) Set(val *Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolume(val *Volume) *NullableVolume {
	return &NullableVolume{value: val, isSet: true}
}

func (v NullableVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


