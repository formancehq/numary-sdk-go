/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// Script struct for Script
type Script struct {
	Plain string `json:"plain"`
	Vars *map[string]interface{} `json:"vars,omitempty"`
	// Reference to attach to the generated transaction
	Reference *string `json:"reference,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewScript instantiates a new Script object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScript(plain string) *Script {
	this := Script{}
	this.Plain = plain
	return &this
}

// NewScriptWithDefaults instantiates a new Script object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptWithDefaults() *Script {
	this := Script{}
	return &this
}

// GetPlain returns the Plain field value
func (o *Script) GetPlain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plain
}

// GetPlainOk returns a tuple with the Plain field value
// and a boolean to check if the value has been set.
func (o *Script) GetPlainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Plain, true
}

// SetPlain sets field value
func (o *Script) SetPlain(v string) {
	o.Plain = v
}

// GetVars returns the Vars field value if set, zero value otherwise.
func (o *Script) GetVars() map[string]interface{} {
	if o == nil || o.Vars == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Vars
}

// GetVarsOk returns a tuple with the Vars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetVarsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Vars == nil {
		return nil, false
	}
	return o.Vars, true
}

// HasVars returns a boolean if a field has been set.
func (o *Script) HasVars() bool {
	if o != nil && o.Vars != nil {
		return true
	}

	return false
}

// SetVars gets a reference to the given map[string]interface{} and assigns it to the Vars field.
func (o *Script) SetVars(v map[string]interface{}) {
	o.Vars = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Script) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Script) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Script) SetReference(v string) {
	o.Reference = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Script) GetMetadata() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Script) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Script) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Script) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o Script) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plain"] = o.Plain
	}
	if o.Vars != nil {
		toSerialize["vars"] = o.Vars
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableScript struct {
	value *Script
	isSet bool
}

func (v NullableScript) Get() *Script {
	return v.value
}

func (v *NullableScript) Set(val *Script) {
	v.value = val
	v.isSet = true
}

func (v NullableScript) IsSet() bool {
	return v.isSet
}

func (v *NullableScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScript(val *Script) *NullableScript {
	return &NullableScript{value: val, isSet: true}
}

func (v NullableScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


