/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// checks if the AccountsCursorResponseCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountsCursorResponseCursor{}

// AccountsCursorResponseCursor struct for AccountsCursorResponseCursor
type AccountsCursorResponseCursor struct {
	PageSize int64 `json:"pageSize"`
	HasMore bool `json:"hasMore"`
	Previous *string `json:"previous,omitempty"`
	Next *string `json:"next,omitempty"`
	Data []Account `json:"data"`
}

// NewAccountsCursorResponseCursor instantiates a new AccountsCursorResponseCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsCursorResponseCursor(pageSize int64, hasMore bool, data []Account) *AccountsCursorResponseCursor {
	this := AccountsCursorResponseCursor{}
	this.PageSize = pageSize
	this.HasMore = hasMore
	this.Data = data
	return &this
}

// NewAccountsCursorResponseCursorWithDefaults instantiates a new AccountsCursorResponseCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsCursorResponseCursorWithDefaults() *AccountsCursorResponseCursor {
	this := AccountsCursorResponseCursor{}
	return &this
}

// GetPageSize returns the PageSize field value
func (o *AccountsCursorResponseCursor) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *AccountsCursorResponseCursor) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *AccountsCursorResponseCursor) SetPageSize(v int64) {
	o.PageSize = v
}

// GetHasMore returns the HasMore field value
func (o *AccountsCursorResponseCursor) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *AccountsCursorResponseCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *AccountsCursorResponseCursor) SetHasMore(v bool) {
	o.HasMore = v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *AccountsCursorResponseCursor) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountsCursorResponseCursor) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *AccountsCursorResponseCursor) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *AccountsCursorResponseCursor) SetPrevious(v string) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *AccountsCursorResponseCursor) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountsCursorResponseCursor) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *AccountsCursorResponseCursor) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *AccountsCursorResponseCursor) SetNext(v string) {
	o.Next = &v
}

// GetData returns the Data field value
func (o *AccountsCursorResponseCursor) GetData() []Account {
	if o == nil {
		var ret []Account
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccountsCursorResponseCursor) GetDataOk() ([]Account, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AccountsCursorResponseCursor) SetData(v []Account) {
	o.Data = v
}

func (o AccountsCursorResponseCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountsCursorResponseCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pageSize"] = o.PageSize
	toSerialize["hasMore"] = o.HasMore
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAccountsCursorResponseCursor struct {
	value *AccountsCursorResponseCursor
	isSet bool
}

func (v NullableAccountsCursorResponseCursor) Get() *AccountsCursorResponseCursor {
	return v.value
}

func (v *NullableAccountsCursorResponseCursor) Set(val *AccountsCursorResponseCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsCursorResponseCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsCursorResponseCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsCursorResponseCursor(val *AccountsCursorResponseCursor) *NullableAccountsCursorResponseCursor {
	return &NullableAccountsCursorResponseCursor{value: val, isSet: true}
}

func (v NullableAccountsCursorResponseCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsCursorResponseCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


