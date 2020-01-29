# \HostsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllHosts**](HostsApi.md#ListAllHosts) | **Get** /v1/hosts | List all hosts



## ListAllHosts

> []TheDedicatedServerEntitySchema ListAllHosts(ctx, optional)

List all hosts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAllHostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllHostsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchPattern** | **optional.String**|  search pattern | 
 **locationId** | **optional.Int32**|  location | 
 **type_** | **optional.String**|  type | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]TheDedicatedServerEntitySchema**](The_Dedicated_Server_Entity_Schema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

