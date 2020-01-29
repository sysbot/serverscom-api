# \UplinkModelOptionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllUplinksForServerModel**](UplinkModelOptionApi.md#ListAllUplinksForServerModel) | **Get** /v1/locations/{location_id}/order_options/server_models/{server_model_id}/uplink_models | List all uplinks for server model
[**RetrieveAnExistingUplink**](UplinkModelOptionApi.md#RetrieveAnExistingUplink) | **Get** /v1/locations/{location_id}/order_options/server_models/{server_model_id}/uplink_models/{uplink_model_id} | Retrieve an existing uplink



## ListAllUplinksForServerModel

> []V1OrderOptionsUplinkModelsBase ListAllUplinksForServerModel(ctx, locationId, serverModelId, optional)

List all uplinks for server model

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
**serverModelId** | **string**|  | 
 **optional** | ***ListAllUplinksForServerModelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllUplinksForServerModelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **redundancy** | **optional.Bool**|  redundancy | 
 **type_** | **optional.String**|  type | 
 **operatingSystemId** | **optional.Int32**|  operating system | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]V1OrderOptionsUplinkModelsBase**](v1-order_options-uplink_models-_base.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingUplink

> V1OrderOptionsUplinkModelsBase RetrieveAnExistingUplink(ctx, locationId, serverModelId, uplinkModelId)

Retrieve an existing uplink

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
**serverModelId** | **string**|  | 
**uplinkModelId** | **string**|  | 

### Return type

[**V1OrderOptionsUplinkModelsBase**](v1-order_options-uplink_models-_base.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

