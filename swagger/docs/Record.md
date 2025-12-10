# Record

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**CreatedInTransaction** | **int32** | The transaction ID in which the record was created | [optional] [default to null]
**DeletedInTransaction** | **int32** | The transaction ID in which the record was deleted | [optional] [default to null]
**Comment** | **string** | Freeform comment field. Markdown is recommended and may be rendered in certain situations in the UI. | [optional] [default to null]
**Hostname** | **string** |  | [default to null]
**Type_** | **string** | The record type (e.g. A, AAAA, SRV) | [default to null]
**Content** | **string** | What the value is for the record | [default to null]
**Ttl** | **int32** | TTL of the record | [optional] [default to null]
**PrioEtAl** | **string** | Prio and related fields. See [documentation](https://networks.it.ox.ac.uk/university/ipam/help/record-fields#prio-et-al)  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

