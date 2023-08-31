/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// Posting struct for Posting
type Posting struct {
	Amount int64 `json:"amount"`
	Asset string `json:"asset"`
	Destination string `json:"destination"`
	Source string `json:"source"`
}

// NewPosting instantiates a new Posting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPosting(amount int64, asset string, destination string, source string) *Posting {
	this := Posting{}
	this.Amount = amount
	this.Asset = asset
	this.Destination = destination
	this.Source = source
	return &this
}

// NewPostingWithDefaults instantiates a new Posting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostingWithDefaults() *Posting {
	this := Posting{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Posting) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Posting) GetAmountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Posting) SetAmount(v int64) {
	o.Amount = v
}

// GetAsset returns the Asset field value
func (o *Posting) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *Posting) GetAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *Posting) SetAsset(v string) {
	o.Asset = v
}

// GetDestination returns the Destination field value
func (o *Posting) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *Posting) GetDestinationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *Posting) SetDestination(v string) {
	o.Destination = v
}

// GetSource returns the Source field value
func (o *Posting) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Posting) GetSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Posting) SetSource(v string) {
	o.Source = v
}

func (o Posting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullablePosting struct {
	value *Posting
	isSet bool
}

func (v NullablePosting) Get() *Posting {
	return v.value
}

func (v *NullablePosting) Set(val *Posting) {
	v.value = val
	v.isSet = true
}

func (v NullablePosting) IsSet() bool {
	return v.isSet
}

func (v *NullablePosting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePosting(val *Posting) *NullablePosting {
	return &NullablePosting{value: val, isSet: true}
}

func (v NullablePosting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePosting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


