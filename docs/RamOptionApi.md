# \RamOptionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllAvailableRamOptionsForServerModel**](RamOptionApi.md#ListAllAvailableRamOptionsForServerModel) | **Get** /v1/locations/{location_id}/order_options/server_models/{server_model_id}/ram | List all available ram options for server model



## ListAllAvailableRamOptionsForServerModel

> []TheItemsSchema4 ListAllAvailableRamOptionsForServerModel(ctx, locationId, serverModelId, optional)

List all available ram options for server model

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
**serverModelId** | **string**|  | 
 **optional** | ***ListAllAvailableRamOptionsForServerModelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllAvailableRamOptionsForServerModelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to ram]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]TheItemsSchema4**](The_Items_Schema_4.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

