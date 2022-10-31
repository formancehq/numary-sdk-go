/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// AccountWithVolumesAndBalances struct for AccountWithVolumesAndBalances
type AccountWithVolumesAndBalances struct {
	Address string `json:"address"`
	Type *string `json:"type,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	Volumes *map[string]map[string]int32 `json:"volumes,omitempty"`
	Balances *map[string]int32 `json:"balances,omitempty"`
}

// NewAccountWithVolumesAndBalances instantiates a new AccountWithVolumesAndBalances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountWithVolumesAndBalances(address string) *AccountWithVolumesAndBalances {
	this := AccountWithVolumesAndBalances{}
	this.Address = address
	return &this
}

// NewAccountWithVolumesAndBalancesWithDefaults instantiates a new AccountWithVolumesAndBalances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithVolumesAndBalancesWithDefaults() *AccountWithVolumesAndBalances {
	this := AccountWithVolumesAndBalances{}
	return &this
}

// GetAddress returns the Address field value
func (o *AccountWithVolumesAndBalances) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AccountWithVolumesAndBalances) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AccountWithVolumesAndBalances) SetAddress(v string) {
	o.Address = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountWithVolumesAndBalances) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountWithVolumesAndBalances) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountWithVolumesAndBalances) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountWithVolumesAndBalances) SetType(v string) {
	o.Type = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AccountWithVolumesAndBalances) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountWithVolumesAndBalances) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AccountWithVolumesAndBalances) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AccountWithVolumesAndBalances) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *AccountWithVolumesAndBalances) GetVolumes() map[string]map[string]int32 {
	if o == nil || o.Volumes == nil {
		var ret map[string]map[string]int32
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountWithVolumesAndBalances) GetVolumesOk() (*map[string]map[string]int32, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *AccountWithVolumesAndBalances) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given map[string]map[string]int32 and assigns it to the Volumes field.
func (o *AccountWithVolumesAndBalances) SetVolumes(v map[string]map[string]int32) {
	o.Volumes = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *AccountWithVolumesAndBalances) GetBalances() map[string]int32 {
	if o == nil || o.Balances == nil {
		var ret map[string]int32
		return ret
	}
	return *o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountWithVolumesAndBalances) GetBalancesOk() (*map[string]int32, bool) {
	if o == nil || o.Balances == nil {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *AccountWithVolumesAndBalances) HasBalances() bool {
	if o != nil && o.Balances != nil {
		return true
	}

	return false
}

// SetBalances gets a reference to the given map[string]int32 and assigns it to the Balances field.
func (o *AccountWithVolumesAndBalances) SetBalances(v map[string]int32) {
	o.Balances = &v
}

func (o AccountWithVolumesAndBalances) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if o.Balances != nil {
		toSerialize["balances"] = o.Balances
	}
	return json.Marshal(toSerialize)
}

type NullableAccountWithVolumesAndBalances struct {
	value *AccountWithVolumesAndBalances
	isSet bool
}

func (v NullableAccountWithVolumesAndBalances) Get() *AccountWithVolumesAndBalances {
	return v.value
}

func (v *NullableAccountWithVolumesAndBalances) Set(val *AccountWithVolumesAndBalances) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountWithVolumesAndBalances) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountWithVolumesAndBalances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountWithVolumesAndBalances(val *AccountWithVolumesAndBalances) *NullableAccountWithVolumesAndBalances {
	return &NullableAccountWithVolumesAndBalances{value: val, isSet: true}
}

func (v NullableAccountWithVolumesAndBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountWithVolumesAndBalances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


