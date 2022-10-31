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

// ScriptResult struct for ScriptResult
type ScriptResult struct {
	Details *string `json:"details,omitempty"`
	ErrorCode *string `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Transaction *Transaction `json:"transaction,omitempty"`
}

// NewScriptResult instantiates a new ScriptResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptResult() *ScriptResult {
	this := ScriptResult{}
	return &this
}

// NewScriptResultWithDefaults instantiates a new ScriptResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptResultWithDefaults() *ScriptResult {
	this := ScriptResult{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ScriptResult) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResult) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ScriptResult) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ScriptResult) SetDetails(v string) {
	o.Details = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ScriptResult) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResult) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ScriptResult) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ScriptResult) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ScriptResult) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResult) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ScriptResult) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ScriptResult) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *ScriptResult) GetTransaction() Transaction {
	if o == nil || o.Transaction == nil {
		var ret Transaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResult) GetTransactionOk() (*Transaction, bool) {
	if o == nil || o.Transaction == nil {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *ScriptResult) HasTransaction() bool {
	if o != nil && o.Transaction != nil {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given Transaction and assigns it to the Transaction field.
func (o *ScriptResult) SetTransaction(v Transaction) {
	o.Transaction = &v
}

func (o ScriptResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Transaction != nil {
		toSerialize["transaction"] = o.Transaction
	}
	return json.Marshal(toSerialize)
}

type NullableScriptResult struct {
	value *ScriptResult
	isSet bool
}

func (v NullableScriptResult) Get() *ScriptResult {
	return v.value
}

func (v *NullableScriptResult) Set(val *ScriptResult) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptResult) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptResult(val *ScriptResult) *NullableScriptResult {
	return &NullableScriptResult{value: val, isSet: true}
}

func (v NullableScriptResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


