/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.8.0-rc.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// GetTransaction400Response struct for GetTransaction400Response
type GetTransaction400Response struct {
	ErrorCode string `json:"error_code"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewGetTransaction400Response instantiates a new GetTransaction400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransaction400Response(errorCode string) *GetTransaction400Response {
	this := GetTransaction400Response{}
	this.ErrorCode = errorCode
	return &this
}

// NewGetTransaction400ResponseWithDefaults instantiates a new GetTransaction400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransaction400ResponseWithDefaults() *GetTransaction400Response {
	this := GetTransaction400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *GetTransaction400Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *GetTransaction400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *GetTransaction400Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GetTransaction400Response) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransaction400Response) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GetTransaction400Response) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *GetTransaction400Response) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o GetTransaction400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransaction400Response struct {
	value *GetTransaction400Response
	isSet bool
}

func (v NullableGetTransaction400Response) Get() *GetTransaction400Response {
	return v.value
}

func (v *NullableGetTransaction400Response) Set(val *GetTransaction400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransaction400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransaction400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransaction400Response(val *GetTransaction400Response) *NullableGetTransaction400Response {
	return &NullableGetTransaction400Response{value: val, isSet: true}
}

func (v NullableGetTransaction400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransaction400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


