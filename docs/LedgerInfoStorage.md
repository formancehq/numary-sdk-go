# LedgerInfoStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Driver** | Pointer to **string** |  | [optional] 
**Migrations** | Pointer to [**[]MigrationInfo**](MigrationInfo.md) |  | [optional] 

## Methods

### NewLedgerInfoStorage

`func NewLedgerInfoStorage() *LedgerInfoStorage`

NewLedgerInfoStorage instantiates a new LedgerInfoStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerInfoStorageWithDefaults

`func NewLedgerInfoStorageWithDefaults() *LedgerInfoStorage`

NewLedgerInfoStorageWithDefaults instantiates a new LedgerInfoStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriver

`func (o *LedgerInfoStorage) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *LedgerInfoStorage) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *LedgerInfoStorage) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *LedgerInfoStorage) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetMigrations

`func (o *LedgerInfoStorage) GetMigrations() []MigrationInfo`

GetMigrations returns the Migrations field if non-nil, zero value otherwise.

### GetMigrationsOk

`func (o *LedgerInfoStorage) GetMigrationsOk() (*[]MigrationInfo, bool)`

GetMigrationsOk returns a tuple with the Migrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrations

`func (o *LedgerInfoStorage) SetMigrations(v []MigrationInfo)`

SetMigrations sets Migrations field to given value.

### HasMigrations

`func (o *LedgerInfoStorage) HasMigrations() bool`

HasMigrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


