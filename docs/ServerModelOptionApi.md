# \ServerModelOptionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllServerModelsForLocation**](ServerModelOptionApi.md#ListAllServerModelsForLocation) | **Get** /v1/locations/{location_id}/order_options/server_models | List all server models for location
[**RetrieveAnExistingServerModel**](ServerModelOptionApi.md#RetrieveAnExistingServerModel) | **Get** /v1/locations/{location_id}/order_options/server_models/{server_model_id} | Retrieve an existing server model



## ListAllServerModelsForLocation

> []TheItemsSchema9 ListAllServerModelsForLocation(ctx, locationId, optional)

List all server models for location

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
 **optional** | ***ListAllServerModelsForLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllServerModelsForLocationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchPattern** | **optional.String**|  search pattern | 
 **hasRaidController** | **optional.Bool**|  has raid controller | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]TheItemsSchema9**](The_Items_Schema_9.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingServerModel

> TheItemsSchema10 RetrieveAnExistingServerModel(ctx, locationId, serverModelId)

Retrieve an existing server model

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
**serverModelId** | **string**|  | 

### Return type

[**TheItemsSchema10**](The_Items_Schema_10.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

