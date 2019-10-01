# \LoadBalancerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateANewL4LoadBalancers**](LoadBalancerApi.md#CreateANewL4LoadBalancers) | **Post** /v1/load_balancers/l4 | Create a new L4 load balancers
[**CreateANewL7LoadBalancers**](LoadBalancerApi.md#CreateANewL7LoadBalancers) | **Post** /v1/load_balancers/l7 | Create a new L7 load balancers
[**DeleteAnExisitingL4LoadBalancer**](LoadBalancerApi.md#DeleteAnExisitingL4LoadBalancer) | **Delete** /v1/load_balancers/l4/{id} | Delete an exisiting L4 load balancer
[**DeleteAnExistingL7LoadBalancer**](LoadBalancerApi.md#DeleteAnExistingL7LoadBalancer) | **Delete** /v1/load_balancers/l7/{id} | Delete an existing L7 load balancer
[**LoadBalancers**](LoadBalancerApi.md#LoadBalancers) | **Get** /v1/load_balancers | Load balancers
[**RetrieveAnExistingL4LoadBalancer**](LoadBalancerApi.md#RetrieveAnExistingL4LoadBalancer) | **Get** /v1/load_balancers/l4/{id} | Retrieve an existing L4 load balancer
[**RetrieveAnExistingL7LoadBalancer**](LoadBalancerApi.md#RetrieveAnExistingL7LoadBalancer) | **Get** /v1/load_balancers/l7/{id} | Retrieve an existing L7 load balancer
[**UpdateAnExistingL4LoadBalancer**](LoadBalancerApi.md#UpdateAnExistingL4LoadBalancer) | **Put** /v1/load_balancers/l4/{id} | Update an existing L4 load balancer
[**UpdateAnExistingL7LoadBalancer**](LoadBalancerApi.md#UpdateAnExistingL7LoadBalancer) | **Put** /v1/load_balancers/l7/{id} | Update an existing L7 load balancer



## CreateANewL4LoadBalancers

> InlineResponse201 CreateANewL4LoadBalancers(ctx, optional)
Create a new L4 load balancers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateANewL4LoadBalancersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateANewL4LoadBalancersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject5** | [**optional.Interface of InlineObject5**](InlineObject5.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateANewL7LoadBalancers

> InlineResponse2011 CreateANewL7LoadBalancers(ctx, optional)
Create a new L7 load balancers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateANewL7LoadBalancersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateANewL7LoadBalancersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject6** | [**optional.Interface of InlineObject6**](InlineObject6.md)|  | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnExisitingL4LoadBalancer

> DeleteAnExisitingL4LoadBalancer(ctx, id)
Delete an exisiting L4 load balancer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnExistingL7LoadBalancer

> DeleteAnExistingL7LoadBalancer(ctx, id)
Delete an existing L7 load balancer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadBalancers

> []TheItemsSchema2 LoadBalancers(ctx, optional)
Load balancers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LoadBalancersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoadBalancersOpts struct


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

[**[]TheItemsSchema2**](The Items Schema_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingL4LoadBalancer

> InlineResponse201 RetrieveAnExistingL4LoadBalancer(ctx, id)
Retrieve an existing L4 load balancer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingL7LoadBalancer

> InlineResponse2011 RetrieveAnExistingL7LoadBalancer(ctx, id)
Retrieve an existing L7 load balancer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnExistingL4LoadBalancer

> InlineResponse201 UpdateAnExistingL4LoadBalancer(ctx, id, optional)
Update an existing L4 load balancer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***UpdateAnExistingL4LoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAnExistingL4LoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject7** | [**optional.Interface of InlineObject7**](InlineObject7.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnExistingL7LoadBalancer

> InlineResponse2011 UpdateAnExistingL7LoadBalancer(ctx, id, optional)
Update an existing L7 load balancer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***UpdateAnExistingL7LoadBalancerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAnExistingL7LoadBalancerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject8** | [**optional.Interface of InlineObject8**](InlineObject8.md)|  | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

