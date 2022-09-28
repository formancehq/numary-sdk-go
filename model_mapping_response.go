/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.8.0-rc.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// MappingResponse struct for MappingResponse
type MappingResponse struct {
	Data NullableMapping `json:"data,omitempty"`
}

// NewMappingResponse instantiates a new MappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappingResponse() *MappingResponse {
	this := MappingResponse{}
	return &this
}

// NewMappingResponseWithDefaults instantiates a new MappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappingResponseWithDefaults() *MappingResponse {
	this := MappingResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MappingResponse) GetData() Mapping {
	if o == nil || o.Data.Get() == nil {
		var ret Mapping
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MappingResponse) GetDataOk() (*Mapping, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *MappingResponse) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableMapping and assigns it to the Data field.
func (o *MappingResponse) SetData(v Mapping) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *MappingResponse) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *MappingResponse) UnsetData() {
	o.Data.Unset()
}

func (o MappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMappingResponse struct {
	value *MappingResponse
	isSet bool
}

func (v NullableMappingResponse) Get() *MappingResponse {
	return v.value
}

func (v *NullableMappingResponse) Set(val *MappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMappingResponse(val *MappingResponse) *NullableMappingResponse {
	return &NullableMappingResponse{value: val, isSet: true}
}

func (v NullableMappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


