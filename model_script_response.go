/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// ScriptResponse struct for ScriptResponse
type ScriptResponse struct {
	ErrorCode *ErrorsEnum `json:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	Details *string `json:"details,omitempty"`
	Transaction *Transaction `json:"transaction,omitempty"`
}

// NewScriptResponse instantiates a new ScriptResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptResponse() *ScriptResponse {
	this := ScriptResponse{}
	return &this
}

// NewScriptResponseWithDefaults instantiates a new ScriptResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptResponseWithDefaults() *ScriptResponse {
	this := ScriptResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ScriptResponse) GetErrorCode() ErrorsEnum {
	if o == nil || o.ErrorCode == nil {
		var ret ErrorsEnum
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResponse) GetErrorCodeOk() (*ErrorsEnum, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ScriptResponse) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given ErrorsEnum and assigns it to the ErrorCode field.
func (o *ScriptResponse) SetErrorCode(v ErrorsEnum) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ScriptResponse) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ScriptResponse) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ScriptResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ScriptResponse) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResponse) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ScriptResponse) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ScriptResponse) SetDetails(v string) {
	o.Details = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *ScriptResponse) GetTransaction() Transaction {
	if o == nil || o.Transaction == nil {
		var ret Transaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptResponse) GetTransactionOk() (*Transaction, bool) {
	if o == nil || o.Transaction == nil {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *ScriptResponse) HasTransaction() bool {
	if o != nil && o.Transaction != nil {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given Transaction and assigns it to the Transaction field.
func (o *ScriptResponse) SetTransaction(v Transaction) {
	o.Transaction = &v
}

func (o ScriptResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Transaction != nil {
		toSerialize["transaction"] = o.Transaction
	}
	return json.Marshal(toSerialize)
}

type NullableScriptResponse struct {
	value *ScriptResponse
	isSet bool
}

func (v NullableScriptResponse) Get() *ScriptResponse {
	return v.value
}

func (v *NullableScriptResponse) Set(val *ScriptResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptResponse(val *ScriptResponse) *NullableScriptResponse {
	return &NullableScriptResponse{value: val, isSet: true}
}

func (v NullableScriptResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


