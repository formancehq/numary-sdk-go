/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// BalancesCursorResponseCursor struct for BalancesCursorResponseCursor
type BalancesCursorResponseCursor struct {
	PageSize int64 `json:"pageSize"`
	HasMore bool `json:"hasMore"`
	Previous *string `json:"previous,omitempty"`
	Next *string `json:"next,omitempty"`
	Data []map[string]map[string]int64 `json:"data"`
}

// NewBalancesCursorResponseCursor instantiates a new BalancesCursorResponseCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalancesCursorResponseCursor(pageSize int64, hasMore bool, data []map[string]map[string]int64) *BalancesCursorResponseCursor {
	this := BalancesCursorResponseCursor{}
	this.PageSize = pageSize
	this.HasMore = hasMore
	this.Data = data
	return &this
}

// NewBalancesCursorResponseCursorWithDefaults instantiates a new BalancesCursorResponseCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalancesCursorResponseCursorWithDefaults() *BalancesCursorResponseCursor {
	this := BalancesCursorResponseCursor{}
	return &this
}

// GetPageSize returns the PageSize field value
func (o *BalancesCursorResponseCursor) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *BalancesCursorResponseCursor) GetPageSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *BalancesCursorResponseCursor) SetPageSize(v int64) {
	o.PageSize = v
}

// GetHasMore returns the HasMore field value
func (o *BalancesCursorResponseCursor) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *BalancesCursorResponseCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *BalancesCursorResponseCursor) SetHasMore(v bool) {
	o.HasMore = v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *BalancesCursorResponseCursor) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancesCursorResponseCursor) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *BalancesCursorResponseCursor) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *BalancesCursorResponseCursor) SetPrevious(v string) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *BalancesCursorResponseCursor) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancesCursorResponseCursor) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *BalancesCursorResponseCursor) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *BalancesCursorResponseCursor) SetNext(v string) {
	o.Next = &v
}

// GetData returns the Data field value
func (o *BalancesCursorResponseCursor) GetData() []map[string]map[string]int64 {
	if o == nil {
		var ret []map[string]map[string]int64
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BalancesCursorResponseCursor) GetDataOk() (*[]map[string]map[string]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BalancesCursorResponseCursor) SetData(v []map[string]map[string]int64) {
	o.Data = v
}

func (o BalancesCursorResponseCursor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pageSize"] = o.PageSize
	}
	if true {
		toSerialize["hasMore"] = o.HasMore
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBalancesCursorResponseCursor struct {
	value *BalancesCursorResponseCursor
	isSet bool
}

func (v NullableBalancesCursorResponseCursor) Get() *BalancesCursorResponseCursor {
	return v.value
}

func (v *NullableBalancesCursorResponseCursor) Set(val *BalancesCursorResponseCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancesCursorResponseCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancesCursorResponseCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancesCursorResponseCursor(val *BalancesCursorResponseCursor) *NullableBalancesCursorResponseCursor {
	return &NullableBalancesCursorResponseCursor{value: val, isSet: true}
}

func (v NullableBalancesCursorResponseCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancesCursorResponseCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


