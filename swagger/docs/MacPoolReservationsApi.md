# {{classname}}

All URIs are relative to *https://www.networks.it.ox.ac.uk/university/ipam/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMacPoolReservations**](MacPoolReservationsApi.md#ListMacPoolReservations) | **Get** /mac-reservations/pools | 
[**MacReservationsPoolBulkPost**](MacPoolReservationsApi.md#MacReservationsPoolBulkPost) | **Post** /mac-reservations/pool/bulk | 
[**MacReservationsPoolsPost**](MacPoolReservationsApi.md#MacReservationsPoolsPost) | **Post** /mac-reservations/pools | 

# **ListMacPoolReservations**
> []MacPoolReservation ListMacPoolReservations(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MacPoolReservationsApiListMacPoolReservationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MacPoolReservationsApiListMacPoolReservationsOpts struct
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

