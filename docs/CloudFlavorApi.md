# \CloudFlavorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCloudFlavors**](CloudFlavorApi.md#ListCloudFlavors) | **Get** /v1/cloud_computing/regions/{region_id}/flavors | List cloud flavors



## ListCloudFlavors

> []TheItemsSchema1 ListCloudFlavors(ctx, regionId, optional)

List cloud flavors

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **string**|  | 
 **optional** | ***ListCloudFlavorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCloudFlavorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]

### Return type

[**[]TheItemsSchema1**](The_Items_Schema_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

