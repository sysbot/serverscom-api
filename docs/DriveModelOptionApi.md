# \DriveModelOptionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllDriveModelsOptionsForServerModel**](DriveModelOptionApi.md#ListAllDriveModelsOptionsForServerModel) | **Get** /v1/locations/{location_id}/order_options/server_models/{server_model_id}/drive_models | List all drive models options for server model
[**RetrieveAnExistingDriveModelOption**](DriveModelOptionApi.md#RetrieveAnExistingDriveModelOption) | **Get** /v1/locations/{location_id}/order_options/server_models/{server_model_id}/drive_models/{drive_model_id} | Retrieve an existing drive model option



## ListAllDriveModelsOptionsForServerModel

> []V1OrderOptionsDriveModelsBase ListAllDriveModelsOptionsForServerModel(ctx, locationId, serverModelId, optional)

List all drive models options for server model

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
**serverModelId** | **string**|  | 
 **optional** | ***ListAllDriveModelsOptionsForServerModelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllDriveModelsOptionsForServerModelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchPattern** | **optional.String**|  search pattern | 
 **mediaType** | **optional.String**|  media type | 
 **interface_** | **optional.String**|  interface | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]V1OrderOptionsDriveModelsBase**](v1-order_options-drive_models-_base.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingDriveModelOption

> V1OrderOptionsDriveModelsBase RetrieveAnExistingDriveModelOption(ctx, locationId, serverModelId, driveModelId)

Retrieve an existing drive model option

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
**serverModelId** | **string**|  | 
**driveModelId** | **string**|  | 

### Return type

[**V1OrderOptionsDriveModelsBase**](v1-order_options-drive_models-_base.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

