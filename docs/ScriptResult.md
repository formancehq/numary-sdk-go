# ScriptResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **string** |  | [optional] 
**Err** | Pointer to **string** |  | [optional] 

## Methods

### NewScriptResult

`func NewScriptResult() *ScriptResult`

NewScriptResult instantiates a new ScriptResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptResultWithDefaults

`func NewScriptResultWithDefaults() *ScriptResult`

NewScriptResultWithDefaults instantiates a new ScriptResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ScriptResult) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ScriptResult) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ScriptResult) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ScriptResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetErr

`func (o *ScriptResult) GetErr() string`

GetErr returns the Err field if non-nil, zero value otherwise.

### GetErrOk

`func (o *ScriptResult) GetErrOk() (*string, bool)`

GetErrOk returns a tuple with the Err field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErr

`func (o *ScriptResult) SetErr(v string)`

SetErr sets Err field to given value.

### HasErr

`func (o *ScriptResult) HasErr() bool`

HasErr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


