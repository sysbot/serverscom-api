# \SSHKeyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNewSshKey**](SSHKeyApi.md#AddNewSshKey) | **Post** /v1/ssh_keys | Add new ssh key
[**DeleteSshKey**](SSHKeyApi.md#DeleteSshKey) | **Delete** /v1/ssh_keys/{fingerprint} | Delete ssh key
[**ListAllSshKeys**](SSHKeyApi.md#ListAllSshKeys) | **Get** /v1/ssh_keys | List all ssh keys
[**ShowSshKey**](SSHKeyApi.md#ShowSshKey) | **Get** /v1/ssh_keys/{fingerprint} | Show ssh key
[**UpdateTheNameOfSshKey**](SSHKeyApi.md#UpdateTheNameOfSshKey) | **Put** /v1/ssh_keys/{fingerprint} | Update the name of ssh key



## AddNewSshKey

> TheSshKeySchema AddNewSshKey(ctx, name, publicKey)

Add new ssh key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  name | 
**publicKey** | **string**|  public key | 

### Return type

[**TheSshKeySchema**](The_Ssh_key_Schema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSshKey

> DeleteSshKey(ctx, fingerprint)

Delete ssh key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fingerprint** | **string**|  | 

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


## ListAllSshKeys

> []TheSshKeySchema ListAllSshKeys(ctx, optional)

List all ssh keys

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAllSshKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllSshKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to created_at]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]TheSshKeySchema**](The_Ssh_key_Schema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowSshKey

> TheSshKeySchema ShowSshKey(ctx, fingerprint)

Show ssh key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fingerprint** | **string**|  | 

### Return type

[**TheSshKeySchema**](The_Ssh_key_Schema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTheNameOfSshKey

> TheSshKeySchema UpdateTheNameOfSshKey(ctx, name, fingerprint)

Update the name of ssh key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  name | 
**fingerprint** | **string**|  | 

### Return type

[**TheSshKeySchema**](The_Ssh_key_Schema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

