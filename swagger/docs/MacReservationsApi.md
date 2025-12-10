# {{classname}}

All URIs are relative to *https://www.networks.it.ox.ac.uk/university/ipam/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMacReservation**](MacReservationsApi.md#CreateMacReservation) | **Post** /mac-reservations | 
[**ListMacPoolReservations**](MacReservationsApi.md#ListMacPoolReservations) | **Get** /mac-reservations/pools | 
[**ListMacReservations**](MacReservationsApi.md#ListMacReservations) | **Get** /mac-reservations | 
[**MacReservationsBulkPost**](MacReservationsApi.md#MacReservationsBulkPost) | **Post** /mac-reservations/bulk | 
[**MacReservationsIdDelete**](MacReservationsApi.md#MacReservationsIdDelete) | **Delete** /mac-reservations/{id} | 
[**MacReservationsIdGet**](MacReservationsApi.md#MacReservationsIdGet) | **Get** /mac-reservations/{id} | 
[**MacReservationsIdPut**](MacReservationsApi.md#MacReservationsIdPut) | **Put** /mac-reservations/{id} | 
[**MacReservationsPoolBulkPost**](MacReservationsApi.md#MacReservationsPoolBulkPost) | **Post** /mac-reservations/pool/bulk | 
[**MacReservationsPoolsPost**](MacReservationsApi.md#MacReservationsPoolsPost) | **Post** /mac-reservations/pools | 

# **CreateMacReservation**
> MacReservation CreateMacReservation(ctx, )


Creates a new mac reservation 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MacReservation**](MacReservation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMacPoolReservations**
> []MacPoolReservation ListMacPoolReservations(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MacReservationsApiListMacPoolReservationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MacReservationsApiListMacPoolReservationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the records returned. If this is not supplied, defaults to 5000 records. There is a hard-coded limit of 500,001 above which the number reverts to 5000.  | 
 **page** | **optional.Int32**| Start the list returned at the page provided. | 
 **order** | **optional.String**| Comma separated string of columns that you wish to use for sorting. Default is big_endian_labels. Invalid columns are ignored. To reverse the search, prefix with a &#x27;-&#x27;  | 
 **q** | **optional.String**| Constrain the records returned to match the provided search parameter. Format is based off [Search::QueryParser](https://metacpan.org/pod/Search::QueryParser). Multiple instances of this parameter may be supplied and they will be logically anded. [Documentation is available](https://networks.it.ox.ac.uk/itss/ipam/help/searching)  | 

### Return type

[**[]MacPoolReservation**](MacPoolReservation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMacReservations**
> []MacReservation ListMacReservations(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MacReservationsApiListMacReservationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MacReservationsApiListMacReservationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the records returned. If this is not supplied, defaults to 5000 records. There is a hard-coded limit of 500,001 above which the number reverts to 5000.  | 
 **page** | **optional.Int32**| Start the list returned at the page provided. | 
 **order** | **optional.String**| Comma separated string of columns that you wish to use for sorting. Default is big_endian_labels. Invalid columns are ignored. To reverse the search, prefix with a &#x27;-&#x27;  | 
 **q** | **optional.String**| Constrain the records returned to match the provided search parameter. Format is based off [Search::QueryParser](https://metacpan.org/pod/Search::QueryParser). Multiple instances of this parameter may be supplied and they will be logically anded. [Documentation is available](https://networks.it.ox.ac.uk/itss/ipam/help/searching)  | 

### Return type

[**[]MacReservation**](MacReservation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MacReservationsBulkPost**
> []MacReservation MacReservationsBulkPost(ctx, body)


Bulk update records in one transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MacreservationsBulkBody**](MacreservationsBulkBody.md)| Array of mac reservations that need to be created/updated. Items with an ID are treated as an update. Entries without an ID are treated as an insert. Any fields that are not valid for editing are stripped and do not generate an error, fields such as created_by and created_in_transaction. To delete a record, add a key &#x60;action&#x60; with value \&quot;delete\&quot;. In this case only the \&quot;id\&quot; field needs to be present in that object.  | 

### Return type

[**[]MacReservation**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MacReservationsIdDelete**
> MacReservationsIdDelete(ctx, id)


Delete (or archive) the mac reservation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the reservation being requested | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MacReservationsIdGet**
> MacReservation MacReservationsIdGet(ctx, id)


Get the reservation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the reservation being requested | 

### Return type

[**MacReservation**](MacReservation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MacReservationsIdPut**
> MacReservation MacReservationsIdPut(ctx, id)


Update the mac reservation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the reservation being requested | 

### Return type

[**MacReservation**](MacReservation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MacReservationsPoolBulkPost**
> []MacPoolReservation MacReservationsPoolBulkPost(ctx, body)


Bulk update pool reservations in one transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)| This does not behave the same as other resources. The table is a virtual table and as such there are differences to the underlying data structure:    1. There is no ID field   2. Upserting is preferred over erroring.  If a (hw_address,subnet) pair already exists, a record creation will be skipped, *even if the record has a different comment*.  Input for records can also be CSV, as described in Hydra&#x27;s UI.  | 

### Return type

[**[]MacPoolReservation**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MacReservationsPoolsPost**
> MacPoolReservation MacReservationsPoolsPost(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MacPoolReservation**](MacPoolReservation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

