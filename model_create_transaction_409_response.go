/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.7.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// CreateTransaction409Response struct for CreateTransaction409Response
type CreateTransaction409Response struct {
	ErrorCode string `json:"error_code"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewCreateTransaction409Response instantiates a new CreateTransaction409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransaction409Response(errorCode string) *CreateTransaction409Response {
	this := CreateTransaction409Response{}
	this.ErrorCode = errorCode
	return &this
}

// NewCreateTransaction409ResponseWithDefaults instantiates a new CreateTransaction409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransaction409ResponseWithDefaults() *CreateTransaction409Response {
	this := CreateTransaction409Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *CreateTransaction409Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *CreateTransaction409Response) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *CreateTransaction409Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CreateTransaction409Response) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransaction409Response) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CreateTransaction409Response) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CreateTransaction409Response) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o CreateTransaction409Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTransaction409Response struct {
	value *CreateTransaction409Response
	isSet bool
}

func (v NullableCreateTransaction409Response) Get() *CreateTransaction409Response {
	return v.value
}

func (v *NullableCreateTransaction409Response) Set(val *CreateTransaction409Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransaction409Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransaction409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransaction409Response(val *CreateTransaction409Response) *NullableCreateTransaction409Response {
	return &NullableCreateTransaction409Response{value: val, isSet: true}
}

func (v NullableCreateTransaction409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransaction409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


