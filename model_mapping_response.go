/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// MappingResponse struct for MappingResponse
type MappingResponse struct {
	Data Mapping `json:"data"`
}

// NewMappingResponse instantiates a new MappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappingResponse(data Mapping) *MappingResponse {
	this := MappingResponse{}
	this.Data = data
	return &this
}

// NewMappingResponseWithDefaults instantiates a new MappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappingResponseWithDefaults() *MappingResponse {
	this := MappingResponse{}
	return &this
}

// GetData returns the Data field value
func (o *MappingResponse) GetData() Mapping {
	if o == nil {
		var ret Mapping
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MappingResponse) GetDataOk() (*Mapping, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MappingResponse) SetData(v Mapping) {
	o.Data = v
}

func (o MappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
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


