/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.0-rc.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
	"time"
)

// Transaction struct for Transaction
type Transaction struct {
	Timestamp time.Time `json:"timestamp"`
	Postings []Posting `json:"postings"`
	Reference *string `json:"reference,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Txid int64 `json:"txid"`
	PreCommitVolumes *map[string]map[string]Volume `json:"preCommitVolumes,omitempty"`
	PostCommitVolumes *map[string]map[string]Volume `json:"postCommitVolumes,omitempty"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(timestamp time.Time, postings []Posting, txid int64) *Transaction {
	this := Transaction{}
	this.Timestamp = timestamp
	this.Postings = postings
	this.Txid = txid
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *Transaction) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Transaction) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetPostings returns the Postings field value
func (o *Transaction) GetPostings() []Posting {
	if o == nil {
		var ret []Posting
		return ret
	}

	return o.Postings
}

// GetPostingsOk returns a tuple with the Postings field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetPostingsOk() (*[]Posting, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Postings, true
}

// SetPostings sets field value
func (o *Transaction) SetPostings(v []Posting) {
	o.Postings = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Transaction) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Transaction) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Transaction) SetReference(v string) {
	o.Reference = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetMetadata() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Transaction) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Transaction) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetTxid returns the Txid field value
func (o *Transaction) GetTxid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTxidOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *Transaction) SetTxid(v int64) {
	o.Txid = v
}

// GetPreCommitVolumes returns the PreCommitVolumes field value if set, zero value otherwise.
func (o *Transaction) GetPreCommitVolumes() map[string]map[string]Volume {
	if o == nil || o.PreCommitVolumes == nil {
		var ret map[string]map[string]Volume
		return ret
	}
	return *o.PreCommitVolumes
}

// GetPreCommitVolumesOk returns a tuple with the PreCommitVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetPreCommitVolumesOk() (*map[string]map[string]Volume, bool) {
	if o == nil || o.PreCommitVolumes == nil {
		return nil, false
	}
	return o.PreCommitVolumes, true
}

// HasPreCommitVolumes returns a boolean if a field has been set.
func (o *Transaction) HasPreCommitVolumes() bool {
	if o != nil && o.PreCommitVolumes != nil {
		return true
	}

	return false
}

// SetPreCommitVolumes gets a reference to the given map[string]map[string]Volume and assigns it to the PreCommitVolumes field.
func (o *Transaction) SetPreCommitVolumes(v map[string]map[string]Volume) {
	o.PreCommitVolumes = &v
}

// GetPostCommitVolumes returns the PostCommitVolumes field value if set, zero value otherwise.
func (o *Transaction) GetPostCommitVolumes() map[string]map[string]Volume {
	if o == nil || o.PostCommitVolumes == nil {
		var ret map[string]map[string]Volume
		return ret
	}
	return *o.PostCommitVolumes
}

// GetPostCommitVolumesOk returns a tuple with the PostCommitVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetPostCommitVolumesOk() (*map[string]map[string]Volume, bool) {
	if o == nil || o.PostCommitVolumes == nil {
		return nil, false
	}
	return o.PostCommitVolumes, true
}

// HasPostCommitVolumes returns a boolean if a field has been set.
func (o *Transaction) HasPostCommitVolumes() bool {
	if o != nil && o.PostCommitVolumes != nil {
		return true
	}

	return false
}

// SetPostCommitVolumes gets a reference to the given map[string]map[string]Volume and assigns it to the PostCommitVolumes field.
func (o *Transaction) SetPostCommitVolumes(v map[string]map[string]Volume) {
	o.PostCommitVolumes = &v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["postings"] = o.Postings
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["txid"] = o.Txid
	}
	if o.PreCommitVolumes != nil {
		toSerialize["preCommitVolumes"] = o.PreCommitVolumes
	}
	if o.PostCommitVolumes != nil {
		toSerialize["postCommitVolumes"] = o.PostCommitVolumes
	}
	return json.Marshal(toSerialize)
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


