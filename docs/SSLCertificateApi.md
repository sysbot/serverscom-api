# \SSLCertificateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateANewCustomSslCertificate**](SSLCertificateApi.md#CreateANewCustomSslCertificate) | **Post** /v1/ssl_certificates/custom | Create a new custom SSL certificate
[**RetrieveAnExistingCustomSslCertificate**](SSLCertificateApi.md#RetrieveAnExistingCustomSslCertificate) | **Get** /v1/ssl_certificates/custom/{id} | Retrieve an existing custom ssl certificate
[**SSLCertificates**](SSLCertificateApi.md#SSLCertificates) | **Get** /v1/ssl_certificates | SSL Certificates



## CreateANewCustomSslCertificate

> TheItemsSchema4 CreateANewCustomSslCertificate(ctx, optional)
Create a new custom SSL certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateANewCustomSslCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateANewCustomSslCertificateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject9** | [**optional.Interface of InlineObject9**](InlineObject9.md)|  | 

### Return type

[**TheItemsSchema4**](The Items Schema_4.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingCustomSslCertificate

> TheItemsSchema4 RetrieveAnExistingCustomSslCertificate(ctx, id)
Retrieve an existing custom ssl certificate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**TheItemsSchema4**](The Items Schema_4.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SSLCertificates

> []TheItemsSchema4 SSLCertificates(ctx, optional)
SSL Certificates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SSLCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SSLCertificatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchPattern** | **optional.String**|  search pattern | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]TheItemsSchema4**](The Items Schema_4.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

