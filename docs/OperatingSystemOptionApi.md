# \OperatingSystemOptionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllOperatingSystemsForServerModel**](OperatingSystemOptionApi.md#ListAllOperatingSystemsForServerModel) | **Get** /v1/locations/{location_id}/order_options/server_models/{server_model_id}/operating_systems | List all operating systems for server model
[**RetrieveAnExstingOperatingSystem**](OperatingSystemOptionApi.md#RetrieveAnExstingOperatingSystem) | **Get** /v1/locations/{location_id}/order_options/server_models/{server_model_id}/operating_systems/{operating_system_id} | Retrieve an exsting operating system



## ListAllOperatingSystemsForServerModel

> []V1OrderOptionsOperatingSystemsBase ListAllOperatingSystemsForServerModel(ctx, locationId, serverModelId, optional)

List all operating systems for server model

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
**serverModelId** | **string**|  | 
 **optional** | ***ListAllOperatingSystemsForServerModelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllOperatingSystemsForServerModelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]V1OrderOptionsOperatingSystemsBase**](v1-order_options-operating_systems-_base.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExstingOperatingSystem

> V1OrderOptionsOperatingSystemsBase RetrieveAnExstingOperatingSystem(ctx, locationId, serverModelId, operatingSystemId)

Retrieve an exsting operating system

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
**serverModelId** | **string**|  | 
**operatingSystemId** | **string**|  | 

### Return type

[**V1OrderOptionsOperatingSystemsBase**](v1-order_options-operating_systems-_base.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

