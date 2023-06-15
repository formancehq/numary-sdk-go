/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.4-beta.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// checks if the AggregateBalancesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregateBalancesResponse{}

// AggregateBalancesResponse struct for AggregateBalancesResponse
type AggregateBalancesResponse struct {
	Data map[string]int64 `json:"data"`
}

// NewAggregateBalancesResponse instantiates a new AggregateBalancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregateBalancesResponse(data map[string]int64) *AggregateBalancesResponse {
	this := AggregateBalancesResponse{}
	this.Data = data
	return &this
}

// NewAggregateBalancesResponseWithDefaults instantiates a new AggregateBalancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateBalancesResponseWithDefaults() *AggregateBalancesResponse {
	this := AggregateBalancesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AggregateBalancesResponse) GetData() map[string]int64 {
	if o == nil {
		var ret map[string]int64
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AggregateBalancesResponse) GetDataOk() (*map[string]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AggregateBalancesResponse) SetData(v map[string]int64) {
	o.Data = v
}

func (o AggregateBalancesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregateBalancesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAggregateBalancesResponse struct {
	value *AggregateBalancesResponse
	isSet bool
}

func (v NullableAggregateBalancesResponse) Get() *AggregateBalancesResponse {
	return v.value
}

func (v *NullableAggregateBalancesResponse) Set(val *AggregateBalancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateBalancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateBalancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateBalancesResponse(val *AggregateBalancesResponse) *NullableAggregateBalancesResponse {
	return &NullableAggregateBalancesResponse{value: val, isSet: true}
}

func (v NullableAggregateBalancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateBalancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


