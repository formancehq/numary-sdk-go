/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// ConfigInfoResponse struct for ConfigInfoResponse
type ConfigInfoResponse struct {
	Data ConfigInfo `json:"data"`
}

// NewConfigInfoResponse instantiates a new ConfigInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigInfoResponse(data ConfigInfo) *ConfigInfoResponse {
	this := ConfigInfoResponse{}
	this.Data = data
	return &this
}

// NewConfigInfoResponseWithDefaults instantiates a new ConfigInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigInfoResponseWithDefaults() *ConfigInfoResponse {
	this := ConfigInfoResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ConfigInfoResponse) GetData() ConfigInfo {
	if o == nil {
		var ret ConfigInfo
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ConfigInfoResponse) GetDataOk() (*ConfigInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ConfigInfoResponse) SetData(v ConfigInfo) {
	o.Data = v
}

func (o ConfigInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableConfigInfoResponse struct {
	value *ConfigInfoResponse
	isSet bool
}

func (v NullableConfigInfoResponse) Get() *ConfigInfoResponse {
	return v.value
}

func (v *NullableConfigInfoResponse) Set(val *ConfigInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigInfoResponse(val *ConfigInfoResponse) *NullableConfigInfoResponse {
	return &NullableConfigInfoResponse{value: val, isSet: true}
}

func (v NullableConfigInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


