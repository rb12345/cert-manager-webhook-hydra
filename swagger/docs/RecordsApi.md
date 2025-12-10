# {{classname}}

All URIs are relative to *https://www.networks.it.ox.ac.uk/university/ipam/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecord**](RecordsApi.md#CreateRecord) | **Post** /records | 
[**ListRecords**](RecordsApi.md#ListRecords) | **Get** /records | 
[**RecordsBulkPost**](RecordsApi.md#RecordsBulkPost) | **Post** /records/bulk | 
[**RecordsDelete**](RecordsApi.md#RecordsDelete) | **Delete** /records | 
[**RecordsIdDelete**](RecordsApi.md#RecordsIdDelete) | **Delete** /records/{id} | 
[**RecordsIdGet**](RecordsApi.md#RecordsIdGet) | **Get** /records/{id} | 
[**RecordsIdPut**](RecordsApi.md#RecordsIdPut) | **Put** /records/{id} | 

# **CreateRecord**
> Record CreateRecord(ctx, hostname, type_, content, comment, ttl, prioEtAl)


Creates a new record. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostname** | **string**|  | 
  **type_** | **string**|  | 
  **content** | **string**|  | 
  **comment** | **string**|  | 
  **ttl** | **int32**|  | 
  **prioEtAl** | **string**|  | 

### Return type

[**Record**](Record.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRecords**
> []Record ListRecords(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecordsApiListRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordsApiListRecordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the records returned. If this is not supplied, defaults to 5000 records. There is a hard-coded limit of 500,001 above which the number reverts to 5000.  | 
 **page** | **optional.Int32**| Start the list returned at the page provided. | 
 **order** | **optional.String**| Comma separated string of columns that you wish to use for sorting. Default is big_endian_labels. Invalid columns are ignored. To reverse the search, prefix with a &#x27;-&#x27;  | 
 **q** | **optional.String**| Constrain the records returned to match the provided search parameter. Format is based off [Search::QueryParser](https://metacpan.org/pod/Search::QueryParser). Multiple instances of this parameter may be supplied and they will be logically anded. [Documentation is available](https://networks.it.ox.ac.uk/itss/ipam/help/searching)  | 
 **show** | **optional.String**| By default, the results returned are only records that you have permission to edit. To show more than this subset, give &#x27;all&#x27; here. Alternatively, to see deleted records, supply &#x27;deleted&#x27;  | 

### Return type

[**[]Record**](Record.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordsBulkPost**
> []Record RecordsBulkPost(ctx, body)


Bulk update records in one transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]Object**](.md)| Array of records that need to be created/updated. Items with an ID are treated as an update. Records without are treated as an insert. Any fields that are not valid for editing are stripped and do not generate an error, fields such as created_by and created_in_transaction. To delete a record, add a key action with value \&quot;delete\&quot;. In this case only the \&quot;id\&quot; field needs to be present in that object.  | 

### Return type

[**[]Record**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordsDelete**
> []Record RecordsDelete(ctx, q, optional)


Delete records based on their values in the database 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**|  | 
 **optional** | ***RecordsApiRecordsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordsApiRecordsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of records to delete. Must be smaller than 5. | 

### Return type

[**[]Record**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordsIdDelete**
> RecordsIdDelete(ctx, id, optional)


Delete (or archive) the record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the record being requested | 
 **optional** | ***RecordsApiRecordsIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordsApiRecordsIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockPassword** | [**optional.Interface of string**](.md)| Unlock a hostname for the period of a transaction | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordsIdGet**
> Record RecordsIdGet(ctx, id)


Get the record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the record being requested | 

### Return type

[**Record**](Record.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordsIdPut**
> Record RecordsIdPut(ctx, id, optional)


Update the record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the record being requested | 
 **optional** | ***RecordsApiRecordsIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordsApiRecordsIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockPassword** | [**optional.Interface of string**](.md)| Unlock a hostname for the period of a transaction | 

### Return type

[**Record**](Record.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

