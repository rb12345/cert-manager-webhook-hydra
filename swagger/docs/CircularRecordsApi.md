# {{classname}}

All URIs are relative to *https://www.networks.it.ox.ac.uk/university/ipam/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteCircularRecords**](CircularRecordsApi.md#BulkDeleteCircularRecords) | **Delete** /circular-records | 
[**CircularRecordsBulkPost**](CircularRecordsApi.md#CircularRecordsBulkPost) | **Post** /circular-records/bulk | 
[**CircularRecordsIdDelete**](CircularRecordsApi.md#CircularRecordsIdDelete) | **Delete** /circular-records/{id} | 
[**CircularRecordsIdGet**](CircularRecordsApi.md#CircularRecordsIdGet) | **Get** /circular-records/{id} | 
[**CircularRecordsIdPut**](CircularRecordsApi.md#CircularRecordsIdPut) | **Put** /circular-records/{id} | 
[**CircularRecordsThingamyDelete**](CircularRecordsApi.md#CircularRecordsThingamyDelete) | **Delete** /circular-records/{thingamy} | 
[**CircularRecordsThingamyGet**](CircularRecordsApi.md#CircularRecordsThingamyGet) | **Get** /circular-records/{thingamy} | 
[**CircularRecordsThingamyPut**](CircularRecordsApi.md#CircularRecordsThingamyPut) | **Put** /circular-records/{thingamy} | 
[**CreateCircularRecord**](CircularRecordsApi.md#CreateCircularRecord) | **Post** /circular-records | 
[**ListCircularRecords**](CircularRecordsApi.md#ListCircularRecords) | **Get** /circular-records | 

# **BulkDeleteCircularRecords**
> []Record BulkDeleteCircularRecords(ctx, q, optional)


Delete circular records based on their values in the database 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**|  | 
 **optional** | ***CircularRecordsApiBulkDeleteCircularRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CircularRecordsApiBulkDeleteCircularRecordsOpts struct
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

# **CircularRecordsBulkPost**
> []Record CircularRecordsBulkPost(ctx, body)


Bulk update records in one transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CircularrecordsBulkBody**](CircularrecordsBulkBody.md)| Array of circular records, under the records key, that need to be created/updated. Items with an ID are treated as an update. Records without are treated as an insert. Any fields that are not valid for editing are stripped and do not generate an error, fields such as created_by and created_in_transaction. To delete a record, add a key action with value \&quot;delete\&quot;. In this case only the \&quot;id\&quot; field needs to be present in that object.  | 

### Return type

[**[]Record**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircularRecordsIdDelete**
> CircularRecordsIdDelete(ctx, id)


Delete (or archive) the circular record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the record being requested | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircularRecordsIdGet**
> CircularRecord CircularRecordsIdGet(ctx, id)


Get the record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the record being requested | 

### Return type

[**CircularRecord**](CircularRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircularRecordsIdPut**
> CircularRecord CircularRecordsIdPut(ctx, id)


Update the circular record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the record being requested | 

### Return type

[**CircularRecord**](CircularRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircularRecordsThingamyDelete**
> CircularRecord CircularRecordsThingamyDelete(ctx, thingamy)


Delete circular record by the unique key of IP (hostname is not unique so that will give a 404 error). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **thingamy** | **string**| Something that can uniquely define a circular-record, like an IP or a hostname. Hostname can return two records, but for updates it will be whittled down to one based on the ip address supplied in the request. For GET requests you cannot use hostname.  | 

### Return type

[**CircularRecord**](CircularRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircularRecordsThingamyGet**
> CircularRecordsThingamyGet(ctx, thingamy)


Get a circular record based off its IP. Note that because hostname is not uniquely identifying, thingamy cannot be a hostname in this context and any attempts will result in 404 errors 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **thingamy** | **string**| Something that can uniquely define a circular-record, like an IP or a hostname. Hostname can return two records, but for updates it will be whittled down to one based on the ip address supplied in the request. For GET requests you cannot use hostname.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircularRecordsThingamyPut**
> CircularRecord CircularRecordsThingamyPut(ctx, thingamy)


Insert or update the circular record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **thingamy** | **string**| Something that can uniquely define a circular-record, like an IP or a hostname. Hostname can return two records, but for updates it will be whittled down to one based on the ip address supplied in the request. For GET requests you cannot use hostname.  | 

### Return type

[**CircularRecord**](CircularRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCircularRecord**
> CircularRecord CreateCircularRecord(ctx, )


Creates a new circular record, which is represented in BIND as an {A,AAAA}+PTR pair. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CircularRecord**](CircularRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCircularRecords**
> []CircularRecord ListCircularRecords(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CircularRecordsApiListCircularRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CircularRecordsApiListCircularRecordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the records returned. If this is not supplied, defaults to 5000 records. There is a hard-coded limit of 500,001 above which the number reverts to 5000.  | 
 **page** | **optional.Int32**| Start the list returned at the page provided. | 
 **order** | **optional.String**| Comma separated string of columns that you wish to use for sorting. Default is big_endian_labels. Invalid columns are ignored. To reverse the search, prefix with a &#x27;-&#x27;  | 
 **q** | **optional.String**| Constrain the records returned to match the provided search parameter. Format is based off [Search::QueryParser](https://metacpan.org/pod/Search::QueryParser). Multiple instances of this parameter may be supplied and they will be logically anded. [Documentation is available](https://networks.it.ox.ac.uk/itss/ipam/help/searching)  | 
 **show** | **optional.String**| By default, the results returned are only records that you have permission to edit. To show more than this subset, give &#x27;all&#x27; here. Alternatively, to see deleted records, supply &#x27;deleted&#x27;  | 

### Return type

[**[]CircularRecord**](CircularRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

