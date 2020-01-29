# \LocationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Locations**](LocationApi.md#Locations) | **Get** /v1/locations | Locations
[**RetrieveAnExisitingLocation**](LocationApi.md#RetrieveAnExisitingLocation) | **Get** /v1/locations/{location_id} | Retrieve an exisiting location



## Locations

> []TheItemsSchema3 Locations(ctx, optional)

Locations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LocationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchPattern** | **optional.String**|  search pattern | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExisitingLocation

> TheItemsSchema3 RetrieveAnExisitingLocation(ctx, locationId)

Retrieve an exisiting location

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

