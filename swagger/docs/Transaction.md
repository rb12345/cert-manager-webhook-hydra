# Transaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Txid** | **int32** |  | [optional] [default to null]
**Username** | **string** | The user who started the transaction | [optional] [default to null]
**UsernameSession** | **string** | The session user who started the transaction. Will be the same as username unless masquerading is occurring. Masquerading is not common.  | [optional] [default to null]
**RecordsCreated** | **int32** | Records created in this transaction. An update is a delete and a create | [optional] [default to null]
**RecordsDeleted** | **int32** | Records deleted in this transaction. An update is a delete and a create | [optional] [default to null]
**At** | [**time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

