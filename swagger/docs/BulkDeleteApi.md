# {{classname}}

All URIs are relative to *https://www.networks.it.ox.ac.uk/university/ipam/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteCircularRecords**](BulkDeleteApi.md#BulkDeleteCircularRecords) | **Delete** /circular-records | 
[**RecordsDelete**](BulkDeleteApi.md#RecordsDelete) | **Delete** /records | 

# **BulkDeleteCircularRecords**
> []Record BulkDeleteCircularRecords(ctx, q, optional)


Delete circular records based on their values in the database 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**|  | 
 **optional** | ***BulkDeleteApiBulkDeleteCircularRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkDeleteApiBulkDeleteCircularRecordsOpts struct
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

# **RecordsDelete**
> []Record RecordsDelete(ctx, q, optional)


Delete records based on their values in the database 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**|  | 
 **optional** | ***BulkDeleteApiRecordsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkDeleteApiRecordsDeleteOpts struct
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

