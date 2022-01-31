/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// AccountCursor struct for AccountCursor
type AccountCursor struct {
	HasMore bool `json:"has_more"`
	Next string `json:"next"`
	PageSize int32 `json:"page_size"`
	Previous string `json:"previous"`
	RemainingResults int32 `json:"remaining_results"`
	Total int32 `json:"total"`
	Data []Account `json:"data"`
}

// NewAccountCursor instantiates a new AccountCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCursor(hasMore bool, next string, pageSize int32, previous string, remainingResults int32, total int32, data []Account) *AccountCursor {
	this := AccountCursor{}
	this.HasMore = hasMore
	this.Next = next
	this.PageSize = pageSize
	this.Previous = previous
	this.RemainingResults = remainingResults
	this.Total = total
	this.Data = data
	return &this
}

// NewAccountCursorWithDefaults instantiates a new AccountCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCursorWithDefaults() *AccountCursor {
	this := AccountCursor{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *AccountCursor) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *AccountCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *AccountCursor) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNext returns the Next field value
func (o *AccountCursor) GetNext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *AccountCursor) GetNextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *AccountCursor) SetNext(v string) {
	o.Next = v
}

// GetPageSize returns the PageSize field value
func (o *AccountCursor) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *AccountCursor) GetPageSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *AccountCursor) SetPageSize(v int32) {
	o.PageSize = v
}

// GetPrevious returns the Previous field value
func (o *AccountCursor) GetPrevious() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
func (o *AccountCursor) GetPreviousOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Previous, true
}

// SetPrevious sets field value
func (o *AccountCursor) SetPrevious(v string) {
	o.Previous = v
}

// GetRemainingResults returns the RemainingResults field value
func (o *AccountCursor) GetRemainingResults() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemainingResults
}

// GetRemainingResultsOk returns a tuple with the RemainingResults field value
// and a boolean to check if the value has been set.
func (o *AccountCursor) GetRemainingResultsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RemainingResults, true
}

// SetRemainingResults sets field value
func (o *AccountCursor) SetRemainingResults(v int32) {
	o.RemainingResults = v
}

// GetTotal returns the Total field value
func (o *AccountCursor) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AccountCursor) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AccountCursor) SetTotal(v int32) {
	o.Total = v
}

// GetData returns the Data field value
func (o *AccountCursor) GetData() []Account {
	if o == nil {
		var ret []Account
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccountCursor) GetDataOk() (*[]Account, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AccountCursor) SetData(v []Account) {
	o.Data = v
}

func (o AccountCursor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["has_more"] = o.HasMore
	}
	if true {
		toSerialize["next"] = o.Next
	}
	if true {
		toSerialize["page_size"] = o.PageSize
	}
	if true {
		toSerialize["previous"] = o.Previous
	}
	if true {
		toSerialize["remaining_results"] = o.RemainingResults
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCursor struct {
	value *AccountCursor
	isSet bool
}

func (v NullableAccountCursor) Get() *AccountCursor {
	return v.value
}

func (v *NullableAccountCursor) Set(val *AccountCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCursor(val *AccountCursor) *NullableAccountCursor {
	return &NullableAccountCursor{value: val, isSet: true}
}

func (v NullableAccountCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


