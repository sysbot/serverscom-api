# \CloudCredentialsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowCloudCredentialsToOpenstack**](CloudCredentialsApi.md#ShowCloudCredentialsToOpenstack) | **Get** /v1/cloud_computing/regions/{region_id}/credentials | Show cloud credentials to OpenStack



## ShowCloudCredentialsToOpenstack

> TheItemsSchema ShowCloudCredentialsToOpenstack(ctx, regionId)

Show cloud credentials to OpenStack

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **string**|  | 

### Return type

[**TheItemsSchema**](The_Items_Schema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

