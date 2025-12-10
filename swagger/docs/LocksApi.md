# {{classname}}

All URIs are relative to *https://www.networks.it.ox.ac.uk/university/ipam/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLock**](LocksApi.md#CreateLock) | **Post** /locks | 
[**ListLocks**](LocksApi.md#ListLocks) | **Get** /locks | 
[**LocksIdDelete**](LocksApi.md#LocksIdDelete) | **Delete** /locks/{id} | 
[**LocksIdGet**](LocksApi.md#LocksIdGet) | **Get** /locks/{id} | 

# **CreateLock**
> Lock CreateLock(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LockEditableFields**](LockEditableFields.md)| Lock to create | 

### Return type

[**Lock**](Lock.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application:x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLocks**
> []Lock ListLocks(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocksApiListLocksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocksApiListLocksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the records returned. If this is not supplied, defaults to 5000 records. There is a hard-coded limit of 500,001 above which the number reverts to 5000.  | 
 **page** | **optional.Int32**| Start the list returned at the page provided. | 
 **order** | **optional.String**| Comma separated string of columns that you wish to use for sorting. Default is big_endian_labels. Invalid columns are ignored. To reverse the search, prefix with a &#x27;-&#x27;  | 
 **q** | **optional.String**| Constrain the records returned to match the provided search parameter. Format is based off [Search::QueryParser](https://metacpan.org/pod/Search::QueryParser). Multiple instances of this parameter may be supplied and they will be logically anded. [Documentation is available](https://networks.it.ox.ac.uk/itss/ipam/help/searching)  | 
 **show** | **optional.String**| By default, the results returned are only records that you have permission to edit. To show more than this subset, give &#x27;all&#x27; here. Alternatively, to see deleted records, supply &#x27;deleted&#x27;  | 

### Return type

[**[]Lock**](Lock.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocksIdDelete**
> LocksIdDelete(ctx, id, optional)


Delete the lock. Unless the lock has been unlocked globally for a time period, lock_password is required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the lock being requested | 
 **optional** | ***LocksApiLocksIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocksApiLocksIdDeleteOpts struct
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

# **LocksIdGet**
> Lock LocksIdGet(ctx, id)


Get the lock

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The UUID for the lock being requested | 

### Return type

[**Lock**](Lock.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

