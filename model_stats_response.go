/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.0-rc.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// StatsResponse struct for StatsResponse
type StatsResponse struct {
	Data Stats `json:"data"`
}

// NewStatsResponse instantiates a new StatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsResponse(data Stats) *StatsResponse {
	this := StatsResponse{}
	this.Data = data
	return &this
}

// NewStatsResponseWithDefaults instantiates a new StatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsResponseWithDefaults() *StatsResponse {
	this := StatsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *StatsResponse) GetData() Stats {
	if o == nil {
		var ret Stats
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StatsResponse) GetDataOk() (*Stats, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StatsResponse) SetData(v Stats) {
	o.Data = v
}

func (o StatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStatsResponse struct {
	value *StatsResponse
	isSet bool
}

func (v NullableStatsResponse) Get() *StatsResponse {
	return v.value
}

func (v *NullableStatsResponse) Set(val *StatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsResponse(val *StatsResponse) *NullableStatsResponse {
	return &NullableStatsResponse{value: val, isSet: true}
}

func (v NullableStatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


