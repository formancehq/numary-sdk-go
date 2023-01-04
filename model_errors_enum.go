/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.0-rc.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
	"fmt"
)

// ErrorsEnum the model 'ErrorsEnum'
type ErrorsEnum string

// List of ErrorsEnum
const (
	INTERNAL ErrorsEnum = "INTERNAL"
	INSUFFICIENT_FUND ErrorsEnum = "INSUFFICIENT_FUND"
	VALIDATION ErrorsEnum = "VALIDATION"
	CONFLICT ErrorsEnum = "CONFLICT"
	NO_SCRIPT ErrorsEnum = "NO_SCRIPT"
	COMPILATION_FAILED ErrorsEnum = "COMPILATION_FAILED"
	METADATA_OVERRIDE ErrorsEnum = "METADATA_OVERRIDE"
)

// All allowed values of ErrorsEnum enum
var AllowedErrorsEnumEnumValues = []ErrorsEnum{
	"INTERNAL",
	"INSUFFICIENT_FUND",
	"VALIDATION",
	"CONFLICT",
	"NO_SCRIPT",
	"COMPILATION_FAILED",
	"METADATA_OVERRIDE",
}

func (v *ErrorsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorsEnum(value)
	for _, existing := range AllowedErrorsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorsEnum", value)
}

// NewErrorsEnumFromValue returns a pointer to a valid ErrorsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorsEnumFromValue(v string) (*ErrorsEnum, error) {
	ev := ErrorsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorsEnum: valid values are %v", v, AllowedErrorsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorsEnum) IsValid() bool {
	for _, existing := range AllowedErrorsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorsEnum value
func (v ErrorsEnum) Ptr() *ErrorsEnum {
	return &v
}

type NullableErrorsEnum struct {
	value *ErrorsEnum
	isSet bool
}

func (v NullableErrorsEnum) Get() *ErrorsEnum {
	return v.value
}

func (v *NullableErrorsEnum) Set(val *ErrorsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsEnum(val *ErrorsEnum) *NullableErrorsEnum {
	return &NullableErrorsEnum{value: val, isSet: true}
}

func (v NullableErrorsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

