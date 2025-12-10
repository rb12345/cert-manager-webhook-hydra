# RecordEditableFields

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | [default to null]
**Type_** | **string** | The record type (e.g. A, AAAA, SRV) | [default to null]
**Content** | **string** | What the value is for the record | [default to null]
**Comment** | **string** | Freeform field | [optional] 
**Ttl** | **int32** | TTL of the record | [optional] [default to null]
**PrioEtAl** | **string** | Prio and related fields. See [documentation](https://networks.it.ox.ac.uk/university/ipam/help/record-fields#prio-et-al)  | [optional] [default to null]
**LockPassword** | **string** | If the hostname is locked, an unlock is attempted with this password | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

