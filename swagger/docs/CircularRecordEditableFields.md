# CircularRecordEditableFields

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The hostname of the record. Validation is record type specific. | [default to null]
**Comment** | **string** | Freeform comment field. Markdown is recommended and may be rendered in certain situations in the UI. | [optional] [default to null]
**LockPassword** | **string** | If the hostname is locked, an unlock is attempted with this password | [optional] [default to null]
**Ttl** | **int32** | TTL of the record | [optional] [default to null]
**Ip** | **string** | IP address of the record. Can be any IP address within the subnet you wish to add the reservation to. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

