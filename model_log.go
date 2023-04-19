/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.0-beta.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
	"time"
)

// Log struct for Log
type Log struct {
	Id int64 `json:"id"`
	Type string `json:"type"`
	Data map[string]interface{} `json:"data"`
	Hash string `json:"hash"`
	Date time.Time `json:"date"`
}

// NewLog instantiates a new Log object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLog(id int64, type_ string, data map[string]interface{}, hash string, date time.Time) *Log {
	this := Log{}
	this.Id = id
	this.Type = type_
	this.Data = data
	this.Hash = hash
	this.Date = date
	return &this
}

// NewLogWithDefaults instantiates a new Log object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogWithDefaults() *Log {
	this := Log{}
	return &this
}

// GetId returns the Id field value
func (o *Log) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Log) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Log) SetId(v int64) {
	o.Id = v
}

// GetType returns the Type field value
func (o *Log) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Log) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Log) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *Log) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Log) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Log) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetHash returns the Hash field value
func (o *Log) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *Log) GetHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *Log) SetHash(v string) {
	o.Hash = v
}

// GetDate returns the Date field value
func (o *Log) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Log) GetDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Log) SetDate(v time.Time) {
	o.Date = v
}

func (o Log) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableLog struct {
	value *Log
	isSet bool
}

func (v NullableLog) Get() *Log {
	return v.value
}

func (v *NullableLog) Set(val *Log) {
	v.value = val
	v.isSet = true
}

func (v NullableLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLog(val *Log) *NullableLog {
	return &NullableLog{value: val, isSet: true}
}

func (v NullableLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


