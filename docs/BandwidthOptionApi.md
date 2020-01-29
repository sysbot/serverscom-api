# \BandwidthOptionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllBandwidthForUplink**](BandwidthOptionApi.md#ListAllBandwidthForUplink) | **Get** /v1/locations/{location_id}/order_options/server_models/{server_model_id}/uplink_models/{uplink_model_id}/bandwidth | List all bandwidth for uplink
[**RetrieveAnExistingBandwidth**](BandwidthOptionApi.md#RetrieveAnExistingBandwidth) | **Get** /v1/locations/{location_id}/order_options/server_models/{server_model_id}/uplink_models/{uplink_model_id}/bandwidth/{bandwidth_id} | Retrieve an existing bandwidth



## ListAllBandwidthForUplink

> []V1OrderOptionsBandwidthBase ListAllBandwidthForUplink(ctx, locationId, serverModelId, uplinkModelId, optional)

List all bandwidth for uplink

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
**serverModelId** | **string**|  | 
**uplinkModelId** | **string**|  | 
 **optional** | ***ListAllBandwidthForUplinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllBandwidthForUplinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **optional.String**|  type | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]V1OrderOptionsBandwidthBase**](v1-order_options-bandwidth-_base.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingBandwidth

> V1OrderOptionsBandwidthBase RetrieveAnExistingBandwidth(ctx, locationId, serverModelId, uplinkModelId, bandwidthId)

Retrieve an existing bandwidth

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string**|  | 
**serverModelId** | **string**|  | 
**uplinkModelId** | **string**|  | 
**bandwidthId** | **string**|  | 

### Return type

[**V1OrderOptionsBandwidthBase**](v1-order_options-bandwidth-_base.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

